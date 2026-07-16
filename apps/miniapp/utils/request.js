const Config = require('./config')

// 串行化的重新登录：多个请求同时 401 时只触发一次 wx.login
let reloginPromise = null

function relogin() {
    if (reloginPromise) {
        return reloginPromise
    }
    reloginPromise = new Promise(function (resolve, reject) {
        wx.login({
            success: function (loginRes) {
                wx.request({
                    url: Config.domain + Config.prefix + '/account/loginWechat',
                    method: 'POST',
                    data: { code: loginRes.code },
                    header: { 'Content-Type': 'application/json' },
                    success: function (res) {
                        if (res.statusCode == 200 && res.data && res.data.code == 0) {
                            let data = res.data.data
                            wx.setStorageSync('openid', data.openid)
                            wx.setStorageSync('token', data.token)
                            wx.setStorageSync('unionid', data.unionid)
                            let app = getApp()
                            if (app && app.globalData) {
                                app.globalData.token = data.token
                            }
                            resolve(data.token)
                        } else {
                            reject(res)
                        }
                    },
                    fail: reject,
                })
            },
            fail: reject,
        })
    })
    // 无论成功失败都置空，允许下次重新触发
    reloginPromise.then(function () {
        reloginPromise = null
    }, function () {
        reloginPromise = null
    })
    return reloginPromise
}

function doRequest(params) {
    return new Promise(function (resolve, reject) {
        let token = wx.getStorageSync('token')
        wx.request({
            url: Config.domain + Config.prefix + params.url,
            method: params.method || 'GET',
            data: params.data || {},
            header: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            timeout: 15000,
            success(res) {
                if (res.statusCode == 200) {
                    if (res.data.code == 0) {
                        resolve(res.data);
                    } else {
                        wx.showToast({
                            title: res.data.message || '系统错误',
                            showCancel: false,
                            icon: 'error',
                        })
                        reject(res);
                    }
                } else if (res.statusCode == 401) {
                    // token 失效：交给上层重新登录后重试，此处不弹提示
                    reject({ __needRelogin: true, res: res });
                } else {
                    wx.showToast({
                        title: '网络请求超时！',
                        icon: 'error',
                    })
                    reject();
                }
            },
            fail(err) {
                wx.showToast({
                    title: '网络错误！',
                    icon: 'error',
                })
                reject(err)
            }
        })
    })
}

function request(params) {
    return doRequest(params).catch(function (err) {
        // 仅在 token 失效且本次尚未重试时，重新登录并重试一次
        if (err && err.__needRelogin) {
            if (params.__retried) {
                wx.showToast({ title: '登录已过期，请重试', icon: 'none' })
                return Promise.reject(err.res)
            }
            return relogin().then(function () {
                let retryParams = Object.assign({}, params, { __retried: true })
                return request(retryParams)
            }, function () {
                wx.showToast({ title: '登录失败，请稍后重试', icon: 'none' })
                return Promise.reject(err.res)
            })
        }
        return Promise.reject(err)
    })
}

module.exports = {
    request,
}
