import {{ .ProjectName }} from '../'
const { SDK, entities } = {{ .ProjectName }}

describe('#Search', () => {
  test('Should fetch all TODO items', async () => {
    const config = {
      token: 'XPTO-123',
      url: 'http://localhost:8001'
    }
    const sdk = new SDK(config)
    const client = sdk.getTodoClient()
    try {
      const res = await client.search()
      const items = res.getTodoitemsList()
      expect(items.length).toBe(1)
      expect('My awesome TODO').toBe(items[0].getTitle())
    } catch (err) {
      console.log('[ERROR]', err.message)
    }

  })
})
