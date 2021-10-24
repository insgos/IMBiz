<template>
  <div>
    <section class="todoapp" v-cloak>
      <header class="header">
        <h1>todos</h1>
        <input class="new-todo" autofocus autocomplete="off" placeholder="What needs to be done?">
      </header>
      <section class="main">
        <ul class="todo-list">
          <li class="todo">
            <div class="view">
              <label>A todo item</label>
            </div>
          </li>
        </ul>
      </section>
    </section>
  </div>
</template>

<script>
import "./assets/css/base.css"
import "./assets/css/app.css"
export default {
  name: "app",
  data() {
    return {
      todos: [],
    }
  }
};
</script>
