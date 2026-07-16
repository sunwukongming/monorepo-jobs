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
        minDate: 0,
        maxDate: 0,
        form: {
            company: '',
            industry: '',
            startTime: '',
            endTime: '',
            position: '',
            content: '',
            skills: '',
            performance: '',
        },
        popup: {
            industry: false,
            startTime: false,
            endTime: false,
        },
        industryParent: [],
        industries: [],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        const { action, index } = options
        if (action == 'edit') {
            let user = wx.getStorageSync('user')
            let items = user.works
            let item = items[index]
            this.data.item = item
            let form = {
                company: item.company,
                industry: item.industry,
                startTime: item.startTime,
                endTime: item.endTime,
                position: item.position,
                content: item.content,
                skills: item.skills,
                performance: item.performance,
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
        this.getIndustry()
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

    getIndustry() {
        let that = this
        request({
            url: '/dictionary/industries',
        }).then(res => {
            const {
                list
            } = res.data
            let parents = [],
                industries = {}
            for (let item of list) {
                parents.push({
                    id: item.id,
                    text: item.name,
                })
                let children = []
                if (item.children && item.children.length > 0) {
                    for (let child of item.children) {
                        children.push({
                            id: child.id,
                            text: child.name
                        })
                    }
                }
                industries[item.name] = children
            }
            let industryParent = [{
                    values: parents,
                    defaultIndex: 0,
                },
                {
                    values: industries[parents[0].text],
                    defaultIndex: 0,
                }
            ]
            that.setData({
                industryParent,
                industries,
            })
        })
    },

    showPopupIndustry() {
        this.setData({
            ['popup.industry']: true,
        })
    },

    closePopupIndustry() {
        this.setData({
            ['popup.industry']: false,
        })
    },

    confirmPopupIndustry(e) {
        const value = e.detail.value
        let industry = value[0].text + '/' + (value[1] && value[1].text)
        this.setData({
            ['popup.industry']: false,
            ['form.industry']: industry,
        })
    },

    onChangeIndustry(e) {
        const value = e.detail.value
        let parents = this.data.industryParent
        parents[1].values = this.data.industries[value[0].text]
        this.setData({
            industryParent: parents,
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
            url = '/user/createWork'
        } else {
            url = '/user/updateWork'
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
                        url: '/user/deleteWork',
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