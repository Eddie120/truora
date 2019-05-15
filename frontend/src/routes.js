export default  [
    {
        path: '/',
        redirect: '/llaves'
    },
    {
        path: '/llaves',
        name: 'listar-llaves',
        component: () => import('./views/ListarLlaves')
    },
    {
        path: '/llaves/crear',
        name: 'crear-llave',
        component: () => import('./views/CrearLlave.vue'),
 
    },
]