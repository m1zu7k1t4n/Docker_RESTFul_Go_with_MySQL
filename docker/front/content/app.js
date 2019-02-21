const vm = new Vue({
  el: '#app',
  data: {
    results: []
  },
  mounted() {
    axios.get("http://golang_host:8080/api/v1/todos/")
      .then(response => { this.results = response.data.data })
  }
});
