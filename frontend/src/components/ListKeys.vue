<template>
    <div v-if="keys.length">

    <b-card>
        <b-table
                striped hover small
                :items="keys"
                :fields="fields">

            <template slot="actions" slot-scope="row">
                <b-button size="sm" @click="view(row.item, row.index, $event.target)" class="mr-1">
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
                 ok-only @hide="resetModal(5)"
                 @show="resetModal(5)">

            <b-row>
                <b-col cols="12">

                    <form ref="formEncrypt" @submit.stop.prevent="handleSubmitEncrypt" @reset="resetModal(1)">
                        <b-form-group
                                label="Text"
                                label-for="name-input"
                                invalid-feedback="Text required">
                            <b-form-input
                                    id="name-input"
                                    v-model="_encrypt.text"
                                    :state="nameStateformEncrypt"
                                    required
                                    :aucomplete="false">

                            </b-form-input>
                        </b-form-group>

                        <b-form-textarea
                                id="textarea-no-resize"
                                v-model="_encrypt.encryptedText"
                                rows="10"
                                no-resize
                                :readonly="true"
                                :aucomplete="false">

                        </b-form-textarea>

                        <b-row>
                            <b-col cols="4"></b-col>
                            <b-col cols="4">
                                <b-button  type="submit" class="mt-2 mr-4" variant="success">Encrypt</b-button>
                                <b-button  type="reset" class="mt-2" variant="secondary">Reset</b-button>
                            </b-col>
                            <b-col cols="4"></b-col>
                        </b-row>

                    </form>

                </b-col>
            </b-row>

            <b-row class="mt-5">
                <b-col cols="12">

                    <form ref="formDecrypt" @submit.stop.prevent="handleSubmitDecrypt" @reset="resetModal(2)">

                        <b-form-textarea
                                v-model="_decrypt.text"
                                rows="10"
                                no-resize
                                :state="nameStateformDecrypt"
                                :aucomplete="false"
                                required>

                        </b-form-textarea>

                        <b-form-group
                                label="Original Text"
                                label-for="name-input"
                                invalid-feedback="Original Text">
                            <b-form-input
                                    v-model="_decrypt.originalText"
                                    :readonly="true"
                                    :aucomplete="false">

                            </b-form-input>
                        </b-form-group>


                        <b-row>
                            <b-col cols="4"></b-col>
                            <b-col cols="4">
                                <b-button  type="submit"  class="mt-2 mr-4" variant="danger">Decrypt</b-button>
                                <b-button  type="reset" class="mt-2" variant="secondary">Reset</b-button>
                            </b-col>
                            <b-col cols="4"></b-col>
                        </b-row>


                    </form>

                </b-col>
            </b-row>

        </b-modal>

    </div>
    <b-alert v-else variant="info" show>No keys available</b-alert>
</template>

<script>
    import { mapActions, mapState } from 'vuex'
    export default {
        components: {
        },
        props: {
            keys: {
                type: Array,
                required: true
            }
        },
        data() {
            return {
                fields: [
                  { key: 'id', label: 'Id', sortable: true ,class: 'text-center'},
                  { key: 'name', label: 'Nombre', sortable: true, class: 'text-center'},
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
            ...mapState('KeyPairModule',['_encrypt', '_decrypt']),
        },
        methods: {
            ...mapActions('KeyPairModule',['encrypt', 'decrypt', '_resetFormEncrypt', '_resetFormDecrypt', '_setKey']),
            view(item, index, button) {

                this.resetModal(5)
                this.infoModal.title = item.name

                this._setKey(item.id)
                this.$root.$emit('bv::show::modal', this.infoModal.id, button)
            },
            resetModal(type) {

                switch (type) {
                    case 5:

                        this.$refs.formEncrypt.reset()
                        this.$refs.formDecrypt.reset()

                        this._resetFormEncrypt()
                        this._resetFormDecrypt()

                        this.nameStateformEncrypt = null
                        this.nameStateformDecrypt = null

                        break;
                    case 1:

                        this.$refs.formEncrypt.reset()
                        this._resetFormEncrypt()

                        this.nameStateformEncrypt = null

                        break;
                    case 2:

                        this.$refs.formDecrypt.reset()
                        this._resetFormDecrypt()

                        this.nameStateformDecrypt = null

                        break;
                }

                this.nameStateformEncrypt = null
                this.nameStateformDecrypt = null
            },
           async handleSubmitEncrypt() {
                if (!this.checkFormValidity(1)) {
                    return
                }

               await this.encrypt(this._encrypt)
            },
            checkFormValidity(type) {

                let valid = null
                if(type == 1) {

                    valid = this.$refs.formEncrypt.checkValidity()
                    this.nameStateformEncrypt = valid ? 'valid' : 'invalid'

                } else if( type == 2) {

                    valid = this.$refs.formDecrypt.checkValidity()
                    this.nameStateformDecrypt = valid ? 'valid' : 'invalid'

                }

                return valid
            },
          async  handleSubmitDecrypt() {
                if (!this.checkFormValidity(2)) {
                    return
                }
               await this.decrypt(this._decrypt)
            },
        }
    }
</script>

<style scoped>

</style>