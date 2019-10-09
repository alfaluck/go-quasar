import axios from 'axios'

const handler = axios.create()
handler.interceptors.request.use(
  function (config) {
    console.log(config)
    if (localStorage.getItem(id_token)) {
      
    }
    return config
  },
  function (error) {
    console.log(error)
    return Promise.reject(error)
  }
)

class RPC {
  constructor (opts) {
    this._opts = { ...opts }
    this.currId = 1
    handler.defaults.baseURL = opts.host
    handler.defaults.headers.post['Content-Type'] = 'application/json'
  }

  call = (method, params = null) =>
    axios.post('/', {
      jsonrpc: '2.0',
      id: ++this.currId,
      method,
      params: typeof params === 'string' ? [params] : params
    }).then(({ data: { /* id, */ result} }) => result);
}

// leave the export, even if you don't use it
export default ({ app, router, Vue }) => {
  const rpc = axios.create()
  rpc.prototype.call = function (method, params) {
    this.currId = 1
    return this.post('/rpc', {
      jsonrpc: '2.0',
      id: ++this.currId,
      method,
      params: typeof params === 'string' ? [params] : params
    })
  }
}
