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
            degree: '本科',
            profession: '',
            startTime: '',
            endTiem: '',
            experience: '',
        },
        popup: {
            degree: false,
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
            let items = user.educations
            let item = items[index]
            this.data.item = item
            let form = {
                name: item.name,
                degree: item.degree,
                profession: item.profession,
                startTime: item.startTime,
                endTime: item.endTime,
                experience: item.experience,
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

    showPopupDegree() {
        this.setData({
            ['popup.degree']: true,
        })
    },

    closePopupDegree() {
        this.setData({
            ['popup.degree']: false,
        })
    },

    confirmPopupDegree(e) {
        const value = e.detail.value
        this.setData({
            ['popup.degree']: false,
            ['form.degree']: value,
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
        let items = date.split('-')
        items.pop()
        date = items.join('-')
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
        let items = date.split('-')
        items.pop()
        date = items.join('-')
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
            url = '/user/createEducation'
        } else {
            url = '/user/updateEducation'
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
                        url: '/user/deleteEducation',
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