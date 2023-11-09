<template>
  <div class="ready">
    <div class="introduce">
      <div class="style"></div>
      <div class="carsouel">
        <el-carousel height="300px">
          <el-carousel-item v-for="item in images" :key="item">
            <img :src="item" alt="">
          </el-carousel-item>
        </el-carousel>
      </div>
    </div>

    <div class="content">
      <div class="style"></div>
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup lang="ts">



const images = ["http://s3sk3sf71.hb-bkt.clouddn.com/IMG_20231109_170634.png",
"http://s3sk3sf71.hb-bkt.clouddn.com/IMG_20231109_170614.png",
"http://s3sk3sf71.hb-bkt.clouddn.com/IMG_20231109_170558.png"]
</script>

<style lang="scss" scoped>
.ready {
  position: absolute;
  display: flex;
  width: 1000px;
  height: 500px;
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  background-color: #fff;

  .introduce {
    height: 100%;
    flex-grow: 1;
    position: relative;
    overflow: hidden;
    justify-content: center;
    align-items: center;

    .carsouel {
      position: absolute;
      top:100px;
      left: 100px;
      height: 60%;
      width: 80%;
      img{
        width: 100%;
      }
    }

    .style {
      position: relative;
      height: 200px;
      width: 200px;
      border: 30px solid #8cc63e;
      right: 100px;
      top: 400px;
      border-radius: 50%;
    }
  }

  .content {
    position: relative;
    justify-content: center;
    align-items: center;
    height: 100%;
    width: 500px;
    overflow: hidden;

    .style {
      position: relative;
      height: 1000px;
      width: 1000px;
      border: 80px solid #008efe;
      left: 200px;
      top: -600px;
      border-radius: 50%;
    }
  }

}</style>