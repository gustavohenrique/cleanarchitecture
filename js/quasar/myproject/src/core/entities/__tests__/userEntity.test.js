import chai from 'chai'
import sinon from 'sinon'

import UserEntity from '../userEntity'

const { expect } = chai

describe('UserEntity', () => {
  const crypto = {
    pbkdf2: {
      hashIt: sinon.fake.returns({ hash: 'hash', salt: '123' })
    },
    base64: {
      encode: sinon.fake.returns('base64')
    }
  }
  const gravatar = sinon.fake.returns('http://gravatar/me')

  const entity = new UserEntity({
    id: '1234',
    password: 'strongpass123',
    fullName: 'Gustavo Henrique',
    email: 'me@mail.com'
  }, { crypto, gravatar })

  describe('#toJSON', () => {
    it('Should return an object with all user attributes', async () => {
      const expected = {
        id: entity.getId(),
        createdAt: '',
        fullName: entity.getFullName(),
        email: entity.getEmail(),
        salt: '',
        password: '',
        picture: 'http://gravatar/me'
      }
      const user = entity.toJSON()
      expect(user).to.deep.equal(expected)
    })
  })
})
