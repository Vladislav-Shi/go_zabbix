<template>
    <div>
        <h2>Логи</h2>
        <div>
            <div class="row">
                <div class="col">
                    <h4>История bash комманд</h4>
                </div>
                <div class="col">
                    <button v-on:click="getBashHistory" class="btn btn-primary">Обновить</button>
                </div>
            </div>
            <div>
                <table class="table">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">path</th>
                        <th scope="col">command</th>
                        <th scope="col">output</th>
                        <th scope="col">status</th>
                        <th scope="col">date</th>
                    </tr>
                </thead>
                <tbody class="">
                    <tr v-for="(item, index) in bashHistory" :key="index">
                        <th scope="row">{{ item['id'] }}</th>
                        <td>{{ item['path'] }}</td>
                        <td>{{ item['command'] }}</td>
                        <td>{{ item['output'] }}</td>
                        <td>{{ item['status'] }}</td>
                        <td>{{ item['created'] }}</td>
                    </tr>

                </tbody>
            </table>
            </div>
        </div>
    </div>
</template>
      
<script>
export default {
    name: 'LogView',
    data() {
        return {
            bashHistory: null,
        };
    },
    methods: {
        getBashHistory() {
            this.axios.get('http://localhost:8081/log/bash')
                .then(response => (this.bashHistory = response.data))
        },
    }
}
</script>
      