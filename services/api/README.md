# 接口文档

## 用户相关

### 更新用户信息

- post: /api/user/update
- 请求参数

```
{
    "name": "gm",
    "email": "444@ww.com",
    "mobile": "111",
    "company": "测试公司",
    "position": "string",
    "wechat": "wwwwwww",
	"industry": "行业",
	"university": "学校",
	"description": "xxxx",
	"tags": "a b c",
	"apply": { //在之前的接口之上增加apply数据
			"currentCity": "北京",  //当前城市
			"currentCompany": "腾讯", //当前公司
			"currentPosition": "老总", //当前职位
			"currentState": "在职", //当前状态
			"currentIndustry": "互联网", //当前行业
			"description": "啦啦啦啦啦", //描述
			"destCity": "西安", //期望城市
			"destCompany": "大唐", //期望公司
			"destIndustry": "互联网", //期望行业
			"destPosition": "老总", //期望职位
			"destSalary": "20w", //期望薪水
			"education": "本科", //最高学历
			"university": "北京大学" //毕业院校
	}
}
```

- 响应

```
{
	"code": 0,
	"data": {},
	"latency": 2.492459735,
	"message": "ok",
	"timestamp": 1666260401
}
```

### 获取用户信息

- get: /api/user/info
- 响应

```
{
	"code": 0,
	"data": {
		"apply": { //在之前的接口之上增加apply数据
			"createdTime": "1666260401", //创建时间
			"currentCity": "北京", //当前城市
			"currentCompany": "腾讯", //当前公司
			"currentIndustry": "互联网", //当前行业
			"currentPosition": "老总", //当前职位
			"currentState": "在职", //当前状态
			"description": "啦啦啦啦啦", //经历及能力描述
			"destCity": "西安", //期望城市
			"destCompany": "大唐", //期望公司
			"destIndustry": "互联网", //期望行业
			"destPosition": "老总", //期望职位
			"destSalary": "20w", //期望薪水
			"education": "本科", // 最高学历
			"id": "1",
			"university": "北京大学", //最高学府
			"updatedTime": "1666260401"
		},
		"company": "测试公司",
		"createdTime": "0",
		"description": "xxxx",
		"email": "444@ww.com",
		"id": "1",
		"industry": "行业",
		"mobile": "111",
		"name": "gm",
		"openid": "o5ADs0PtkRSyYrfzbgq4KxitUpiY",
		"position": "string",
		"tags": "a b c",
		"unionid": "",
		"university": "学校",
		"updatedTime": "0",
		"wechat": "wwwwwww"
	},
	"latency": 0.421465193,
	"message": "ok",
	"timestamp": 1666260914
}
```

### 求职列表

- post: /api/apply/list

- 请求参数

```
{
	"keyword": "", //关键字
	"destPosition": "", //期望职位
	"destIndustry": "", //期望行业
	"destCity": "" //期望地区
}
```

- 响应

```
{
	"code": 0,
	"data": {
		"currentPage": "1",
		"lastPage": "1",
		"list": [
			{
				"createdTime": "1666260401", //创建时间
                "currentCity": "北京", //当前城市
                "currentCompany": "腾讯", //当前公司
                "currentIndustry": "互联网", //当前行业
                "currentPosition": "老总", //当前职位
                "currentState": "在职", //当前状态 选项来自字典选项接口
                "description": "啦啦啦啦啦", //经历及能力描述
                "destCity": "西安", //期望城市
                "destCompany": "大唐", //期望公司
                "destIndustry": "互联网", //期望行业
                "destPosition": "老总", //期望职位
                "destSalary": "20w", //期望薪水
                "education": "本科", // 最高学历
                "id": "1",
                "university": "北京大学", //最高学府
                "updatedTime": "1666260401"
			}
		],
		"perPage": "20",
		"total": "1"
	},
	"latency": 0.441223913,
	"message": "ok",
	"timestamp": 1666255627
}
```

### 求职详情

- post: /api/apply/detail

- 请求参数

```
{
	"id": "1" //求职id
}
```

- 响应

```
{
	"code": 0,
	"data": {
		"createdTime": "1666260401", //创建时间
        "currentCity": "北京", //当前城市
        "currentCompany": "腾讯", //当前公司
        "currentIndustry": "互联网", //当前行业
        "currentPosition": "老总", //当前职位
        "currentState": "在职", //当前状态
        "description": "啦啦啦啦啦", //经历及能力描述
        "destCity": "西安", //期望城市
        "destCompany": "大唐", //期望公司
        "destIndustry": "互联网", //期望行业
        "destPosition": "老总", //期望职位
        "destSalary": "20w", //期望薪水
        "education": "本科", // 最高学历
        "id": "1",
        "university": "北京大学", //最高学府
        "updatedTime": "1666260401"
		"similarApplies": [
			{
				"createdTime": "1666260401", //创建时间
                "currentCity": "北京", //当前城市
                "currentCompany": "腾讯", //当前公司
                "currentIndustry": "互联网", //当前行业
                "currentPosition": "老总", //当前职位
                "currentState": "在职", //当前状态
                "description": "啦啦啦啦啦", //经历及能力描述
                "destCity": "西安", //期望城市
                "destCompany": "大唐", //期望公司
                "destIndustry": "互联网", //期望行业
                "destPosition": "老总", //期望职位
                "destSalary": "20w", //期望薪水
                "education": "本科", // 最高学历
                "id": "1",
                "university": "北京大学", //最高学府
                "updatedTime": "1666260401"
			}
		],
		"university": "北京大学",
		"updatedTime": "1666255420"
	},
	"latency": 0.71307087,
	"message": "ok",
	"timestamp": 1666255805
}
```

### 求职收藏

- post: /api/apply/like

- 请求参数

```
{
	"id": "", //职位id
}
```

- 响应

```
{
	"code": 0,
	"data": {},
	"latency": 2.492459735,
	"message": "ok",
	"timestamp": 1666260401
}
```

### 求职取消收藏

- post: /api/apply/unlike

- 请求参数

```
{
	"id": "", //职位id
}
```

- 响应

```
{
	"code": 0,
	"data": {},
	"latency": 2.492459735,
	"message": "ok",
	"timestamp": 1666260401
}
```

### 求职收藏列表

- post: /api/apply/listLike

- 请求参数

```
{
	"keyword": "", //关键字
	"destPosition": "", //期望职位
	"destIndustry": "", //期望行业
	"destCity": "" //期望地区
}
```

- 响应

```
{
	"code": 0,
	"data": {
		"currentPage": "1",
		"lastPage": "1",
		"list": [
			{
				"createdTime": "1666260401", //创建时间
                "currentCity": "北京", //当前城市
                "currentCompany": "腾讯", //当前公司
                "currentIndustry": "互联网", //当前行业
                "currentPosition": "老总", //当前职位
                "currentState": "在职", //当前状态 选项来自字典选项接口
                "description": "啦啦啦啦啦", //经历及能力描述
                "destCity": "西安", //期望城市
                "destCompany": "大唐", //期望公司
                "destIndustry": "互联网", //期望行业
                "destPosition": "老总", //期望职位
                "destSalary": "20w", //期望薪水
                "education": "本科", // 最高学历
                "id": "1",
                "university": "北京大学", //最高学府
                "updatedTime": "1666260401"
			}
		],
		"perPage": "20",
		"total": "1"
	},
	"latency": 0.441223913,
	"message": "ok",
	"timestamp": 1666255627
}
```

## 数据相关

### 字典选项

- get: /api/dictionary/data
- 响应

```

{
"code": 0,
"data": {
"currentState": { //currentState 是选项
"key": "currentState",
"name": "当前职位状态",
"list": [
{
"id": 105,
"fid": 52,
"sort": 1,
"remark": "在职, 有好的机会欢迎聊聊"
},
{
"id": 106,
"fid": 52,
"sort": 2,
"remark": "在职, 不看机会"
},
{
"id": 107,
"fid": 52,
"sort": 3,
"remark": "不在职, 随时可以看机会"
},
{
"id": 108,
"fid": 52,
"sort": 4,
"remark": "不在职, 不看机会"
}
]
}
},
"latency": 0.000037535,
"message": "ok",
"timestamp": 1666260375
}

```

```

```

```

```
