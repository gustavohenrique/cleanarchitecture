import PBKDF2 from './pbkdf2'
import Base64 from './base64'

/*
class Base64 {
  encode (str) {
    return helpers.strToBase64(str)
  }

  decode (b64) {
    return helpers.base64ToStr(b64)
  }
}
*/

export default {
  pbkdf2: new PBKDF2(),
  base64: new Base64()
}
