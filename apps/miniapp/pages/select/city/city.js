const {
    request
} = require('../../../utils/request')

Page({

    /**
     * 页面的初始数据
     */
    data: {
        items: [],
        mainActiveIndex: 0,
        activeId: 0,
        from: 'home',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let from = options.from || 'home'
        let isCompany = options.isCompany || false
        this.data.from = from
        this.data.isCompany = isCompany
        wx.setStorageSync('lastPage', 'city')
        this.getList()
    },

    getList: function () {
        let pages = getCurrentPages()
        let parentPage = pages[pages.length - 2]
        let fromJob = false
        if (parentPage.route == 'pages/job/job') {
            fromJob = true
        }
        request({
            url: '/dictionary/cities',
        }).then(res => {
            const {
                data
            } = res
            let items = data.list
            let cities = []
            for (let item of items) {
                let cityName = ''
                if (item.name === '全部地区') {
                    cityName = '全部地区'
                } else {
                    cityName = '全' + item.name
                }
                let children = [{
                    text: cityName,
                    id: item.id + '-',
                }]

                if (!fromJob) {
                    for (let child of item.children) {
                        children.push({
                            text: child.name,
                            id: child.cityId + '-' + child.id
                        })
                    }
                }
                cities.push({
                    text: item.name,
                    id: item.id,
                    children,
                })
            }
            let selectedCity = wx.getStorageSync('selectedCity')
            this.setData({
                items: cities,
                activeId: selectedCity && selectedCity.id
            })
        })
    },

    /**
     * 单击左侧列表项
     * @param {*} param0 
     */
    onClickNav({
        detail = {}
    }) {
        this.setData({
            mainActiveIndex: detail.index || 0,
        })
    },

    /**
     * 选中对应的选项
     * @param {*} param0 
     */
    onClickItem({
        detail = {}
    }) {
        const activeId = this.data.activeId === detail.id ? null : detail.id
        this.setData({
            activeId
        })
        if (this.data.from == 'home') {
            if (this.data.isCompany == 1) {
                wx.setStorageSync('selectedHomeCompanyCity', detail)
            } else {
                wx.setStorageSync('selectedHomeCity', detail)
            }
        } 
        else {
            wx.setStorageSync('selectedJobCity', detail)
        }
        wx.navigateBack({
            delta: 0,
        })
    },
})