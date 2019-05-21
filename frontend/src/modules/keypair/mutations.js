export function setKeys (state, keys) {
    state.keys = keys
}

export function setKey (state, key) {
    state.key = key
}

export function setEncryptedText (state, encryptedText) {
    state._encrypt.encryptedText = encryptedText
}

export function setOriginalText (state, originalText) {
    state._decrypt.originalText = originalText
}

export function resetFormEncrypt(state) {
    state._encrypt.text = ''
    state._encrypt.encryptedText = ''
}


export function resetFormDecrypt(state) {
    state._decrypt.text = ''
    state._decrypt.originalText = ''
}

export function setIdKey(state, id) {
    state._encrypt.id = id
    state._decrypt.id = id
}

export function setFirstId(state, id) {
    state.firstId = id
}

export function setLastId(state, id) {
    state.lastId = id
}
