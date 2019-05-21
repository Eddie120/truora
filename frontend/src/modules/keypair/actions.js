import Vue from 'vue'

export async function loadKeys({commit}, params) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})

       let url = '/keys?perPage=' + params.perPage
        if(params.filter != '') {
            url += '&text='+params.filter
        }

        const {data} = await Vue.axios({
            url: url
        })

        commit('setKeys', data)
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}


export async function create({commit}, key) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})
        const response = await Vue.axios({
            method: 'POST',
            url: '/key',
            data: key
        })
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function encrypt({commit}, content) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})
        const respuesta = await Vue.axios({
            method: 'POST',
            url: '/key/encrypt',
            data: content
        })
        if(respuesta) {
            commit('setEncryptedText', respuesta.data)
        }
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export async function decrypt({commit}, content) {
    try {
        commit('setError', false, {root: true})
        commit('setLoading', true, {root: true})
        const respuesta = await Vue.axios({
            method: 'POST',
            url: '/key/decrypt',
            data: content
        })
        if(respuesta) {
            commit('setOriginalText', respuesta.data)
        }
    } catch (error) {
        commit('setError', {modulo: 'Key Pair', error}, {root: true})
    }finally{
        commit('setLoading',false, { root: true })
    }
}

export function _resetFormEncrypt({commit}) {
    commit('resetFormEncrypt')
}

export function _resetFormDecrypt({commit}) {
    commit('resetFormDecrypt')
}

export function _setKey({commit}, idkey) {
    commit('setIdKey', idkey)
}

export function _setFirstId({commit}, id) {
    commit('setFirstId', id)
}

export function _setLastId({commit}, id) {
    commit('setLastId', id)
}