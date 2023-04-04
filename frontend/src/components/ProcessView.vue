<template>
    <div>
        <h3>Системная информацмя</h3>
        <div class="row m-2">
            <div class="card m-1" style="width: 25rem">
                <h4>Информация о хосте</h4>
                <button class="btn btn-primary" v-on:click="getHostInfo">Обновить</button>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item" v-for="(item, index) in hostInfo" :key="index">
                        <b>{{ index }}</b>: {{ item }}
                    </li>
                </ul>
            </div>

            <div class="card m-1" style="width: 25rem">
                <h4>Информация о системе</h4>
                <button class="btn btn-primary" v-on:click="getCpuInfo">Обновить</button>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item" v-for="(item, index) in cpuInfo" :key="index">
                        <b>{{ index }}</b>: {{ item }}
                    </li>
                </ul>
            </div>

            <div class="card m-1" style="width: 25rem">
                <h4>Информация о памяти</h4>
                <button class="btn btn-primary" v-on:click="getMemoryInfo">Обновить</button>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item" v-for="(item, index) in memoryInfo" :key="index">
                        <b>{{ index }}</b>: {{ item }}
                    </li>
                </ul>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <h4>Процессы</h4>
            </div>
            <div class="col">
                <button v-on:click="getProcessInfo" class="btn btn-primary">Обновить</button>
            </div>
        </div>
        <div class="processes">
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col">pid</th>
                        <th scope="col">memoryPercent</th>
                        <th scope="col">cpuPercent</th>
                        <th scope="col">cwd</th>
                        <th scope="col">username</th>
                        <th scope="col">name</th>
                        <th scope="col">status</th>
                        <th scope="col">info</th>
                    </tr>
                </thead>
                <tbody class="">
                    <tr v-for="(item, index) in processInfo" :key="index">
                        <td>{{ item['pid'] }}</td>
                        <td>{{ parseFloat(item['memoryPercent']).toFixed(4) }}</td>
                        <td>{{ parseFloat(item['cpuPercent']).toFixed(4) }}</td>
                        <td>{{ item['cwd'] }}</td>
                        <td>{{ item['username'] }}</td>
                        <td>{{ item['name'] }}</td>
                        <td>{{ item['status'] }}</td>

                        <td><router-link :to="{ name: 'proc', params: { pid: item['pid'] }}" class="btn btn-outline-success btn-sm"
                                href="#">O</router-link></td>
                    </tr>

                </tbody>
            </table>
        </div>
    </div>
</template>
      
<script>
export default {
    name: 'ProcessView',
    data() {
        return {
            hostInfo: null,
            cpuInfo: null,
            memoryInfo: null,
            processInfo: null,
            lastProcessUpdate: null
        };
    },
    methods: {
        getHostInfo() {
            this.axios.get('http://localhost:8081/host')
                .then(response => (this.hostInfo = response.data,
                    this.hostInfo['update'] = new Date().toLocaleTimeString()))
        },
        getCpuInfo() {
            this.axios.get('http://localhost:8081/cpu')
                .then(response => (this.cpuInfo = response.data,
                    this.cpuInfo['update'] = new Date().toLocaleTimeString()))
        },
        getMemoryInfo() {
            this.axios.get('http://localhost:8081/mem')
                .then(response => (this.memoryInfo = response.data,
                    this.memoryInfo['update'] = new Date().toLocaleTimeString()))
        },
        getProcessInfo() {
            this.axios.get('http://localhost:8081/proc')
                .then(response => (this.processInfo = response.data))
            this.lastProcessUpdate = new Date().toLocaleTimeString()

        }
    }
}
</script>
