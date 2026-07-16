// pages/job/canvas/canvas.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        config: {},
        item: {},
        img: '',
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        wx.showLoading({
            title: '海报生成中'
        })
        let that = this
        let job = wx.getStorageSync('job')
        that.setData({
            item: job,
        })
        setTimeout(() => {
            that.updateData()
        }, 1000)
    },

    onImgOK(e) {
        wx.hideLoading({
          success: (res) => {},
        })
        this.setData({
            img: e.detail.path
        })
    },

    saveImage() {
        let that = this
        wx.getSetting({
            success(res) {
                if (!res.authSetting['scope.writePhotosAlbum']) {
                    wx.authorize({
                        scope: 'scope.writePhotosAlbum',
                        success() {
                            wx.saveImageToPhotosAlbum({
                                filePath: that.data.img,
                                success(res) {
                                    wx.showToast({
                                        title: '保存成功',
                                    })
                                },
                            })
                        },
                        fail() {
                            wx.showModal({
                                title: '提示',
                                content: '请先允许添加到相册权限',
                                confirmText: '去授权',
                                success(res) {
                                    if (res.confirm) {
                                        wx.openSetting({
                                          withSubscriptions: true,
                                        })
                                    } else if (res.cancel) {
                                    }
                                }
                            })
                        },
                    })
                } else {
                    wx.saveImageToPhotosAlbum({
                        filePath: that.data.img,
                        success(res) {
                            wx.showToast({
                                title: '保存成功',
                            })
                        },
                    })
                }
            }
        })
    },

    updateData() {
        let item = this.data.item
        let views = [
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/job-bg.png',
                css: {
                    width: '750px',
                    height: '1624px',
                    top: 0,
                    left: 0,
                    borderRadius: '50px',
                }
            },
            {
                type: 'rect',
                css: {
                    width: '686px',
                    height: '1224px',
                    background: '#fff',
                    borderRadius: '30px',
                    color: '#fff',
                    top: '110px',
                    left: '32px',
                }
            },
            {
                type: 'text',
                text: '求职目标：' + item.destPosition,
                css: {
                    color: '#101010',
                    fontSize: '40px',
                    fontWeight: 'bold',
                    lineHeight: '56px',
                    top: '156px',
                    left: '72px',
                    maxLines: 1,
                }
            },
            {
                type: 'text',
                text: '目前状态：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#80808B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '230px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.currentState,
                css: {
                    width: '612px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '230px',
                    left: '210px',
                }
            },
            {
                type: 'text',
                text: '目标城市：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#80808B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '270px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.destCity,
                css: {
                    width: '612px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '270px',
                    left: '210px',
                }
            },
            {
                type: 'text',
                text: '目标公司：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#80808B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '310px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.destCompany,
                css: {
                    width: '612px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '310px',
                    left: '210px',
                }
            },
            {
                type: 'text',
                text: '期望年薪：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#80808B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '350px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.destSalary,
                css: {
                    width: '612px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '350px',
                    left: '210px',
                }
            },
            {
                type: 'text',
                text: '协助奖金：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#80808B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '390px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.helpRewardToC == 0 ? '(暂无协助奖金预算)' : `${item.helpRewardToC} 元`,
                css: {
                    width: '612px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '390px',
                    left: '210px',
                }
            },
            {
                type: 'text',
                text: '个人概述',
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#181818',
                    fontSize: '34px',
                    lineHeight: '48px',
                    fontWeight: 'bold',
                    top: '472px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.description,
                css: {
                    width: '612px',
                    height: '606px',
                    color: '#7B7B7B',
                    fontSize: '28px',
                    lineHeight: '48px',
                    fontWeight: '400',
                    top: '540px',
                    left: '72px',
                    maxLines: '13',
                }
            },
            {
                type: 'image',
                url: 'https://mini.bolejiang.com/api/wechat/barcode?page=pages%2Fjob%2Fdetail%2Fdetail&scene=' + item.id,
                css: {
                    width: '180px',
                    height: '180px',
                    top: '1370px',
                    left: '150px',
                    borderRadius: '90px',
                }
            },
        ]
        let config = {
            width: '750px',
            height: '1624px',
            background: '',
            views: views,
        }
        this.setData({
            config,
        })
    },
})