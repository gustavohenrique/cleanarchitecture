export default class {
  constructor (storage) {
    this.$storage = storage
  }

  async findUserByEmail (email) {
    await new Promise(resolve => setTimeout(resolve, 500))
    if (email.indexOf('signup') >= 0) {
      throw new Error('User not found')
    }
    const responseBody = {
      id: 'xpto-1-2',
      email,
      fullName: 'Gustavo Henrique'
    }
    if (email.indexOf('admin') > 0) {
      responseBody.id = 'admin-123'
      responseBody.fullName = 'Admin'
    }
    return responseBody
  }

  async signUp (item) {
    await new Promise(resolve => setTimeout(resolve, 500))
    this.$storage.set('token', 'token-123')
    this.$storage.set('id', item.id)
    return item
  }

  async signIn (item) {
    await new Promise(resolve => setTimeout(resolve, 500))
    this.$storage.set('token', 'another-token-123')
    this.$storage.set('id', item.id)
    return item
  }

  getId () {
    return this.$storage.get('id')
  }
}
