export default {
    methods: {
        getIdsKeys: function (array_keys) {

            const length = array_keys.length;
            const firstId = array_keys[0].id;
            const lastId = array_keys[length - 1].id;

            return [firstId, lastId]
        }
    }
}