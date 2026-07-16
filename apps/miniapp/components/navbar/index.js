// components/navbar/index.js
Component({
    /**
     * 组件的属性列表
     */
    properties: {
        title: {
            type: String,
            value: '标题',
        },
        showBackBtn: {
            type: Number,
            value: 1,
        },
        isTransparent: {
            type: Number,
            value: 0,
        },
    },

    /**
     * 组件的初始数据
     */
    data: {
        statusBarHeight: wx.getStorageSync('statusBarHeight'),
        navBarHeight: wx.getStorageSync('navBarHeight'),
        showNavbar: wx.getStorageSync('showNavbar'),
        showHomeBtn: true,
    },

    lifetimes: {
        // 生命周期函数，可以为函数，或一个在methods段中定义的方法名
        ready: function () { 
            const pages = getCurrentPages()
            const currentPage = pages[pages.length - 1]
            const url = `/${currentPage.route}`
            if (url.includes('home/home')) {
                // 首页无需显示 home 按钮
                this.setData({
                    showHomeBtn: false,
                })
            }
        },
    },

    /**
     * 组件的方法列表
     */
    methods: {
        back() {
            wx.navigateBack({
                delta: 0,
                fail(err) {
                    wx.switchTab({
                      url: '/pages/home/home',
                    })
                }
            })
        },
        home() {
            wx.switchTab({
              url: '/pages/home/home',
            })
        },
    }
})
