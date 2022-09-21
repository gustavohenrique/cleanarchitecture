import { wrap } from 'comlink'

function makeWorkerFromFile (file) {
  const url = URL.createObjectURL(new File([file], import.meta.url))
  const worker = new SharedWorker(url)
  return wrap(worker.port)
}

export default function () {
  return {
    Account: makeWorkerFromFile('accountWorker.js')
  }
}
