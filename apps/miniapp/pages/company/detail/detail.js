const { request } = require('../../../utils/request')
const { formatTime } = require('../../../utils/util')
const Config = require('../../../utils/config')

Page({

    /**
     * 页面的初始数据
     */
    data: {
        id: 0,
        item: {
            desc: '',
        },
        showRemark: '展开',
        activeTab: 1,
        teams: [],
        events: [],
        activeEvent: 0,
        jobs: [],
        jobHasMore: true,
        jobPage: 1,
        articles: [],
        articleHasMore: true,
        articlePage: 1,
        showMemberModal: false,
        member: {},
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        const { id, activeTab } = options
        const that = this
        that.setData({
            id,
        })
        if (activeTab) {
            that.setData({
                activeTab,
            })
        }
        wx.setStorageSync('lastPage', 'companyDetail')
        that.getDetail()
        that.getTeams()
        that.getEvents()
        that.getArticles()
        that.getJobs()
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    },

    onPullDownRefresh() {
        wx.showLoading({
            title: '刷新中',
        })
        this.setData({
            jobs: [],
            jobHasMore: true,
            jobPage: 1,
        })
        this.getJobs()
    },

    getDetail() {
        const that = this
        const { id } = that.data
        request({
            url: '/rootCompany/get',
            method: 'POST',
            data: {
                id,
            },
        }).then(res => {
            wx.stopPullDownRefresh()
            wx.hideLoading()
            const item = res.data
            item.desc = that.handleDesc(item)
            item.icon = 'https://admin.bolejiang.com' + item.icon
            if (item.tags.length > 0) {
                item.tags = item.tags.split(',')
            } else {
                item.tags = []
            }
            that.setData({
                item,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    toggleDesc() {
        const that = this
        let { item, showRemark } = that.data
        let { desc, companysynopsis } = item
        if (showRemark == '展开') {
            item.desc = item.companysynopsis
            showRemark = '收起'
        } else {
            item.desc = that.handleDesc(item)
            showRemark = '展开'
        }
        that.setData({
            item,
            showRemark,
        })
    },

    handleDesc(item) {
        const remark = item.companysynopsis
        let desc = ''
        if (remark.length > 120) {
            desc = remark.slice(0, 70) + '...'
        } else {
            desc = remark
            this.setData({
                showRemark: '收起',
            })
        }
        return desc
    },

    onChangeTab(e) {
        const { index } = e.currentTarget.dataset
        this.setData({
            activeTab: index,
        })
    },

    getTeams() {
        const that = this
        const { id } = that.data
        request({
            url: '/rootCompany/listMember',
            method: 'POST',
            data: {
                id,
            },
        }).then(res => {
            wx.hideLoading()
            that.setData({
                teams: res.data.list,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    getEvents() {
        const that = this
        const { id } = that.data
        request({
            url: '/rootCompany/listEvent',
            method: 'POST',
            data: {
                id,
            },
        }).then(res => {
            wx.hideLoading()
            const items = res.data.list
            for (const item of items) {
                const { eventDate } = item
                const [first, second] = eventDate.split(' ')
                const [year, month, day] = first.split('-')
                item.year = year
            }
            that.setData({
                events: items,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    getJobs() {
        let that = this
        let data = that.data
        if (data.jobHasMore) {
            let params = {
                page: data.jobPage,
                pageSize: 100,
                rootCompanyId: data.id,
            }
            request({
                url: '/passage/list',
                method: 'POST',
                data: params,
            }).then(res => {
                let { data } = res
                let items = data.list
                let currentPage = parseInt(data.currentPage)
                let page = that.data.jobPage
                if (currentPage < parseInt(data.lastPage)) {
                    page += 1
                }
                else {
                    that.setData({
                        jobHasMore: false
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
                    if (rootCompany && rootCompany.tags) {
                        item.tags = item.tags.concat(item.rootCompany.tags.split(','))
                    }
                    item.tags.push('自荐礼包')
                }
                that.setData({
                    jobs: that.data.jobs.concat(items),
                    jobPage: page,
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

    getArticles() {
        let that = this
        let data = that.data
        if (data.articleHasMore) {
            let params = {
                id: data.id,
                page: data.articlePage,
                pageSize: 10,
                rootCompanyId: data.id,
            }
            request({
                url: '/rootCompany/listArticle',
                method: 'POST',
                data: params,
            }).then(res => {
                let { data } = res
                let items = data.list
                let currentPage = parseInt(data.currentPage)
                let page = that.data.articlePage
                if (currentPage < parseInt(data.lastPage)) {
                    page += 1
                }
                else {
                    that.setData({
                        articleHasMore: false
                    })
                }
                for (let item of items) {
                    item.updateAt = formatTime(item.timeUpdate * 1000)
                    for (let menu of Config.discoverMenus) {
                        if (menu.type == item.type) {
                            item.menuId = menu.id
                            continue
                        }
                    }
                }
                that.setData({
                    articles: that.data.articles.concat(items),
                    articlePage: page,
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

    openMemberModal(e) {
        const { index } = e.currentTarget.dataset
        this.setData({
            showMemberModal: true,
            member: this.data.teams[index]
        })
    },

    closeMemberModal() {
        this.setData({
            showMemberModal: false,
        })
    },

    toJobDetail(e) {
        const { id } = e.currentTarget.dataset
        wx.navigateTo({
            url: `/pages/detail/detail?id=${id}`,
        })
    },

    toArticleDetail(e) {
        const { id, menu } = e.currentTarget.dataset
        wx.navigateTo({
            url: `/pages/discover/detail/detail?id=${id}&menuId=${menu}`,
        })
    },
})