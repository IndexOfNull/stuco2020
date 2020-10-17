<template>
    <div class="border-solid border-2 rounded p-2 mb-4 border-white">
        <h1 class="text-6xl m-2" :class=" { 'm-2': (subtext != '') } ">{{ title }}</h1>
        <p class="m-3" v-if="subtext != ''">{{ subtext }}</p>
    </div>
</template>

<script>
export default {
    name: "TitleCard",
    props: {
        title: String,
        subtext: String
    }
}
</script>