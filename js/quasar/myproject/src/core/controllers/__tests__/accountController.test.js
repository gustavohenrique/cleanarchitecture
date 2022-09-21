import chai from 'chai'
import sinon from 'sinon'
import sinonChai from 'sinon-chai'

import AccountController from '../accountController'

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

const services = {
  Account: {
    findUserByEmail: sinon.fake.returns(fakeUser),
    signUp: sinon.fake.returns(fakeUser),
    signIn: sinon.fake.returns(fakeUser)
  }
}
const workers = {
  Account: {
    getHi: sinon.fake.returns('running tests'),
    sayHello: sinon.spy()
  }
}

function getInstance () {
  return new AccountController({ services, workers })
}

describe('AccountController', () => {
  const $controller = getInstance()

  describe('#findUserByEmail', () => {
    it('Should return an user found by email', async () => {
      const found = await $controller.findUserByEmail(fakeUser.email)
      expect(services.Account.findUserByEmail).to.have.been.called
      expect(found.id).to.equal(fakeUser.id)
      expect(found.email).to.equal(fakeUser.email)
    })
  })

  describe('#signUp', () => {
    it('Should create a pbkdf2 hash encoded in base64 using createdAt as key', async () => {
      await $controller.signUp(fakeUser)
      const want = services.Account.signUp.getCall(0).args[0].password
      expect(want).to.equal(fakeUser.password)
    })
  })

  describe('#signIn', () => {
    it('Should encrypt the raw password using the previous salt stored with user account', async () => {
      await $controller.signIn(fakeUser)
      const want = services.Account.signIn.getCall(0).args[0].password
      expect(want).to.equal(fakeUser.password)
    })
  })
})
