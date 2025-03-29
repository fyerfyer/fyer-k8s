<template>
  <div class="app-container">
    <app-header />
    <div class="main-container">
      <sidebar class="sidebar" />
      <div class="content-container">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import AppHeader from '@/components/layout/Header.vue'
import Sidebar from '@/components/layout/Sidebar.vue'

export default {
  name: 'App',
  components: {
    AppHeader,
    Sidebar
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  min-height: 100vh;
}

.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-container {
  display: flex;
  margin-top: 60px;
  /* Header height */
  min-height: calc(100vh - 60px);
}

.sidebar {
  min-height: calc(100vh - 60px);
  background-color: #304156;
  transition: width 0.3s;
  overflow: hidden;
}

.content-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  background-color: #f0f2f5;
}
</style>