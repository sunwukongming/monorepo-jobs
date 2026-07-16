const { request } = require('../../../utils/request')
const { formatDatetime, formatTime, handleEscapeChar } = require('../../../utils/util')
const Config = require('../../../utils/config.js')
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        id: null,
        menuId: null,
        menu: {},
        baseUrl: null,
        url: null,
        menus: [],
        article: {},
        similarItems: [],
        showShareModal: false,
        fromTimeLine: false,
        banner: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/buttonbanner.PNG',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        const { id, menuId } = options
        this.data.id = id
        this.data.menuId = menuId
        this.initMenus()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        let launchOptions = wx.getLaunchOptionsSync()
        const { scene } = launchOptions
        if (scene == 1154) {
            this.setData({
                fromTimeLine: true,
            })
        }
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {
        const { article } = this.data
        this.setData({
            showShareModal: false
        })
        return {
            title: article.title || '上奖聘 找工作找人才找资金',
            // imageUrl: '',
        }
    },

    onShareTimeline() {
        const { article } = this.data
        this.setData({
            showShareModal: false
        })
        return {
            title: article.title || '上奖聘 找工作找人才找资金',
            imageUrl: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/squarelogo.png',
        }
    },

    initMenus() {
        let that = this
        let menus = []
        for (let item of Config.discoverMenus) {
            let menu = Object.assign({}, item)
            menu.icon = '/assets/discover/' + menu.icon + '.png'
            menus.push(menu)
            if (item.id == that.data.menuId) {
                let url = item.url
                let baseUrl = url.substr(0, url.length - 1)
                if (item.detailUrl) {
                    that.data.url = item.detailUrl
                } else {
                    that.data.url = baseUrl + 'Detail'
                }
                if (item.baseUrl) {
                    that.data.baseUrl = item.baseUrl
                } else {
                    that.data.baseUrl = baseUrl
                }
                that.setData({
                    menu: item,
                })
            }
        }
        that.setData({
            menus,
        })
        app.whenReady().then(() => {
            that.getDetail()
        })
    },

    onClickMenu(e) {
        const { id, index } = e.detail
        app.globalData.discoverMenu = id
        wx.switchTab({
          url: '../discover',
        })
    },

    getDetail() {
        let that = this
        const { id, url } = that.data
        wx.showLoading({
          title: '加载中...',
        })
        request({
            url: url,
            method: 'POST',
            data: {
                id,
            },
        }).then(res => {
            wx.hideLoading()
            let rspData = res.data
            rspData.datetime = formatDatetime(rspData.timeUpdate * 1000)
            if (that.data.menuId == Config.expertMenuId) {
                rspData.content = rspData.introduction
            }
            const similarItems = []
            if (rspData.passages) {
                for (const item of rspData.passages) {
                    item.datetime = formatTime(item.mtime * 1000)
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
                    similarItems.push(item)
                }
            }
            that.setData({
                article: rspData,
                similarItems,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    toHome() {
        wx.switchTab({
          url: '../../home/home',
        })
    },

    openShareModal() {
        let that = this
        that.setData({
            showShareModal: true,
        })
    },

    closeShareModal() {
        let that = this
        that.setData({
            showShareModal: false,
        })
    },

    previewImage() {
        const { data } = this
        wx.previewImage({
            urls: [data.banner],
        })
    },

    toDetail(e) {
        let that = this
        const id = e.currentTarget.dataset.id
        const { article, menu } = that.data
        wx.navigateTo({
            url: `/pages/detail/detail?id=${id}&articleType=${menu.type}&articleId=${article.id}`
        })
    },

    copyTitle() {
        let that = this
        let { article } = that.data
        // let content = handleEscapeChar(article.title)
        let content = article.title
        wx.setClipboardData({
            data: content,
            success() {
                wx.showToast({
                    title: '复制成功',
                })
            },
            fail(err) {
                wx.showToast({
                    icon: 'error',
                    title: '复制失败',
                })
            },
        })
    },

    copyContent() {
        let that = this
        let { article } = that.data
        let content = ''
        if (article.editContent) {
            content = handleEscapeChar(article.editContent)
        } else {
            content = handleEscapeChar(article.content)
            content = content.replace(/<br>/g, '\n')
            content = content.replace(/<\/div>/g, '\n')
            content = content.replace(/<\/p>/g, '\n')
            content = content.replace(/<\/?[^>]*>/g, '')
            content = content.replace(/(\n[\s\t]*\r*\n)/g, '\n')
        }
        wx.setClipboardData({
            data: content,
            success() {
                wx.showToast({
                    title: '复制成功',
                })
            },
            fail(err) {
                wx.showToast({
                    icon: 'error',
                    title: '复制失败',
                })
            },
        })
    },

    likeOrUnlike() {
        let that = this
        const { id, article, baseUrl } = that.data
        let url = '', count = article.likeCount
        count *= 1
        if (article.isLike) {
            url =  `${baseUrl}Unlike`
            count = count - 1
        } else {
            url = `${baseUrl}Like`
            count = count + 1
        }
        request({
            url,
            method: 'POST',
            data: {
                id,
            },
        }).then(res => {
            that.setData({
                [`article.isLike`]: !article.isLike,
                [`article.likeCount`]: count,
            })
        })
    },
})