/**
 * 将多级对象的key值铺平
 *
 * @param {object} obj 想要铺平的对象
 * @param {object} result 目标对象，初始值是空对象
 * @return {object} 返回铺平对象
 */
const flapObject = (obj, result = {}) => {
    if (Object.keys(obj).length === 0) return result;
    Object.keys(obj).forEach(item => {
        if (typeof obj[item] === 'object') {
            result = flapObject(obj[item], result);
        } else {
            result[item] = obj[item];
        }
    });
    return result;
};

module.exports = flapObject
