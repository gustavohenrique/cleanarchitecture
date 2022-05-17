const arrayBufferToText = (buf) => {
  const byteArray = new Uint8Array(buf)
  let str = ''
  for (let i = 0; i < byteArray.byteLength; i++) {
    str += decodeURIComponent(String.fromCharCode(byteArray[i]))
  }
  return str
}

const textToArrayBuffer = str => {
  const buf = unescape(encodeURIComponent(str)) // 2 bytes for each char
  const bufView = new Uint8Array(buf.length)
  for (let i = 0; i < buf.length; i++) {
    bufView[i] = buf.charCodeAt(i)
  }
  return bufView
}

const hexToArrayBuffer = (hexString) => {
  if (hexString.length % 2 !== 0) {
    throw new Error('Invalid hexString.')
  }
  const arrayBuffer = new Uint8Array(hexString.length / 2)

  for (let i = 0; i < hexString.length; i += 2) {
    const byteValue = parseInt(hexString.substr(i, 2), 16)
    if (isNaN(byteValue)) {
      throw new Error('Byte value is not a number.')
    }
    arrayBuffer[i / 2] = byteValue
  }
  return arrayBuffer
}

const arrayBufferToHex = (b) => {
  if (!b) {
    throw new Error('No bytes to convert to Hex')
  }
  const bytes = new Uint8Array(b)
  const hexBytes = []
  for (let i = 0; i < bytes.length; ++i) {
    let byteString = bytes[i].toString(16)
    if (byteString.length < 2) {
      byteString = `0${byteString}`
    }
    hexBytes.push(byteString)
  }
  return hexBytes.join('')
}

const randomIv = (size) => {
  const bytes = window.crypto.getRandomValues(new Uint8Array(size / 2))
  const hexBytes = []
  for (let i = 0; i < bytes.length; ++i) {
    let byteString = bytes[i].toString(16)
    if (byteString.length < 2) {
      byteString = `0${byteString}`
    }
    hexBytes.push(byteString)
  }
  return hexBytes.join('')
}

export default {
  arrayBufferToText,
  textToArrayBuffer,
  hexToArrayBuffer,
  arrayBufferToHex,
  randomIv
}
