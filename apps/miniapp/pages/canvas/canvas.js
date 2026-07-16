const { handleEscapeChar } = require('../../utils/util')

Page({

    /**
     * 页面的初始数据
     */
    data: {
        recommendId: '',
        config: {},
        item: {},
        img: '',
        type: 1,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        wx.showLoading({
            title: '海报生成中'
        })
        let { type, recommendId } = options                  // type: 1：无二维码，2：有二维码
        let that = this
        let offer = wx.getStorageSync('offer')
        that.setData({
            item: offer,
            type,
            recommendId: recommendId || '0',
        })
        setTimeout(() => {
            that.updateData()
        }, 2000)
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

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
        const { item, recommendId } = this.data
        let views1 = [
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/bg1.png',
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
                text: item.title,
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
                text: item.liangdian,
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#7B7B7B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    top: '230px',
                    left: '72px',
                    maxLines: '3',
                }
            },
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/salary.png',
                css: {
                    width: '28px',
                    height: '32px',
                    top: '354px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: '职位年薪：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '350px',
                    left: '120px',
                }
            },
            {
                type: 'text',
                text: `${item.salaryMin}-${item.salaryMax}万`,
                css: {
                    width: '230px',
                    height: '45px',
                    color: '#248379',
                    fontSize: '32px',
                    lineHeight: '45px',
                    fontWeight: 'bold',
                    top: '348px',
                    left: '260px',
                }
            },
            {
                type: 'text',
                text: '需求详情',
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#181818',
                    fontSize: '34px',
                    lineHeight: '48px',
                    fontWeight: 'bold',
                    top: '432px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: handleEscapeChar(item.editContent),
                css: {
                    width: '612px',
                    height: '606px',
                    color: '#7B7B7B',
                    fontSize: '28px',
                    lineHeight: '48px',
                    fontWeight: '400',
                    top: '500px',
                    left: '72px',
                    maxLines: '13',
                }
            },
            {
                type: 'text',
                text: '工作地点',
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#181818',
                    fontSize: '34px',
                    lineHeight: '48px',
                    fontWeight: 'bold',
                    top: '1164px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.address,
                css: {
                    width: '612px',
                    height: '100px',
                    color: '#7B7B7B',
                    fontSize: '26px',
                    lineHeight: '48px',
                    fontWeight: '400',
                    top: '1230px',
                    left: '72px',
                    maxLines: 2,
                }
            },
        ]
        let views2 = [
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/bg2.png',
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
                text: item.title,
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
                text: item.liangdian,
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#7B7B7B',
                    fontSize: '28px',
                    lineHeight: '40px',
                    top: '230px',
                    left: '72px',
                    maxLines: '3',
                }
            },
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/salary.png',
                css: {
                    width: '40px',
                    height: '40px',
                    top: '360px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: '职位年薪：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '360px',
                    left: '120px',
                }
            },
            {
                type: 'text',
                text: `${item.salaryMin}-${item.salaryMax}万`,
                css: {
                    width: '230px',
                    height: '45px',
                    color: '#248379',
                    fontSize: '32px',
                    lineHeight: '45px',
                    fontWeight: 'bold',
                    top: '360px',
                    left: '260px',
                }
            },
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/trophy.png',
                css: {
                    width: '40px',
                    height: '40px',
                    top: '420px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: '自荐奖金：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '420px',
                    left: '120px',
                }
            },
            {
                type: 'text',
                text: `${item.interviewReward}`,
                css: {
                    width: '230px',
                    height: '45px',
                    color: '#248379',
                    fontSize: '32px',
                    lineHeight: '45px',
                    fontWeight: 'bold',
                    top: '420px',
                    left: '260px',
                }
            },
            {
                type: 'image',
                url: 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/bonus.png',
                css: {
                    width: '40px',
                    height: '40px',
                    top: '484px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: '推荐奖金：',
                css: {
                    width: '140px',
                    height: '40px',
                    color: '#181818',
                    fontSize: '28px',
                    lineHeight: '40px',
                    fontWeight: '500',
                    top: '484px',
                    left: '120px',
                }
            },
            {
                type: 'text',
                text: `${item.successReward}`,
                css: {
                    width: '230px',
                    height: '45px',
                    color: '#248379',
                    fontSize: '32px',
                    lineHeight: '45px',
                    fontWeight: 'bold',
                    top: '484px',
                    left: '260px',
                }
            },
            {
                type: 'text',
                text: '需求详情',
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#181818',
                    fontSize: '34px',
                    lineHeight: '48px',
                    fontWeight: 'bold',
                    top: '560px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: handleEscapeChar(item.editContent),
                css: {
                    width: '612px',
                    height: '606px',
                    color: '#7B7B7B',
                    fontSize: '28px',
                    lineHeight: '48px',
                    fontWeight: '400',
                    top: '620px',
                    left: '72px',
                    maxLines: '12',
                }
            },
            {
                type: 'text',
                text: '工作地点',
                css: {
                    width: '612px',
                    height: '120px',
                    color: '#181818',
                    fontSize: '34px',
                    lineHeight: '48px',
                    fontWeight: 'bold',
                    top: '1208px',
                    left: '72px',
                }
            },
            {
                type: 'text',
                text: item.address,
                css: {
                    width: '612px',
                    height: '100px',
                    color: '#7B7B7B',
                    fontSize: '26px',
                    lineHeight: '48px',
                    fontWeight: '400',
                    top: '1266px',
                    left: '72px',
                    maxLines: 2,
                }
            },
            {
                type: 'image',
                url: 'https://mini.bolejiang.com/api/wechat/barcode?page=pages%2Fdetail%2Fdetail&scene=' + 
                    item.id + `_${recommendId}`,
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
            views: this.data.type == 1 ? views1 : views2,
        }
        this.setData({
            config,
        })
    }
})