import chai from 'chai'

import UserAdapter from '../userAdapter'

const { expect } = chai

describe('UserAdapter', () => {
  const $adapter = new UserAdapter()

  describe('#toJSON', () => {
    it('Should return an object with all user attributes empty', async () => {
      const expected = {
        id: '',
        createdAt: '',
        fullName: '',
        email: '',
        salt: '',
        password: '',
        picture: ''
      }
      const user = $adapter.toJSON()
      expect(user).to.deep.equal(expected)
    })
  })
})
