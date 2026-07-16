const {
    request
} = require('../../../utils/request')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        description: '',
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
        this.setData({
            description: user.description,
        })
    },

    changeValue(e) {
        const { value } = e.detail
        this.setData({
            description: value,
        })
    },
    
    goBack() {
        wx.navigateBack()
    },

    submit() {
        let that = this
        const { description } = that.data
        wx.showLoading({
          title: '保存中...',
        })
        request({
            url: '/user/update',
            method: 'POST',
            data: {
                description,
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
})