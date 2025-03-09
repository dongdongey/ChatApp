<script setup lang="ts">
import {nextTick, onMounted, reactive, useTemplateRef} from "vue";
import {useRoute, useRouter} from "vue-router";
import Chat from "./Chat.vue";
import {DestroyClient, Receive, Send} from "../../wailsjs/go/main/App";
import * as runtime from "../../wailsjs/runtime/runtime"

const data = reactive({
  content : "",
  receivedContent: [] as { Name: string, Content: string }[]
})

const chatLog = useTemplateRef<HTMLInputElement>('chat-log')

const routes = useRoute()
const router = useRouter()

async function send() {
  if (data.content === ""){
    return
  }
  console.log("name: " + routes.params.name)
  console.log("content: " + data.content)
  await Send(data.content)
  data.content = ""
}

async function logout(){
  await DestroyClient()
  await router.push("/")
}

Receive()
runtime.EventsOn('receive-message', (msg: string) => {
  const message: { Name: string, Content: string } = JSON.parse(msg)
  data.receivedContent.push(message)

  nextTick(() => {
    if (chatLog.value === null) return
    chatLog.value.scrollTop = chatLog.value.scrollHeight
  })
})
</script>

<template>
  <main>
    <div>
      <button id="logout" class="btn" @click="logout">logout</button>
      <span>you : {{routes.params.name}} </span>
    </div>
    <div class="messages" ref="chat-log">
      <Chat v-for="message in data.receivedContent " :name="message.Name" :content="message.Content" />
    </div>
    <div id="input" class="input-box">
      <input id="message" v-model="data.content" autocomplete="off" class="input" type="text" @keydown.enter="send" />
      <button id="send" class="btn" @click="send">Send</button>
    </div>
  </main>
</template>

<style scoped>
.messages{
  height: 500px;
  overflow-x: hidden;
  overflow-y: scroll;
  width: 250px;
  padding: 10px;
  margin: 10px auto;
}

.messages::-webkit-scrollbar {
  display: none;
}

#logout{
  width: 60px;
  height: 20px;
  line-height: 10px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 10px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 10px;
  padding: 0 8px;
  cursor: pointer;
}

.btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
  margin: 0 10px 0 0;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>