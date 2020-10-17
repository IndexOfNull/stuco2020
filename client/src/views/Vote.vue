<template>
  <div class="h-full flex flex-col justify-center items-center">
    <div class="sm:w-1/2 w-full">
        <Loader></Loader>
        <p class="text-lg">Please wait while we find your code</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import Loader from "@/components/Loader.vue"

export default {
    name: "Vote",
    components: {
        Loader
    },
    resolveCode() {
        return false //just for now
    },
    beforeMount () {
        axios
        .get('https://api.coindesk.com/v1/bpi/currentprice.json')
        .then(response => (console.log(response))) //Just testing axios
        setTimeout(() => {
          this.$store.commit("error", "Error! Here's proof that we're changing: " + Math.trunc(Math.random()*10) + ". Also heres the code passed back from /vote: " + this.$store.state.code)
          this.$router.push("/")
        }, 2000)
    }
}
</script>