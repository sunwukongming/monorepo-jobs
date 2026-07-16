const { request } = require('../../utils/request')
const { formatTime, handleEscapeChar } = require('../../utils/util')
import Dialog from '../../components/vant/dialog/dialog'
const { domain, prefix } = require('../../utils/config')
const app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        id: '',
        deliverId: '',
        recommendId: '',                   // 当前登录用户的推荐ID
        parentPassageRecommendId: '',      // 推荐人的id
        showTips: true,
        showModal: false,
        canShowCompany: false,
        showCompanyModal: false,
        companyModalTitle: '',
        companyModalContent: '',
        item: {},
        items: [],
        showAction: false,
        actionTitle: '请选择您的操作',
        actions: [
            {name: '呼叫'},
            {name: '复制内容'},
        ],
        showAwardModal: false,
        showBonusModal: false,
        awardRules: [
            '推荐奖金是奖聘用于激励用户积极推荐符合岗位要求的候选人而发放的激励奖金。以下是具体介绍：',
            '1、推荐奖金额度标准：以被推荐人入职后的转正月薪对应金额作为推荐奖金(税前)。',
            '2、推荐奖金成立条件：被推荐人认可推荐人的参与，不领取自荐奖金；推荐的简历必须真实有效，一年内未投递过职位所属公司，通过面试且成功入职。',
            '3、推荐奖金发放方式：由推荐人提供发票及银行账户信息，线上支付。奖金分两次发放，第一次在被推荐人入职后一个月内发放前50%；第二次在被推荐人通过试用期后一个月内发放后50%，试用期一般为3个月，部分为6个月，视情况而定。若被推荐人未通过试用期或在试用期内离职，后50%奖金不再发放，并退还已支付奖金的50%。',
            '4、推荐奖金发票及缴税：推荐人应提供合法增值税发票方式领取奖金；推荐人可通过电子税局（或税务微信公众号）申请线上代开或到税务局现场申请代开增值税发票，并自行申报个税；发票的代开申请人（也是纳税人）应当与推荐人是同一人；推荐人代开发票时必须选择自行申报个税，即推荐人提供的发票备注栏应载明“个人所得税由纳税人依法自行申报缴纳”字样；若推荐人提供的发票没有选择自行申报，发票备注栏有“个人所得税由支付方代扣代缴”字样， 奖聘将拒收该类发票。',
            '5、如有任何疑问及建议，可加奖聘客服微信沟通：284785123(请注明沟通奖金规则',
        ],
        bonusRules: [
            '自荐奖金是奖聘用于激励符合岗位要求的用户积极应聘平台岗位而发放的激励奖金。以下是具体介绍：',
            '1、自荐奖金额度标准：以自荐人入职后的转正月薪对应金额作为自荐奖金(税前)。',
            '2、自荐奖金成立条件：自荐人提供的简历必须真实有效，一年内未投递过职位所属公司，通过面试且成功入职；没有其他推荐人参与推荐。',
            '3、自荐奖金发放方式：由自荐人提供银行账户信息，线上支付。奖金分两次发放，第一次在自荐人入职后一个月内发放前50%；第二次在自荐人通过试用期后一个月内发放后50%，试用期一般为3个月，部分为6个月，视情况而定。若自荐人未通过试用期或在试用期内离职，后50%奖金不再发放，并退还已支付奖金的50%。',
            '4、自荐奖金发票及缴税：自荐人应提供合法增值税发票方式领取奖金；自荐人可通过电子税局（或税务微信公众号）申请线上代开或到税务局现场申请代开增值税发票，并自行申报个税；发票的代开申请人（也是纳税人）应当与自荐人是同一人；自荐人代开发票时必须选择自行申报个税，即自荐人提供的发票备注栏应载明“个人所得税由纳税人依法自行申报缴纳”字样；若自荐人提供的发票没有选择自行申报，发票备注栏有“个人所得税由支付方代扣代缴”字样， 奖聘将拒收该类发票。',
            '5、如有任何疑问及建议，可加奖聘客服微信沟通：284785123(请注明沟通奖金规则）',
        ],
        user: {},
        resumeModal: false,
        mobileModal: false,
        articleType: '',
        articleId: '',
        showRecommendModal: false,
        recommendContent: [],
        showRecommendCancelButton: true,
        recommendConfirmButtonText: '关闭',
        showUploadModal: false,
        uploadProgress: 0,
        uploadTips: '处理中',
        showLogin: false,
        selfConfirmModal: false,
        showCallout: true,
        shareTitle: '选择协作操作方式',
        showHelpModal: false,
        helpForm: {
            name: '',
            mobile: '',
            email: '',
            resumeUrl: '',
            recommendComment: '',
        },
        textarea: {
            minHeight: 80,
            maxHeight: 100,
        },
        shareOptions: [
            { name: '转发', icon: '/assets/share/share.jpg', openType: 'share', },
            { name: '推荐简历', icon: '/assets/icons/resume.jpg' },
            { name: '生成专属分享卡片', icon: '/assets/icons/no-award.png' },
            { name: '复制专属链接及JD', icon: '/assets/share/copy.jpg' },
        ],
        showMenuTip: true,
        canShowCompanyBak: false,
        from: {},
    },

    /**
     * 生命周期函数--监听页面加载
     */ 
    onLoad(options) {
        let that = this
        setTimeout(() => {
            that.setData({
                showTips: false,
            })
        }, 8000)
        if ('scene' in options) {
            const scene = decodeURIComponent(options.scene)
            const ids = scene.split('_')
            options.id = ids[0]
            options.parent = ids[1]
        }

        const deliverId = options.deliverId || ''

        wx.setStorageSync('lastPage', 'detail')
        this.setData({
            id: options.id,
            deliverId,
            parentPassageRecommendId: options.parent || '',
            articleType: options.articleType || '',
            articleId: options.articleId || '',
        })
        
        const from = wx.getLaunchOptionsSync()
        that.setData({
            from,
        })
        if (from.mode == 'singlePage') {
            // 单页模式，从朋友圈进入
            that.getDetail()
        } else {
            app.whenReady().then(() => {
                let user = wx.getStorageSync('user')
                that.setData({
                    user,
                    canShowCompany: !!(user.isAllies * 1)
                })
                that.getRecommendId()
                if (!user.mobile) {
                    let shareOptions = that.data.shareOptions
                    shareOptions[0].openType = ''
                    that.setData({
                        shareOptions,
                    })
                }
            })
        }
        wx.onCopyUrl((result) => {
            const { id, recommendId } = that.data
            that.data.canShowCompanyBak = that.data.canShowCompany
            that.setData({
                showModal: false,
                canShowCompany: false,
            })
            return {
                query: `id=${id}&parent=${recommendId}`
            }
        })
    },

    onUnload() {
        wx.offCopyUrl()
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        if (this.data.canShowCompanyBak) {
            this.setData({
                canShowCompany: this.data.canShowCompanyBak
            })
        }
    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {
        let that = this
        const { id, recommendId } = that.data
        that.data.canShowCompanyBak = that.data.canShowCompany
        that.setData({
            showModal: false,
            canShowCompany: false,
        })
        const path = `/pages/detail/detail?id=${id}&parent=${recommendId}`
        return {
            title: that.data.item.title || '奖聘职位推荐',
            path,
            imageUrl: '',
        }
    },

    onShareTimeline() {
        const that = this
        const { item, from } = that.data
        const { title, rootCompany, isAnonymous, salaryMin, salaryMax, cityName, successReward } = item
        let shareTitle = ''
        shareTitle = title + `|${salaryMin}-${salaryMax}万` + `|${cityName}`
        let company = ''
        if (isAnonymous == 1) {
            company = rootCompany.fakeName
        } else {
            company = rootCompany.simpleName
        }
        shareTitle += `|${company}(推荐奖${successReward})`
        let imageUrl = 'https://bolejiang.oss-cn-beijing.aliyuncs.com/images/mizhao.png'
        if (isAnonymous == 0) {
            imageUrl = rootCompany.icon
        }
        return {
            title: shareTitle,
            imageUrl,
        }
    },

    onPageScroll() {
        this.setData({
            showMenuTip: false,
        })
    },

    /**
     * 打开弹窗
     */
    openModal(e) {
        const { currentTarget } = e
        const { title } = currentTarget.dataset
        let that = this
        that.setData({
            showModal: true,
            shareTitle: title,
        })
    },

    /**
     * 关闭弹窗
     */
    closeModal() {
        let that = this
        that.setData({
            showModal: false,
            showMenuTip: false,
        })
    },

    openMobileModal() {
        let that = this
        that.setData({
            mobileModal: true,
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
            wx.hideLoading()
            that.setData({
                mobileModal: false,
            })
            that.getUserInfo()
        }).catch(() => {
            wx.hideLoading()
        })
    },

    getDetail() {
        let that = this
        const { id } = that.data
        request({
            url: '/passage/getOrigin',
            method: 'POST',
            data: {
                id,
            },
        }).then(res => {
            let item = res.data
            let content = item.editContent
            let sentences = content.split('\n')
            let formatterItems = []
            for (let sentence of sentences) {
                let item = handleEscapeChar(sentence)
                formatterItems.push(item)
            }
            item.formatContent = handleEscapeChar(item.editContent)
            item.sentences = formatterItems
            item.updateAt = formatTime(item.mtime * 1000)
            const rootCompany = item.rootCompany
            rootCompany.icon = 'https://admin.bolejiang.com' + rootCompany.icon
            rootCompany.tags = rootCompany.tags.split(',')
            const remark = rootCompany.companysynopsis
            rootCompany.desc = remark
            if (remark.length > 50) {
                rootCompany.desc = remark.slice(0, 50) + '...'
            }
            that.setData({
                item,
                rootCompany,
            })
            that.getRelativeList()
        })
    },

    getData() {
        let that = this
        const { id, parentPassageRecommendId, articleId, articleType } = that.data
        request({
            url: '/passage/get',
            method: 'POST',
            data: {
                id,
                passageRecommendId: parentPassageRecommendId,
                articleType,
                articleId,
            },
        }).then(res => {
            let item = res.data
            let content = item.editContent
            let sentences = content.split('\n')
            let formatterItems = []
            for (let sentence of sentences) {
                let item = handleEscapeChar(sentence)
                formatterItems.push(item)
            }
            item.formatContent = handleEscapeChar(item.editContent)
            item.sentences = formatterItems
            item.updateAt = formatTime(item.mtime * 1000)
            const rootCompany = item.rootCompany
            rootCompany.icon = 'https://admin.bolejiang.com' + rootCompany.icon
            rootCompany.tags = rootCompany.tags.split(',')
            const remark = rootCompany.companysynopsis
            rootCompany.desc = remark
            if (remark.length > 50) {
                rootCompany.desc = remark.slice(0, 50) + '...'
            }
            that.setData({
                item,
                rootCompany,
            })
            wx.setStorageSync('offer', that.data.item)
            that.getRelativeList()
            if (item.status != 0) {
                Dialog.alert({
                      title: '温馨提示',
                      message: '该职位已下架',
                }).then(() => {
                    // on close
                })
            }
            let recommendContent = []
            let shareCount = 0
            let shareCountL2 = 0
            let recommendCount = 0
            let recommendCountL2 = 0
            const { selfAccount } = item 
            if (selfAccount.shareCount) {
                shareCount = selfAccount.shareCount
                shareCountL2 = selfAccount.shareCountL2
                recommendCount = selfAccount.recommendCount
                recommendCountL2 = selfAccount.recommendCountL2
            }
            if (item.recommendAccount && item.recommendAccount.id != 0) {
                let account = item.recommendAccount
                let { name } = account
                if (!name) {
                    name = account.mobile.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
                }
                recommendContent.push(`该职位来自用户【${name}】的分享，职位传递链条已被记录，欢迎自荐简历，成功入职后您将获得相应的自荐礼包。`)
                that.setData({
                    showRecommendCancelButton: false,
                })
            } else {
                recommendContent.push('您当前打开的职位没有上游分享者，传递关系已被记录。若您在该职位自荐成功，将获得对应的自荐礼包以及一笔额外的自荐奖金。')
                that.setData({
                    showRecommendCancelButton: false,
                })
            }
            recommendContent.push('您也可以转发该职位给合适的朋友或分享到社群里，有朋友通过您的链接应聘成功，您将获得相应的推荐奖金。')
            recommendContent.push(`当下，您已将该职位传递给了【${shareCount}】位一级用户，一级有【${recommendCount}】人应聘；传递给了【${shareCountL2}】位二级用户，二级有【${recommendCountL2}】人应聘。`)
            that.setData({
                recommendContent,
            })
        })
    },

    openRecommendModal() {
        this.setData({
            showRecommendModal: true,
        })
    },

    getRelativeList() {
        let that = this
        request({
            url: '/passage/list',
            method: 'POST',
            data: {
                positionTagPath: that.data.item.positionTagPath || '',
                pageSize: '5',
            },
        }).then(res => {
            let { data } = res
            let items = data.list
            for (let item of items) {
                item.updateAt = formatTime(item.mtime * 1000)
            }
            that.setData({
                items,
            })
        })
    },

    onCollect() {
        let that = this
        let item = that.data.item
        let url = ''
        if (item.isLike) {
            url = '/passage/unlike'
        }
        else {
            url = '/passage/like'
        }
        request({
            url: url,
            method: 'POST',
            data: {
                id: that.data.id,
            },
        }).then(res => {
            that.getData()
        })
    },

    toCanvas(type) {
        let that = this
        const { item, recommendId, user } = that.data
        if (!user.mobile) {
            that.setData({
                mobileModal: true,
            })
            return false
        }
        wx.setStorageSync('offer', item)
        wx.navigateTo({
            url: `/pages/canvas/canvas?type=${type}&recommendId=${recommendId}` ,
        })
    },

    toDetail: function (e) {
        let that = this
        let id = e.currentTarget.dataset.id
        wx.navigateTo({
            url: 'detail?id=' + id,
        })
        // that.setData({
        //     id,
        // })
        // that.onLoad({id})
        // wx.pageScrollTo({
        //   duration: 100,
        //   scrollTop: 0,
        // })
    },

    getQrCode: function () {
        let that = this
        let item = that.data.item
        request({
            url: '/wechat/barcode',
            data: {
                page: 'pages/detail/detail',
                scene: item.id,
            },
            success(res) {
            }
        })
    },

    clickPhoneNumber: function (e) {
        let that = this
        let item = that.data.item
        that.setData({
            actionTitle: `${item.contact} 可能是一个电话、微信号或，你可以`,
            showAction: true,
        })
    },

    onActionClick(e) {
        let that = this
        let item = that.data.item
        let value = e.detail.name
        switch (value) {
            case '呼叫': {
                wx.makePhoneCall({
                    phoneNumber: item.contact,
                })
                break
            }
            case '复制内容': {
                wx.setClipboardData({
                    data: item.contact,
                    success() {
                        wx.showToast({
                            title: '内容复制成功',
                        })
                    },
                    fail(err) {
                        wx.showToast({
                            title: '内容复制失败',
                        })
                    },
                })
            }
        }
        that.setData({
            showAction: false,
        })
    },

    closeAction() {
        let that = this
        that.setData({
            showAction: false,
        })
    },

    copyContent(e) {
        let that = this
        const { item } = that.data
        let formatContent = item.formatContent
        let editContent = item.title + '\r\n' + item.liangdian + '\r\n' + formatContent
        wx.setClipboardData({
            data: editContent,
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

    clickShowCompany() {
        const { item } = this.data
        if (this.data.canShowCompany || item.isAnonymous == 0) {
            let items1 = []
            let sentences = item.positionResearch.split(/\r\n/g)
            if (item.positionResearch) {
                for (let sentence of sentences) {
                    let item = handleEscapeChar(sentence)
                    items1.push(item)
                }
            }
            let items2 = []
            sentences = item.companyRemark.split(/\r\n/g)
            for (let sentence of sentences) {
                let item = handleEscapeChar(sentence)
                items2.push(item)
            }
            this.setData({
                companyModalTitle: item.outName,
                companyModalContent1: items1,
                companyModalContent2: items2,
            })
        } else {
            this.setData({
                companyModalTitle: '温馨提示',
                companyModalContent: ['尊敬的用户您好，该职位出于保密需要，公司及项目信息仅向有自荐或推荐意愿的用户开放，欢迎通过下方按钮或手机/微信方式与我们联系，感谢支持！！'],
            })
        }
        this.setData({
            showCompanyModal: true,
        })
    },

    onCloseCompanyDialog() {
        this.setData({
            showCompanyModal: false,
        })
    },

    copyCompanyRemark1() {
        let that = this
        let { item } = that.data
        let content = handleEscapeChar(item.positionResearch)
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

    copyCompanyRemark2() {
        let that = this
        let { item } = that.data
        let content = handleEscapeChar(item.companyRemark)
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

    onShowAwardModal() {
        // this.setData({
        //     showAwardModal: true,
        // })
        wx.navigateTo({
            url: '/pages/discover/detail/detail?id=2&menuId=7',
        })
    },

    onCloseAwardDialog() {
        this.setData({
            showAwardModal: false,
        })
    },

    onShowBonusModal() {
        // this.setData({
        //     showBonusModal: true,
        // })
        wx.navigateTo({
            url: '/pages/discover/detail/detail?id=1&menuId=7',
        })
    },

    onCloseBonusDialog() {
        this.setData({
            showBonusModal: false,
        })
    },

    copyAwardRemark() {
        let that = this
        let { awardRules } = that.data
        const content = awardRules.join('\n')
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

    copyBonusRemark() {
        let that = this
        let { bonusRules } = that.data
        const content = bonusRules.join('\n')
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

    copyCompanyDesc() {
        let that = this
        let { item } = that.data
        const content = item.title + '：' + item.liangdian
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

    getRecommendId() {
        let that = this
        const { id, parentPassageRecommendId } = that.data
        request({
            url: '/passage/recommend',
            method: 'POST',
            data: {
                id,
                parentPassageRecommendId: parentPassageRecommendId,
            },
        }).then(res => {
            const rspData = res.data
            that.setData({
                recommendId: rspData.passageRecommendId
            })
            that.getData()
        })
    },

    checkResumeBySelf() {
        let that = this
        const { user } = that.data
        if (!user.resumeUrl) {
            wx.showModal({
                title: '',
                content: '感谢加入奖聘，请上传您的附件简历，即可完成自荐，获取丰厚的自荐礼包',
                cancelText: '暂时跳过',
                confirmText: '去上传',
                complete: (res) => {
                    if (res.confirm) {
                        if (user.mobile) {
                            that.uploadResume()
                        } else {
                            that.setData({
                                showLogin: true,
                            })
                        }
                    }
                },
            })
        } else {
            that.setData({
                resumeModal: true,
            })
        }
    },

    uploadResume() {
        let that = this
        wx.chooseMessageFile({
            count: 1,
            type: 'file',
            success(res) {
                that.setData({
                    showUploadModal: true,
                    uploadProgress: 0,
                    uploadTips: '识别中',
                })
                const { tempFiles } = res
                let file = null
                if (tempFiles.length > 0) {
                    file = res.tempFiles[0]
                } else {
                    return false
                }
                let token = wx.getStorageSync('token')
                const uploadTask = wx.uploadFile({
                    header: {
                        'Authorization': `Bearer ${token}`,
                    },
                    filePath: file.path,
                    name: 'file',
                    url: `${domain}${prefix}/resume/upload`,
                    timeout: 60000,
                    success(res) {
                        if (res.statusCode != 200) {
                            wx.showToast({
                              title: '简历上传失败',
                              icon: 'error',
                            })
                            return false
                        }

                        const rsp = res.data
                        const rspData = JSON.parse(rsp)
                        if (rspData.code != 0) {
                            wx.showToast({
                                icon: 'error',
                                title: '上传失败',
                            })
                        } else {
                            wx.showToast({
                                title: '上传成功',
                            })
                            that.getUserInfo()
                        }
                    },
                    fail(err) {
                        wx.showToast({
                            icon: 'error',
                            title: '上传失败',
                        })
                    },
                    complete() {
                        wx.hideLoading()
                        uploadTask.offProgressUpdate()
                        that.setData({
                            showUploadModal: false,
                        })
                    }
                })
                uploadTask.onProgressUpdate((res) => {
                    const { progress } = res
                    let uploadTips = `${progress}%`
                    if (progress == 100) {
                        uploadTips = '文件识别中'
                    }
                    that.setData({
                        uploadProgress: progress,
                        uploadTips,
                    })
                })
            },
        })
    },

    downResume() {
        let that = this
        const { user } = that.data
        that.setData({
            resumeModal: false,
        })
        wx.downloadFile({
            url: user.resumeUrl,
            success(res) {
                wx.openDocument({
                    filePath: res.tempFilePath,
                })
            },
            fail(err) {
                wx.showToast({
                    icon: 'error',
                    title: '下载失败',
                })
            }
        })
    },

    getUserInfo() {
        let that = this
        request({
            url: '/user/info',
        }).then(res => {
            let data = res.data
            wx.setStorageSync('user', data)
            app.globalData.userInfo = data
            that.setData({
                user: data,
            })
        })
    },

    deliverResume() {
        let that = this
        const { id, parentPassageRecommendId } = that.data
        that.setData({
            resumeModal: false,
            selfConfirmModal: false,
        })
        request({
            url: '/passage/deliver',
            method: 'POST',
            data: {
                id,
                passageRecommendId: parentPassageRecommendId, 
                isReal: 1,
            },
        }).then(res => {
            wx.showToast({
              title: '自荐成功',
            })
        })
    },

    closeResumeModal() {
        this.setData({
            resumeModal: false,
        })
    },

    onCancelRecommendDialog() {
        wx.switchTab({
            url: '../home/home',
        })
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
          title: '登录中...',
        })
        request({
            url: '/account/bindWechatMobile',
            method: 'POST',
            data: {
                code,
            },
        }).then(res => {
            wx.hideLoading()
            this.getUserInfo()
        }).catch(() => {
            wx.hideLoading()
        })
    },

    openLogin() {
        this.setData({
            showLogin: true,
        })
    },

    hideLogin() {
        this.setData({
            showLogin: false,
        })
    },

    getUserInfo: function () {
        let that = this
        request({
            url: '/user/info',
        }).then(res => {
            that.setData({
                showLogin: false,
            })
            let data = res.data
            wx.setStorageSync('user', data)
            app.globalData.userInfo = data
            that.setData({
                user: data,
            })
        })
    },

    showSelfConfirmModal() {
        this.setData({
            resumeModal: false,
            selfConfirmModal: true,
        })
    },

    hideSelfConfirmModal() {
        this.setData({
            selfConfirmModal: false,
        })
    },

    copyShareLink() {
        let that = this
        const { item, recommendId } = that.data
        wx.showLoading({
          title: '生成中',
        })
        request({
            // url: '/wechat/urlLink',
            url: '/wechat/urlSchema',
            data: {
                path: '/pages/detail/detail',
                query: `scene=${item.id}_${recommendId}`,
                version: 'release',
            },
        }).then(res => {
            wx.hideLoading()
            const { data } = res.data
            const { url } = data
            let { item } = that.data
            let content = `急招【${item.title}】${item.liangdian} 年包 ${item.salaryMin} - ${item.salaryMax}万。\r\n\r\n`
            content += `JD及投递入口👉 ${url}\r\n\r\n`
            // content += `通过上方的奖聘专属链接应聘入职有丰厚礼包！生成专属链接发给朋友应聘入职获${item.successReward}作推荐奖金！`
            content += `通过上方奖聘专属链接应聘入职获免费旅行机票/住宿/ 咖啡/特产等大礼包！生成专属链接发给朋友应聘入职获${item.successReward}作推荐奖金！`
            wx.setClipboardData({
                data: content,
                success (res) {
                    wx.getClipboardData({
                        success (res) {
                            wx.showToast({
                              title: '复制成功',
                            })
                        },
                    })
                },
                fail (err) {
                    wx.showToast({
                        icon: 'error',
                        title: '复制失败，请检查权限',
                    })
                }
            })
        }).catch(() => {
            wx.hideLoading()
        })
    },

    openHelpModal() {
        const that = this
        const { id, deliverId, helpForm } = that.data
        if (deliverId) {
            request({
                url: '/deliver/detail',
                method: 'POST',
                data: {
                    id: deliverId,
                },
            }).then(res => {
                const rspData = res.data.deliver
                helpForm.name = rspData.name
                helpForm.mobile = rspData.mobile
                helpForm.email = rspData.email 
                helpForm.resumeUrl = rspData.resumeUrl 
                helpForm.recommendComment = rspData.recommendComment 
                that.setData({
                    showHelpModal: true,
                    showModal: false,
                    helpForm,
                })
            })
        }
        else {
            that.setData({
                showHelpModal: true,
                showModal: false,
                helpForm: {
                    name: '',
                    mobile: '',
                    email: '',
                    resumeUrl: '',
                    recommendComment: '',
                },
            })
        }
    },

    closeHelpModal() {
        this.setData({
            showHelpModal: false,
        })
    },

    onChangeHelpField(e) {
        const that = this
        const { detail } = e
        const { helpForm } = that.data
        helpForm[e.currentTarget.dataset.field] = detail
        that.setData({
            helpForm,
        })
    },

    submitHelpModal() {
        const that = this
        const { id, recommendId, helpForm } = that.data
        const user = wx.getStorageSync('user')
        if (!user.mobile) {
            that.setData({
                showLogin: true,
            })
            return false
        }
        request({
            url: '/deliver/createManual',
            method: 'POST',
            data: {
                passageId: id,
                passageRecommendId: recommendId,
                ...helpForm,
            },
        }).then(res => {
            that.setData({
                showHelpModal: false,
            })
            wx.showToast({
              title: '提交成功',
            })
        })
    },

    uploadHelpResume() {
        let that = this
        wx.chooseMessageFile({
            count: 1,
            type: 'file',
            success(res) {
                that.setData({
                    showUploadModal: true,
                    uploadProgress: 0,
                    uploadTips: '识别中',
                })
                const { tempFiles } = res
                let file = null
                if (tempFiles.length > 0) {
                    file = res.tempFiles[0]
                } else {
                    return false
                }
                let token = wx.getStorageSync('token')
                const uploadTask = wx.uploadFile({
                    header: {
                        'Authorization': `Bearer ${token}`,
                    },
                    filePath: file.path,
                    name: 'file',
                    url: `${domain}${prefix}/resume/upload`,
                    timeout: 60000,
                    success(res) {
                        if (res.statusCode != 200) {
                            wx.showToast({
                              title: '简历上传失败',
                              icon: 'error',
                            })
                            return false
                        }

                        const rsp = res.data
                        const rspData = JSON.parse(rsp)
                        if (rspData.code != 0) {
                            wx.showToast({
                                icon: 'error',
                                title: '上传失败',
                            })
                        } else {
                            wx.showToast({
                                title: '上传成功',
                            })
                            const { helpForm } = that.data
                            helpForm.resumeUrl = rspData.data.url
                            that.setData({
                                helpForm,
                            })
                        }
                    },
                    fail(err) {
                        wx.showToast({
                            icon: 'error',
                            title: '上传失败',
                        })
                    },
                    complete() {
                        wx.hideLoading()
                        uploadTask.offProgressUpdate()
                        that.setData({
                            showUploadModal: false,
                        })
                    }
                })
                uploadTask.onProgressUpdate((res) => {
                    const { progress } = res
                    let uploadTips = `${progress}%`
                    if (progress == 100) {
                        uploadTips = '文件识别中'
                    }
                    that.setData({
                        uploadProgress: progress,
                        uploadTips,
                    })
                })
            },
        })
    },

    onSelectOptions(e) {
        const that = this
        that.closeModal()
        const { user } = that.data
        const { detail } = e
        const { index } = detail
        if (index == 0) {
            if (!user.mobile) {
                that.openLogin()
            } else {
                let shareOptions = that.data.shareOptions
                shareOptions[0].openType = 'share'
                that.setData({
                    shareOptions,
                })
            }
            return true
        }
        if (index == 1) {
            if (!user.mobile) {
                that.openLogin()
            } else {
                that.openHelpModal()
            }
            return true
        }
        if (index == 2) {
            if (!user.mobile) {
                that.openLogin()
            } else {
                that.toCanvas(2)
            }
            return true
        }
        if (index == 3) {
            if (!user.mobile) {
                // that.openMobileModal()
                that.openLogin()
            } else {
                that.copyShareLink()
            }
            return true
        }
    },

    toCompany(e) {
        const { id } = e.currentTarget.dataset
        wx.navigateTo({
            url: `/pages/company/detail/detail?id=${id}`
        })
    },

    toCompanyJob(e) {
        const { id } = e.currentTarget.dataset
        wx.navigateTo({
            url: `/pages/company/detail/detail?id=${id}&activeTab=1`
        })
    },
})