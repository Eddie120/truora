<template>
    <b-container fluid>
        <div class="mt-5">
            <b-row>
                <b-col cols="1"></b-col>
                <b-col cols="10">
                    <b-form-group class="mb-4">

                        <b-input-group>
                            <b-form-input v-model="filter" v-on:change="buscar" placeholder="Buscar ..."></b-form-input>
                            <b-input-group-append>
                                <b-button :disabled="!filter" @click="filter = ''">Limpiar</b-button>
                            </b-input-group-append>
                        </b-input-group>

                    </b-form-group>
                    <ListarLlaves v-if="llaves" :llaves="llaves"></ListarLlaves>
                    <b-alert v-else variant="info" show>No hay llaves disponibles</b-alert>
                </b-col>
                <b-col cols="1"></b-col>
            </b-row> 
        </div>
    </b-container>
</template>

<script>
    import { mapActions, mapState } from 'vuex'
    import ListarLlaves from '../components/ListarLlaves'
    export default {
        components: {
            ListarLlaves
        },
        data() {
            return {
                filter: null,
            }
        },
        computed: {
            ...mapState('KeyPairModule',['llaves']),
        },
        methods: {
            ...mapActions('KeyPairModule',['cargarLlaves']),
            buscar() {
                if(this.filter.length >= 3) {
                    this.cargarLlaves(this.filter)
                } else if(this.filter.length < 3) {
                    this.cargarLlaves()
                }
            }
        },
        async mounted() {
           await this.cargarLlaves()
        },
    }
</script>

<style scoped>

</style>