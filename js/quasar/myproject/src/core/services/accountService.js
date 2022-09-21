import UserEntity from '../entities/userEntity'

export default class {
  constructor ({ gateways }) {
    this.$accountStorage = gateways.AccountStorage
    this.$accountDb = gateways.AccountDb
  }

  async findUserByEmail (email) {
    const entity = new UserEntity({ email })
    try {
      const found = await this.$accountStorage.findUserByEmail(email)
      return new UserEntity().fromStorage(found).toJSON()
    } catch (err) {
      return entity.toJSON()
    }
  }

  async getLocalUser () {
    const user = new UserEntity().toJSON()
    try {
      const id = this.$accountStorage.getId()
      if (!id) {
        return user
      }
      const found = await this.$accountDb.getById(id)
      return new UserEntity().fromDb(found).toJSON()
    } catch (err) {
      return user
    }
  }

  async signUp (vo) {
    const entity = new UserEntity(vo)
    const saved = await this.$accountStorage.signUp(entity.toJSON())
    await this.$accountDb.save(saved)
    return new UserEntity().fromStorage(saved).toJSON()
  }

  async signIn (vo) {
    const entity = new UserEntity(vo)
    const saved = await this.$accountStorage.signIn(entity.toJSON())
    await this.$accountDb.save(saved)
    return new UserEntity().fromStorage(saved).toJSON()
  }
}
