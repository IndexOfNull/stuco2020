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
                        v-model="selected"
                        :disabled="max != -1 && selected.length >= max && !selected.includes(student.id)"
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
        }
    }
}
</script>