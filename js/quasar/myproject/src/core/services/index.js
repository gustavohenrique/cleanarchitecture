import AccountService from './accountService'

export default function ({ gateways }) {
  return {
    Account: new AccountService({ gateways })
  }
}
