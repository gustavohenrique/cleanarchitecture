import chai from 'chai'
import sinon from 'sinon'
import sinonChai from 'sinon-chai'

import AccountController from '../accountController'
import UserAdapter from '../../adapters/userAdapter'
import makeUser from '../../entities/userEntity'

chai.use(sinonChai)
const { expect } = chai

const fakeUser = {
  id: '1',
  createdAt: '2022-01-01T13:00:00',
  email: 'me@mail.com',
  salt: '1234567812345678',
  rawPassword: 'strongpass123',
  password: 'NzQwODc3ZmU0OTc1YTk5NzM5NmY1OTMxNzA4M2VmODM2OWRhNGU4NWJiYTg4NTRkNTQ1ODBhOTA4YWQ5YTc5NA=='
}

const $authService = {
  findUserByEmail: sinon.fake.returns(makeUser(fakeUser)),
  signUp: sinon.spy(),
  signIn: sinon.fake.returns(makeUser(fakeUser))
}

function getInstance () {
  const params = {
    $adapters: {
      $userAdapter: new UserAdapter()
    },
    $services: {
      $authService
    }
  }
  return new AccountController(params)
}

describe('AccountController', () => {
  const $controller = getInstance()

  describe('#findUserByEmail', () => {
    it('Should return an user found by email', async () => {
      const found = await $controller.findUserByEmail(fakeUser.email)
      expect($authService.findUserByEmail).to.have.been.called
      expect(found.id).to.equal(fakeUser.id)
      expect(found.email).to.equal(fakeUser.email)
    })
  })

  describe('#signUp', () => {
    it('Should create a pbkdf2 hash encoded in base64 using createdAt as key', async () => {
      await $controller.signUp(fakeUser)
      expect($authService.signUp.getCall(0).args[0].getPassword()).to.equal(fakeUser.password)
    })
  })

  describe('#signIn', () => {
    it('Should encrypt the raw password using the previous salt stored with user account', async () => {
      await $controller.signIn(fakeUser)
      expect($authService.signIn.getCall(0).args[0].getPassword()).to.equal(fakeUser.password)
    })
  })
})
