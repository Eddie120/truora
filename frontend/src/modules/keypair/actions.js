import Vue from 'vue'

export async function cargarLlaves({commit}, term = null) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})

        let url = '/llaves'
        if(term != null) {
            url = '/llaves?texto='+term
        }

        const {data} = await Vue.axios({
            url: url
        })

        commit('setLlaves', data)
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}


export async function crearLlave({commit}, llave) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})
        const respuesta = await Vue.axios({
            method: 'POST',
            url: '/llave',
            data: llave
        })
         if(respuesta) {
            console.log("LLave creada con exito")
        }
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function encriptar({commit}, content) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})
        const respuesta = await Vue.axios({
            method: 'POST',
            url: '/llave/encriptar',
            data: content
        })
        if(respuesta) {
            console.log(respuesta)
            commit('setSalidaTextoEncriptado', respuesta.data)
        }
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function desencriptar({commit}, content) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})
        const respuesta = await Vue.axios({
            method: 'POST',
            url: '/llave/desencriptar',
            data: content
        })
        if(respuesta) {
            console.log(respuesta)
            commit('setSalidaTextoOrigininal', respuesta.data)
        }
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export function _limpiarFormularioEncriptar({commit}) {
    commit('limpiarFormularioEncriptar')
}

export function _limpiarFormularioDesencriptar({commit}) {
    commit('limpiarFormularioDesencriptar')
}

export function _setLlave({commit}, idLlave) {
    commit('setIdLlave', idLlave)
}