const {
    request
} = require('../../../../utils/request')
const {
    formatDate
} = require('../../../../utils/util')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        action: 'add',
        item: {},
        form: {
            name: '',
            role: '',
            startTime: '',
            endTime: '',
            description: '',
            performance: '',
            link: '',
        },
        popup: {
            startTime: false,
            endTime: false,
        },
        minDate: 0,
        maxDate: 0,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        const { action, index } = options
        if (action == 'edit') {
            let user = wx.getStorageSync('user')
            let items = user.projects
            let item = items[index]
            this.data.item = item
            let form = {
                name: item.name,
                role: item.role,
                startTime: item.startTime,
                endTime: item.endTime,
                description: item.description,
                performance: item.performance,
                link: item.link,
            }
            this.setData({
                form,
            })
        }
        this.setData({
            action,
        })
        let date = new Date()
        let maxDate = date.getTime()
        date.setFullYear(date.getFullYear() - 43)
        let minDate = date.getTime()
        this.setData({
            minDate,
            maxDate,
        })
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    changeValue (e) {
        let key = e.currentTarget.dataset.name
        let value = e.detail.value
        this.setData({
            [`form.${key}`]: value
        })
    },

    showPopupStartTime() {
        this.setData({
            ['popup.startTime']: true,
        })
    },

    closePopupStartTime() {
        this.setData({
            ['popup.startTime']: false,
        })
    },

    confirmPopupStartTime(e) {
        const timestamp = e.detail
        let date = formatDate(timestamp)
        this.setData({
            ['popup.startTime']: false,
            ['form.startTime']: date,
        })
    },

    showPopupEndTime() {
        this.setData({
            ['popup.endTime']: true,
        })
    },

    closePopupEndTime() {
        this.setData({
            ['popup.endTime']: false,
        })
    },

    confirmPopupEndTime(e) {
        const timestamp = e.detail
        let date = formatDate(timestamp)
        this.setData({
            ['popup.endTime']: false,
            ['form.endTime']: date,
        })
    },

    submit() {
        let that = this
        const {
            action,
            form,
            item,
        } = that.data
        let url = '',
            params = {
                ...form,
            }
        if (action == 'add') {
            url = '/user/createProject'
        } else {
            url = '/user/updateProject'
            params.id = item.id
        }
        wx.showLoading({
            title: '加载中...',
        })
        request({
            url: url,
            method: 'post',
            data: params,
        }).then(res => {
            wx.hideLoading()
            wx.showToast({
                title: '保存成功',
            })
            setTimeout(function() {
                wx.navigateBack()
            }, 1000)
        }).catch(() => {
            wx.hideLoading()
        })
    },

    del() {
        let that = this
        if (that.data.action == 'add') {
            wx.navigateBack()
            return false
        }
        wx.showModal({
            title: '温馨提示',
            content: '删除操作不可撤销，是否继续',
            confirmText: '继续',
            complete: (res) => {
                if (res.confirm) {
                    wx.showLoading({
                        title: '删除中...',
                    })
                    request({
                        url: '/user/deleteProject',
                        method: 'post',
                        data: {
                            id: that.data.item.id
                        },
                    }).then(res => {
                        wx.hideLoading()
                        wx.showToast({
                            title: '删除成功',
                        })
                        setTimeout(() => {
                            wx.navigateBack()
                        }, 1000)
                    }).catch(() => {
                        wx.hideLoading()
                    })
                }
            }
        })
    },
})