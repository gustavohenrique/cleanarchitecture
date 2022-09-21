import useStorage from './infra/drivers/storage'
import useIndexdb from './infra/drivers/indexdb'
import useHttpClient from './infra/drivers/httpClient'

import makeGateways from './gateways'
import makeServices from './services'
import makeControllers from './controllers'
import makeWorkers from './workers'

const gateways = makeGateways({
  storage: useStorage(),
  indexdb: useIndexdb({ dbName: '{{ .ProjectName }}' }),
  httpClient: useHttpClient({ baseUrl: 'http://' })
})
const services = makeServices({ gateways })
const workers = makeWorkers()
const controllers = makeControllers({ services, workers })

export default function () {
  return {
    $controllers: controllers
  }
}
