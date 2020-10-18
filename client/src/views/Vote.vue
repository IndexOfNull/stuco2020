<template>
  <div class="h-full flex flex-col justify-center items-center">
    <transition name="fade" mode="out-in">
      <div class="sm:w-1/2 w-full" v-if="code.code == undefined">
        <Loader></Loader>
        <p class="text-lg">Please wait while we find your code</p>
      </div>
      <div class="sm:w-1/2 w-full" v-else>
        <header class="mb-4">
          <h1 class="text-5xl">Hello, {{ code.student.firstname }}</h1>
          <p>Your ballot is below</p>
        </header>
        <form>
          <div v-for="c in classes" :key="c.id">
            <div class="p-4 m-2 w bg-lightbg rounded-lg text-black text-left">
              <header class="mb-2">
                <h1 class="sm:text-4xl text-3xl">{{ c.name }}</h1>
                <p class="italic">(Pick {{ c.vote_count }})</p>
              </header>
              <MultiCheckbox :options="c.students" :max=c.vote_count />
            </div>
          </div>
        </form>
      </div>
    </transition>
  </div>
</template>

<script>
import axios from "axios";
import Loader from "@/components/Loader.vue";
import MultiCheckbox from "@/components/MultiCheckbox.vue";

export default {
  name: "Vote",
  data () {
    return {
    code: Object,
    response: Object,
    classes: Object
    }
  },
  components: {
    Loader,
    MultiCheckbox
  },
  beforeMount() {
    this.$store.commit("code", this.$route.params.code) //Push the code to the store in case of direct navigation

    let s = this.$store.state
    axios.get(s.globalSettings.backendUrl + "code/" + s.code + "").then((response) => {

      //Validation of code, throw error if necessary
      if (response.data.code == null) {
        throw 'The entered code does not exist'
      } else if (response.data.code.times_used >= response.data.code.max_uses) {
        throw 'The entered code has no more uses left'
      } else if (response.data.code.active == false) {
        throw 'The entered code is currently deactivated'
      }

      //If the code is valid
      let data = response.data

      for (const key in response.data.classes) {
        data.classes[key].students = []
      }

      console.log(response.data)
      data.candidates.forEach(student => {
        if (student.class_id != undefined) {
          data.classes[student.class_id].students.push(student)
        }
      })
      this.classes = data.classes
      this.response = response.data
      setTimeout(() => {
        this.code = response.data.code
      }, 500) //Our UI is too fast otherwise lol
    }).catch((e) => {
      setTimeout(() => {
        console.log(e.type)
        if (e.message == "Network Error") {
          this.$store.commit("error", "Something went wrong while contacting the voting service")
        } else {
          this.$store.commit("error", e)
        }
        this.$router.push("/")
      }, 500) //Our UI is too fast otherwise lol
    })
  }
};
</script>
