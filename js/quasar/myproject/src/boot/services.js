const requireFile = require.context(
  '../application/services',
  false,
  /[\w-]+\.js$/
)

export function getServices (repositories) {
  const services = {}
  requireFile.keys().forEach(fileName => {
    const conf = requireFile(fileName)
    const name = fileName
      .replace(/^\.\//, '')
      .replace(/^\.\/_/, '')
      .replace(/\.\w+$/, '')
    const Service = conf.default || conf
    services[`$${name}`] = new Service(repositories)
  })
  return services
}
