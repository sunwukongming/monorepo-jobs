const { request } = require('../../utils/request')
const { formatTime } = require('../../utils/util')
const Config = require('../../utils/config')
const app = getApp()
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
        city: {
            text: '全部地区',
            id: '0-'
        },
        industry: {
            path: '',
            name: '全部行业',
        },
        profession: {
            path: '',
            name: '全部职位',
        },
        page: 1,
        pageSize: 10,
        keyword: '',
        hasMore: true,
        items: [],
        user: {},
        userJob: '传简历',
        helpPlan: '全部类型',
        helpPlanItems: ['全部类型', '提供协助奖金', '无需他人协助'],
        popup: {
            helpPlan: false,
        },
        showLogin: false,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let city = wx.getStorageSync('selectedCity') || {text: '全部地区', id: '0',}
        let industry = wx.getStorageSync('selectedIndustry') || {path: '', name: '全部行业',}
        let profession = wx.getStorageSync('selectedProfession') || {path: '', name: '全部职位',}
        this.setData({
            city,
            industry,
            profession,
        })
        
        let that = this
        app.whenReady().then(function () {
            that.getList()
        })
        this.getSwipers()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        wx.setStorageSync('tabPage', '/pages/job/job')
        let user = wx.getStorageSync('user')
        this.setData({
            user,
        })
        if (user.applies.length > 0) {
            let userJob = '传简历'
            for (let i = 0; i < user.applies.length; i++) {
                const item = user.applies[i]
                if (item.isFirst == 1) {
                    userJob = item.destPosition + ' | ' + item.destSalary
                    break
                }
            }
            this.setData({
                userJob: userJob,
            })
        }
        let lastPage = wx.getStorageSync('lastPage')
        if (lastPage == 'city' || lastPage == 'industry' || lastPage == 'profession') {
            let city = wx.getStorageSync('selectedJobCity') || {text: '全部地区', id: '0',}
            let industry = wx.getStorageSync('selectedJobIndustry') || {path: '', name: '全部行业',}
            let profession = wx.getStorageSync('selectedJobProfession') || {path: '', name: '全部职位',}
            this.setData({
                city,
                industry,
                profession,
                items: [],
                page: 1,
                hasMore: true,
            })
            this.getList()
        }
        this.initMenu()
        wx.stopPullDownRefresh()
    },

    
    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {
        wx.showLoading({
            title: '刷新中',
        })
        this.setData({
            items: [],
            page: 1,
            hasMore: true,
        })
        this.getList()
    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {
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

    /**
     * 轮播图事件
     * @param {*} e 
     */
    swiperChange: function (e) { //指示图标
        this.setData({
            ['swiper.current']: e.detail.current
        })
    },

    swiperToPage: function (e) {
        let that = this
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
            const { menus } = that.data
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

    /**
     * 城市选择
     */
    toSelectCity: function () {
        wx.navigateTo({
          url: '../select/city/city?from=job',
        })
    },

    /**
     * 行业选择
     */
    toSelectIndustry: function () {
        wx.navigateTo({
          url: '../select/industry/industry?from=job',
        })
    },

    /**
     * 职位选择
     */
    toSelectProfession: function () {
        wx.navigateTo({
          url: '../select/profession/profession?from=job',
        })
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
        wx.showLoading({
          title: '加载中',
        })
        const keyword = e.detail.value
        this.setData({
            page: 1,
            keyword,
            hasMore: true,
            items: [],
        })
        this.getList()
    },

    clearSearchInput: function () {
        wx.showLoading({
          title: '加载中',
        })
        this.setData({
            keyword: '',
            items: [],
            page: 1,
            hasMore: true,
        })
        this.getList()
    },

    toDetail: function (e) {
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
          url: './detail/detail?id=' + id,
        })
    },

    getList: function () {
        let that = this
        let data = that.data
        if (data.hasMore) {
            let params = {
                page: data.page,
                pageSize: data.pageSize,
                keyword: data.keyword,
                destIndustryPath: data.industry.path,
                destPositionTagPath: data.profession.path,
            }
            let ids = that.data.city.id.split('-')
            params.destCityId = ids[0]
            // if (ids.length > 1 && ids[1]) {
            //     params.destCityId = ids[1]
            // }
            if (data.helpPlan == '提供协助奖金') {
                params.isHelpRewardVisible = 1
            }
            if (data.helpPlan == '无需他人协助') {
                params.isHelpRewardVisible = 0
            }
            request({
                url: '/apply/list',
                method: 'POST',
                data: params,
            }).then(res => {
                let { data } = res
                let items = data.list
                let currentPage = parseInt(data.currentPage)
                let page = that.data.page
                if (currentPage < parseInt(data.lastPage)) {
                    page += 1
                }
                else {
                    that.setData({
                        hasMore: false
                    })
                }
                for (let item of items) {
                    item.description = item.description.substring(0, 56) + '...'
                    item.updateAt = formatTime(item.updatedTime * 1000)
                }
                that.setData({
                    items: that.data.items.concat(items),
                    page,
                })
                wx.hideLoading({
                  success: (res) => {},
                })
                wx.stopPullDownRefresh()
            })
        } 
        else {
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
                if (item.type == 2) {
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

    initMenu() {
        let that = this
        let menus = []
        for (let item of Config.discoverMenus) {
            menus.push(item)
        }
        that.setData({
            menus,
        })
    },

    toRule() {
        // wx.navigateTo({
        //     url: '/pages/discover/detail/detail?id=5&menuId=7'
        // })
        let user = wx.getStorageSync('user')
        if (!user.mobile) {
            this.setData({
                showLogin: true,
            })
            return false
        }
        wx.navigateTo({
            url: '/pages/user/resume/online/online'
        })
    },

    showPopupHelpPlan: function () {
        this.setData({
            ['popup.helpPlan']: true,
        })
    },

    closePopupHelpPlan: function () {
        this.setData({
            ['popup.helpPlan']: false,
        })
    },

    confirmPopupHelpPlan: function (e) {
        const state = e.detail.value
        this.setData({
            ['popup.helpPlan']: false,
            helpPlan: state,
            page: 1,
            hasMore: true,
            items: [],
        })
        this.getList()
    },

    openPrivacy() {
        wx.openPrivacyContract({
            success: () => {}, // 打开成功
            fail: () => {}, // 打开失败
            complete: () => {}
        })
    },

    getPhoneNumber (e) {
        let that = this
        let code = e.detail.code
        if (!code) {
            wx.showToast({
                icon: 'none',
                title: '获取手机号失败',
            })
            return false
        }
        wx.showLoading({
          title: '登录中...',
        })
        request({
            url: '/account/bindWechatMobile',
            method: 'POST',
            data: {
                code,
            },
        }).then(res => {
            that.getUserInfo()
            that.setData({
                showLogin: false,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    getUserInfo: function () {
        let that = this
        request({
            url: '/user/info',
        }).then(res => {
            wx.hideLoading()
            let data = res.data
            wx.setStorageSync('user', data)
            app.globalData.userInfo = data
            let hasMobile = false
            if (data.mobile) {
                hasMobile = true
            }
            that.setData({
                userInfo: data,
                targets: data.applies,
                showLogin: !hasMobile,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    handleAgreePrivacyAuthorization() {
        // 用户同意隐私协议事件回调
    },

    loginCancel() {
        this.setData({
            showLogin: false,
        })
    },
})