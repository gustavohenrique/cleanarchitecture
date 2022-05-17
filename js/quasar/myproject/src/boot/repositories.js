import HttpClient from '../application/infra/httpClient'
import Jwt from '../application/infra/jwt'

const requireFile = require.context(
  '../application/repositories',
  false,
  /[\w-]+\.js$/
)

const params = {
  $httpClient: new HttpClient(),
  $jwt: new Jwt()
}

export function getRepositories () {
  const repositories = {}
  requireFile.keys().forEach(fileName => {
    const conf = requireFile(fileName)
    const name = fileName
      .replace(/^\.\//, '')
      .replace(/^\.\/_/, '')
      .replace(/\.\w+$/, '')
    const Repository = conf.default || conf
    repositories[`$${name}`] = new Repository(params)
  })
  return repositories
}
