import { offlineFallback, warmStrategyCache, googleFontsCache } from 'workbox-recipes'
import { CacheFirst, StaleWhileRevalidate } from 'workbox-strategies'
import { registerRoute, Route } from 'workbox-routing'
import { CacheableResponsePlugin } from 'workbox-cacheable-response'
import { ExpirationPlugin } from 'workbox-expiration'
import { precacheAndRoute } from 'workbox-precaching'

const DAY = 24 * 60 * 60

precacheAndRoute(self.__WB_MANIFEST)
googleFontsCache()

/*
 * Cache First checks the cache for a response first and uses it if available.
 * If the request isn't in the cache, the network is used and any valid
 * response is added to the cache before being passed to the browser.
 */
const pageCache = new CacheFirst({
  cacheName: 'page-cache',
  plugins: [
    new CacheableResponsePlugin({
      statuses: [0, 200]
    }),
    new ExpirationPlugin({
      maxAgeSeconds: DAY
    })
  ]
})

warmStrategyCache({
  urls: ['/index.html', '/'],
  strategy: pageCache
})
registerRoute(({ request }) => request.mode === 'navigate', pageCache)

async function handleRequestNetworkFirst (cacheName, request) {
  const key = JSON.stringify(request.clone())
  const cache = await caches.open(cacheName)
  try {
    const networkResponse = await fetch(request)
    const fetchedResponse = networkResponse.clone()
    if (fetchedResponse.ok) {
      cache.put(key, fetchedResponse)
      return networkResponse
    }
  } catch (err) {}
  return await cache.match(key)
}

addEventListener('fetch', async event => {
  const request = event.request.clone()
  const services = ['/api/some/endpoint']
  const path = request.url.replace(request.referrer, '/')
  if (services.includes(path)) {
    event.respondWith(handleRequestNetworkFirst('api-cache', request))
  }
})

registerRoute(({ url }) => url.hostname.indexOf('googleusercontent.com') > 0, new StaleWhileRevalidate({
  cacheName: 'google-avatar-cache',
  plugins: [
    new ExpirationPlugin({
      maxAgeSeconds: 3 * DAY,
      maxEntries: 3
    })
  ]
}))

registerRoute(({ request }) => request.destination === 'image', new CacheFirst({
  cacheName: 'images-cache',
  plugins: [
    new ExpirationPlugin({
      maxAgeSeconds: 3 * DAY
    })
  ]
}))

/*
 * Stale While Revalidate uses a cached response for a request if it's
 * available and updates the cache in the background with a response from the
 * network. Therefore, if the asset isn't cached, it will wait for the network
 * response and use that. It's a fairly safe strategy, as it regularly updates
 * cache entries that rely on it. The downside is that it always requests an
 * asset from the network in the background.
 */
const isAsset = ({ request }) => {
  return ['style', 'script', 'worker'].includes(request.destination)
}
const assetsRoute = new Route(isAsset, new StaleWhileRevalidate({
  cacheName: 'assets-cache',
  plugins: [
    new CacheableResponsePlugin({
      statuses: [0, 200]
    }),
    new ExpirationPlugin({
      maxAgeSeconds: 3 * DAY
    })
  ]
}))
registerRoute(assetsRoute)

// Set up offline fallback
offlineFallback({
  pageFallback: '/index.html/#/offline'
})
