<script lang="ts" setup>
import {reactive} from 'vue'
import {useRouter} from "vue-router";
import {NewClient} from "../../wailsjs/go/main/App";

const data = reactive({
  name: "",
  serverIP: "",
  errorMessage: "Your Name?"
})

const router = useRouter();

async function logIn() {
  if (data.name === "") return
  const ok = await NewClient(data.name, data.serverIP)
  if (!ok) {
    data.errorMessage = `We can't find this server : ${data.serverIP}`
    setTimeout(() => {
      data.errorMessage = "Your Name?"
    }, 3000)
    return
  }

  await router.push({ name: "chat", params: { name: data.name } })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.errorMessage }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text" placeholder="Name..."/>
      <input id="serverIP" v-model="data.serverIP" autocomplete="off" class="input" type="text" placeholder="Server IP..."/>
      <button class="btn" @click="logIn">Log in</button>
    </div>
  </main>
</template>

<style scoped>
main {
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
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

.input-box .btn:hover {
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
