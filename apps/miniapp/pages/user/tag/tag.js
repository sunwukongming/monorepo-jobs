const {
    request
} = require('../../../utils/request')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        items: [],
        tag: '',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        let user = wx.getStorageSync('user')
        let items = user.tags.split(' ')
        this.setData({
            items,
        })
    },

    changeValue(e) {
        const { value } = e.detail
        this.setData({
            tag: value,
        })
    },

    addTag() {
        let { items, tag } = this.data
        items.push(tag)
        this.setData({
            items,
            tag: '',
        })
    },

    delTag(e) {
        const { dataset } = e.currentTarget
        const { index } = dataset
        let { items } = this.data
        items.splice(index, 1)
        this.setData({
            items,
        })
    },

    submit(currentState) {
        let that = this
        const { items } = that.data
        wx.showLoading({
          title: '保存中...',
        })
        request({
            url: '/user/update',
            method: 'POST',
            data: {
                tags: items.join(' '),
            },
        }).then(res => {
            wx.hideLoading()
            wx.showToast({
              title: '保存成功',
            })
            setTimeout(() => {
                wx.navigateBack()
            }, 1000)
        }).catch(() => {
            wx.hideLoading()
        })
    },

    goBack() {
        wx.navigateBack()
    },
})