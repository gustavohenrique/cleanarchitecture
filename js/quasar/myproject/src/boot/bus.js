import { boot } from 'quasar/wrappers'

import bus from '../application/shared/bus'

export default boot(async ({ app }) => {
  app.config.globalProperties.$publish = bus.publish
  app.config.globalProperties.$subscribe = bus.subscribe
  app.config.globalProperties.$unsubscribe = bus.unsubscribe
})
