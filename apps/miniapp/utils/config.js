module.exports = {
    domain: 'https://mini.bolejiang.com',
    prefix: '/api',
    discoverMenus: [
        { 
            id: 1, name: '行业资讯', icon: 'news', active: 'active', url: '/article/industryInfos', 
            btn: '私聊小编', type: 'industryInfo',
        },
        { 
            id: 2, name: '投资/并购', icon: 'demand', active: '', url: '/article/investmentDemands', 
            btn: '立即联系', type: 'investmentDemand',
        },
        { 
            id: 3, name: '融资/求助', icon: 'finance', active: '', url: '/article/financingDemands', 
            btn: '立即联系', type: 'financingDemand',
        },
        { 
            id: 5, name: '推荐专家', icon: 'expert', active: '', url: '/profService/list', 
            detailUrl: '/profService/get',
            btn: '立即联系', type: 'profService',
        },
        { 
            id: 4, name: '业务合作', icon: 'cooperation', active: '', url: '/article/cooperations', 
            detailUrl: '/article/cooperationDetail',
            btn: '立即联系', type: 'cooperation',
        },
        { 
            id: 6, name: '行业活动', icon: 'activity', active: '', url: '/article/meetings', 
            btn: '立即报名', type: 'meeting',
        },
        // { id: 6, name: '行业社群', icon: 'people', active: '', url: '/article/industryAssociations', },
        { 
            id: 7, name: '公告&指南', icon: 'question', active: '', url: '/article/topicQas', 
            btn: '咨询更多', type: 'topicQa',
        },
    ],
    expertMenuId: 5,
}