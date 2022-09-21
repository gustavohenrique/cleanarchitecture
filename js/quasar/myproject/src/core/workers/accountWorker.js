import { expose } from 'comlink'

class AccountWorker {
  sayHello (msg) {
    console.log('>>> Hello!', msg)
  }

  getHi () {
    return 'Hi!!'
  }
}

self.onconnect = (evt) => {
  expose(new AccountWorker(), evt.ports[0])
}
