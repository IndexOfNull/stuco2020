<template>
    <div>
        <div>
            <div v-for="student in options" :key="student.id" class="text-left">
                <div class="mb-1">
                    <input
                        type="checkbox"
                        class="h-8 w-8 align-middle rounded form-checkbox text-primary"
                        :id="student.id"
                        :value="student.id"
                        :checked="modelValue.includes(student.id)"
                        :disabled="(max != -1 && modelValue.length >= max && !modelValue.includes(student.id)) || options.length == max"
                        v-on:change="onChange"
                    >
                    <label class="text-lg ml-2 align-middle" :for="student.id">{{ student.firstname}} {{ student.lastname }}</label>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "MultiCheckbox",
    data () {
        return {
            selected: []
        }
    },
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
                return []
            }
        }
    },
    methods: {
        onChange (e) {
            let newVal = [...this.modelValue]
            if (e.target.checked) {
                newVal.push(parseInt(e.target.value))
            } else {
                newVal.splice(newVal.indexOf(e.target.value), 1)
            }

            this.$emit('update:modelValue', newVal)
        }
    },
    beforeMount () {
        console.log(this.modelValue)
    }
}
</script>