import makeUser from '../entities/userEntity'

export default class {
  constructor (params) {
    this.$jwt = params.$jwt
    this.$httpClient = params.$httpClient.useToken(params.$jwt)
  }

  async findUserByEmail (entity) {
    await new Promise(resolve => setTimeout(resolve, 500))
    const email = entity.getEmail()
    if (email.indexOf('novo') >= 0) {
      throw new Error('Usuario nao existe')
    }
    const responseBody = {
      id: 'xpto-1-2',
      email,
      fullName: 'Usuário comum'
    }
    if (email.indexOf('admin') > 0) {
      responseBody.id = 'admin-123'
      responseBody.fullName = 'Admin'
    }
    return makeUser(responseBody)
    // return this.$httpClient.get(`/auth/${entity.email}`) */
  }

  async signUp (entity) {
    await new Promise(resolve => setTimeout(resolve, 500))
    this.saveTokenInLocalStorage('token123')
    entity.setId('some-id')
    return entity
  }

  async signIn (entity) {
    await new Promise(resolve => setTimeout(resolve, 500))
    this.saveTokenInLocalStorage('token123')
    entity.setId('some-id')
    return entity
  }

  saveTokenInLocalStorage (token) {
    this.$jwt.setToken(token)
  }
}
