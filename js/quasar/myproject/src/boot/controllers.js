import { boot } from 'quasar/wrappers'

import bus from '../application/shared/bus'
import gravatar from '../application/shared/gravatar'
import { getAdapters } from './adapters'
import { getRepositories } from './repositories'
import { getServices } from './services'

const requireFile = require.context(
  '../application/controllers',
  false,
  /[\w-]+\.js$/
)

const deps = {
  $bus: bus,
  $gravatar: gravatar,
  $adapters: getAdapters(),
  $services: getServices(getRepositories())
}

export function getControllers (params) {
  const controllers = {}
  requireFile.keys().forEach(fileName => {
    const conf = requireFile(fileName)
    const name = fileName
      .replace(/^\.\//, '')
      .replace(/^\.\/_/, '')
      .replace(/\.\w+$/, '')
    const Controller = conf.default || conf
    controllers[`$${name}`] = new Controller(params)
  })
  return controllers
}

export default boot(async ({ app, router, store }) => {
  if (process.env.DEBUG) {
    console.log('[DEBUG] mode on')
  }
  const params = {
    ...deps,
    $store: store,
    $router: router
  }
  const controllers = getControllers(params)
  Object.keys(controllers).forEach(name => {
    app.config.globalProperties[name] = controllers[name]
  })
})
