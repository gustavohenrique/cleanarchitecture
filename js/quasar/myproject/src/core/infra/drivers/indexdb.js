import { openDB } from 'idb'

const modes = {
  ReadOnly: 'readonly',
  Write: 'readwrite'
}

class Database {
  constructor (params) {
    this._params = params
    this._table = params.table
    this._db = null
  }

  setTable (table) {
    this._table = table
  }

  async set (item) {
    const tx = await this.tx(modes.Write)
    await tx.objectStore(this._table).put(item)
  }

  async getById ({ id }) {
    const tx = await this.tx(modes.ReadOnly)
    return await tx.objectStore(this._table).get(id)
  }

  async get ({ index, key }) {
    return await this
      .tx(modes.ReadOnly)
      .getFromIndex(this._table, index, key)
  }

  async tx (m) {
    const mode = m || modes.Write
    const db = await this.db()
    return await db.transaction(this._table, mode)
  }

  async db () {
    if (!this._db) {
      await this._createdb(this._params)
    }
    return this._db
  }

  async _createdb ({ dbName = '', table = '', keyPath = 'id', indexes = [] }) {
    const db = await openDB(dbName, 1, {
      upgrade (db) {
        const store = db.createObjectStore(table, {
          keyPath,
          autoIncrement: false
        })
        indexes.forEach(index => store.createIndex(index, index))
      }
    })
    this._db = db
  }
}

export default function ({ dbName }) {
  return function (params) {
    return new Database({ ...params, dbName })
  }
}
