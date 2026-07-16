const { request } = require('../../../utils/request')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        items: [],
        children: [],
        showPopup: false,
        statusBarHeight: wx.getStorageSync('statusBarHeight'),
        navBarHeight: wx.getStorageSync('navBarHeight'),
        selectedItem: {},
        selectedChild: {},
        from: 'home',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let from = options.from || 'home'
        let isCompany = options.isCompany || false
        this.data.from = from
        this.data.isCompany = isCompany
        wx.setStorageSync('lastPage', 'industry')
        this.getList()
    },

    getList: function () {
        let that = this
        request({
            url: '/dictionary/industries',
        }).then(res => {
            const { data } = res
            let items = data.list
            let children = []
            if (items.length > 0) {
                children = items[0].children
            }
            for (let item of items) {
                let children = [{
                    id: item.id,
                    name: '全' + item.name + '领域',
                    path: item.path,
                    pid: item.pid,
                }].concat(item.children)
                item.children = children
            }
            items = [{
                name: '全部行业',
                path: '',
                children: [],
            }].concat(items)
            that.setData({
                items,
                children,
            })
        })
    },

    clickItem: function (e) {
        let that = this
        let id = e.currentTarget.dataset.id
        let items = that.data.items
        let selectedItem = {}
        for (let item of items) {
            if (item.id == id) {
                selectedItem = item
                break;
            }
        }
        that.setData({
            selectedItem,
        })
        if (selectedItem.children.length > 0) {
            that.setData({
                showPopup: true,
                children: selectedItem.children,
            })
        } 
        else {
            if (that.data.from === 'home') {
                if (that.data.isCompany == 1) {
                    wx.setStorageSync('selectedHomeCompanyIndustry', selectedItem)
                } else {
                    wx.setStorageSync('selectedHomeIndustry', selectedItem)
                }
            } 
            else {
                wx.setStorageSync('selectedJobIndustry', selectedItem)
            }
            wx.navigateBack({
              delta: 0,
            })
        }
    },

    clickChild: function (e) {
        let that = this
        let id = e.currentTarget.dataset.id
        let child = {}
        for (let item of that.data.children) {
            if (item.id == id) {
                child = item
                break;
            }
        }
        if (this.data.from == 'home') {
            if (that.data.isCompany == 1) {
                wx.setStorageSync('selectedHomeCompanyIndustry', child)
            } else {
                wx.setStorageSync('selectedHomeIndustry', child)
            }
        } else {
            wx.setStorageSync('selectedJobIndustry', child)
        }
        wx.navigateBack({
          delta: 0,
        })
    },

    closePopup: function (e) {
        this.setData({
            showPopup: false,
        })
    },
})