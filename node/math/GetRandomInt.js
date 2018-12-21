/**
 * Get a random number between min ~ max
 *
 * @param {number} max
 * @param {number} min
 */
const GetRandomInt = (min, max) => {
    return Math.floor(Math.random() * (max - min + 1) + min);
};

module.exports = GetRandomInt;
