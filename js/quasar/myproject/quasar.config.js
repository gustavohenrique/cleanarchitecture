/* eslint-env node */

/*
 * This file runs in a Node context (it's NOT transpiled by Babel), so use only
 * the ES6 features that are supported by your Node version. https://node.green/
 */

// Configuration for your app
// https://v2.quasar.dev/quasar-cli-webpack/quasar-config-js

const ESLintPlugin = require('eslint-webpack-plugin')
const { configure } = require('quasar/wrappers')

module.exports = configure(function (ctx) {
  require('dotenv').config({ path: require('path').join(__dirname, `./.env.${process.env.NODE_ENV}`)})
  return {
    supportTS: false,
    // https://v2.quasar.dev/quasar-cli-webpack/prefetch-feature
    // preFetch: true,
    // app boot file (/src/boot)
    // --> boot files are part of "main.js"
    boot: [
      'components',
      'bus',
      'logger',
      'i18n',
      'controllers'
    ],

    css: [
      'app.css'
    ],

    extras: [
      'roboto-font',
      'material-icons'
    ],

    htmlVariables: {
      buildAt: new Date().toISOString(),
      version: process.env.VERSION || '0.0.1'
    },

    build: {
      vueRouterMode: 'hash', // available values: 'hash', 'history'
      env: {
        VERSION: process.env.VERSION,
        BUILD_AT: process.env.BUILD_AT,
        DEBUG: process.env.DEBUG,
        API_BASE_URL: process.env.API_BASE_URL,
        API_TIMEOUT: process.env.API_TIMEOUT
      },

      chainWebpack (chain) {
        chain.plugin('eslint-webpack-plugin')
          .use(ESLintPlugin, [{ extensions: ['js', 'vue'] }])
      }
    },

    devServer: {
      server: {
        type: 'http'
      },
      port: 8080,
      open: false // opens browser window automatically
    },

    // https://v2.quasar.dev/quasar-cli-webpack/quasar-config-js#Property%3A-framework
    framework: {
      config: {
        brand: {
          primary: '#3232F8',
          secondary: '#7474FA',
          accent: '#7732F8',
          positive: '#18bd5c',
          negative: '#E81f46',
          info: '#FF8A3C',
          warning: '#FFC22B'
        },
        screen: {
          bodyClasses: true
        }
      },
      plugins: [
        'Dialog',
        'Notify'
      ]
    },

    // animations: 'all', // --- includes all animations
    // https://quasar.dev/options/animations
    animations: [],

    ssr: {
      pwa: false,
      prodPort: 3000, // The default port that the production server should use
      maxAge: 1000 * 60 * 60 * 24 * 30, // Tell browser when a file from the server should expire from cache (in ms)
      chainWebpackWebserver (chain) {
        chain.plugin('eslint-webpack-plugin')
          .use(ESLintPlugin, [{ extensions: ['js'] }])
      },
      middlewares: [
        ctx.prod ? 'compression' : '',
        'render' // keep this as last one
      ]
    },

    pwa: {
      workboxPluginMode: 'GenerateSW', // 'GenerateSW' or 'InjectManifest'
      workboxOptions: {}, // only for GenerateSW

      // for the custom service worker ONLY (/src-pwa/custom-service-worker.[js|ts])
      // if using workbox in InjectManifest mode
      chainWebpackCustomSW (chain) {
        chain.plugin('eslint-webpack-plugin')
          .use(ESLintPlugin, [{ extensions: ['js'] }])
      },

      manifest: {
        name: 'myproject',
        short_name: 'myproject',
        description: '',
        display: 'standalone',
        orientation: 'portrait',
        background_color: '#ffffff',
        theme_color: '#027be3',
        icons: [
          {
            src: 'icons/icon-128x128.png',
            sizes: '128x128',
            type: 'image/png'
          },
          {
            src: 'icons/icon-192x192.png',
            sizes: '192x192',
            type: 'image/png'
          },
          {
            src: 'icons/icon-256x256.png',
            sizes: '256x256',
            type: 'image/png'
          },
          {
            src: 'icons/icon-384x384.png',
            sizes: '384x384',
            type: 'image/png'
          },
          {
            src: 'icons/icon-512x512.png',
            sizes: '512x512',
            type: 'image/png'
          }
        ]
      }
    },

    cordova: {
    },

    capacitor: {
      hideSplashscreen: true
    },

    electron: {
      bundler: 'packager', // 'packager' or 'builder'

      packager: {
      },

      builder: {
        appId: 'frontend'
      },

      chainWebpackMain (chain) {
        chain.plugin('eslint-webpack-plugin')
          .use(ESLintPlugin, [{ extensions: ['js'] }])
      },

      chainWebpackPreload (chain) {
        chain.plugin('eslint-webpack-plugin')
          .use(ESLintPlugin, [{ extensions: ['js'] }])
      }
    }
  }
})
