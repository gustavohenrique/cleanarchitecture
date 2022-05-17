import axios from 'axios'

function getAxios (params) {
  const instance = axios.create(params)
  instance.defaults.headers.common['Content-Type'] = 'application/json'
  const baseUrl = params.baseUrl || process.env.API_BASE_URL
  if (baseUrl) {
    instance.defaults.baseURL = baseUrl
  }
  return instance
}

export default class {
  constructor (params = {}) {
    this.params = params
    this.$axios = getAxios(params)
  }

  withHeaders (headers) {
    const instance = getAxios(this.params)
    const { common } = instance.defaults.headers
    instance.defaults.headers.common = {
      ...common,
      headers
    }
    return instance
  }

  useToken (instance) {
    const header = instance.getHeader()
    return this.withHeaders(header)
  }
}
