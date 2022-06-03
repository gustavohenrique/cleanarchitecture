class TodoItem {
  constructor() {
    this._id = ''
    this._title = ''
  }

  setId(id) {
    this._id = id
  }

  setTitle(title) {
    this._title = title
  }

  getId() {
    return this._id
  }

  getTitle() {
    return this._title
  }
}

export const makeTodoClient = ({ httpClient }) => {
  return {
    async search() {
      try {
        const res = await httpClient.get('/todo')
        const { data } = res.data
        return {
          getTodoitemsList() {
            return data.map(i => {
              const item = new TodoItem()
              item.setId(i.id)
              item.setTitle(i.title)
              return item
            })
          }
        }
      } catch(err) {
        console.log('[ERROR]', err.response.data)
        return err.response.data
      }
    }
  }
}

export default {
  makeTodoClient,
  TodoItem
}
