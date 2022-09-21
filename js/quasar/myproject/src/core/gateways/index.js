import AccountLocalStorageGateway from './accountLocalStorageGateway'
import AccountDbGateway from './accountDbGateway'

export default function ({ storage, indexdb }) {
  return {
    AccountStorage: new AccountLocalStorageGateway(storage),
    AccountDb: new AccountDbGateway(indexdb)
  }
}
