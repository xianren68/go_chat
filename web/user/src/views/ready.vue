<template>
    <div class="ready">
        <div class="introduce">
            <div class="style"></div>
        </div>
        <div class="content">
          <div class="style"></div>
            <router-view></router-view>

        </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style lang="scss" scoped>
.ready{
    position: absolute;
    display:flex;
    width: 1000px;
    height: 500px;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    margin: auto;
    background-color: #fff;
    .introduce{
        height:100%;
        flex-grow:1;
        overflow: hidden;
        .style{
          position: relative;
          height: 200px;
          width: 200px;
          border:30px solid #8cc63e;
          right: 100px;
          top:400px;
          border-radius: 50%;
        }
    }
    .content{
      position: relative;
        justify-content: center;
        align-items: center;
        height:100%;
        width:500px;
        overflow: hidden;
      .style{
        position: relative;
        height: 1000px;
        width: 1000px;
        border:80px solid #008efe;
        left: 200px;
        top: -600px;
        border-radius: 50%;
      }
    }

}
</style>