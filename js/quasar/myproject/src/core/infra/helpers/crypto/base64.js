const keystr = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'

function atobLookup (chr) {
  const index = keystr.indexOf(chr)
  return index < 0 ? undefined : index
}

function btoaLookup (index) {
  if (index >= 0 && index < 64) {
    return keystr[index]
  }
  return undefined
}

export default class {
  encode (data) {
    return this._btoa(data)
  }

  decode (data) {
    return this._atob(data)
  }

  _btoa (s) {
    if (arguments.length === 0) {
      throw new TypeError('1 argument required, but only 0 present.')
    }

    let i
    // String conversion as required by Web IDL.
    s = `${s}`
    // 'The btoa() method must throw an 'InvalidCharacterError' DOMException if
    // data contains any character whose code point is greater than U+00FF.'
    for (i = 0; i < s.length; i++) {
      if (s.charCodeAt(i) > 255) {
        return null
      }
    }
    let out = ''
    for (i = 0; i < s.length; i += 3) {
      const groupsOfSix = [undefined, undefined, undefined, undefined]
      groupsOfSix[0] = s.charCodeAt(i) >> 2
      groupsOfSix[1] = (s.charCodeAt(i) & 0x03) << 4
      if (s.length > i + 1) {
        groupsOfSix[1] |= s.charCodeAt(i + 1) >> 4
        groupsOfSix[2] = (s.charCodeAt(i + 1) & 0x0f) << 2
      }
      if (s.length > i + 2) {
        groupsOfSix[2] |= s.charCodeAt(i + 2) >> 6
        groupsOfSix[3] = s.charCodeAt(i + 2) & 0x3f
      }
      for (let j = 0; j < groupsOfSix.length; j++) {
        if (typeof groupsOfSix[j] === 'undefined') {
          out += '='
        } else {
          out += btoaLookup(groupsOfSix[j])
        }
      }
    }
    return out
  }

  _atob (data) {
    if (arguments.length === 0) {
      throw new TypeError('1 argument required, but only 0 present.')
    }

    // Web IDL requires DOMStrings to just be converted using ECMAScript
    // ToString, which in our case amounts to using a template literal.
    data = `${data}`
    // 'Remove all ASCII whitespace from data.'
    data = data.replace(/[ \t\n\f\r]/g, '')
    // 'If data's length divides by 4 leaving no remainder, then: if data ends
    // with one or two U+003D (=) code points, then remove them from data.'
    if (data.length % 4 === 0) {
      data = data.replace(/==?$/, '')
    }
    if (data.length % 4 === 1 || /[^+/0-9A-Za-z]/.test(data)) {
      return null
    }
    // 'Let output be an empty byte sequence.'
    let output = ''
    let buffer = 0
    let accumulatedBits = 0
    for (let i = 0; i < data.length; i++) {
      buffer <<= 6
      buffer |= atobLookup(data[i])
      accumulatedBits += 6
      if (accumulatedBits === 24) {
        output += String.fromCharCode((buffer & 0xff0000) >> 16)
        output += String.fromCharCode((buffer & 0xff00) >> 8)
        output += String.fromCharCode(buffer & 0xff)
        buffer = accumulatedBits = 0
      }
    }
    if (accumulatedBits === 12) {
      buffer >>= 4
      output += String.fromCharCode(buffer)
    } else if (accumulatedBits === 18) {
      buffer >>= 2
      output += String.fromCharCode((buffer & 0xff00) >> 8)
      output += String.fromCharCode(buffer & 0xff)
    }
    return output
  }
}
