ELEMENT.locale(ELEMENT.lang.ja)
var app = new Vue({
  el: '#app',
  ready: {
  },
  data: {
    tasks: [],
    newTask: "",
    loading: false,
  },
  created: function() {
    this.loading = true;
    axios.get('/tasks')
      .then(function (response) {
        console.log(response);
        app.$data.tasks = response.data.items;
        app.$data.loading = false;
      })
      .catch(function (error) {
        console.log(error);
        app.$data.loading = false;
      });
  },
  methods: {
    addTask: function(task) {
      app.$data.loading = true;
      let params = new URLSearchParams();
      params.append('body', app.$data.newTask);
      axios.post('/tasks', params)
        .then(function (response) {
          app.$data.tasks.loading = false;
          app.$data.tasks.push(response.data);
          app.$data.newTask = "";
          app.$data.loading = false;
        })
        .catch(function (error) {
          console.log(error);
          app.$data.loading = false;
        });
    },
    doneTask: function(task) {
      app.$data.loading = true;
      let params = new URLSearchParams();
      params.append('done', !task.done);
      axios.put('/tasks/' + task.id, params)
        .then(function (response) {
          console.log(response);
          app.$data.loading = false;
        })
        .catch(function (error) {
          console.log(error);
          app.$data.loading = false;
        });
    } 
  }
})
