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
        this.data.from = from
        wx.setStorageSync('lastPage', 'profession')
        this.getList()
    },

    getList: function () {
        let that = this
        request({
            url: '/dictionary/positionTags',
        }).then(res => {
            const { data } = res
            let items = data.list
            for (let item of items) {
                let children = [{
                    id: item.id,
                    name: '全部' + item.name,
                    path: item.path,
                    pid: item.pid,
                }].concat(item.children)
                item.children = children
            }
            items = [{
                name: '全部职位',
                path: '',
                children: [],
            }].concat(items)
            that.setData({
                items,
            })
        })
    },

    clickItem: function (e) {
        let that = this
        let id = e.currentTarget.dataset.id
        let selectedItem = {}
        for (let item of that.data.items) {
            if (id == item.id) {
                selectedItem = item
                break
            }
        }
        that.setData({
            selectedItem,
        })
        if (selectedItem.children.length > 0) {
            that.setData({
                children: selectedItem.children,
                showPopup: true,
            })
        }
        else {
            if (that.data.from === 'home') {
                wx.setStorageSync('selectedHomeProfession', selectedItem)
            } else if (that.data.from === 'userinfo') {
                let pages = getCurrentPages()
                let prevPage = pages[pages.length - 2]
                prevPage.setData({
                    targetPosition: selectedItem.name
                })
            } else {
                wx.setStorageSync('selectedJobProfession', selectedItem)
            }
            wx.navigateBack({
              delta: 0,
            })
        }
    },

    clickChild: function (e) {
        let that = this
        let id = e.currentTarget.dataset.id
        let selectedChild = {}
        for (let item of that.data.children) {
            if (id == item.id) {
                selectedChild = item
                break
            }
        }
        if (that.data.from === 'home') {
            wx.setStorageSync('selectedHomeProfession', selectedChild)
        } else if (that.data.from === 'userinfo') {
            let pages = getCurrentPages()
            let prevPage = pages[pages.length - 2]
            prevPage.setData({
                targetPosition: selectedChild.name
            })
        } else {
            wx.setStorageSync('selectedJobProfession', selectedChild)
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