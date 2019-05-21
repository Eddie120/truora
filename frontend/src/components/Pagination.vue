<template>

    <nav aria-label="Page navigation example">
        <ul class="pagination justify-content-end">
            <li class="page-item">
                <a class="page-link" href="javascript:void(0);" v-on:click="previous" tabindex="-1">Anterior</a>
            </li>
            <li class="page-item">
                <a class="page-link" href="javascript:void(0);" v-on:click="next">Siguiente</a>
            </li>
        </ul>
    </nav>

</template>

<script>
    import { mapActions, mapState } from 'vuex'
    export default {
        props: {
            keys: {
                type: Array,
                required: true
            },
            params: {
                type: Object,
                required: true
            }
        },
        data() {
            return {

            }
        },
        computed: {
            ...mapState('KeyPairModule',['firstId', 'lastId']),
        },
        methods: {
            ...mapActions('KeyPairModule',['_setFirstId', '_setLastId']),
            previous() {
                console.log("antes")
            },
            next() {
                console.log("despues")
            }
        },
        mounted() {
            const length = this.keys.length

            const firstId = this.keys[0].id
            const lastId = this.keys[  length - 1  ].id

            this._setFirstId(firstId)
            this._setLastId(lastId)
        }
    }
</script>

<style scoped>

</style>