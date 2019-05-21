<template>
    <b-container fluid>
        <div class="mt-5">
            <b-row>
                <b-col cols="1"></b-col>
                <b-col cols="10">
                    <b-form-group class="mb-4">

                        <b-input-group>
                            <b-form-input v-model="params.filter" v-on:change="search" placeholder="Buscar ..."></b-form-input>
                            <b-input-group-append>
                                <b-button :disabled="!params.filter" @click="params.filter = ''">Limpiar</b-button>
                            </b-input-group-append>
                        </b-input-group>

                    </b-form-group>

                    <b-row>
                        <b-col class="ml-5">
                            <b-form-group label-cols-sm="10">
                                <b-form-select v-model="params.perPage" :options="params.pageOptions" v-on:change="search"></b-form-select>
                            </b-form-group>
                        </b-col>
                    </b-row>

                    <ListKeys v-if="keys" :keys="keys"></ListKeys>
                    <b-alert v-else variant="info" show>No hay llaves disponibles</b-alert>
                </b-col>
                <b-col cols="1"></b-col>
            </b-row> 
        </div>
    </b-container>
</template>

<script>
    import { mapActions, mapState } from 'vuex'
    import ListKeys from '../components/ListKeys'
    export default {
        components: {
            ListKeys
        },
        data() {
            return {
                params: {
                    filter: '',
                    perPage: 5,
                    pageOptions: [5, 10, 15]
                }
            }
        },
        computed: {
            ...mapState('KeyPairModule',['keys']),
        },
        methods: {
            ...mapActions('KeyPairModule',['loadKeys']),
            search() {

                if(this.params.filter.length >= 3) {
                    this.loadKeys(this.params)
                } else if(this.params.filter.length == 0) {
                    this.loadKeys(this.params)
                }
            }
        },
        async mounted() {
            await this.loadKeys(this.params)
        },
    }
</script>

<style scoped>

</style>