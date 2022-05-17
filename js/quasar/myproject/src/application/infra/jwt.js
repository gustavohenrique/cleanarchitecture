const defaultStorage = {
  get: key => localStorage.getItem(key),
  set: (key, value) => localStorage.setItem(key, value)
}

export default class {
  constructor (params = {}) {
    const storage = params.$storage || defaultStorage
    this.$storage = storage
    this.tokenKey = params.tokenKey || 'token'
  }

  getToken () {
    return this.$storage.get(this.tokenKey)
  }

  setToken (token) {
    return this.$storage.set(this.tokenKey, token)
  }

  getHeader () {
    const token = this.getToken()
    return {
      Authorization: `Bearer ${token}`
    }
  }
}
