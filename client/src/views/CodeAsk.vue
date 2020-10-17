<template>
  <div class="h-full flex flex-col justify-center items-center">
    <div class="sm:w-1/2 w-full">
      <TitleCard
        title="👋 Hi"
        subtext="Enter your code below to begin voting"
      ></TitleCard>
      <div
        v-if="error != ''"
        class="p-2 rounded border-red-500 bg-red-400 text-red-900 border-2 mb-2"
      >
        {{ error }}
      </div>
      <div class="flex justify-center">
        <CodeInput class="w-full"></CodeInput>
      </div>
    </div>
  </div>
</template>

<script>
import CodeInput from "@/components/CodeInput.vue";
import TitleCard from "@/components/TitleCard.vue";

export default {
  name: "CodeAsk",
  components: {
    CodeInput,
    TitleCard
  },
  computed: {
    error() {
      return this.$store.state.redirectError;
    }
  }
};
</script>
