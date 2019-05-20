<template>
  <div id="app">
    <loading :active.sync="loading"></loading>
  	<b-container>
      <b-card>
        <b-nav justified tabs>
          <b-nav-item exact to="/keys">Listado de llaves</b-nav-item>
          <b-nav-item exact to="/keys/create">Formulario para crear una llave</b-nav-item>
        </b-nav>


        <!-- Renderiza el contenido de las rutas -->
        <router-view/>

      </b-card>


    </b-container>
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import Loading from 'vue-loading-overlay'
  import 'vue-loading-overlay/dist/vue-loading.css'
export default {
  components: {
      Loading
    },
    computed: {
      ...mapState(['loading']),
    }
}
</script>
