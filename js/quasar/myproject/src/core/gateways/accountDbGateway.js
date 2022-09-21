export default class {
  constructor (useIndexDb) {
    this.$db = useIndexDb({ table: 'users' })
  }

  async save (item) {
    await this.$db.set(item)
  }

  async getById (id) {
    return await this.$db.getById({ id })
  }

  async findByEmail (email) {
    return await this.$db.get({ index: 'email', key: email })
  }
}
