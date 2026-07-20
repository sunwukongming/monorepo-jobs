/**
 * 使用基础库拆分后的系统信息 API（2.20.1+ / 单页模式等均推荐）。
 * 不再调用已停止维护的 wx.getSystemInfo / wx.getSystemInfoSync。
 *
 * platform 常见值：ios / android / ohos / ohos_pc / windows / mac / devtools
 */
function getSystemInfoCompat() {
  const appBase =
    typeof wx.getAppBaseInfo === 'function' ? wx.getAppBaseInfo() : {}
  const device =
    typeof wx.getDeviceInfo === 'function' ? wx.getDeviceInfo() : {}
  const windowInfo =
    typeof wx.getWindowInfo === 'function' ? wx.getWindowInfo() : {}
  const systemSetting =
    typeof wx.getSystemSetting === 'function' ? wx.getSystemSetting() : {}

  return Object.assign({}, systemSetting, appBase, device, windowInfo)
}

function isHarmonyOS(platform) {
  return platform === 'ohos' || platform === 'ohos_pc'
}

module.exports = {
  getSystemInfoCompat,
  isHarmonyOS,
}
