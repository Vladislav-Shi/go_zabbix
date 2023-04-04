<template>
    <div>
        <h3>Выполнение Bash</h3>
        <div>
            <br>
            <div class="row justify-content-md-center">
                <div class="col-sm-3 my-1">
                    <input type="text" class="form-control" placeholder="/home/vladislav" v-model="bashPath">
                    <small id="passwordHelpBlock" class="form-text text-muted">
                        путь выполнения команды
                    </small>
                    <button type="button" class="btn btn-primary btn-sm mt-2"
                        v-on:click="setDefaultPath">по-умолчанию</button>
                </div>
                <div class="col-sm-3 my-1">
                    <input type="text" class="form-control" placeholder="Команда" v-model="bashCommand">
                    <small id="passwordHelpBlock" class="form-text text-muted">
                        команда выполнения
                    </small>
                </div>
                <div class="col-auto my-1">
                    <button class="btn btn-primary" v-on:click="runBash">Отправить</button>
                </div>
            </div>
        </div>
        <div>
            <h5>Вывод</h5>
            <div>
                <table class="table">
                    <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col"></th>
                        </tr>
                    </thead>
                    <tbody class="">
                        <tr v-for="(item, index) in output" :key="index">
                            <td>{{ index }}</td>
                            <td>{{ item }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>


        </div>

    </div>
</template>
      
<script>
export default {
    name: 'BashView',
    data() {
        return {
            bashPath: "/home/vladislav",
            bashCommand: "",
            output: null
        }
    },
    methods: {
        setDefaultPath() {
            this.bashPath = "/home/vladislav"
        },
        runBash() {
            console.log('bashCommand', this.bashCommand)
            this.axios.post('http://localhost:8081/run', {
                "path": this.bashPath, "command": this.bashCommand
            })
                .then(response => (this.output = response.data.split('\n')))
        },
    }
}
</script>
