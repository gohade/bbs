<template>
  <div class="notfound">
    <el-card>
      <h2>页面找不到了</h2>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "notfound",
  data() {
    return {
      model: {
        username: "",
        password: ""
      },
      loading: false,
    };
  },
  methods: {
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.notfound {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 240px;
}

</style>
