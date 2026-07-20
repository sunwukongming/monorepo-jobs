const { request } = require('./utils/request.js')
const { getSystemInfoCompat, isHarmonyOS } = require('./utils/system.js')
App({
    onLaunch() {
        let that = this

        // 计算自定义的标题栏高度（含 HarmonyOS / ohos 平台判断）
        let systemInfo = getSystemInfoCompat()
        that.systemInfo = systemInfo
        that.globalData.platform = systemInfo.platform || ''
        that.globalData.isHarmonyOS = isHarmonyOS(systemInfo.platform)
        let rect = null
        try {
            rect = wx.getMenuButtonBoundingClientRect ? wx.getMenuButtonBoundingClientRect() : null
            if (rect === null) {
                throw 'getMenuButtonBoundingClientRect error'
            }
            //取值为0的情况
            if (!rect.width) {
                throw 'getMenuButtonBoundingClientRect error'
            }
        } catch (error) {
            let gap = '' //胶囊按钮上下间距 使导航内容居中
            let width = 96 //胶囊的宽度，android大部分96，ios为88
            if (systemInfo.platform === 'android' || isHarmonyOS(systemInfo.platform)) {
                // HarmonyOS 手机端按 Android 间距处理更稳妥
                gap = 8
                width = 96
            } else if (systemInfo.platform === 'devtools') {
                if ((systemInfo.system || '').includes('iOS')) {
                    gap = 5.5 //开发工具中ios手机
                } else {
                    gap = 7.5 //开发工具中android和其他手机
                }
            } else {
                gap = 4
                width = 88
            }
            if (!systemInfo.statusBarHeight) {
                //开启wifi的情况下修复statusBarHeight值获取不到
                systemInfo.statusBarHeight = systemInfo.screenHeight - systemInfo.windowHeight - 20
            }
            rect = {
                //获取不到胶囊信息就自定义重置一个
                bottom: systemInfo.statusBarHeight + gap + 32,
                height: 32,
                left: systemInfo.windowWidth - width - 10,
                right: systemInfo.windowWidth - 10,
                top: systemInfo.statusBarHeight + gap,
                width: width
            }
        }
        let navBarHeight = (function () { //导航栏高度
            let gap = rect.top - systemInfo.statusBarHeight //动态计算每台手机状态栏到胶囊按钮间距
            return 2 * gap + rect.height
        })()
        wx.setStorageSync('statusBarHeight', systemInfo.statusBarHeight)
        wx.setStorageSync('navBarHeight', navBarHeight)

        // 登录 gate：页面通过 app.whenReady() 等待登录完成后再加载数据，
        // 避免在 token / user 就绪前发请求（替代此前各页面的 setInterval 轮询）
        that.readyPromise = new Promise(resolve => {
            that.readyResolve = resolve
        })
        wx.login({
            success: res => {
                request({
                    url: '/account/loginWechat',
                    method: 'POST',
                    data: {
                        code: res.code,
                    }
                }).then(res => {
                    let { data } = res
                    wx.setStorageSync('openid', data.openid)
                    wx.setStorageSync('token', data.token)
                    wx.setStorageSync('unionid', data.unionid)
                    that.globalData.token = data.token
                    // 拉取用户信息后再放行，保证依赖 user 的页面能拿到数据
                    request({
                        url: '/user/info',
                    }).then(res => {
                        let data = res.data
                        wx.setStorageSync('user', data)
                        that.globalData.userInfo = data
                    }).catch(() => {}).then(() => {
                        that.readyResolve(that.globalData.token)
                    })
                    request({
                        url: '/status',
                    }).then(res => {
                        let data = res.data
                        wx.setStorageSync('showWechat', !data.check)
                    })
                }).catch(() => {
                    // 登录失败也放行，后续请求各自报错提示，避免页面永久卡住
                    that.readyResolve(that.globalData.token)
                })
            },
            fail: () => {
                that.readyResolve(that.globalData.token)
            }
        })

        // 基本设置
        wx.setStorageSync('lastPage', '')
    },
    // 返回登录完成的 Promise；登录已完成时立即 resolve
    whenReady() {
        return this.readyPromise || Promise.resolve(this.globalData.token)
    },
    globalData: {
        userInfo: null,
        openid: null,
        token: null,
        discoverMenu: 1,
        platform: '',
        isHarmonyOS: false,
    }
})
