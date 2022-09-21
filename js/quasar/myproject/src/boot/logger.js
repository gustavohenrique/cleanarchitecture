import { boot } from 'quasar/wrappers'

import Logger from '../core/infra/helpers/logger'

export default boot(async ({ app }) => {
  app.config.globalProperties.$logger = new Logger()
})
