const {
    request
} = require('../../../../utils/request')
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        editSize: '40rpx',
        addSize: '60rpx',
        arrowSize: '30rpx',
        autoHeight: 40,
        showWx: false,
        targets: [],
        works: [],
        projects: [],
        educations: [],
        user: {},
        statuses: [],
        popup: {
            status: false,
        },
        firstTargetIndex: 1,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let showWx = wx.getStorageSync('showWechat')
        this.setData({
            showWx,
        })
        this.getStatus()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        this.getUserInfo()
    },

    getUserInfo: function () {
        let that = this
        request({
            url: '/user/info',
        }).then(res => {
            let user = res.data
            let description = ''
            let index = 0
            let firstTargetIndex = 0
            for (const item of user.applies) {
                if (item.isFirst) {
                    description = item.description
                    break
                }
                index += 1
            }
            firstTargetIndex = index
            user.descriptions = description.split('\n')
            wx.setStorageSync('user', user)
            app.globalData.userInfo = user
            let targets = user.applies
            let works = user.works
            let projects = user.projects
            let educations = user.educations
            user.workYear = this.calcYear(user.workday)
            user.age = this.calcYear(user.birthday)
            user.tagItems = user.tags.split(' ')
            let degree = ''
            for (let item of educations) {
                if (item.degree == '博士') {
                    degree = item.degree
                    break
                }
                if (item.degree == '研究生') {
                    degree = item.degree
                    break
                }
                if (item.degree == '本科') {
                    degree = item.degree
                    break
                }
                if (item.degree == '专科') {
                    degree = item.degree
                    break
                }
            }
            user.degree = degree
            that.setData({
                targets,
                works,
                projects,
                educations,
                user,
                firstTargetIndex,
            })
        })
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
            ['user.currentState']: state
        })
        this.updateState(state)
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

    calcYear(date) {
        if (!date) {
            return 0
        }
        date = date.replace(/\./g, '-')
        let now = new Date()
        let before = new Date(date)
        let nowYear = now.getFullYear()
        let beforeYear = before.getFullYear()
        return nowYear - beforeYear
    },

    showResume() {
        const that = this
        const { user } = that.data
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
    },

    uploadResume() {
        let that = this
        const { data } = that
        const user = data.user
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
})