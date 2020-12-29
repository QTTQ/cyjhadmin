<template>
  <div>
    <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px"
      label-position="left">
      <el-form-item label="单行文本" prop="mobile">
        <el-input v-model="formData.mobile" placeholder="请输入手机号" :maxlength="11" show-word-limit clearable
          prefix-icon='el-icon-mobile' :style="{width: '100%'}"></el-input>
      </el-form-item>
      <el-form-item label="呵呵呵" prop="field101">
        <el-input v-model="formData.field101" placeholder="请输入呵呵呵" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item label="单行文本" prop="field102">
        <el-input v-model="formData.field102" placeholder="请输入单行文本" clearable :style="{width: '100%'}">
        </el-input>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getTest } from '@/api/test'
export default {
  components: {},
  props: [],
  data() {
    return {
      formData: {
        mobile: undefined,
        field101: undefined,
        field102: undefined,
      },
      rules: {
        mobile: [{
          required: true,
          message: '请输入手机号',
          trigger: 'blur'
        }, {
          pattern: /^1(3|4|5|7|8|9)\d{9}$/,
          message: '手机号格式错误',
          trigger: 'blur'
        }],
        field101: [{
          required: true,
          message: '请输入呵呵呵',
          trigger: 'blur'
        }],
        field102: [{
          required: true,
          message: '请输入单行文本',
          trigger: 'blur'
        }],
      },
    }
  },
  computed: {},
  watch: {},
  created() {},
  mounted() {
        console.log("bbbbbbbbbbbbbbbbbbbbbbb")
  },
  methods: {
    submitForm() {
         getTest(1).then(response => {
        console.log(response,"bbbbbbbbbbbbbbbbbbbbbbb")
      })
        console.log("aaaaaaaaaaaaaaaaaaaaaaa")
        this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单

      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
  }
}

</script>
<style>
</style>
