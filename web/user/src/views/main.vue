<template>
    <div class="content">
        <div class="bg"></div>
        <main>
            <asideVue></asideVue>
            
        </main>
    </div>
</template>

<script lang="ts" setup>
import asideVue from '../components/aside.vue';

</script>

<style lang="scss" scoped>
    .content{
        height: 100%;
        position: relative;
        .bg{
            position: fixed;
            top:0;
            left: 0;
            width: 100vw;
            height: 100vh;
            background: url("../assets/img/texture-rain-bg.png");
            background-size: cover;
            background-position: 100%;
            filter: blur(5px);
        }
        main{
            position: fixed;
            display: flex;
            border-radius: 15px;
            top:0;
            left: 0;
            right: 0;
            bottom: 0;
            margin: auto;
            width: 1200px;
            height: 580px;
            background-color: aliceblue;
        }
    }
</style>