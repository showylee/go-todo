<template>
  <div>
    <input type="text" v-model="item">
    <button v-on:click="addItem">add</button>
    <div>
      <p v-for="(item, index) in items" v-bind:key="index" >
        {{ item.item }}
        <button v-on:click="deleteItem">delete</button>
        <button>update</button>
      </p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Todo',
  props: {
    msg: String
  },
  data () {
    return {
      items: null,
      item: ""
    }
  },
  mounted () {
    axios
      .get('http://localhost:8888/api/v1/todo')
      .then(response => {
        this.items = response.data
        alert(JSON.stringify(response))
      }).catch(error => alert(error))
  },
  methods: {
    addItem: function () {
      axios
        .post('http://localhost:8888/api/v1/todo', {
          item: this.item
        })
        .then(response => {
          this.items = response
          alert(response)
        }).catch(error => alert(error))
    },
    deleteItem: function () {
      axios.delete('http://localhost:8888')
    }

  }
}
</script>

<style>

</style>

