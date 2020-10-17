<template>
  <!--div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </div-->
  <div class="bg-dark text-white p-2 h-full">
    <p
      class="opacity-25 absolute m-2 bottom-0 right-0 w-full text-center sm:text-right"
    >
      Made with ❤️ by IoN
    </p>

    <router-view v-slot="{ Component }">
      <transition name="fade" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}

.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.15s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
