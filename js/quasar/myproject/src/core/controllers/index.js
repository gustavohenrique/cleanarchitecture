import AccountController from './accountController'

export default function (deps) {
  return {
    Account: new AccountController(deps)
  }
}
