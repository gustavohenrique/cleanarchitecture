export default class {
  constructor (params) {
    this.$userAdapter = params.$adapters.$userAdapter
    this.$authService = params.$services.$authService
  }

  async findUserByEmail (email) {
    const entity = this.$userAdapter.toEntity({ email })
    if (!entity.isValidEmail()) {
      throw new Error('Invalid email')
    }
    try {
      const found = await this.$authService.findUserByEmail(entity)
      return this.$userAdapter.toJSON(found)
    } catch (err) {
      return this.$userAdapter.toJSON(entity)
    }
  }

  async signUp (vo) {
    const entity = this.$userAdapter.toEntity(vo)
    await entity.encryptPassword(vo.rawPassword)
    const updated = await this.$authService.signUp(entity)
    return this.$userAdapter.toJSON(updated)
  }

  async signIn (vo) {
    const entity = this.$userAdapter.toEntity(vo)
    await entity.encryptPassword(vo.rawPassword)
    const found = await this.$authService.signIn(entity)
    return this.$userAdapter.toJSON(found)
  }

  getAvailableMenusFor (user) {
    return [
      {
        routeName: 'account',
        icon: 'manage_accounts'
      },
      {
        routeName: 'dns',
        icon: 'dns'
      },
      {
        routeName: 'insights',
        icon: 'analytics'
      },
      {
        routeName: 'help',
        icon: 'live_help'
      }
    ]
  }
}
