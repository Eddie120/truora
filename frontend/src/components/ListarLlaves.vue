<template>
    <div v-if="llaves.length">

    <b-card>



        <b-table
                striped hover small
                :items="llaves"
                :fields="fields">

            <template slot="actions" slot-scope="row">
                <b-button size="sm" @click="ver(row.item, row.index, $event.target)" class="mr-1">
                    Ver
                </b-button>
            </template>
        </b-table>
    </b-card>


        <!-- modal -->
        <b-modal :id="infoModal.id"
                 ref="modal"
                 size="lg"
                 :title="infoModal.title"
                 ok-only @hide="limipiarModal(5)"
                 @show="limipiarModal(5)">

            <b-row>
                <b-col cols="12">

                    <form ref="formEncrypt" @submit.stop.prevent="handleSubmitEncrypt" @reset="limipiarModal(1)">
                        <b-form-group
                                label="Texto"
                                label-for="name-input"
                                invalid-feedback="El texto es requerido">
                            <b-form-input
                                    id="name-input"
                                    v-model="_encriptar.texto"
                                    :state="nameStateformEncrypt"
                                    required
                                    :aucomplete="false">

                            </b-form-input>
                        </b-form-group>

                        <b-form-textarea
                                id="textarea-no-resize"
                                v-model="_encriptar.salidaTextoEncriptado"
                                rows="10"
                                no-resize
                                :readonly="true"
                                :aucomplete="false">

                        </b-form-textarea>

                        <b-row>
                            <b-col cols="4"></b-col>
                            <b-col cols="4">
                                <b-button  type="submit" class="mt-2 mr-4" variant="success">Encriptar</b-button>
                                <b-button  type="reset" class="mt-2" variant="secondary">Limpiar</b-button>
                            </b-col>
                            <b-col cols="4"></b-col>
                        </b-row>

                    </form>

                </b-col>
            </b-row>

            <b-row class="mt-5">
                <b-col cols="12">

                    <form ref="formDecrypt" @submit.stop.prevent="handleSubmitDecrypt" @reset="limipiarModal(2)">

                        <b-form-textarea
                                v-model="_desencriptar.texto"
                                rows="10"
                                no-resize
                                :state="nameStateformDecrypt"
                                :aucomplete="false"
                                required>

                        </b-form-textarea>

                        <b-form-group
                                label="Texto original"
                                label-for="name-input"
                                invalid-feedback="Texto original">
                            <b-form-input
                                    v-model="_desencriptar.textoOriginal"
                                    :readonly="true"
                                    :aucomplete="false">

                            </b-form-input>
                        </b-form-group>


                        <b-row>
                            <b-col cols="4"></b-col>
                            <b-col cols="4">
                                <b-button  type="submit"  class="mt-2 mr-4" variant="danger">Desencriptar</b-button>
                                <b-button  type="reset" class="mt-2" variant="secondary">Limpiar</b-button>
                            </b-col>
                            <b-col cols="4"></b-col>
                        </b-row>


                    </form>

                </b-col>
            </b-row>

        </b-modal>

    </div>
    <b-alert v-else variant="info" show>No hay llaves disponibles</b-alert>
</template>

<script>
    import { mapActions, mapState } from 'vuex'
    export default {
        components: {
        },
        props: {
            llaves: {
                type: Array,
                required: true
            }
        },
        data() {
            return {
                fields: [
                  { key: 'id', label: 'Id', sortable: true ,class: 'text-center'},
                  { key: 'nombre', label: 'Nombre de la llave', sortable: true, class: 'text-center'},
                  { key: 'actions', label: 'Opción', sortable: false ,class: 'text-center' }
                ],
                infoModal: {
                    id: 'info-modal',
                    title: '',
                },
                nameStateformEncrypt: null,
                nameStateformDecrypt: null,

            }
        },
        computed: {
            ...mapState('KeyPairModule',['_encriptar', '_desencriptar']),
        },
        methods: {
            ...mapActions('KeyPairModule',['encriptar', 'desencriptar', '_limpiarFormularioEncriptar', '_limpiarFormularioDesencriptar', '_setLlave', 'cargarLlaves']),
            ver(item, index, button) {
                this.limipiarModal(5)
                this.infoModal.title = item.nombre

                this._setLlave(item.id)
                this.$root.$emit('bv::show::modal', this.infoModal.id, button)
            },
            limipiarModal(tipoLimpieza) {

                switch (tipoLimpieza) {
                    case 5: // limpiar todo el formulario
                        this.$refs.formEncrypt.reset()
                        this.$refs.formDecrypt.reset()

                        this._limpiarFormularioEncriptar()
                        this._limpiarFormularioDesencriptar()
                        this.nameStateformEncrypt = null
                        this.nameStateformDecrypt = null
                        break;
                    case 1: // limpiar formulario para encriptar
                        this.$refs.formEncrypt.reset()
                        this._limpiarFormularioEncriptar()
                        this.nameStateformEncrypt = null
                        break;
                    case 2: // limpiar formulario pasa desencriptar
                        this.$refs.formDecrypt.reset()
                        this._limpiarFormularioDesencriptar()
                        this.nameStateformDecrypt = null
                        break;
                }

                this.nameStateformEncrypt = null
                this.nameStateformDecrypt = null
            },
            handleSubmitEncrypt() {
                if (!this.checkFormValidity(1)) {
                    return
                }

                this.encriptar(this._encriptar).then((respuesta) => {
                   console.log("Encriptacion finalizada")
                })
            },
            checkFormValidity(tipo) {
                let valid = null
                if(tipo == 1) {
                    valid = this.$refs.formEncrypt.checkValidity()
                    this.nameStateformEncrypt = valid ? 'valid' : 'invalid'
                } else if( tipo == 2) {
                    valid = this.$refs.formDecrypt.checkValidity()
                    this.nameStateformDecrypt = valid ? 'valid' : 'invalid'
                }

                return valid
            },
            handleSubmitDecrypt() {
                if (!this.checkFormValidity(2)) {
                    return
                }
                this.desencriptar(this._desencriptar).then((respuesta) => {
                    console.log("Finalizo el proceso para desencriptar la cadena cifrada")
                })
            },
        }
    }
</script>

<style scoped>

</style>