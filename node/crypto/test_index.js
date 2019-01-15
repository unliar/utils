const crypto = require('./index');
const assert = require('assert');
const path = require('path');

const data = `2019-01-15`;

const EncryptedResult = crypto.PublicKeyEncrypt(
    data,
    path.join(__dirname, './pub.key')
);

const DecryptResult = crypto.PrivateKeyDecrypt(
    EncryptedResult,
    path.join(__dirname, './pri.key')
);

assert.strictEqual('2019-01-15', DecryptResult);
console.log('Encrypt DecryptResult, passing');
