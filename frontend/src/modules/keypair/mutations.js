export function setLlaves (state, llaves) {
    state.llaves = llaves
}

export function setLlave (state, llave) {
    state.llave = llave
}

export function setSalidaTextoEncriptado (state, textoEncriptado) {
    state._encriptar.salidaTextoEncriptado = textoEncriptado
}

export function setSalidaTextoOrigininal (state, textoOriginal) {
    state._desencriptar.textoOriginal = textoOriginal
}

export function limpiarFormularioEncriptar(state) {
    state._encriptar.texto = ''
    state._encriptar.salidaTextoEncriptado = ''
}


export function limpiarFormularioDesencriptar(state) {
    state._desencriptar.texto = ''
    state._desencriptar.textoOriginal = ''
}

export function setIdLlave(state, id) {
    state._encriptar.id = id
    state._desencriptar.id = id
}
