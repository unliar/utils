const crypto = require('crypto');
const fs = require('fs');
const path = require('path');

/**
 * 使用公钥加密数据得到base64数据
 * @param {any} data 数据
 * @param {string} keyPath key路径
 */
exports.PublicKeyEncrypt = (data, keyPath) => {
  const keyData = fs.readFileSync(keyPath);
  const publicKey = keyData.toString();
  const buf = Buffer.from(data, 'utf8');
  const result = crypto.publicEncrypt(
      {key: publicKey, padding: crypto.constants.RSA_PKCS1_PADDING}, buf);
  return result.toString('base64');
};

/**
 * 私钥解密数据
 * @param {any} data 数据
 * @param {string} keyPath key路径
 */
exports.PrivateKeyDecrypt = (base64Data, keyPath) => {
  const keyData = fs.readFileSync(keyPath);
  const PrivateKey = keyData.toString();
  const buf = Buffer.from(base64Data, 'base64');
  const result = crypto.privateDecrypt(
      {key: PrivateKey, padding: crypto.constants.RSA_PKCS1_PADDING}, buf);
  return result.toString();
};
