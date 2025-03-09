<script setup lang="ts">
const message = defineProps(
    {
      name : String,
      content : String,
    }
)
</script>

<template>
  <div class="chat">
    <p id="name">{{message.name}}</p>
    <p id="content">{{message.content}}</p>
  </div>
</template>

<style scoped>
p {
  margin: 0;
}

#name {
  background-color: bisque;
  border-radius: 5px;
  padding: 3px;
  width: fit-content;
  font-weight: bold;
}

#content{
  margin-top: 3px;
  white-space: pre-wrap;
  user-select: text;
  -webkit-user-select: text;
  word-break: break-word;
}

.chat {
  width: 100%;
  background-color: beige;

  outline: #12101d solid 3px;

  border-radius: 5px;
  color: #333333;
  padding: 10px;
  text-align: left;
  margin: 10px auto;
  align-self: flex-start;
}
</style>