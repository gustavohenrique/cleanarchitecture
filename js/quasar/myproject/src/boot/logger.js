import { boot } from 'quasar/wrappers'

import Logger from '../application/shared/logger'

export default boot(async ({ app }) => {
  app.config.globalProperties.$logger = new Logger()
})
