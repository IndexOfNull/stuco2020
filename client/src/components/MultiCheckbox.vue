<template>
  <div>
    <div>
      <div v-for="student in options" :key="student.id" class="text-left">
        <div class="mb-1">
          <span v-if="student.image != undefined" class="mr-2">
            <img class="inline w-20 h-20 rounded" :src="student.image">
          </span>
          <span class="w-20" v-else>
            <p class="inline bg-gray-300 mr-2 rounded p-2 w-full h-full text-sm">No Image</p>
          </span>
          <input
            type="checkbox"
            class="h-10 w-10 align-middle rounded form-checkbox text-primary"
            :id="student.id"
            :value="student.id"
            :checked="modelValue.includes(student.id)"
            :disabled="
              (max != -1 &&
                modelValue.length >= max &&
                !modelValue.includes(student.id)) ||
                options.length == max
            "
            v-on:change="onChange"
          />
          <label class="text-xl ml-2 align-middle" :for="student.id"
            >{{ student.firstname }} {{ student.lastname }}</label
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "MultiCheckbox",
  props: {
    max: {
      type: Number,
      default: -1
    },
    options: {
      type: Object
    },
    modelValue: {
      type: Array,
      default: function() {
        return [];
      }
    }
  },
  methods: {
    onChange(e) {
      let newVal = [...this.modelValue];
      if (e.target.checked) {
        newVal.push(parseInt(e.target.value));
      } else {
        newVal.splice(newVal.indexOf(parseInt(e.target.value)), 1);
      }

      this.$emit("update:modelValue", newVal);
    }
  }
};
</script>
