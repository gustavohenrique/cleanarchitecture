import {{ .ProjectName }} from '../'
const { SDK, entities } = {{ .ProjectName }}

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
      expect(res.length).toBe(1)
      expect(res[0].getTitle()).toBe('My awesome TODO')
    } catch (err) {
      console.log('[ERROR]', err.message)
      throw err
    }

  })
})
