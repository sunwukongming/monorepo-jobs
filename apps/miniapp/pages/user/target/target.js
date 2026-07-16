const {
    request
} = require('../../../utils/request')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        item: {},
        action: 'add',
        cities: [],
        popup: {
            target: false,
            city: false,
            industry: false,
        },
        destCity: [],
        industryParent: [],
        industries: [],
        industry: {},
        form: {
            target: '全职',
            destCity: '',
            destPosition: '',
            destSalary: '',
            destIndustry: '',
            destCompany: '',
            isFirst: false,
            isPublic: false,
            helpReward: '',
            description: '',
        },
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        const {
            action,
            index
        } = options
        this.setData({
            action,
        })
        if (action == 'edit') {
            let user = wx.getStorageSync('user')
            let items = user.applies
            let item = items[index]
            this.data.item = item
            let cities = item.destCity.split('、')
            let form = {
                target: '全职',
                destCity: item.destCity,
                destPosition: item.destPosition,
                destSalary: item.destSalary,
                destIndustry: item.destIndustry,
                destCompany: item.destCompany,
                isFirst: item.isFirst == 1 ? true : false,
                isPublic: item.isPublic == 1 ? false : true,
                helpReward: item.helpReward == 0 ? '' : item.helpReward,
                description: item.description,
            }
            this.setData({
                form,
                destCity: cities,
            })
        }
        this.getCity()
        this.getIndustry()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    changeValue (e) {
        let key = e.currentTarget.dataset.name
        let value = e.detail.value
        this.setData({
            [`form.${key}`]: value
        })
    },

    showPopupTarget() {
        this.setData({
            ['popup.target']: true,
        })
    },

    closePopupTarget() {
        this.setData({
            ['popup.target']: false,
        })
    },

    confirmPopupTarget(e) {
        const value = e.detail.value
        this.setData({
            ['popup.target']: false,
            ['form.target']: value,
        })
    },

    getCity() {
        let that = this
        request({
            url: '/dictionary/cities',
        }).then(res => {
            const {
                list
            } = res.data
            let cities = []
            for (let item of list) {
                if (item.name != '全部地区') {
                    cities.push({
                        id: item.id,
                        name: item.name,
                        checked: false,
                    })
                }
            }
            that.setData({
                cities,
            })
        })
    },

    showPopupCity: function () {
        this.setData({
            ['popup.city']: true,
        })
    },

    closePopupCity: function () {
        this.setData({
            ['popup.city']: false,
        })
    },

    confirmPopupCity: function () {
        this.setData({
            ['popup.city']: false,
            ['form.destCity']: this.data.destCity.join('、')
        })
    },

    onChangeCity(e) {
        const value = e.detail
        this.setData({
            destCity: value,
        })
    },

    getIndustry: function () {
        let that = this
        request({
            url: '/dictionary/industries',
        }).then(res => {
            const {
                list
            } = res.data
            let parents = [],
                industries = {}
            for (let item of list) {
                parents.push({
                    id: item.id,
                    text: item.name,
                })
                let children = []
                if (item.children && item.children.length > 0) {
                    for (let child of item.children) {
                        children.push({
                            id: child.id,
                            text: child.name
                        })
                    }
                }
                industries[item.name] = children
            }
            let industryParent = [{
                    values: parents,
                    defaultIndex: 0,
                },
                {
                    values: industries[parents[0].text],
                    defaultIndex: 0,
                }
            ]
            that.setData({
                industryParent,
                industries,
            })
        })
    },

    showPopupIndustry: function () {
        this.setData({
            ['popup.industry']: true,
        })
    },

    closePopupIndustry: function () {
        this.setData({
            ['popup.industry']: false,
        })
    },

    confirmPopupIndustry: function (e) {
        const value = e.detail.value
        let industry = value[0].text + '/' + (value[1] && value[1].text)
        this.setData({
            ['popup.industry']: false,
            ['form.destIndustry']: industry,
        })
    },

    onChangeIndustry(e) {
        const value = e.detail.value
        let parents = this.data.industryParent
        parents[1].values = this.data.industries[value[0].text]
        this.setData({
            industryParent: parents,
        })
    },

    onChangeCheckbox(e) {
        this.setData({
            ['form.isFirst']: e.detail
        })
    },

    onChangePublicCheckbox(e) {
        this.setData({
            ['form.isPublic']: e.detail
        })
    },

    submit() {
        let that = this
        const {
            action,
            form,
            item,
        } = that.data
        form.isFirst = form.isFirst ? 1 : 0
        let url = '',
            params = {
                ...form,
                isPublic: form.isPublic ? 0 : 1,
            }
        if (form.helpReward && form.helpReward < 5000) {
            wx.showToast({
                icon: 'error',
                title: '协作奖金不能小于5000',
            })
            return false
        }
        if (!form.destCity) {
            wx.showToast({
                icon: 'error',
                title: '求职城市必选',
            })
            return false
        }
        if (!form.destPosition) {
            wx.showToast({
                icon: 'error',
                title: '目标职位必选',
            })
            return false
        }
        if (!form.destSalary) {
            wx.showToast({
                icon: 'error',
                title: '薪资要求必填',
            })
            return false
        }
        if (!form.destIndustry) {
            wx.showToast({
                icon: 'error',
                title: '目标行业必选',
            })
            return false
        }
        if (action == 'add') {
            url = '/user/createApply'
        } else {
            url = '/user/updateApply'
            params.id = item.id
        }
        wx.showLoading({
            title: '加载中...',
        })
        request({
            url: url,
            method: 'post',
            data: params,
        }).then(res => {
            wx.hideLoading()
            wx.showToast({
                title: '保存成功',
            })
            setTimeout(function() {
                wx.navigateBack()
            }, 1000)
        }).catch(() => {
            wx.hideLoading()
        })
    },

    del() {
        let that = this
        if (that.data.action == 'add') {
            wx.navigateBack()
            return false
        }
        wx.showModal({
            title: '温馨提示',
            content: '删除操作不可撤销，是否继续',
            confirmText: '继续',
            complete: (res) => {
                if (res.confirm) {
                    wx.showLoading({
                        title: '删除中...',
                    })
                    request({
                        url: '/user/deleteApply',
                        method: 'post',
                        data: {
                            id: that.data.item.id
                        },
                    }).then(res => {
                        wx.hideLoading()
                        wx.showToast({
                            title: '删除成功',
                        })
                        setTimeout(() => {
                            wx.navigateBack()
                        }, 1000)
                    }).catch(() => {
                        wx.hideLoading()
                    })
                }
            }
        })
    },

    toHelpPlanRule() {
        wx.navigateTo({
            url: '/pages/discover/detail/detail?id=27&menuId=7',
        })
    },
})