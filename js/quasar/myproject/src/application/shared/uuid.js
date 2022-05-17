export default function () {
  const rnds = new Uint8Array(16)
  crypto.getRandomValues(rnds)
  rnds[6] = (rnds[6] & 0x0f) | 0x40
  rnds[8] = (rnds[8] & 0x3f) | 0x80

  function bytesToUuid (buf) {
    const byteToHex = []
    for (let i = 0; i < 256; ++i) {
      byteToHex[i] = (i + 0x100).toString(16).substr(1)
    }
    let i = 0
    const bth = byteToHex
    return bth[buf[i++]] + bth[buf[i++]] +
      bth[buf[i++]] + bth[buf[i++]] + '-' +
      bth[buf[i++]] + bth[buf[i++]] + '-' +
      bth[buf[i++]] + bth[buf[i++]] + '-' +
      bth[buf[i++]] + bth[buf[i++]] + '-' +
      bth[buf[i++]] + bth[buf[i++]] +
      bth[buf[i++]] + bth[buf[i++]] +
      bth[buf[i++]] + bth[buf[i++]]
  }
  return bytesToUuid(rnds)
}
