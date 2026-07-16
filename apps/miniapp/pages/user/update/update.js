const { request } = require('../../../utils/request')
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        textareaHolder: '简短描述：描述时请尽可能保护隐私，该部分内容会对外展示，用来给奖聘经纪人及HR了解大致信息，企业有真实兴趣后我们会向您索要完整简历做推荐',
        // 基本信息
        name: '',
        mobile: '',
        email: '',

        // 求职意向及当前信息
        isPublic: true,
        currentState: '',
        company: '',
        position: '',
        industry: '',
        nowCity: '',
        nowCities: [],
        targetCity: '',
        targetCities: [],
        targetCompany: '',
        targetIndustry: '',
        targetPosition: '',
        destPositionTag: '',
        destSalary: '',
        
        // 教育背景
        university: '',
        degree: '本科',
        
        // 经验及能力描述
        description: '',

        // 标签
        tagItems: [],
        tag: '',

        wechat: '',
        canSubmit: true,
        showDialog: false,
        wxUser: {},
        showWx: false,
        popup: {
            status: false,
            degree: false,
            nowCity: false,
            targetCity: false,
            industry: false,
            targetIndustry: false,
        },
        columns: {
            status: [],
            degree: ['专科', '本科', '研究生', '博士'],
            nowCities: [],
            targetCities: [],
        },
        industryParent: [],
        targetIndustryParent: [],
        industries: [],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let that = this
        let showWx = wx.getStorageSync('showWechat')
        that.setData({
            showWx,
        })
        setTimeout(function () {
            let user = wx.getStorageSync('user')
            let tags = user.tags
            let tagItems = []
            if (tags) {
                tagItems = tags.split(' ')
            }
            let nowCity = user.apply.currentCity
            let nowCities = nowCity.split('、')
            if (nowCities.length > 0 && nowCities[0][nowCities[0].length - 1] == '区') {
                nowCities.pop()
            }
            let targetCity = user.apply.destCity
            let targetCities = targetCity.split('、')
            if (targetCities.length > 0 && targetCities[0][targetCities[0].length - 1] == '区') {
                targetCities.pop()
            }
            let isPublic = true
            if (user.apply.id != 0) {
                if (user.apply.isPublic == 0) {
                    isPublic = false
                }
            }
            that.setData({
                name: user.name,
                mobile: user.mobile,
                email: user.email,

                // 求职意向及当前信息
                isPublic: isPublic,
                currentState: user.apply.currentState,
                company: user.company,
                position: user.position,
                industry: user.industry,
                nowCity,
                nowCities,
                targetCity,
                targetCities,
                targetCompany: user.apply.destCompany,
                targetIndustry: user.apply.destIndustry,
                targetPosition: user.apply.destPositionTag,
                destPositionTag: user.apply.destPosition,
                destSalary: user.apply.destSalary,

                degree: user.apply.education,
                university: user.university,
                wechat: user.wechat,
                description: user.description,

                tagItems,
            })
        }, 1000)
        let wxUser = wx.getStorageSync('wxUser')
        that.setData({
            wxUser,
        })
        that.getStatus()
        that.getCity()
        that.getIndustry()
    },

    getStatus() {
        let that = this
        request({
            url: '/dictionary/data',
        }).then(res => {
            const { data } = res
            const { currentState } = data
            const { list } = currentState
            let items = []
            for (let item of list) {
                items.push(item.remark)
            }
            that.setData({
                ['columns.status']: items
            })            
        })
    },

    getCity() {
        let that = this
        request({
            url: '/dictionary/cities',
        }).then(res => {
            const { list } = res.data
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
                ['columns.nowCities']: cities,
                ['columns.targetCities']: cities,
            })
        })
    },

    getIndustry: function () {
        let that = this
        request({
            url: '/dictionary/industries',
        }).then(res => {
            const { list } = res.data
            let parents = [], industries = {}
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
            let industryParent = [
                {
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
                targetIndustryParent: industryParent,
                industries,
            })
        })
    },

    checkboxChange: function (e) {
        this.data.tag = e.detail.value
    },

    submit: function () {
        let that = this
        let data = that.data
        let targetCity = data.targetCity
        let firstLetter = targetCity.substr(0, 1)
        if (firstLetter === '、') {
            targetCity = targetCity.slice(1)
        }
        let params = {
            name: data.name,
            email: data.email,
            mobile: data.mobile,
            company: data.company,
            position: data.position,
            wechat: data.wechat,
            industry: data.industry,
            university: data.university,
            description: data.description,
            tags: data.tagItems.join(' '),
            apply: {
                currentCity: data.nowCity,              // 当前城市
                currentCompany: data.company,           // 当前公司
                currentPosition: data.position,         // 当前职位
                currentState: data.currentState,        // 当前状态
                currentIndustry: data.industry,         // 当前行业
                description: data.description,          // 描述
                isPublic: data.isPublic ? '1' : '0',    // 是否对外展示
                destCity: targetCity,                   // 期望城市
                destCompany: data.targetCompany,        // 期望公司
                destIndustry: data.targetIndustry,      // 期望行业
                destPosition: data.destPositionTag,     // 期望职位
                destPositionTag: data.targetPosition,
                destSalary: data.destSalary,            // 期望薪水
                education: data.degree,                 // 最高学历
                university: data.university,            // 毕业院校
            }
        }
        request({
            url: '/user/update',
            method: 'POST',
            data: params,
        }).then(res => {
            let user = wx.getStorageSync('user')
            user = {
                ...user,
                ...params,
            }
            wx.setStorageSync('user', user)
            wx.showToast({
              title: '更新成功',
            })
            setTimeout(function () {
                wx.navigateBack({
                  delta: 0,
                })
            }, 500)
        })
    },

    changeValue: function (e) {
        let key = e.currentTarget.dataset.name
        let value = e.detail.value
        let params = {}
        params[key] = value
        this.setData(params)
    },

    checkEmail: function (e) {
        let that = this
        let reg = /^([a-zA-Z0-9]+[_|-|.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|-|.]?)*[a-zA-Z0-9]+.[a-zA-Z]{2,3}$/gi
        if (!reg.test(value)) {
            wx.showToast({
                title: '邮箱格式不正确',
            })
            that.data.canSubmit = false
        }
        else {
            that.data.canSubmit = true
        }
    },

    onChangePublic: function (e) {
        this.setData({
            isPublic: e.detail
        })
    },

    onShowDialog: function () {
        this.setData({
            showDialog: true,
        })
    },

    tagInput: function (e) {
        this.data.tag = e.detail.value
    },

    delTag: function (e) {
        let index = e.currentTarget.dataset.index
        let items = this.data.tagItems
        items.splice(index, 1)
        this.setData({
            tagItems: items,
        })
    },

    onCloseDialog: function () {
        let items = this.data.tagItems
        let tag = this.data.tag
        items.push(tag)
        this.setData({
            tagItems: items,
        })
    },

    toPosition: function () {
        wx.navigateTo({
          url: '../../select/profession/profession?from=userinfo',
        })
    },

    showPopupStatus: function () {
        this.setData({
            ['popup.status']: true,
        })
    },

    closePopupStatus: function () {
        this.setData({
            ['popup.status']: false,
        })
    },

    confirmPopupStatus: function (e) {
        this.setData({
            ['popup.status']: false,
            currentState: e.detail.value
        })
    },

    showPopupDegree: function () {
        this.setData({
            ['popup.degree']: true,
        })
    },

    closePopupDegree: function () {
        this.setData({
            ['popup.degree']: false,
        })
    },

    confirmPopupDegree: function (e) {
        this.setData({
            ['popup.degree']: false,
            degree: e.detail.value
        })
    },

    showPopupNowCity: function () {
        this.setData({
            ['popup.nowCity']: true,
        })
    },

    closePopupNowCity: function () {
        this.setData({
            ['popup.nowCity']: false,
        })
    },

    confirmPopupNowCity: function () {
        let value = this.data.nowCities
        const items = value.filter(item => item != '')
        this.setData({
            ['popup.nowCity']: false,
            nowCity: items.join('、')
        })
    },

    onChangeNowCity(e) {
        const value = e.detail
        this.setData({
            nowCities: value,
        })
    },

    showPopupTargetCity: function () {
        this.setData({
            ['popup.targetCity']: true,
        })
    },

    closePopupTargetCity: function () {
        this.setData({
            ['popup.targetCity']: false,
        })
    },

    confirmPopupTargetCity: function () {
        const value = this.data.targetCities
        const items = value.filter(item => item != '')
        this.setData({
            ['popup.targetCity']: false,
            targetCity: items.join('、'),
        })
    },

    onChangeTargetCity(e) {
        const value = e.detail
        this.setData({
            targetCities: value,
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
            industry,
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

    showPopupTargetIndustry: function () {
        this.setData({
            ['popup.targetIndustry']: true,
        })
    },

    closePopupTargetIndustry: function () {
        this.setData({
            ['popup.targetIndustry']: false,
        })
    },

    confirmPopupTargetIndustry: function (e) {
        const value = e.detail.value
        let targetIndustry = value[0].text + '/' + (value[1] && value[1].text)
        this.setData({
            ['popup.targetIndustry']: false,
            targetIndustry,
        })
    },

    onChangeTargetIndustry(e) {
        const value = e.detail.value
        let parents = this.data.targetIndustryParent
        parents[1].values = this.data.industries[value[0].text]
        this.setData({
            targetIndustryParent: parents,
        })
    },
})