export default  [
    {
        path: '/',
        redirect: '/keys'
    },
    {
        path: '/keys',
        name: 'keys',
        component: () => import('./views/ListKeys')
    },
    {
        path: '/keys/create',
        name: 'create-key',
        component: () => import('./views/CreateKey.vue'),
 
    },
]