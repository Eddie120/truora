<template>
    <b-container fluid>
        <div class="mt-5">
            <b-row>
                <b-col cols="2"></b-col>
                <b-col cols="8">
                    <FormCreateKey :text="text"
                                   :_key="key"
                                   :action="action">

                    </FormCreateKey>
                </b-col>
                <b-col cols="2"></b-col>
            </b-row> 
        </div>
    </b-container>
   </template>

<script>
    import FormCreateKey from '../components/FormCreateKey'
    export default {
        components: {
            FormCreateKey
        },
        data() {
            return {
                key: {
                    name: null,
                    privateKey: null,
                    publicKey: null
                },
                text: "Crear llave",
                action: false
            }
        },
    }
</script>

<style scoped>

</style>