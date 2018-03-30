ELEMENT.locale(ELEMENT.lang.ja)
var app = new Vue({
  el: '#app',
  data: {
    tasks: [],
    newTask: "",
    loading: false,
  },
  created: function() {
    this.$data.loading = true;
    axios.get('/tasks')
      .then((response) => {
        console.log(response);
        app.$data.tasks = response.data.items;
        app.$data.loading = false;
      })
      .catch((error) => {
        console.log(error);
        app.$data.loading = false;
      });
  },
  methods: {
    addTask: (task) => {
      app.$data.loading = true;
      let params = new URLSearchParams();
      params.append('body', app.$data.newTask);
      axios.post('/tasks', params)
        .then((response) => {
          app.$data.loading = false;
          app.$data.tasks.push(response.data);
          app.$data.newTask = "";
          app.$data.loading = false;
        })
        .catch((error) => {
          console.log(error);
          app.$data.loading = false;
        });
    },
    doneTask: (task) => {
      app.$data.loading = true;
      let params = new URLSearchParams();
      params.append('done', !task.done);
      axios.put('/tasks/' + task.id, params)
        .then(function (response) {
          console.log(response);
          app.$data.loading = false;
        })
        .catch((error) => {
          console.log(error);
          app.$data.loading = false;
        });
    } 
  }
})
