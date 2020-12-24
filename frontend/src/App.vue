<template>
  <div id="app">
    <Narbar />
    <b-container>
      <router-view />
    </b-container>
  </div>
</template>

<script>
import Narbar from './views/layout/Navbar.vue';

export default {
  components: { Narbar },
  data() {
    return {};
  },
};
</script>

<style lang="scss">
</style>
