const {
    request
} = require('../../../utils/request')

const {
    formatDate
} = require('../../../utils/util')

Page({

    /**
     * 页面的初始数据
     */
    data: {
        items: [],
        page: 1,
        keyword: '',
        hasMore: true,
        showClose: false,
        navBarHeight: 0,
        statusBarHeight: 0,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.getData()
        const navBarHeight = wx.getStorageSync('navBarHeight')
        const statusBarHeight = wx.getStorageSync('statusBarHeight')
        this.setData({
            navBarHeight,
            statusBarHeight,
        })
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {
        this.getData()
    },

    getData() {
        let that = this
        let { hasMore, page, items } = that.data
        if (hasMore) {
            wx.showLoading({
              title: '加载中...',
            })
            request({
                url: '/passage/listRecommend',
                method: 'POST',
                data: {
                    keyword: '',
                    isSelf: 1,
                }
            }).then(res => {
                wx.hideLoading()
                const rspData = res.data
                const results = rspData.list
                for (let result of results) {
                    if (result.deliver && result.deliver.accountMobile) {
                        result.phone = result.deliver.accountMobile.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
                    } else {
                        result.phone = ''
                    }
                    if (result.deliver && result.deliver.deliverTime) {
                        result.date = formatDate(result.deliver.deliverTime * 1000)
                    } else {
                        result.date = ''
                    }
                }
                if (parseInt(rspData.currentPage) < parseInt(rspData.lastPage)) {
                    page += 1
                } else {
                    that.setData({
                        hasMore: false
                    })
                }
                that.setData({
                    items: items.concat(results),
                    page,
                })
            }).catch(() => {
                wx.hideLoading()
            })
        }
    },
    
    keywordSearch(e) {
        let keyword = e.detail.value
        this.setData({
            keyword,
            page: 1,
            hasMore: true,
            items: [],
        })
        this.getOffers()
    },

    onSearchInput: function (e) {
        const keyword = e.detail.value
        if (keyword) {
            this.setData({
                showClose: true,
            })
        }
        else {
            this.setData({
                showClose: false,
            })
        }
    },

    clearSearchInput () {
        this.setData({
            keyword: '',
            page: 1,
            hasMore: true,
            items: [],
            showClose: false,
        })
        this.getData()
    },

    toDetail(e) {
        const { dataset } = e.currentTarget
        const { id } = dataset
        wx.navigateTo({
          url: `/pages/detail/detail?id=${id}`,
        })
    },
})