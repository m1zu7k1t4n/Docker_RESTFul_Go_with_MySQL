const vm = new Vue({
  el: '#app',
  data: {
    results: []
  },
  mounted() {
    axios.get("http://localhost:8080/api/v1/todos/")
      .then(response => { this.results = response.data.data })
  }
});
