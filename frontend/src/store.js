import Vue from 'vue'
import Vuex from 'vuex'
import KeyPairModule from './modules/keypair/index'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    loading: false,
     error: {
      module: '',
      status: false,
      message: ''
    },
  },
   mutations: {
    setLoading(state, loading) {
      state.loading = loading
    },
    setError(state, data) {

        if(! data) {
            state.error.status = false
            state.error.module = ''
            state.error.message = ''

            return
        }

        state.error.status = true
        state.error.module = data.modulo
        if(!data.error.response){
            state.error.message = "El servidor no responde, por favor intente mas tarde"
        }else if(data.error.response.status === 401){
            state.error.message = "Las credenciales proporcionadas son incorrectas"
        }else if(data.error.response.status === 500){
            state.error.message = "Error interno en el servidor"
        } else {
            state.error.message = data.error.message
        }

    }
  },
  modules: {
    KeyPairModule
  }
})
