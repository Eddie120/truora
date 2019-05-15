<template>
		<div>
			<b-form @submit.prevent="onSubmit">

			 <b-form-group
			        id="input-group-1"
			        label="Nombre Llave:"
			        label-for="input-1"
			        description="Como se va identificar la key pair generada por el sistema">

		     <b-form-input
		          id="input-1"
		          type="text"
		          v-model="llave.nombre"
		          v-validate="'required|min:4'" name="NombreLLave"
		          placeholder="Ingrese el nombre de la llave"
		        ></b-form-input>
		        <i v-show="errors.has('NombreLLave')" class="fa fa-warning"></i>
                <span v-show="errors.has('NombreLLave')" class="help is-danger">{{ errors.first('NombreLLave') }}</span>

      		</b-form-group>

      		<b-button type="submit" variant="primary">{{ texto }}</b-button>
			</b-form>
		</div>
</template>
<script>
 	import { mapActions, mapState } from 'vuex'
	export default {
		props: {
			texto: {
				type: String,
				required: true,
				default: "Crear"
			},
			llave: {
				type: Object,
				required: true
			},
			accion: {
				type: Boolean,
				required: true	
			}
		},
		computed: {
			...mapState(['error']),
		},
		methods: {
		...mapActions('KeyPairModule',['crearLlave']),	
	    async onSubmit() {
	      const result = await this.$validator.validateAll()
	      if(! result) {
	      	return
	      } else {
	      	 // Con esta prop sabemos si la accion es para crear o actualizar
            if(! this.accion) { // false
                await this.crearLlave(this.llave)
            } else {
                console.log("Vamos bien estamos actualizando el nombre de la llave")
            }

            if(! this.error.status) {
                this.$router.push('/llaves')
             } else {
             	console.log("Poseemos problemas marica")
             }
	      }
	    }
     }
	}
</script>
<style scope>
	
</style>