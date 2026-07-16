const {
    request
} = require('../../../utils/request')
Page({

    /**
     * 页面的初始数据
     */
    data: {
        arrowSize: '36rpx',
        showWx: false,
        workTime: [],
        birthday: [],
        popup: {
            sex: false,
            workTime: false,
            birthday: false,
        },
        form: {
            sex: '男',
            workTime: '',
            birthday: '',
            mobile: '',
            wechat: '',
            email: '',
        }
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        let showWx = wx.getStorageSync('showWechat')
        let maxDate = (new Date()).getTime()
        this.setData({
            maxDate,
            showWx,
        })
        this.init()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        let user = wx.getStorageSync('user')
        let form = {
            name: user.name,
            gender: user.gender || '男',
            workday: user.workday,
            birthday: user.birthday,
            mobile: user.mobile,
            wechat: user.wechat,
            email: user.email,
        }
        this.setData({
            form,
        })
    },

    init() {
        let timestamp = new Date()
        let nowYear = timestamp.getFullYear()
        let years = [], months = [], days = []
        for (let i = 1970; i < nowYear; i++) {
            years.push(i)
        }
        for (let i = 0; i < 12; i++) {
            if (i < 9) {
                months.push(`0${i + 1}`)
            } else {
                months.push(`${i + 1}`)
            }
        }
        for (let i = 0; i < 31; i++) {
            if (i < 9) {
                days.push(`0${i + 1}`)
            } else {
                days.push(`${i + 1}`)
            }
        }
        let defaultIndex = years.length - 12
        let workTime = [
            { values: years, defaultIndex: defaultIndex, },
            { values: months, defaultIndex: 0, },
            { values: days, defaultIndex: 0, },
        ]
        let birthday = JSON.parse(JSON.stringify(workTime))
        birthday[0].defaultIndex = years.length - 25
        this.setData({
            workTime,
            birthday,
        })
    },

    changeValue (e) {
        let key = e.currentTarget.dataset.name
        let value = e.detail.value
        this.setData({
            [`form.${key}`]: value
        })
    },

    showPopupSex() {
        this.setData({
            ['popup.sex']: true,
        })
    },

    closePopupSex() {
        this.setData({
            ['popup.sex']: false,
        })
    },

    confirmPopupSex(e) {
        const value = e.detail.value
        this.setData({
            ['popup.sex']: false,
            ['form.gender']: value,
        })
    },

    showPopupWorkTime() {
        this.setData({
            ['popup.workTime']: true,
        })
    },

    closePopupWorkTime() {
        this.setData({
            ['popup.workTime']: false,
        })
    },

    confirmPopupWorkTime(e) {
        const value = e.detail.value
        this.setData({
            ['popup.workTime']: false,
            ['form.workday']: value.join('.')
        })
    },

    changePopupWorkTime(e) {
        const { index, value } = e.detail
        if (index != 2) {
            // 年份、月份更改时，重新生成日期
            let day = this.getLastDay(value[0], value[1])
            let items = []
            for (let i = 0; i < day; i++) {
                if (i < 9) {
                    items.push(`0${i + 1}`)
                } else {
                    items.push(`${i + 1}`)
                }
            }
            let workTime = this.data.workTime
            workTime[2].values = items
            this.setData({
                workTime,
            })
        }
    },

    getLastDay(year, month) {
        const date = new Date(year, month, 0)
        return date.getDate()      
    },

    showPopupBirthday() {
        this.setData({
            ['popup.birthday']: true,
        })
    },

    closePopupBirthday() {
        this.setData({
            ['popup.birthday']: false,
        })
    },

    confirmPopupBirthday(e) {
        const value = e.detail.value
        this.setData({
            ['popup.birthday']: false,
            ['form.birthday']: value.join('.')
        })
    },

    changePopupBirthday(e) {
        const { index, value } = e.detail
        if (index != 2) {
            // 年份、月份更改时，重新生成日期
            let day = this.getLastDay(value[0], value[1])
            let items = []
            for (let i = 0; i < day; i++) {
                if (i < 9) {
                    items.push(`0${i + 1}`)
                } else {
                    items.push(`${i + 1}`)
                }
            }
            let birthday = this.data.birthday
            birthday[2].values = items
            this.setData({
                birthday,
            })
        }
    },

    submit() {
        let that = this
        const { form } = that.data
        wx.showLoading({
          title: '保存中...',
        })
        request({
            url: '/user/update',
            method: 'POST',
            data: form,
        }).then(() => {
            wx.hideLoading()
            wx.showToast({
              title: '保存成功',
            })
            setTimeout(() => {
                wx.navigateBack()
            }, 1000)
        }).catch(() => {
            wx.hideLoading()
        })
    },

    goBack() {
        wx.navigateBack()
    },

    

    getPhoneNumber (e) {
        let code = e.detail.code
        if (!code) {
            wx.showToast({
                icon: 'none',
                title: '获取手机号失败',
            })
            return false
        }
        wx.showLoading({
          title: '保存中...',
        })
        request({
            url: '/account/bindWechatMobile',
            method: 'POST',
            data: {
                code,
            },
        }).then(res => {
            wx.hideLoading()
            this.setData({
                showLogin: false,
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },
})