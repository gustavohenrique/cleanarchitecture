import { TodoItem, SearchRequest } from './main_pb'
import { TodoRpcPromiseClient } from './main_grpc_web_pb'

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
      search: async req => {
        try {
          const res = await todoClient.search(req, metadata)
          return res.getTodoitemsList()
        } catch(err) {
          throw new Error('Cannot search items: '+err)
        }
      }
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
