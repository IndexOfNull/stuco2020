<template>
  <div class="h-full flex flex-col items-center py-2" ref="top">
    <transition name="fade" mode="out-in">
      <div
        class="h-full flex flex-col justify-center"
        v-if="showBallot == false"
      >
        <div class="w-full">
          <Loader></Loader>
          <p class="text-lg">Please wait while we find your code</p>
        </div>
      </div>
      <div class="sm:w-2/3 w-full" v-else>
        <header class="mb-4">
          <h1 class="text-5xl">Hello, {{ student.firstname }}</h1>
          <p>🗳️ Your ballot is below 🗳️</p>
          <p class="italic opacity-50">(Remember: Your vote for a position will not count unless it says "This one's good to go!")</p>
        </header>
        <div class="p-2 rounded border-red-500 bg-red-400 text-red-900 border-2 mb-2" v-if="errors.length > 0">
          <li v-for="(error, index) in errors" :key="`error-${index}`" >{{ error }}</li>
        </div>
        <form class="p-2">
          <div v-for="c in classes" :key="c.id">
            <div 
              class="p-4 mb-2 w bg-white rounded-lg text-black text-left transition-colors duration-100"
              v-if="c.students.length > 0"  
            >
              <header class="mb-2">
                <h1 class="sm:text-4xl text-3xl">{{ c.name }}</h1>
                <!--p class="italic inline mr-1">(Pick {{ c.vote_count }})</p-->
                <div class="italic inline text-red-900">
                  <p v-if="c.selected.length == 0">(Currently abstaining, {{ c.vote_count }} votes needed)</p>
                  <p v-else-if="c.selected.length != 0 && c.selected.length != c.vote_count">(Not enough votes, select {{ c.vote_count - c.selected.length }} more or leave none selected to abstain)</p>
                  <p class="text-tertiary" v-else>This one's good to go!</p>
                </div>
              </header>
              <MultiCheckbox
                v-model="c.selected"
                :options="c.students"
                :max="c.vote_count"
              />
            </div>
          </div>
          <button 
            class="w-full h-12 bg-primary hover:bg-secondary transition-colors duration-100 text-black rounded disabled:opacity-25"
            :disabled="sendingVote == true"
            @click="submitVote"
          >
            <p v-if="sendingVote == true">Working...</p>
            <p v-else>Submit Vote</p>
          </button>
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
  data() {
    return {
      code: Object,
      response: Object,
      student: Object,
      classes: Object,
      showBallot: false,
      sendingVote: false,
      errors: []
    };
  },
  components: {
    Loader,
    MultiCheckbox
  },
  methods: {
    submitVote(e) {
      e.preventDefault();
      this.errors = []
      var payload = {"ballot": []}
      this.classes.forEach(cls => { //Why did I call them 'class'? Now I have to use short abbrevations :(

        //Error checking
        if (cls.selected.length != 0 && cls.selected.length != cls.vote_count) {
          this.errors.push(cls.name + " doesn't have enough votes. Select more candidates or deselect all to abstain.")
        }
        let s = new Set(cls.selected)
        if (s.size != cls.selected.length) {
          this.errors.push(cls.name + " has duplicate candidates selected. You may be able to fix this by reloading your page.")
        }

        //Payload assembly
        payload.ballot.push({
          "id": cls.id,
          "selected": cls.selected
        })

      });

      if (this.errors.length > 0) {
        this.$refs.top.scrollIntoView({behavior: "smooth"});
        return
      }

      let s = this.$store.state;
      this.sendingVote = true;
      axios.post(s.globalSettings.backendUrl + "vote/" + s.code, payload).then(response => {
        if (response.data.error !== undefined) {
          this.errors.push(response.data.error)
          this.$refs.top.scrollIntoView({behavior: "smooth"});
        }
      }).catch(e => {
        let data = e.response.data
        if (data.error !== undefined) {
          this.errors.push("Server error: " + data.error + ". If you have not intentially caused this, please report this immediately!")
          this.$refs.top.scrollIntoView({behavior: "smooth"});
        }
      });
    }
  },
  beforeMount() {
    this.$store.commit("code", this.$route.params.code); //Push the code to the store in case of direct navigation

    let s = this.$store.state;
    axios
      .get(s.globalSettings.backendUrl + "code/" + s.code)
      .then(response => {
        //Validation of code, throw error if necessary
        let data = response.data;
        this.response = data;
        this.student = data.student;
        this.code = data;

        //Check for problems
        if (this.student == null) {
          throw "The entered code does not exist";
        } else if (data.times_used >= data.max_uses) {
          throw "The entered code has no more uses left";
        } else if (data.active == false) {
          throw "The entered code is currently deactivated";
        }

        this.classes = data.student.votes_for;

        Object.keys(this.classes).forEach(key => {
          let c = this.classes[key]
          if (c.students.length == c.vote_count) {
            c.selected = c.students.map(student => student.id)
          } else {
            c.selected = []
          }
        });

        setTimeout(() => {
          this.showBallot = true;
        }, 500); //Our UI is too fast otherwise lol
      })
      .catch(e => {
        setTimeout(() => {
          if (e.message == "Network Error") {
            this.$store.commit(
              "error",
              "Something went wrong while contacting the voting service"
            );
          } else {
            this.$store.commit("error", e);
          }
          this.$router.push("/");
        }, 500); //Our UI is too fast otherwise lol
      });
  }
};
</script>
