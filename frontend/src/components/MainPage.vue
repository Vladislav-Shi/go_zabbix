<template>
<div class="container main">
      <nav class="navbar navbar-expand-lg navbar navbar-dark bg-dark">
        <div class="container-fluid">
          <a class="navbar-brand" href="#">Navbar</a>
          <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
              <li class="nav-item">
                <router-link to="/" class="nav-link active" aria-current="page" href="#"
                  >Мониторинг</router-link
                >
              </li>
              <li class="nav-item">
                <router-link to="/bash" class="nav-link" href="#">bash</router-link>
              </li>
              <li class="nav-item">
                <router-link to="/log" class="nav-link" href="#">Logs</router-link>
              </li>
            </ul>
          </div>
        </div>
      </nav>

      <router-view />

    </div>
</template>
  
  <script>
  export default {
    name: 'MainPage',
  }
  </script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
  <style scoped>
  h3 {
    margin: 40px 0 0;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
    margin: 0 10px;
  }
  a {
    color: #42b983;
  }
  </style>
  