import { TodoItem, SearchRequest } from './{{ .ProjectName }}_pb'
import { TodoRpcPromiseClient } from './{{ .ProjectName }}_grpc_web_pb'

class SDK {
  constructor(config) {
    this._config = config
    this._todoClient = new TodoRpcPromiseClient(config.url, null, null)
  }

  getTodoClient() {
    const todoClient = this._todoClient
    const metadata = {
      ...this.getHeaders(),
      ...this.getDeadline()
    }
    return {
      ...todoClient,
      search: req => todoClient.search(req, metadata)
    }
  }

  getHeaders() {
    const { token } = this._config
    return {
      'X-CSRF-Token': token
    }
  }

  getDeadline() {
    const now = new Date();
    now.setSeconds(now.getSeconds() + this._config.deadline || 10)
    return {
      deadline: now.getTime()
    }
  }
}

const sdk = {
  SDK,
  entities: {
    TodoItem,
    SearchRequest
  }
}
if (window) {
  window.{{ .ProjectName }} = sdk;
}
export default sdk;
