import Vue from 'vue'
import VueResource from 'vue-resource'
Vue.use(VueResource)

const searchx = {
  state: {
    options: [
     {'value': 'GET', 'label': 'GET'},
     {'value': 'PUT', 'label': 'PUT'},
     {'value': 'DELETE', 'label': 'DELETE'},
     {'value': 'POST', 'label': 'POST'},
     {'value': 'HEAD', 'label': 'HEAD'}
    ],
    fields: [],
    result: '{}',
    showSpinner: false
  },
  mutations: {
    SET_RESULT: (state, data) => {
      state.result = data
    },
    SET_FIELDS: (state, data) => {
      state.fields = data
    },
    SET_SHOW_SPINNER: (state, flag) => {
      state.showSpinner = flag
    }
  },
  actions: {
    CURLSerachResult ({commit}, body) {
      Vue.http.post(body.url, body)
      .then(
        response => {
          commit('SET_SHOW_SPINNER', false)
          commit('SET_RESULT', response.body.data)
        },
        error => {
          commit('SET_SHOW_SPINNER', false)
          commit('SET_RESULT', error)
          console.log(error)
        }
      )
    },
    CURLSerachFields ({commit}, body) {
      Vue.http.post(body.url, body)
      .then(
        response => {
          if (response.body.result === 0) {
            commit('SET_FIELDS', response.body.data)
          }
        },
        error => {
          console.log(error)
        }
      )
    },
    SetShowSpinner ({commit}, flag) {
      commit('SET_SHOW_SPINNER', flag)
    }
  }
}
export default searchx
