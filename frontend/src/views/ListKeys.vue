<template>
    <b-container fluid>
        <div class="mt-5">
            <b-row>
                <b-col cols="1"></b-col>
                <b-col cols="10">
                    <b-form-group class="mb-4">

                        <b-input-group>
                            <b-form-input v-model="filter" v-on:change="search" placeholder="Buscar ..."></b-form-input>
                            <b-input-group-append>
                                <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
                            </b-input-group-append>
                        </b-input-group>

                    </b-form-group>
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
                filter: null,
            }
        },
        computed: {
            ...mapState('KeyPairModule',['keys']),
        },
        methods: {
            ...mapActions('KeyPairModule',['loadKeys']),
            search() {
                if(this.filter.length >= 3) {
                    this.loadKeys(this.filter)
                } else if(this.filter.length < 3) {
                    this.loadKeys()
                }
            }
        },
        async mounted() {
           await this.loadKeys()
        },
    }
</script>

<style scoped>

</style>