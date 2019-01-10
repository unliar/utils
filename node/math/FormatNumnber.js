/**
 * 小数截取保留指定位数并且不四舍五入
 *
 * @param {number} num 数值
 * @param {number} [dig=2] 位数
 */
const FormatFloatWithNoRound = (num, dig = 2) => {
    const s = Math.pow(10, dig);
    return parseInt(num * s, 10) / s;
};
/**
 * 格式化数字为千分位字符串
 *
 * @param {number} num 数字
 * @returns {string}
 */
const FormatToThousandthString = num => num.toLocaleString('en-US');

module.exports = {
    FormatFloatWithNoRound,
    FormatToThousandthString
};
