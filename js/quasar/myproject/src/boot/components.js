import { boot } from 'quasar/wrappers'

const upperFirst = string => {
  return string ? string.charAt(0).toUpperCase() + string.slice(1) : ''
}
const wordPattern = new RegExp(['[A-Z][a-z]+', '[A-Z]+(?=[A-Z][a-z])', '[A-Z]+', '[a-z]+', '[0-9]+'].join('|'), 'g')
const camelCase = string => {
  const words = string.match(wordPattern) || []
  return words
    .map((word, index) => (index === 0 ? word.toLowerCase() : upperFirst(word.toLowerCase())))
    .join('')
}

const requireComponent = require.context(
  '../components',
  true,
  /[\w-]+\.vue$/
)

export default boot(async ({ app }) => {
  requireComponent.keys().forEach(fileName => {
    const componentConfig = requireComponent(fileName)
    const componentName = upperFirst(
      camelCase(
        fileName
          .replace(/^\.\/_/, '')
          .replace(/\.\w+$/, '')
          .replace('index', '')
      )
    )
    app.component(componentName, componentConfig.default || componentConfig)
  })
})
