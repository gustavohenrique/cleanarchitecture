import { TodoItem, SearchRequest } from './myproject_pb'
import { TodoRpcPromiseClient } from './myproject_grpc_web_pb'

class SDK {
  constructor(config) {
    this._config = config
    this._todoClient = new TodoRpcPromiseClient(config.url, null, null)
  }

  getTodoClient() {
    const todoClient = this._todoClient
    const metadata = this.getHeaders()
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
}

const sdk = {
  SDK,
  entities: {
    TodoItem,
    SearchRequest
  }
}
if (window) {
  window.myproject = sdk;
}
export default sdk;
