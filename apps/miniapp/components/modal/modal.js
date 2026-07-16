// components/modal/modal.js
Component({
    /**
     * 组件的属性列表
     */
    properties: {
        show: {
            type: Boolean,
            value: false,
        },
        title: {
            type: String,
            value: '提示',
        },
        showClose: {
            type: Boolean,
            value: true,
        },
    },

    /**
     * 组件的初始数据
     */
    data: {

    },

    /**
     * 组件的方法列表
     */
    methods: {
        closeModal: function () {
            let myEventDetail = {} // detail对象，提供给事件监听函数
            let myEventOption = {} // 触发事件的选项
            this.triggerEvent('close', myEventDetail, myEventOption)
        },
    }
})
