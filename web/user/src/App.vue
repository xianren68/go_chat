<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router"
const router = useRouter()
if(router.currentRoute.value.name == 'chat'){
  router.replace('/')
}
</script>

<style lang="scss" scoped>

</style>
