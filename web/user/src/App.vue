<template>
  <router-view></router-view>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router"
import {userStore} from "@/store"
document.oncontextmenu = function (event) {
  event.preventDefault();
};
const userstore = userStore()
userstore.getUser()
const router = useRouter()
  router.replace('/')

</script>

<style lang="scss" scoped>

</style>
