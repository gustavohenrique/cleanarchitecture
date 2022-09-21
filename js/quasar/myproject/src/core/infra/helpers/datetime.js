import { DateTime } from 'luxon'
const defaultPattern = 'dd/LL/yyyy HH:mm'

export default class {
  toString (val, pattern = defaultPattern) {
    if (!val) {
      return ''
    }
    return DateTime.fromISO(val).toFormat(pattern)
  }

  valueOf () {
    return DateTime.now().toMillis()
  }
}
