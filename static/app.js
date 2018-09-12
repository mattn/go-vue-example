ELEMENT.locale(ELEMENT.lang.ja)
var app = new Vue({
  el: '#app',
  data: {
    tasks: [],
    newTask: "",
    loading: false,
  },
  created: function() {
    this.loading = true
    axios.get('/tasks')
      .then((response) => {
        console.log(response)
        this.tasks = response.data.items || []
        this.loading = false
      })
      .catch((error) => {
        console.log(error)
        this.loading = false
      })
  },
  methods: {
    addTask: function(task) {
      this.loading = true
      let params = new URLSearchParams()
      params.append('body', this.newTask)
      params.append('done', false)
      axios.post('/tasks', params)
        .then((response) => {
          this.loading = false
          this.tasks.unshift(response.data)
          this.newTask = ""
          vue.$forceUpdate()
          this.loading = false
        })
        .catch((error) => {
          console.log(error)
          this.loading = false
        })
    },
    doneTask: function(task) {
      this.loading = true
      let params = new URLSearchParams()
      params.append('done', !task.done)
      axios.put('/tasks/' + task.id, params)
        .then((response) => {
          console.log(response)
          this.loading = false
        })
        .catch((error) => {
          console.log(error)
          this.loading = false
        })
    },
    removeTask: function(task) {
      this.loading = true
      axios.delete('/tasks/' + task.id)
        .then((response) => {
          console.log(response)
          Vue.delete(this.tasks, task)
          this.tasks.forEach(function(v, i, a) { if (v.id == task.id) a.splice(i,1) })
          vue.$forceUpdate()
          this.loading = false
        })
        .catch((error) => {
          console.log(error)
          this.loading = false
        })
    }
  }
})
