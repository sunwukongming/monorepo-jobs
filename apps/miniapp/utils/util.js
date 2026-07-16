const formatTime = time => {
    let nowDate = new Date()
    let nowTime = nowDate.getTime()
    let nowDay = nowDate.getDate()

    let oneDay = 24 * 3600 * 1000

    let yesterday = new Date(nowTime - oneDay)
    let yesDay = yesterday.getDate()

    const date = new Date(time)
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    let hour = date.getHours()
    let minute = date.getMinutes()
    const second = date.getSeconds()

    if (hour < 10) {
        hour = `0${hour}`
    }

    if (minute < 10) {
        minute = `0${minute}`
    }

    if (nowTime - time <= oneDay && nowDay === day) {
        // 这是当天更新的信息
        return `今日${hour}:${minute}更新`
    } 
    else if (nowTime - time <= oneDay * 2 && yesDay === day) {
        return `昨日${hour}:${minute}更新`
    }
    else {
        return `${[year, month, day].map(formatNumber).join('-')}`
    }
}

const formatDate = time => {
    const date = new Date(time)
    const year = date.getFullYear()
    const month = (date.getMonth() + 1).toString().padStart(2, '0')
    const day = date.getDate().toString().padStart(2, '0')
    return `${year}-${month}-${day}`
}

const formatDatetime = time => {
    const date = new Date(time)
    const year = date.getFullYear()
    const month = (date.getMonth() + 1).toString().padStart(2, '0')
    const day = date.getDate().toString().padStart(2, '0')
    const hour = date.getHours().toString().padStart(2, '0')
    const minute = date.getMinutes().toString().padStart(2, '0')
    const second = date.getSeconds().toString().padStart(2, '0')
    return `${year}-${month}-${day} ${hour}:${minute}:${second}`
}

const formatNumber = n => {
    n = n.toString()
    return n[1] ? n : `0${n}`
}

const handleEscapeChar = str => {
    const ESCAPE_CHARACTERS = {
        'nbsp': ' ',
        'lt': '<',
        'gt': '>',
        'amp': '&',
        'apos': '\"',
        'ensp': '	 ',
        'emsp': ' ',
        'quot': '"',
        'middot': '·',
        'brvbar': '¦',
        'mdash': '—',
        'ndash': '–',
        'ge': '≥',
        'le': '≤',
        'laquo': '«',
        'raquo': '»',
        'deg': '°',
        'bull': '•',
        'macr': '¯',
        '#64': '@',
        'ldquo': '“',
        'rdquo': '”',
        'rsquo': '‚',
        'lsquo': '‘',
      }

      return str.replace(new RegExp(`&(${ Object.keys(ESCAPE_CHARACTERS).join('|') });`, 'g'), (all, t) => {
        return ESCAPE_CHARACTERS[t]
      })
}

const numForamt =  (num) => {
    num = num.toString().split("."); // 分隔小数点
    let arr = num[0].split("").reverse(); // 转换成字符数组并且倒序排列
    let res = [];
    for (let i = 0, len = arr.length; i < len; i++) {
        if (i % 3 === 0 && i !== 0) {
            res.push(","); // 添加分隔符
        }
        res.push(arr[i]);
    }
    res.reverse(); // 再次倒序成为正确的顺序
    if (num[1]) { // 如果有小数的话添加小数部分
        res = res.join("").concat("." + num[1]);
    } else {
        res = res.join("");
    }
    return res;
}

module.exports = {
    formatTime,
    formatDate,
    formatDatetime,
    handleEscapeChar,
    numForamt,
}
