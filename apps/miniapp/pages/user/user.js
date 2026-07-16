const { request } = require('../../utils/request')
const { domain, prefix } = require('../../utils/config')
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        menuSize: '60rpx',
        statuses: [],
        popup: {
            status: false,
        },
        userInfo: {},
        targets: [],
        showContactModal: false,
        showUploadModal: false,
        uploadProgress: 0,
        uploadTips: '处理中',
        showLogin: false,
        year: '',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        wx.setStorageSync('lastPage', 'user')
        this.getStatus()
        const year = new Date().getFullYear()
        this.setData({
            year,
        })
    },

    onShow: function () {
        let that = this
        this.getUserInfo()
    },

    getUserInfo: function () {
        let that = this
        request({
            url: '/user/info',
        }).then(res => {
            let data = res.data
            wx.setStorageSync('user', data)
            app.globalData.userInfo = data
            let hasMobile = false
            if (data.mobile) {
                hasMobile = true
            }
            that.setData({
                userInfo: data,
                targets: data.applies,
                showLogin: !hasMobile,
            })
        })
    },

    getStatus: function () {
        let that = this
        request({
            url: '/dictionary/data',
        }).then(res => {
            const { data } = res
            const { currentState } = data
            const { list } = currentState
            let items = []
            for (let item of list) {
                items.push(item.remark)
            }
            that.setData({
                statuses: items
            })            
        })
    },

    showPopupStatus: function () {
        this.setData({
            ['popup.status']: true,
        })
    },

    closePopupStatus: function () {
        this.setData({
            ['popup.status']: false,
        })
    },

    confirmPopupStatus: function (e) {
        const state = e.detail.value
        this.setData({
            ['popup.status']: false,
            ['userInfo.currentState']: state
        })
        this.updateState(state)
    },

    toPage(e) {
        const { dataset } = e.currentTarget
        let path = ''
        let query = '?'
        for (const key in dataset) {
            if (key == 'path') {
                path = dataset[key]
            } else {
                query += `${key}=${dataset[key]}&`
            }
        }
        query = query.substr(0, query.length - 1)
        wx.navigateTo({
            url: path + query,
        })
    },

    toContact: function () {
        wx.navigateTo({
          url: '/pages/discover/detail/detail?id=4&menuId=7',
        })
        // this.setData({
        //     showContactModal: true,
        // })
    },

    onCloseContactModal: function () {
        this.setData({
            showContactModal: false,
        })
    },

    updateState(currentState) {
        let that = this
        request({
            url: '/user/update',
            method: 'POST',
            data: {
                currentState,
            },
        }).then(res => {
            // that.getUserInfo()
        })
    },

    showTips() {
        wx.showToast({
          title: '您还没有奖金，请继续努力推荐!',
          icon: 'none',
        })
    },

    targetToTop() {
        let that = this
        const { targets } = that.data
        let item = null
        for(let target of targets) {
            if (target.isFirst == 1) {
                target.isTop = 1
                item = target
                break
            }
        }
        wx.showModal({
            title: '温馨提示',
            content: '该操作将置顶求职目标中的优选目标，是否继续？',
            confirmText: '继续',
            complete: (res) => {
                if (res.confirm) {
                    that.updateTarget(item)
                }
            }
        })
    },

    updateTarget(item) {
        wx.showLoading({
            title: '置顶中...',
        })
        request({
            url: '/user/updateApply',
            method: 'post',
            data: item,
        }).then(res => {
            wx.hideLoading()
            wx.showToast({
                title: '置顶成功',
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    getPhoneNumber (e) {
        let code = e.detail.code
        if (!code) {
            wx.showToast({
                icon: 'none',
                title: '获取手机号失败',
            })
            return false
        }
        wx.showLoading({
          title: '登录中...',
        })
        request({
            url: '/account/bindWechatMobile',
            method: 'POST',
            data: {
                code,
            },
        }).then(res => {
            wx.hideLoading()
            this.setData({
                showLogin: false,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    goBack() {
        let tabPage = wx.getStorageSync('tabPage')
        wx.switchTab({
          url: tabPage,
        })
    },

    handleAgreePrivacyAuthorization() {
        // 用户同意隐私协议事件回调
    },

    openPrivacy() {
        wx.openPrivacyContract({
            success: () => {}, // 打开成功
            fail: () => {}, // 打开失败
            complete: () => {}
        })
    },

    uploadResume() {
        let that = this
        const { data } = that
        const user = data.userInfo
        let content = '请选择您的简历进行上传'
        let cancelText = '取消'
        let flag = false
        if (user.resumeUrl) {
            content = '上传新的简历会将原有的覆盖，是否继续？'
            flag = true
            cancelText = '查看'
        }
        wx.showModal({
            title: '温馨提示',
            content: content,
            cancelText: cancelText,
            confirmText: '去上传',
            complete: (res) => {
                if (res.cancel) {
                    wx.downloadFile({
                        url: user.resumeUrl,
                        success(res) {
                            wx.openDocument({
                                filePath: res.tempFilePath,
                            })
                        },
                        fail(err) {
                            wx.showToast({
                                icon: 'error',
                                title: '下载失败',
                            })
                        }
                    })
                }
                // 上传简历
                if (res.confirm) {
                    wx.chooseMessageFile({
                        count: 1,
                        type: 'file',
                        success(res) {
                            that.setData({
                                showUploadModal: true,
                                uploadProgress: 0,
                                uploadTips: '识别中',
                            })
                            const { tempFiles } = res
                            let file = null
                            if (tempFiles.length > 0) {
                                file = res.tempFiles[0]
                            } else {
                                return false
                            }
                            let token = wx.getStorageSync('token')
                            const uploadTask = wx.uploadFile({
                                header: {
                                    'Authorization': `Bearer ${token}`,
                                },
                                filePath: file.path,
                                name: 'file',
                                url: `${domain}${prefix}/resume/upload`,
                                timeout: 60000,
                                success(res) {
                                    if (res.statusCode != 200) {
                                        wx.showToast({
                                          title: '简历上传失败',
                                          icon: 'error',
                                        })
                                        return false
                                    }
                                    const rsp = res.data
                                    const rspData = JSON.parse(rsp)
                                    if (rspData.code != 0) {
                                        wx.showToast({
                                            icon: 'error',
                                            title: '上传失败',
                                        })
                                    } else {
                                        wx.showToast({
                                            title: '上传成功',
                                        })
                                    }
                                },
                                fail(err) {
                                    wx.showToast({
                                        icon: 'error',
                                        title: '上传失败',
                                    })
                                },
                                complete() {
                                    // wx.hideLoading()
                                    uploadTask.offProgressUpdate()
                                    that.setData({
                                        showUploadModal: false,
                                    })
                                }
                            })
                            uploadTask.onProgressUpdate((res) => {
                                const { progress } = res
                                let uploadTips = `${progress}%`
                                if (progress == 100) {
                                    uploadTips = '文件识别中'
                                }
                                that.setData({
                                    uploadProgress: progress,
                                    uploadTips,
                                })
                            })
                        },
                    })
                }
            }
        })
    },

    toRecommend() {
        wx.navigateTo({
          url: './recommend/recommend',
        })
    },

    toSelfRecommend() {
        wx.navigateTo({
            url: './self-recommend/self-recommend',
        })
    },

    toHelpPlanRule() {
        wx.navigateTo({
            url: '/pages/discover/detail/detail?id=27&menuId=7',
        })
    },
})