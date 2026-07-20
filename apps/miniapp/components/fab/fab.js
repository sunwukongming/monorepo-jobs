Component({
    /**
     * 组件的属性列表
     */
    properties: {
        pattern: {
            type: Object,
            default () {
                return {}
            }
        },
        horizontal: {
            type: String,
            default: 'left'
        },
        vertical: {
            type: String,
            default: 'bottom'
        },
        direction: {
            type: String,
            default: 'horizontal'
        },
        content: {
            type: Array,
            default () {
                return []
            }
        },
        show: {
            type: Boolean,
            default: false
        },
        popMenu: {
            type: Boolean,
            default: true
        }
    },

    /**
     * 组件的初始数据
     */
    data: {
        fabShow: false,
        isShow: false,
        styles: {
            color: '#3c3e49',
            selectedColor: '#007AFF',
            backgroundColor: '#fff',
            buttonColor: '#007AFF',
            iconColor: '#fff',
            icon: 'plusempty'
        },
    },

    ready() {
		const that = this
		const { getSystemInfoCompat } = require('../../utils/system')
		const res = getSystemInfoCompat()
		that.setData({
			y: `${res.windowHeight - 180}px`,
		})
	},

    /**
     * 组件的方法列表
     */
    methods: {
        clickBtn() {
            this.setData({
                isShow: !this.data.isShow
            })
        },

        onClickMenu(e) {
            this.triggerEvent('clickMenu', e.currentTarget.dataset)
        },
    }
})
