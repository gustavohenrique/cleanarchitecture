import axios from 'axios'
import { makeTodoClient, TodoItem } from './todo'

const sdkHeaders = {
  version: {
    key: 'X-SDK-Version',
    value: '1.0.0'
  },
  agent: {
    key: 'X-SDK-Agent',
    value: 'jsbrowser'
  },
  token: {
    key: 'X-CSRF-Token'
  }
}

function getAxios (config) {
  const instance = axios.create(config)
  instance.defaults.headers.common['Content-Type'] = 'application/json'
  instance.defaults.headers.common[sdkHeaders.version.key] = sdkHeaders.version.value
  instance.defaults.headers.common[sdkHeaders.agent.key] = sdkHeaders.agent.value
  const { url } = config
  if (url) {
    instance.defaults.baseURL = url
  }
  return instance
}

class SDK {
  constructor(config) {
    this._config = config
    this.$httpClient = getAxios(config)
    if (config.$axios) {
      const instance = config.$axios
      const { common } = instance.defaults.headers
      instance.defaults.headers.common = {
        ...common,
        [sdkHeaders.version.key]: sdkHeaders.version.value,
        [sdkHeaders.agent.key]: sdkHeaders.agent.version
      }
      this.$httpClient = instance
    }
  }

  getTodoClient() {
    const { token } = this._config
    const httpClient = this._withHeaders({ [sdkHeaders.token.key]: token })
    return makeTodoClient({ httpClient })
  }

  _withHeaders (headers) {
    const instance = getAxios(this._config)
    const { common } = instance.defaults.headers
    instance.defaults.headers.common = {
      ...common,
      ...headers
    }
    return instance
  }
}

const sdk = {
  SDK,
  entities: {
    TodoItem
  }
}
if (window) {
  window.{{ .ProjectName }} = sdk;
}
export default sdk;
