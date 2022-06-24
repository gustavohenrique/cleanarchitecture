import { Notify } from 'quasar'

const getMessage = err => {
  if (err && typeof (err) !== 'string') {
    if (err.message === 'Network Error') {
      return 'Não consegui conectar ao servidor'
    }
    return err.message ? err.message : err.toString()
  }
  return err || ''
}

export default class {
  notify (args) {
    const params = args || {}
    const opts = {
      type: params.type || 'negative',
      timeout: params.timeout || 2000,
      progress: params.progress || true,
      message: params.message || getMessage(this.err)
    }
    Notify.create(opts)
  }
}