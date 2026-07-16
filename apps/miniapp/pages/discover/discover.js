const { request } = require('../../utils/request')
const { formatDatetime } = require('../../utils/util')
const Config = require('../../utils/config')
const app = getApp()

// 每个菜单一份独立的列表状态，避免多个页面实例共享同一数组引用
function initialListState() {
    return {
        page: [1, 1, 1, 1, 1, 1, 1, 1],
        hasMore: [true, true, true, true, true, true, true, true],
        loading: [false, false, false, false, false, false, false, false],
        items: [[], [], [], [], [], [], [], []],
    }
}

Page({

    /**
     * 页面的初始数据
     */
    data: {
        fixedHeight: wx.getStorageSync('statusBarHeight') + wx.getStorageSync('navBarHeight'),
        swiper: {
            items: [],
            current: 0,
            indicatorDots: false,
            autoplay: true,
            interval: 2000,
            duration: 500,
            circular: true,
        },
        menus: Config.discoverMenus,
        activeMenu: {},
        activeIndex: 0,
        keyword: '',
        page: [1, 1, 1, 1, 1, 1, 1, 1],
        hasMore: [true, true, true, true, true, true, true, true],
        loading: [false, false, false, false, false, false, false, false],
        items: [[], [], [], [], [], [], [], []],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.getSwipers()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        wx.setStorageSync('tabPage', '/pages/discover/discover')
        let { discoverMenu } = app.globalData
        if (!discoverMenu) {
            discoverMenu = 1
        }
        this.initMenu(discoverMenu)
        this.getList()
        wx.stopPullDownRefresh()
    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {
        const { activeIndex, page, items, hasMore } = this.data
        items[activeIndex] = []
        page[activeIndex] = 1
        hasMore[activeIndex] = 1
        this.setData({
            items,
            page,
            hasMore,
        })
        this.getList()
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {
        return {
            title: '上奖聘 找工作找人才找资金',
            imageUrl: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/top-banner.PNG',
        }
    },

    getSwipers: function () {
        let that = this
        request({
            url: '/banner/list',
        }).then(res => {
            let data = res.data
            let items = []
            for (let item of data.list) {
                if (item.type == 1) {
                    items.push(item)
                }
            }
            items.sort((a, b) => {
                if (a.sort < b.sort) {
                    return -1
                } else {
                    return 1
                }
            })
            that.setData({
                ['swiper.items']: items,
            })
        })
    },

    /**
     * 轮播图事件
     * @param {*} e 
     */
    swiperChange: function (e) { //指示图标
        this.setData({
            ['swiper.current']: e.detail.current
        })
    },

    initData(keyword) {
        this.setData(Object.assign({ keyword }, initialListState()))
    },

    searchInput: function (e) {
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

    keywordSearch: function (e) {
        let keyword = e.detail.value
        keyword = keyword.split(' ').join('')
        this.initData(keyword)
        this.getList()
    },

    clearSearchInput: function () {
        this.setData({
            showClose: false,
        })
        this.initData('')
        this.getList()
    },

    getList: function () {
        let that = this
        const { activeMenu, activeIndex, page, hasMore, loading, items, keyword } = that.data
        if (hasMore[activeIndex] && !loading[activeIndex]) {
            wx.showLoading({
              title: '加载中...',
            })
            loading[activeIndex] = true
            that.data.loading = loading
            request({
                url: activeMenu.url,
                method: 'post',
                data: {
                    page: page[activeIndex],
                    pageSize: 10,
                    keyword: keyword,
                },
            }).then(res => {
                wx.hideLoading()
                wx.stopPullDownRefresh()
                loading[activeIndex] = false
                const rspData = res.data
                if (page[activeIndex] >= rspData.lastPage) {
                    hasMore[activeIndex] = false
                    that.setData({
                        hasMore,
                    })
                } else {
                    page[activeIndex] += 1
                    that.data.page = page
                }
                let rspItems = []
                for(let item of rspData.list) {
                    item.datetime = formatDatetime(item.timeUpdate * 1000)
                    rspItems.push(item)
                }
                items[activeIndex] = items[activeIndex].concat(rspItems)
                that.setData({
                    items,
                })
            }).catch( err => {
                wx.hideLoading()
                wx.stopPullDownRefresh()
                loading[activeIndex] = false
            })
        }
    },

    initMenu(activeId) {
        let that = this
        let menus = []
        let index = 0
        for (let item of Config.discoverMenus) {
            if (item.id == activeId) {
                item.active = 'active'
                that.setData({
                    activeMenu: item
                })
                that.setData({
                    activeIndex: index,
                })
            } else {
                item.active = ''
            }
            menus.push(item)
            index += 1
        }
        that.setData({
            menus,
        })
    },

    clickMenu: function (e) {
        let that = this
        const { dataset } = e.currentTarget
        const { id, index } = dataset
        let items = []
        for (let item of that.data.menus) {
            if (item.id == id) {
                item.active = 'active'
                that.setData({
                    activeMenu: item,
                })
            } else {
                item.active = ''
            }
            items.push(item)
        }
        that.setData({
            menus: items,
            activeIndex: index,
        })
        app.globalData.discoverMenu = id
        that.getList()
    },

    toDetail(e) {
        const { dataset } = e.currentTarget
        const { id } = dataset
        const { activeMenu } = this.data
        wx.navigateTo({
          url: './detail/detail?id=' + id + `&menuId=${activeMenu.id}`,
        })
    },

    refreshList(e) {
        this.getList()
    },

    bannerToPage (e) {
        let that = this
        const { menus } = that.data
        let items = that.data.swiper.items
        let id = e.currentTarget.dataset.id
        let swiper = {}
        for (let item of items) {
            if (id == item.id) {
                swiper = item
                break
            }
        }
        if (!swiper.jumpUrl) {
            return false
        }
        if (swiper.jumpUrl.includes('http://') || swiper.jumpUrl.includes('https://')) {
            wx.navigateTo({
                url: '../web/web?url=' + swiper.jumpUrl,
            })
        } else {
            let jumpUrl = swiper.jumpUrl
            let url = ''
            if (jumpUrl.includes('discover/detail/detail')) {
                let items = jumpUrl.split('?')
                url = items[0] + '?'
                let queryString = items[1]
                let params = queryString.split('&')
                for (let param of params) {
                    if (param.includes('type=')) {
                        let tmp = param.split('=')
                        let activeIndex = tmp[1] - 1
                        url += 'url=' + menus[activeIndex].url
                        app.globalData.discoverMenu = tmp[1]
                    } else {
                        url += param + '&'
                    }
                }
            }
            wx.navigateTo({
                url,
            })
        }
    },
})