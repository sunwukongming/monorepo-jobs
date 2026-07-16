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
        user: {},
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        const user = wx.getStorageSync('user')
        this.setData({
            user,
        })
        this.getData()
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
        let {
            hasMore,
            page,
            items
        } = that.data
        if (hasMore) {
            wx.showLoading({
                title: '加载中...',
            })
            request({
                url: '/passage/listRecommend',
                method: 'POST',
                data: {
                    keyword: '',
                }
            }).then(res => {
                wx.hideLoading()
                const rspData = res.data
                const results = rspData.list
                for (let result of results) {
                    if (result.deliver && result.deliver.mobile) {
                        result.mobile = result.deliver.mobile.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
                    } else {
                        result.mobile = ''
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
        } else {
            this.setData({
                showClose: false,
            })
        }
    },

    clearSearchInput() {
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
        const {
            dataset
        } = e.currentTarget
        const {
            id,
            deliver,
        } = dataset
        wx.navigateTo({
            url: `/pages/detail/detail?id=${id}&deliverId=${deliver}`,
        })
    },

    showResume(event) {
        const {
            dataset
        } = event.currentTarget
        const {
            index
        } = dataset
        const that = this
        const {
            items,
            user,
        } = that.data
        const { deliver } = items[index]
        if (user.isResumeWatcher == 0 || !deliver.resumeUrl) {
            return false
        }
        wx.downloadFile({
            url: deliver.resumeUrl,
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