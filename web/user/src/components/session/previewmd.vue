<template>
        <MdPreview v-model="textV"></MdPreview>
</template>

<script setup lang="ts">
import  {MdPreview}  from "md-editor-v3"
import {ref} from 'vue'
import 'md-editor-v3/lib/style.css'
const props = defineProps(["text"])
const textV = ref(props.text)

</script>

<style scoped lang="scss">
</style>