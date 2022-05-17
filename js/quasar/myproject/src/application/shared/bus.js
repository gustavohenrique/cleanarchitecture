import mitt from 'mitt'

const bus = mitt()
const loadingEvt = 'global-loading'

export default {
  publish: (evt, ...args) => bus.emit(evt, ...args),
  unsubscribe: (evt, ...args) => bus.off(evt, ...args),
  subscribe: (evt, callback) => {
    if (typeof evt === 'string') {
      bus.on(evt, callback)
    }
    if (Array.isArray(evt)) {
      for (const e of evt) {
        bus.on(e, callback)
      }
    }
  },
  setLoadingOn: () => bus.emit(loadingEvt, true),
  setLoadingOff: () => bus.emit(loadingEvt, false),
  loadingListener: (callback, event) => bus.on(event || loadingEvt, callback),
  destroyLoadingListener: (event) => bus.off(event || loadingEvt)
}
