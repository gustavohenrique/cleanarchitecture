export default class {
  constructor (params) {
    this.$accountRepository = params.$accountRepository
  }

  async findUserByEmail (entity) {
    return await this.$accountRepository.findUserByEmail(entity)
  }

  async signUp (entity) {
    return await this.$accountRepository.signUp(entity)
  }

  async signIn (entity) {
    return await this.$accountRepository.signIn(entity)
  }
}
