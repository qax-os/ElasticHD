import Vue from 'vue'
import VueResource from 'vue-resource'
Vue.use(VueResource)
const mappingx = {
  state: {
    toggleMapping: true,
    mappingInfo: '',
    DeleteMappingError: 'Delete Mapping Error',
    DeleteMappingSuccess: 'Delete Mapping Success.'
  },
  mutations: {
    SET_DELETE_MAPPING_ERROR: (state, text) => {
      state.DeleteMappingError = text
    },
    SET_DELETE_MAPPING_SUCCESS: (state, text) => {
      state.DeleteMappingSuccess = text
    },
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
            return
          }
          if (response.body.result === 0 && body.method === 'DELETE') {
            // commit('SET_TOGGLE_MAPPING', !body.toggle)
            commit('SET_DELETE_MAPPING_SUCCESS', body.indices + ' Delete Success.')
          } else {
            commit('SET_DELETE_MAPPING_ERROR', body.indices + ' Delete Error.')
          }
        },
        error => {
          console.log(error)
          if (body.method === 'GET') {
            commit('SET_MAPPING_INFO', {info: '服务器异常'})
          } else {
            commit('SET_DELETE_MAPPING_ERROR', body.indices + ' Delete Error.')
          }
        }
      )
    }
  }
}
export default mappingx
