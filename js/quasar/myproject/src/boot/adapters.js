const requireFile = require.context(
  '../application/adapters',
  false,
  /[\w-]+\.js$/
)

export function getAdapters () {
  const adapters = {}
  requireFile.keys().forEach(fileName => {
    const conf = requireFile(fileName)
    const name = fileName
      .replace(/^\.\//, '')
      .replace(/^\.\/_/, '')
      .replace(/\.\w+$/, '')
    const Adapter = conf.default || conf
    adapters[`$${name}`] = new Adapter()
  })
  return adapters
}
