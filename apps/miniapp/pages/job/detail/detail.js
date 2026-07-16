const { request } = require('../../../utils/request')
const { formatTime, handleEscapeChar } = require('../../../utils/util')

Page({

    /**
     * 页面的初始数据
     */
    data: {
        id: '',
        showModal: false,
        showPhoneModal: false,
        item: {},
        similarItems: [],
        showAction: false,
        actionTitle: '请选择您的操作',
        actions: [
            {name: '呼叫'},
            {name: '复制内容'},
        ],
        user: {},
        showHelpModal: false,
        helpForm: {
            company: '',
            position: '',
            helpPlan: '',
        },
        textarea: {
            minHeight: 80,
            maxHeight: 100,
        },
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        if ('scene' in options) {
            const scene = decodeURIComponent(options.scene)
            options.id = scene
        }
        
        wx.setStorageSync('lastPage', 'job-detail')
        this.setData({
            id: options.id
        })
        this.getData()
    },

    onShow() {
        let user = wx.getStorageSync('user')
        this.setData({
            user,
        })
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {
        let that = this
        that.setData({
            showModal: false
        })
        return {
            title: '求职目标：' + that.data.item.destPosition || '奖聘求职推荐',
            imageUrl: '',
        }
    },

    getData() {
        let that = this
        request({
            url: '/apply/detail',
            method: 'POST',
            data: {
                id: that.data.id
            },
        }).then(res => {
            let item = res.data
            let items = []
            for (let job of item.similarApplies) {
                items.push({
                    ...job,
                    updatedAt: formatTime(job.updatedTime * 1000)
                })
            }
            item.descriptions = item.description.split('\n')
            item.updatedAt = formatTime(item.updatedTime * 1000)
            that.setData({
                item,
                similarItems: items,
            })
            wx.setStorageSync('job', that.data.item)
        })
    },

    onCollect() {
        let that = this
        let item = that.data.item
        let url = ''
        if (item.isLike) {
            url = '/apply/unlike'
        }
        else {
            url = '/apply/like'
        }
        request({
            url: url,
            method: 'POST',
            data: {
                id: that.data.id,
            },
        }).then(res => {
            that.getData()
        })
    },

    toCanvas(e) {
        let that = this
        let item = that.data.item
        wx.setStorageSync('job', item)
        wx.navigateTo({
            url: '../canvas/canvas',
        })
    },

    /**
     * 打开弹窗
     */
    openModal() {
        let that = this
        that.setData({
            showModal: true,
        })
    },

    /**
     * 关闭弹窗
     */
    closeModal() {
        let that = this
        that.setData({
            showModal: false,
        })
    },

    clickPhoneNumber: function (e) {
        let that = this
        let item = that.data.item
        that.setData({
            actionTitle: `${item.contact} 可能是一个电话、微信号或，你可以`,
            showAction: true,
        })
    },

    onActionClick(e) {
        let that = this
        let item = that.data.item
        let value = e.detail.name
        switch (value) {
            case '呼叫': {
                wx.makePhoneCall({
                    phoneNumber: item.contact,
                })
                break
            }
            case '复制内容': {
                wx.setClipboardData({
                    data: item.contact,
                    success() {
                        wx.showToast({
                            title: '内容复制成功',
                        })
                    },
                    fail(err) {
                        wx.showToast({
                            title: '内容复制失败',
                        })
                    },
                })
            }
        }
        that.setData({
            showAction: false,
        })
    },

    closeAction() {
        let that = this
        that.setData({
            showAction: false,
        })
    },

    copyContent(e) {
        let that = this
        let editContent = that.data.item.description
        wx.setClipboardData({
            data: editContent,
            success() {
                wx.showToast({
                    title: '复制成功',
                })
            },
            fail(err) {
                wx.showToast({
                    icon: 'error',
                    title: '复制失败',
                })
            },
        })
    },

    toDetail: function (e) {
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
          url: 'detail?id=' + id,
        })
    },

    showSalary2(e) {
        const { dataset } = e.currentTarget
        wx.showModal({
            title: '提示',
            content: `当前协助奖金为：${dataset.num}`,
            showCancel: false,
            complete: (res) => {
                if (res.confirm) {
                
                }
            }
        })
    },

    showHelpModal() {
        const that = this
        const { id, helpForm } = that.data
        request({
            url: '/help/info',
            method: 'POST',
            data: {
                accountApplyId: id,
            },
        }).then(res => {
            const rspData = res.data
            helpForm.company = rspData.company
            helpForm.position = rspData.position
            helpForm.helpPlan = rspData.helpPlan 
            that.setData({
                showHelpModal: true,
                helpForm,
            })
        })
    },

    closeHelpModal() {
        this.setData({
            showHelpModal: false,
        })
    },

    onChangeHelpField(e) {
        const that = this
        const { detail } = e
        const { helpForm } = that.data
        helpForm[e.currentTarget.dataset.field] = detail
        that.setData({
            helpForm,
        })
    },

    submitHelpModal() {
        const that = this
        const { id, helpForm } = that.data
        const user = wx.getStorageSync('user')
        if (!user.mobile) {
            that.setData({
                showLogin: true,
            })
            return false
        }
        request({
            url: '/help/apply',
            method: 'POST',
            data: {
                accountApplyId: id,
                ...helpForm,
            },
        }).then(res => {
            that.setData({
                showHelpModal: false,
            })
            wx.showToast({
              title: '提交成功',
            })
        })
    },

    toHelpPlanRule() {
        wx.navigateTo({
            url: '/pages/discover/detail/detail?id=27&menuId=7',
        })
    },

    openPrivacy() {
        wx.openPrivacyContract({
            success: () => {}, // 打开成功
            fail: () => {}, // 打开失败
            complete: () => {}
        })
    },

    getPhoneNumber (e) {
        let that = this
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
            that.getUserInfo()
            that.setData({
                showLogin: false,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    getUserInfo: function () {
        let that = this
        request({
            url: '/user/info',
        }).then(res => {
            wx.hideLoading()
            let data = res.data
            wx.setStorageSync('user', data)
            app.globalData.userInfo = data
            let hasMobile = false
            if (data.mobile) {
                hasMobile = true
            }
            that.setData({
                showLogin: !hasMobile,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    handleAgreePrivacyAuthorization() {
        // 用户同意隐私协议事件回调
    },

    loginCancel() {
        this.setData({
            showLogin: false,
            showHelpModal: false,
        })
    },

    showResume() {
        const {
            data
        } = this
        const {
            item
        } = data
        const {
            resumeAuth,
            resumeUrl
        } = item
        if (!resumeAuth) {
            wx.showModal({
                title: '温馨提示',
                content: '尊敬的用户您好，出于对候选人信息的保护，该附件简历仅限已获授权的企业HR及协作用户查看。有需求的用户可点击下方按钮与奖聘联络，候选人同意后即可查看 ，感谢支持！！',
                showCancel: false,
                confirmText: '我知道了',
                complete: (res) => {
                    if (res.confirm) {
                    }
                }
            })
            return false
        }
        if (!resumeUrl) {
            wx.showToast({
                icon: 'error',
                title: '暂无简历',
            })
            return false
        }
        wx.downloadFile({
            url: resumeUrl,
            success(res) {
                wx.openDocument({
                    showMenu: true,
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
    },
})