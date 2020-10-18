<template>
  <div class="text-black">
    <div>
      <input
        class="p-2 mb-2 h-16 rounded w-full text-2xl text-center font-mono border focus:shadow-outline focus:outline-none"
        type="text"
        v-bind:placeholder="placeholder"
        v-bind:value="code"
        @input="updateCode"
      />

      <div>
        <button
          v-on:click="redirect"
          class=" bg-primary rounded p-2 sm:px-4 sm:w-auto w-full hover:bg-secondary transition-colors duration-100 disabled:opacity-25 disabled:no"
          :disabled="code == ''"
        >
          Get Started
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CodeInput",
  computed: {
    code() {
      return this.$store.state.code;
    }
  },
  data() {
    return {
      error: ""
    };
  },
  props: {
    enableScanner: Boolean,
    placeholder: {
      type: String,
      default: "Code"
    }
  },
  methods: {
    redirect: function() {
      if (this.code != "") {
        this.$router.push("/vote/" + this.code);
      }
    },
    updateCode: function(e) {
      this.$store.commit("code", e.target.value);
    }
  }
};
</script>
