<template>
  <input @click="typing = true"
         @blur="typing = false"
         v-model="message"
         type="text"
         placeholder="留下你的精彩评论吧"
         class="input">
</template>

<script>
export default {
  name: "AutoInput",
  mounted() {
  },
  computed: {},
  data() {
    return {
      message: "",
    };
  },
  methods: {
  },
};
</script>

<style scoped lang="less">
.auto-input {
  font-size: 12rem;
  width: 100%;
  max-height: 70rem;
  overflow-y: scroll;
  padding: 0 5rem;
  outline: none;
}

.auto-input::-webkit-scrollbar {
  width: 0 !important;
}

.auto-input:empty::before {
  /*content: "留下你的精彩评论吧";*/
  content: attr(placeholder);
  color: #999999;
}

.input {
  flex: 1;
  outline: none;
  border: none;
  background: #ececec;
}

.auto-input:focus::before {
  content: none;
}
</style>
