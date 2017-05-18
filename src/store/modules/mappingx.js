import Vue from 'vue'
import VueResource from 'vue-resource'
Vue.use(VueResource)
const mappingx = {
  state: {
    'toggleMapping': true,
    'mappingInfo': ''
  },
  mutations: {
    SET_TOGGLE_MAPPING: (state, toggle) => {
      state.toggleMapping = toggle
    },
    SET_MAPPING_INFO: (state, data) => {
      state.mappingInfo = data
    }
  },
  actions: {
    HandleIndexTemplate ({ commit }, body) {
      Vue.http.post(body.url, body)
      .then(
        response => {
          if (body.method === 'GET') {
            commit('SET_MAPPING_INFO', response.body.data)
          }
          if (response.body.result === 0 && body.method === 'DELETE') {
            commit('SET_TOGGLE_MAPPING', !body.toggle)
          }
        },
        error => {
          console.log(error)
          if (body.method === 'GET') {
            commit('SET_MAPPING_INFO', {info: '服务器异常'})
          }
        }
      )
    }
  }
}
export default mappingx
