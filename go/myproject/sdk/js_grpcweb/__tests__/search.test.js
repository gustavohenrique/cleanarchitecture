import myproject from '../'
const { SDK, entities } = myproject

describe('#Search', () => {
  test('Should fetch all TODO items', async () => {
    const config = {
      token: 'XPTO-123',
        deadline: 5, // 5s timeout
      url: 'http://localhost:8001'
    }
    const sdk = new SDK(config)
    const client = sdk.getTodoClient()
    try {
      const req = new entities.SearchRequest()
      const res = await client.search(req)
      const items = res.getTodoitemsList()
      expect(1).toBe(items.length)
      expect('My awesome TODO').toBe(items[0].getTitle())
    } catch (err) {
      console.log('[ERROR]', err.message)
    }

  })
})
