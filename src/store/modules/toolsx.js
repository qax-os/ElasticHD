import Vue from 'vue'
import VueResource from 'vue-resource'
Vue.use(VueResource)
const toolsbarx = {
  state: {
    'sqlBody': '',
    'dslBody': '',
    'currentError': ''
  },
  mutations: {
    SET_DSL_BODY: (state, dsl) => {
      state.dslBody = dsl
    },
    SET_SQL_BODY: (state, sql) => {
      state.sqlBody = sql
    },
    SET_CURRENT_ERROR: (state, errorMessage) => {
      state.currentError = errorMessage
    }
  },
  actions: {
    ConvertSQLToDSL ({ commit }, body) {
      commit('SET_SQL_BODY', body.sql)
      Vue.http.post(body.url, body)
      .then(
        response => {
          if (response.body.result === 0) {
            commit('SET_DSL_BODY', response.body.data)
            commit('SET_CURRENT_ERROR', '')
          } else {
            console.log(222)
            commit('SET_CURRENT_ERROR', response.body.data)
          }
        },
        error => {
          commit('SET_CURRENT_ERROR', error)
          console.log(error)
        }
      )
    }
  }
}
export default toolsbarx
