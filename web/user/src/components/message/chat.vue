<template>
    <div class="chat">
        <div class="top">hello</div>
        <div class="main"></div>
        <!--发送消息-->
        <div class="send">
            <div class="send-main">
                <div class="input">
                    <svg class="icon">
                        <use xlink:href="#icon-huatong"></use>
                    </svg>
                    <input>
                    <svg class="icon">
                        <use xlink:href="#icon-biaoqing"></use>
                    </svg>
                    <svg class="icon">
                        <use xlink:href="#icon-attachment"></use>
                    </svg>
                </div>
                <div class="send-icon">
                    <svg class="icon">
                        <use xlink:href="#icon-send"></use>
                    </svg>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
</script>

<style scoped lang="scss">
.chat {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;

    .top {
        height: 8%;
        display: flex;
        border-bottom: #ccc 1px solid;
        justify-content: center;
        align-items: center;
        font-family: "Fira Code";
    }

    .main {
        height: 80%;
        overflow: scroll;
    }

    .main::-webkit-scrollbar {
        display: none;
    }

    .send {
        flex-grow: 1;

        .send-main {
            display: flex;
            align-items: center;

            .input {
                margin-left: 20px;
                height: 40px;
                width: 655px;
                border: 1px solid #ccc;
                border-radius: 13px;
                display: flex;
                align-items: center;

                .icon {
                    height: 20px;
                    width: 20px;
                    margin-left: 8px;
                    fill: #999;
                }

                input {
                    width: 550px;
                    height: 90%;
                    font-size: 16px;
                    border: none;
                    outline: none;
                    margin-left: 10px;
                }
            }

            .send-icon {
                height: 30px;
                width: 30px;
                margin-left: 20px;
                background-color: #ecedf6;
                display: flex;
                align-items: center;
                border-radius: 50%;

                .icon {
                    height: 20px;
                    width: 20px;
                    margin-left:4px ;
                    fill: #4a499b;
                }
            }

        }

    }
}
</style>