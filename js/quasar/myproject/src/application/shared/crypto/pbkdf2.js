import helpers from './helpers'

const pbkdf2Algo = {
  name: 'PBKDF2'
}

const aesAlgo = {
  name: 'AES-CBC',
  length: 256
}

export default class {
  async digest ({ salt, plaintext }) {
    const key = await crypto.subtle.importKey('raw', helpers.textToArrayBuffer(plaintext), pbkdf2Algo, false, ['deriveKey'])
    const s = salt ? helpers.hexToArrayBuffer(salt) : crypto.getRandomValues(new Uint8Array(16))
    const params = {
      salt: s,
      name: pbkdf2Algo.name,
      iterations: 100,
      hash: 'SHA-256'
    }
    const result = await crypto.subtle.deriveKey(params, key, aesAlgo, true, ['encrypt', 'decrypt'])
    const exported = await crypto.subtle.exportKey('raw', result)
    return {
      salt: helpers.arrayBufferToHex(params.salt),
      hex: helpers.arrayBufferToHex(new Uint8Array(exported))
      // base64: helpers.arrayBufferToBase64(new Uint8Array(exported))
    }
  }

  async hashIt ({ raw, key, salt }) {
    const plaintext = `${raw}${key}`
    const res = await this.digest({ salt, plaintext })
    return {
      hash: res.hex,
      salt: res.salt
    }
  }
}
