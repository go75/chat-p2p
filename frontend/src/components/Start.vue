<script lang="ts" setup>
import {reactive} from 'vue'
import {C} from '../../wailsjs/go/main/App'
import {ref} from 'vue'

var name = ref('')
function c() {
  C(name.value).then(result => {
    if (result != "") {
      console.log(result)
      alert("加入失败, 请切换一个名称后重试")
      return
    }

    console.log(result, "---")
    emitNmae(name.value)
  })
}

const emit = defineEmits(['name'])
function emitNmae(name:string) {
  emit('name',name)
}

</script>

<template>
  <div class="start">
    <br/>
    <input class="input" v-model="name" type="text" placeholder="输入自己的名称">
    <br>
    <button @click="c">加入</button>
  </div>
</template>

<style scoped>
  .start {
    background-image: url(https://img95.699pic.com/photo/40140/9844.gif_wh300.gif!/gifto/false);
    background-repeat: no-repeat;
    background-size:cover;
    color: white;
    height: 100%;
    width: 100%;
  }
  .input {
    width: 250px;
    height: 40px;
    margin-top: 50px;
    margin-bottom: 50px;
    text-align: center;
  }
  button {
    width: 80px;
    height: 40px;
    border-radius: 10px;
    border: 0;
    color: white;
    text-align: center;
    font-size: 20px;
    font-weight: bold;
    background-color: lightskyblue;
  }
</style>
