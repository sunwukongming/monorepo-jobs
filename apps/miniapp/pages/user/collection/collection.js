const { request } = require('../../../utils/request')
const { formatTime } = require('../../../utils/util')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        activeTab: 0,
        offers: [],
        offer: {
            page: 1,
            hasMore: true,
            keyword: '',
        },
        jobs: [],
        job: {
            page: 1,
            hasMore: true,
            keyword: '',
        },
        showClose: false,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let that = this
        getApp().whenReady().then(function () {
            that.getOffers()
            that.getJobs()
        })
    },

    onChangeTab(e) {
        this.setData({
            activeTab: e.detail.index
        })
    },

    getOffers() {
        let that = this
        let { data } = that
        let { offer } = data
        if (offer.hasMore) {
            request({
                url: '/passage/listLike',
                method: 'POST',
                data: {
                    page: offer.page,
                    keyword: offer.keyword
                }
            }).then(res => {
                let rspData = res.data
                let items = rspData.list
                for (let item of items) {
                    item.updateAt = formatTime(item.mtime * 1000)
                }
                if (rspData.currentPage < rspData.lastPage) {
                    that.setData({
                        ['offer.page']: offer.page + 1,
                    })
                }
                else {
                    that.setData({
                        ['offer.hasMore']: false,
                    })
                }
                that.setData({
                    offers: that.data.offers.concat(items)
                })
            })
        }
    },

    toOfferDetail: function (e) {
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
          url: '../../detail/detail?id=' + id,
        })
    },

    keywordSearchOffer(e) {
        let keyword = e.detail.value
        this.setData({
            keyword,
            ['offer.page']: 1,
            ['offer.hasMore']: true,
            offers: [],
        })
        this.getOffers()
    },

    onOfferSearchInput: function (e) {
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

    clearSearchInput: function () {
        let obj = {
            keyword: '',
            page: 1,
            hasMore: true,
        }
        this.setData({
            showClose: false,
        })
        if (this.data.activeTab == 0) {
            this.setData({
                offers: [],
                offer: obj
            })
            this.getOffers()
        } else {
            this.setData({
                jobs: [],
                job: obj,
            })
            this.getJobs()
        }
    },

    getJobs() {
        let that = this
        let { data } = that
        let { job } = data
        if (job.hasMore) {
            request({
                url: '/apply/listLike',
                method: 'POST',
                data: {
                    page: job.page,
                    keyword: job.keyword
                }
            }).then(res => {
                let rspData = res.data
                let items = rspData.list
                for (let item of items) {
                    item.description = item.description.substring(0, 56) + '...'
                    item.updateAt = formatTime(item.updatedTime * 1000)
                }
                if (rspData.currentPage < rspData.lastPage) {
                    that.setData({
                        ['job.page']: job.page + 1,
                    })
                }
                else {
                    that.setData({
                        ['job.hasMore']: false,
                    })
                }
                that.setData({
                    jobs: that.data.jobs.concat(items)
                })
            })
        }
    },

    toJobDetail: function (e) {
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
          url: '../../job/detail/detail?id=' + id,
        })
    },

    keywordSearchJob(e) {
        let keyword = e.detail.value
        this.setData({
            keyword,
            ['job.page']: 1,
            ['job.hasMore']: true,
            jobs: [],
        })
        this.getJobs()
    },

    onJobSearchInput: function (e) {
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
})