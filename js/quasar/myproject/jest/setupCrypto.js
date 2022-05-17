const crypto = require('crypto').webcrypto
const { subtle, getRandomValues } = crypto
const { importKey, exportKey, deriveKey } = subtle

global.crypto = {
  subtle: {
    importKey,
    exportKey,
    deriveKey
  },
  getRandomValues
}
