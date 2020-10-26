<template>
  <div class="h-full flex flex-col justify-center items-center">
    <div class="sm:w-1/2 w-full">
      <TitleCard
        title="✅ Thanks!"
        subtext="Your vote has been recorded"
      ></TitleCard>
      <div class="flex justify-center flex-col" v-if="code.times_used != undefined && code.times_used < code.max_uses">
        <p class="mb-2">It looks like your code has more uses on it. Click the button below to submit another ballot.</p>
        <router-link :to="{ name: 'Vote', params: { code: code.code }}"><button class=" bg-primary rounded p-2 sm:px-4 sm:w-auto w-full hover:bg-secondary transition-colors duration-100 disabled:opacity-25 disabled:no">Vote Again</button></router-link>
      </div>
    </div>
  </div>
</template>

<script>
import TitleCard from "@/components/TitleCard.vue";

export default {
  name: "Thanks",
  components: {
    TitleCard
  },
  computed: {
    code: function() {
      return this.$store.state.code;
    }
  }
};
</script>
