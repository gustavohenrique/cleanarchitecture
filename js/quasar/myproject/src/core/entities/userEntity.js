import crypto from '../infra/helpers/crypto'
import gravatar from '../infra/helpers/gravatar'

export default class {
  constructor (params = {}, deps = {}) {
    this.$crypto = deps.crypto || crypto
    this.$gravatar = deps.gravatar || gravatar

    this.id = params.id || ''
    this.createdAt = params.createdAt || ''
    this.fullName = params.fullName || ''
    this.email = params.email || ''
    this.salt = params.salt || ''
    this.password = params.pasword || ''
    this.picture = params.picture || ''
  }

  getId () {
    return this.id
  }

  setId (data = '') {
    this.id = data
  }

  getCreatedAt () {
    return this.createdAt
  }

  setCreatedAt (data = '') {
    this.createdAt = data
  }

  getFullName () {
    return this.fullName
  }

  setFullName (data = '') {
    this.fullName = data
  }

  getEmail () {
    return this.email
  }

  setEmail (data = '') {
    this.email = data
  }

  getPassword () {
    return this.getEncryptedPassword()
  }

  setPassword (data = '') {
    this.password = data
  }

  getSalt () {
    return this.salt
  }

  setSalt (data = '') {
    this.salt = data
  }

  getPicture () {
    if (this.picture) {
      return this.picture
    }
    return this.$gravatar(this.email)
  }

  setPicture (data = '') {
    this.picture = data
  }

  isValidEmail () {
    const email = this.getEmail()
    return email && email.indexOf('@') > 0
  }

  async getEncryptedPassword () {
    // Password deve ser um hash gerado usando pbkdf2
    // Depois criptografado com a chave publica RSA do usuario
    // Em seguida, codificado como base64
    // E no servidor decodificar e descriptografar usando a chave privada RSA
    // Salvando o hash no banco de dados
    const { $crypto } = this
    const pbkdf2 = await $crypto.pbkdf2.hashIt({
      key: this.getCreatedAt(),
      salt: this.getSalt(),
      raw: this.password
    })
    const { hash, salt } = pbkdf2
    const encoded = $crypto.base64.encode(hash)
    this.setSalt(salt)
    return encoded
  }

  toJSON () {
    return {
      id: this.getId(),
      fullName: this.getFullName(),
      createdAt: this.getCreatedAt(),
      email: this.getEmail(),
      salt: this.getSalt(),
      password: '',
      picture: this.getPicture()
    }
  }

  fromStorage (params) {
    this.id = params.id || ''
    this.createdAt = params.createdAt || ''
    this.fullName = params.fullName || ''
    this.email = params.email || ''
    this.salt = params.salt || ''
    this.password = params.pasword || ''
    this.picture = params.picture || ''
    return this
  }

  fromDb (params) {
    return this.fromStorage(params)
  }
}
