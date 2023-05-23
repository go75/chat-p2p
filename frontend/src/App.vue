<script lang="ts" setup>
  import { Ref, reactive, ref } from "vue";
  import { dataType } from 'element-plus/es/components/table-v2/src/common';
  import Start from './components/Start.vue'
  import {Send, Connect} from '../wailsjs/go/main/App'
  import {EventsOn} from '../wailsjs/runtime/runtime'

  var myname = ref('')
  var peername = ref('')
  var msg = ref('')
  var msgs = ref([''])
  function getName(name:string) {
    myname.value = name
  }

  var sessionList = ref([''])
  var recvList = ref([''])
  function send() {
    Send(msg.value).then(res => {
      if (res != '') {
        alert('发送失败' + res)
      }
    })
  }
  
  
  function connect() {
    Connect(peername.value).then(res => {
      if (res != '') {
        alert('连接失败' + res)
        return
      }
    })
    
  }

  EventsOn('msg', (msg)=>{
    msgs.value.push(msg as string)
  })

  EventsOn('addr', (res)=> {
    if (res == 'err') {
      alert('连接失败')
    }
  })

</script>

<template>
  <Start v-if="myname==''" @name = "getName"></Start>
  <div class="main" v-else>
    我的名称:{{ myname }}
    <br>
    <input id="peerNameInput" v-model="peername" type="text" placeholder="对方名称">
    <button @click="connect">连接</button>
    <br>
    <input id="msgInput" v-model="msg" type="text" placeholder="消息内容">
    <button @click="send">发送</button>
    <div class="msgs">
      <p v-for="msg in msgs">
        {{ msg }}
      </p>
    </div>
  </div>
</template>

<style scoped>
</style>