<template>
    <h3>Информация о процессе {{ PID }}</h3>
    <div v-if="isLife">
        <div class="row justify-content-md-center">
            <div class="card m-1 justify-content-start" style="width: 35rem">
                <ul class="list-group list-group-flush justify-content-start">
                    <li class="list-group-item" v-for="(item, index) in info" :key="index">
                        <b>{{ index }}</b>: {{ item }}
                    </li>
                </ul>
            </div>
        </div>
        <div class="row justify-content-md-center mt-4">
            <div>
                <h3>Дейстия с процессом</h3>
            </div>

            <div class="col-sm-3 my-1" v-if="isActive">
                <button class="btn btn-warning" v-on:click="stopProcess">Остановить</button>
            </div>
            <div class="col-sm-3 my-1" v-else>
                <button class="btn btn-success" v-on:click="resumeProcess">Возобновить</button>
            </div>
            <div class="col-sm-3 my-1">
                <button class="btn btn-danger" v-on:click="killProcess">Убить</button>
            </div>
            <div class="col-sm-3 my-1">
                <div class="input-group mb-2">
                    <div class="input-group-prepend">
                        <button class="btn btn-success input-group-text">Отправить сигнал</button>
                    </div>
                    <input type="text" class="form-control" id="inlineFormInputGroup" placeholder="Username">
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        <div class="alert alert-danger" role="alert">
            Процесс убит! {{ info['name'] }}
        </div>
    </div>
</template>

<script>
export default {
    name: "InfoProcess",
    data() {
        return {
            PID: this.$route.params.pid,
            info: null,
            isLife: true
        }
    },
    computed: {
        isActive() { return ['idle', 'sleep', 'running'].includes(this.info['status']) }
    },
    methods: {
        getInfo() {
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid)
                .then(response => (this.info = response.data,
                    this.info['createTime'] = new Date(Number(this.info['createTime']) * 1000).toLocaleTimeString()))
        },
        stopProcess() {
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid + "/stop")
                .then()
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid)
                .then(response => (this.info = response.data,
                    this.info['createTime'] = new Date(Number(this.info['createTime']) * 1000).toLocaleTimeString()))
        },
        resumeProcess() {
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid + "/run")
                .then()
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid)
                .then(response => (this.info = response.data,
                    this.info['createTime'] = new Date(Number(this.info['createTime']) * 1000).toLocaleTimeString()))
        },
        killProcess() {
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid + "/kill")
                .then()
            this.axios.get('http://localhost:8081/proc/' + this.$route.params.pid)
                .then(response => (this.info = response.data,
                    this.info['createTime'] = new Date(Number(this.info['createTime']) * 1000).toLocaleTimeString()))
            this.isLife = false
        }
    },
    created() {
        this.getInfo()
    }

}
</script>

<style>
body {
    background-color: whitesmoke;
}

.main {
    background-color: white;
    min-height: 100vh;
    height: max-content;
}

.process-list {
    height: 200px;
    overflow: auto;
}

.processes {
    height: 70vh;
    overflow: auto;
}

.list-group-item {
    text-align: left;
}
</style>