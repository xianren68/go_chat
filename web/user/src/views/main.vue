<template>
  <div class="content">
    <div class="bg"></div>
    <main>
      <asideVue></asideVue>
      <router-view></router-view>
    </main>
  </div>
</template>

<script lang="ts" setup>
import asideVue from '@/components/aside.vue'
import { onBeforeMount} from 'vue'
import {init} from '@/utils/init'
onBeforeMount(async () => {
  // 初始化 socket连接，数据库连接
  await init()
})

</script>

<style lang="scss" scoped>
.content {
  height: 100%;
  position: relative;

  .bg {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: linear-gradient(to right, #005aa7, #fffde4)
  }

  main {
    padding: 20px;
    position: fixed;
    display: flex;
    border-radius: 15px;
    box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    margin: auto;
    width: 1000px;
    height: 540px;
    background-color: #f9f9fd;
  }
}
</style>@/utils/init