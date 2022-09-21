export default class {
  constructor ({ services = {}, workers = {} }) {
    this.$accountService = services.Account
    this.$accountWorker = workers.Account
  }

  async findUserByEmail (email) {
    await this.$accountWorker.sayHello(email)
    const answer = await this.$accountWorker.getHi()
    console.log('.answer=', answer)
    return await this.$accountService.findUserByEmail(email)
  }

  async getLocalUser () {
    return await this.$accountService.getLocalUser()
  }

  async signUp (vo) {
    return await this.$accountService.signUp(vo)
  }

  async signIn (vo) {
    return await this.$accountService.signIn(vo)
  }
}
