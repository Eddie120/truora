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
    import {mapActions, mapState} from 'vuex'
    import Mixin from '../mixins/mixin'

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
        mixins: [Mixin],
        computed: {
            ...mapState('KeyPairModule', ['firstId', 'lastId']),
        },
        methods: {
            ...mapActions('KeyPairModule', ['loadKeys', '_setFirstId', '_setLastId']),
            async previous() {

                const params = this.filterParams("previous")
                await this.loadKeys(params)

                const array = this.getIdsKeys(this.keys)

                this._setFirstId(array[0])
                this._setLastId(array[1])
            },
            async next() {

                const params = this.filterParams("next")
                await this.loadKeys(params)

                const array = this.getIdsKeys(this.keys)

                this._setFirstId(array[0])
                this._setLastId(array[1])
            },
            filterParams(type) {

                let params = {}
                params.filter = this.params.filter
                params.perPage = this.params.perPage

                if(type == "previous") {
                    params.firstId = this.firstId
                } else {
                    params.lastId = this.lastId
                }

                return params
            }
        }
    }
</script>

<style scoped>

</style>