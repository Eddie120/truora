<template>
    <b-container fluid>
        <div class="mt-5">
            <b-row>
                <b-col cols="2"></b-col>
                <b-col cols="8">
                    <FormularioLlave :texto="texto" :llave="llave" :accion="accion"></FormularioLlave>
                </b-col>
                <b-col cols="2"></b-col>
            </b-row> 
        </div>
    </b-container>
   </template>

<script>
    import FormularioLlave from '../components/FormularioLlave'
    export default {
        components: {
           FormularioLlave
        },
        data() {
            return {
                llave: {
                    nombre: null,
                    llaveprivada: null,
                    llavepublica: null
                },
                texto: "Crear llave",
                accion: false
            }
        },
        mounted() {
            console.log("aqui estamos")
        },
        methods: {
          
        },
        computed: {
          
        }
    }
</script>

<style scoped>

</style>