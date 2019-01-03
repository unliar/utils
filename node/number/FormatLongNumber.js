/**
 * 用于将数字格式化成千分计数
 * @example
 * 1234.2===>1,234.20
 * @param {number} num 输入的数值
 * @param {number} len 保留的小数位
 * @returns {string} 字符串
 */
const FormatLongNumber = (num, len = 2) => {
  if (num == 0) {
    let s = "0.";
    for (let index = 0; index < len; index++) {
      s = s + "0";
    }
    return s;
  }
  return num.toFixed(len).toLocaleString("en-US");
};

module.exports = FormatLongNumber;
