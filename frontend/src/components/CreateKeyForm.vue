<template>
		<div>
			<b-form @submit.prevent="onSubmit">

			 <b-form-group
			        id="input-group-1"
			        label="Nombre:"
			        label-for="input-1"
			        description="Como se va identificar la key pair generada por el sistema">

		     <b-form-input
		          id="input-1"
		          type="text"
		          v-model="_key.name"
		          v-validate="'required|min:4'"
                  name="Nombre"
		          placeholder="Ingrese el nombre de la llave"
		        ></b-form-input>
		        <i v-show="errors.has('Nombre')" class="fa fa-warning"></i>
                <span v-show="errors.has('Nombre')" class="help is-danger text-danger">{{ errors.first('Nombre') }}</span>

      		</b-form-group>

      		<b-button type="submit" variant="primary">{{ text }}</b-button>
			</b-form>
		</div>
</template>
<script>
 	import { mapActions, mapState } from 'vuex'
	export default {
		props: {
			text: {
				type: String,
				required: true,
				default: "Crear"
			},
            _key: {
				type: Object,
				required: true
			},
			action: {
				type: Boolean,
				required: true	
			}
		},
		computed: {
			...mapState(['error']),
		},
		methods: {
		...mapActions('KeyPairModule',['create']),
	    async onSubmit() {
	      const result = await this.$validator.validateAll()
	      if(! result) {
	      	return
	      } else {

            if(! this.action) {
                await this.create(this._key)
            }

            if(! this.error.status) {
                this.$router.push('/keys')
             } else {
				console.log("Paso un error durante la peticion "+ this.error.message)
             }
	      }
	    }
     },
	}
</script>
<style scope>
	
</style>
