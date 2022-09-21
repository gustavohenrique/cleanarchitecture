class LocalStorage {
  set (key, value) {
    localStorage.setItem(key, value)
  }

  get (key) {
    return localStorage.getItem(key)
  }
}

export default function () {
  return new LocalStorage()
}
