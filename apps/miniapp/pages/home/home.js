const { request } = require('../../utils/request')
const { formatTime } = require('../../utils/util')
const Config = require('../../utils/config')
const JOB_HISTORY_SEARCH  = 'jobHistorySearch'
const COMPANY_HISTORY_SEARCH  = 'companyHistorySearch'
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
            interval: 3500,
            duration: 500,
            circular: true,
        },
        items: [],
        page: 1,
        pageSize: 10,
        jobKeyword: '',
        hasMore: true,
        companyItems: [],
        companyPage: 1,
        companyPageSize: 10,
        companyKeyword: '',
        companyHasMore: true,
        city: {
            text: '全部地区',
            id: '0-'
        },
        industry: {
            path: '',
            name: '全部行业',
        },
        cityCompany: {
            text: '全部地区',
            id: '0-'
        },
        industryCompany: {
            path: '',
            name: '全部行业',
        },
        profession: {
            path: '',
            name: '全部职位',
        },
        showJobClose: false,
        showCompanyClose: false,
        jobHistoryWords: [],
        showJobHistory: false,
        companyHistoryWords: [],
        showCompanyHistory: false,
        jobInfo: '发职位',
        jobTarget: {},
        publicItems: ['全部类型', '实名', '密招'],
        public: '全部类型',
        scaleItems: [],
        scale: {text: '全部规模', id: 0,},
        stageItems: [],
        stage: {text: '全部阶段', id: 0,},
        popup: {
            public: false,
            stage: false,
            scale: false,
        },
        showCallout: false,
        showLogin: false,
        activeTab: 1,
        scrollTop: 0,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let city = wx.getStorageSync('selectedCity') || {text: '全部地区', id: '0',}
        let industry = wx.getStorageSync('selectedIndustry') || {path: '', name: '全部行业',}
        let profession = wx.getStorageSync('selectedProfession') || {path: '', name: '全部职位',}
        // 有意为之：随机首屏显示「职位」或「公司」tab，使两者曝光机会均等（勿改为固定值）
        const randomTab = Math.round(Math.random())
        this.setData({
            city,
            industry,
            profession,
            activeTab: randomTab,
        })
        this.getDictionary()
        this.getSwipers()
        let that = this
        app.whenReady().then(function () {
            if (that.data.activeTab == 1) {
                that.getCompanyList()
            } else {
                that.getList()
            }
        })
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {
        wx.setStorageSync('tabPage', '/pages/home/home')
        let user = wx.getStorageSync('user')
        let jobInfo = '发职位', jobTarget = {}
        if (user.applies && user.applies.length > 0) {
            jobTarget = user.applies[0]
            // jobInfo = ``
            // if (jobTarget.destPosition) {
            //     jobInfo += ' ' + jobTarget.destPosition
            // }
            // if (jobTarget.destSalary) {
            //     jobInfo += ' | ' + jobTarget.destSalary
            // }
        }
        this.setData({
            jobInfo,
            jobTarget,
        })
        let lastPage = wx.getStorageSync('lastPage')
        if (lastPage == 'city' || lastPage == 'industry' || lastPage == 'profession') {
            let city = wx.getStorageSync('selectedHomeCity') || {text: '全部地区', id: '0',}
            let industry = wx.getStorageSync('selectedHomeIndustry') || {path: '', name: '全部行业',}
            let profession = wx.getStorageSync('selectedHomeProfession') || {path: '', name: '全部职位',}
            let cityCompany = wx.getStorageSync('selectedHomeCompanyCity') || {text: '全部地区', id: '0',}
            let industryCompany = wx.getStorageSync('selectedHomeCompanyIndustry') || {path: '', name: '全部行业',}
            this.setData({
                city,
                industry,
                profession,
                cityCompany,
                industryCompany,
                items: [],
                page: 1,
                hasMore: true,
                companyItems: [],
                companyPage: 1,
                companyHasMore: true,
            })
            this.getDataList()
        }
        this.initMenu()
        wx.stopPullDownRefresh()
    },

    /**
     * 下拉刷新
     */
    onPullDownRefresh: function () {
        // wx.startPullDownRefresh()
        wx.showLoading({
          title: '刷新中',
        })
        if (this.data.activeTab == 0) {
            this.setData({
                items: [],
                page: 1,
                hasMore: true,
            })
        } else {
            this.setData({
                companyItems: [],
                companyPage: 1,
                companyHasMore: true,
            })
        }
        this.getDataList()
    },

    onHide() {
        this.closeCallout()
    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {
        this.getDataList()
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {
        return {
            title: '上奖聘 帮HR招人 帮人找工作',
            imageUrl: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/top-banner.PNG',
        }
    },

    onPageScroll(e) {
        this.data.scrollTop = e.scrollTop
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

    onChangeTab(e) {
        const { index } = e.detail 
        this.setData({
            activeTab: index,
        })
        wx.pageScrollTo({
            scrollTop: this.data.scrollTop + 1,
        })
        this.getDataList()
    },

    /**
     * 城市选择
     */
    toSelectCity(e) {
        const { type } = e.currentTarget.dataset
        this.closeCallout()
        const isCompany = type == 'company' ? 1 : 0
        wx.navigateTo({
          url: '../select/city/city?from=home&isCompany=' + isCompany,
        })
    },

    /**
     * 行业选择
     */
    toSelectIndustry(e) {
        const { type } = e.currentTarget.dataset
        this.closeCallout()
        const isCompany = type == 'company' ? 1 : 0
        wx.navigateTo({
          url: '../select/industry/industry?from=home&isCompany=' + isCompany,
        })
    },

    /**
     * 职位选择
     */
    toSelectProfession: function () {
        this.closeCallout()
        wx.navigateTo({
          url: '../select/profession/profession?from=home',
        })
    },

    jobSearchInput: function (e) {
        this.closeCallout()
        const keyword = e.detail.value
        if (keyword) {
            this.setData({
                showJobClose: true,
            })
        }
        else {
            this.setData({
                showJobClose: false,
            })
        }
    },

    jobKeywordSearch: function (e) {
        this.closeCallout()
        let keyword = e.detail.value
        keyword = keyword.split(' ').join('')
        if (keyword.length > 0) {
            let words = wx.getStorageSync(JOB_HISTORY_SEARCH) || []
            if (words.indexOf(keyword) == -1) {
                words.unshift(keyword)
                wx.setStorageSync(JOB_HISTORY_SEARCH, words)
            }
            // return false
        }
        wx.showLoading({
          title: '加载中',
        })
        this.setData({
            page: 1,
            jobKeyword: keyword,
            hasMore: true,
            items: [],
        })
        this.getList()
    },

    clearJobSearchInput: function () {
        this.closeCallout()
        wx.showLoading({
          title: '加载中',
        })
        this.setData({
            jobKeyword: '',
            items: [],
            page: 1,
            hasMore: true,
        })
        this.getList()
    },

    toDetail: function (e) {
        this.closeCallout()
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
          url: '/pages/detail/detail?id=' + id,
        })
    },

    toCompanyDetail(e) {
        this.closeCallout()
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
          url: '/pages/company/detail/detail?id=' + id,
        })
    },

    getDataList() {
        const that = this
        const { data } = that
        if (data.activeTab == 1) {
            that.getCompanyList()
        } else {
            that.getList()
        }
    },

    getList: function () {
        let that = this
        let data = that.data
        if (data.hasMore) {
            let isAnonymous = ''
            if (data.public == '实名') {
                isAnonymous = 0
            }
            if (data.public == '密招') {
                isAnonymous = 1
            }
            let params = {
                page: data.page,
                pageSize: data.pageSize,
                keyword: data.jobKeyword,
                industryPath: data.industry.path,
                positionTagPath: data.profession.path,
                isAnonymous,
            }
            let ids = that.data.city.id.split('-')
            if (ids.length > 1 && ids[1]) {
                params.districtId = ids[1]
            }
            params.cityId = ids[0]
            request({
                url: '/passage/list',
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
                    item.updateAt = formatTime(item.mtime * 1000)
                    const rootCompany = item.rootCompany
                    item.tags = []
                    if (item.isAnonymous == 1) {
                        item.tags.push(rootCompany.fakeName)
                    } else {
                        item.tags.push(rootCompany.simpleName)
                    }
                    item.tags.push(rootCompany.companyscaleName)
                    if (rootCompany && rootCompany.tags) {
                        item.tags = item.tags.concat(item.rootCompany.tags.split(','))
                    }
                    item.tags.push('自荐礼包')
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

    getCompanyList() {
        let that = this
        let data = that.data
        if (data.companyHasMore) {
            let params = {
                page: data.companyPage,
                pageSize: data.companyPageSize,
                keyword: data.companyKeyword,
                industryPath: data.industry.path,
                positionTagPath: data.profession.path,
                stage: data.stage.id,
                scale: data.scale.id
            }
            let ids = that.data.cityCompany.id.split('-')
            if (ids.length > 1 && ids[1]) {
                params.districtId = ids[1]
            }
            params.cityId = ids[0]
            request({
                url: '/rootCompany/list',
                method: 'POST',
                data: params,
            }).then(res => {
                let { data } = res
                let items = data.list
                let currentPage = parseInt(data.currentPage)
                let page = that.data.companyPage
                if (currentPage < parseInt(data.lastPage)) {
                    page += 1
                }
                else {
                    that.setData({
                        companyHasMore: false
                    })
                }
                for (let item of items) {
                    // item.icon = Config.domain + item.icon
                    item.icon = 'https://admin.bolejiang.com' + item.icon
                    item.tags = item.tags.split(',')
                }
                that.setData({
                    companyItems: that.data.companyItems.concat(items),
                    companyPage: page,
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

    toPage: function (e) {
        this.closeCallout()
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

    jobFocusSearch: function (e) {
        this.closeCallout()
        let that = this
        let words = wx.getStorageSync(JOB_HISTORY_SEARCH)
        if (words.length > 0) {
            that.setData({
                jobHistoryWords: words,
                showJobHistory: true,
            })
        }
    },

    jobBlurSearch: function (e) {
        setTimeout(() => {
            this.setData({
                showJobHistory: false,
            })
        }, 100)
    },

    quickSearchJob(e) {
        const { keyword } = e.currentTarget.dataset
        let words = wx.getStorageSync(JOB_HISTORY_SEARCH)
        let index = words.indexOf(keyword)
        words.splice(index, 1)
        words.unshift(keyword)
        wx.setStorageSync(JOB_HISTORY_SEARCH, words)
        this.setData({
            page: 1,
            jobKeyword: keyword,
            hasMore: true,
            items: [],
        })
        this.getList()
    },

    companySearchInput(e) {
        this.closeCallout()
        const keyword = e.detail.value
        if (keyword) {
            this.setData({
                showCompanyClose: true,
            })
        }
        else {
            this.setData({
                showCompanyClose: false,
            })
        }
    },

    companyKeywordSearch(e) {
        this.closeCallout()
        let keyword = e.detail.value
        keyword = keyword.split(' ').join('')
        if (keyword.length > 0) {
            let words = wx.getStorageSync(COMPANY_HISTORY_SEARCH) || []
            if (words.indexOf(keyword) == -1) {
                words.unshift(keyword)
                wx.setStorageSync(COMPANY_HISTORY_SEARCH, words)
            }
            // return false
        }
        wx.showLoading({
          title: '加载中',
        })
        this.setData({
            companyKeyword: keyword,
            companyItems: [],
            companyPage: 1,
            companyHasMore: true,
        })
        this.getCompanyList()
    },

    clearCompanySearchInput() {
        this.closeCallout()
        wx.showLoading({
          title: '加载中',
        })
        this.setData({
            companyKeyword: '',
            companyItems: [],
            companyPage: 1,
            companyHasMore: true,
        })
        this.getCompanyList()
    },

    companyFocusSearch(e) {
        this.closeCallout()
        let that = this
        let words = wx.getStorageSync(COMPANY_HISTORY_SEARCH)
        if (words.length > 0) {
            that.setData({
                companyHistoryWords: words,
                showCompanyHistory: true,
            })
        }
    },

    companyBlurSearch(e) {
        setTimeout(() => {
            this.setData({
                showCompanyHistory: false,
            })
        }, 100)
    },

    quickSearchCompany(e) {
        const { keyword } = e.currentTarget.dataset
        let words = wx.getStorageSync(COMPANY_HISTORY_SEARCH)
        let index = words.indexOf(keyword)
        words.splice(index, 1)
        words.unshift(keyword)
        wx.setStorageSync(COMPANY_HISTORY_SEARCH, words)
        this.setData({
            page: 1,
            companyKeyword,
            hasMore: true,
            items: [],
        })
        this.getList()
    },

    toJobTarget() {
        this.closeCallout()
        // let user = wx.getStorageSync('user')
        // if (!user.mobile) {
        //     this.setData({
        //         showLogin: true,
        //     })
        //     return false
        // }
        if (this.data.jobInfo == '发职位') {
            wx.navigateTo({
                url: '/pages/discover/detail/detail?id=5&menuId=7',
            })
        }
        else {
            wx.navigateTo({
                url: '../user/resume/online/online'
            })
        }
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

    showPopupPublic: function () {
        this.closeCallout()
        this.setData({
            ['popup.public']: true,
        })
    },

    closePopupPublic: function () {
        this.setData({
            ['popup.public']: false,
        })
    },

    confirmPopupPublic: function (e) {
        const state = e.detail.value
        this.setData({
            ['popup.public']: false,
            public: state,
            page: 1,
            hasMore: true,
            items: [],
        })
        this.getList()
    },

    closeCallout() {
        this.setData({
            showCallout: false,
        })
    },

    toRule() {
        this.closeCallout()
        wx.navigateTo({
            url: '/pages/discover/detail/detail?id=27&menuId=7'
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

    toHome() {
        this.setData({
            showLogin: false,
        })
    },
    
    handleAgreePrivacyAuthorization() {
        // 用户同意隐私协议事件回调
    },

    openPrivacy() {
        wx.openPrivacyContract({
            success: () => {}, // 打开成功
            fail: () => {}, // 打开失败
            complete: () => {}
        })
    },

    showPopupStage: function () {
        this.closeCallout()
        this.setData({
            ['popup.stage']: true,
        })
    },

    closePopupStage: function () {
        this.setData({
            ['popup.stage']: false,
        })
    },

    confirmPopupStage: function (e) {
        const state = e.detail.value
        this.setData({
            ['popup.stage']: false,
            stage: state,
            companyPage: 1,
            companyHasMore: true,
            companyItems: [],
        })
        this.getCompanyList()
    },

    showPopupScale: function () {
        this.closeCallout()
        this.setData({
            ['popup.scale']: true,
        })
    },

    closePopupScale: function () {
        this.setData({
            ['popup.scale']: false,
        })
    },

    confirmPopupScale: function (e) {
        const state = e.detail.value
        this.setData({
            ['popup.scale']: false,
            scale: state,
            companyPage: 1,
            companyHasMore: true,
            companyItems: [],
        })
        this.getCompanyList()
    },

    getDictionary() {
        const that = this
        request({
            url: '/dictionary/data',
            method: 'POST',
            data: {},
        }).then(res => {
            const { data } = res
            const scaleItems = [{id: 0, remark: '全部规模'}].concat(data.scale.list)
            for (const item of scaleItems) {
                item.text = item.remark
            }
            const stageItems = data.stage.list
            for (const item of stageItems) {
                item.text = item.remark
                if (item.text == '不限') {
                    item.text = '全部阶段'
                    item.id = 0
                }
            }
            that.setData({
                scaleItems,
                stageItems,
            })
        })
    },
})