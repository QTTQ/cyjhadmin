
<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form
          ref="queryForm"
          :model="queryParams"
          :inline="true"
          label-width="68px"
        >
          <el-form-item label="创建时间" prop="createdAt"
            ><el-input
              v-model="queryParams.createdAt"
              placeholder="请输入创建时间"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="标签索引" prop="tagIndex"
            ><el-input
              v-model="queryParams.tagIndex"
              placeholder="请输入标签索引"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="标签名" prop="tagName"
            ><el-select
              v-model="queryParams.tagName"
              placeholder="标签标签名"
              clearable
              size="small"
            >
              <el-option
                v-for="dict in tagNameOptions"
                :key="dict.dictValue"
                :label="dict.dictLabel"
                :value="dict.dictValue"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="标签索引" prop="typeIndex"
            ><el-input
              v-model="queryParams.typeIndex"
              placeholder="请输入标签索引"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>
          <el-form-item label="类型名" prop="typeName"
            ><el-input
              v-model="queryParams.typeName"
              placeholder="请输入类型名"
              clearable
              size="small"
              @keyup.enter.native="handleQuery"
            />
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              icon="el-icon-search"
              size="mini"
              @click="handleQuery"
              >搜索</el-button
            >
            <el-button icon="el-icon-refresh" size="mini" @click="resetQuery"
              >重置</el-button
            >
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:tags:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:tags:edit']"
              type="success"
              icon="el-icon-edit"
              size="mini"
              :disabled="single"
              @click="handleUpdate"
              >修改
            </el-button>
          </el-col>
          <el-col :span="1.5">
            <el-button
              v-permisaction="['admin:tags:remove']"
              type="danger"
              icon="el-icon-delete"
              size="mini"
              :disabled="multiple"
              @click="handleDelete"
              >删除
            </el-button>
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="tagsList"
          @selection-change="handleSelectionChange"
        >
          <el-table-column
            type="selection"
            width="55"
            align="center"
          /><el-table-column
            label="创建时间"
            align="center"
            prop="createdAt"
            :show-overflow-tooltip="true"
          /><el-table-column
            label="标签名"
            align="center"
            prop="tagName"
            :formatter="tagNameFormat"
            width="100"
          >
            <template slot-scope="scope">
              {{ tagNameFormat(scope.row) }}
            </template> </el-table-column
          ><el-table-column
            label="类型名"
            align="center"
            prop="typeName"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
              <el-button
                v-permisaction="['admin:tags:edit']"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
                >修改
              </el-button>
              <el-button
                v-permisaction="['admin:tags:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
                >删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />

        <!-- 添加或修改对话框 -->
        <el-dialog :title="title" :visible.sync="open" width="500px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="标签索引" prop="tagIndex">
              <el-input v-model="form.tagIndex" placeholder="标签索引" />
            </el-form-item>
            <el-form-item label="标签名" prop="tagName">
              <el-select v-model="form.tagName" placeholder="请选择">
                <el-option
                  v-for="dict in tagNameOptions"
                  :key="dict.dictValue"
                  :label="dict.dictLabel"
                  :value="dict.dictValue"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="标签索引" prop="typeIndex">
              <el-input v-model="form.typeIndex" placeholder="标签索引" />
            </el-form-item>
            <el-form-item label="类型名" prop="typeName">
              <el-input v-model="form.typeName" placeholder="类型名" />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
        <FileChoose
          ref="fileChoose"
          :dialog-form-visible="fileOpen"
          @confirm="getImgList"
          @close="fileClose"
        />
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { addTags, delTags, getTags, listTags, updateTags } from '@/api/tags'
import FileChoose from '@/components/FileChoose'

export default {
  name: 'Tags',
  components: {
    FileChoose
  },
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      fileOpen: false,
      fileIndex: undefined,
      // 类型数据字典
      typeOptions: [],
      tagsList: [],
      tagNameOptions: [],
      // 关系表类型

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        createdAt: undefined,
        tagIndex: undefined,
        tagName: undefined,
        typeIndex: undefined,
        typeName: undefined,

      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: {
        createdAt:
          [
            { required: true, message: '创建时间不能为空', trigger: 'blur' }
          ],
        tagIndex:
          [
            { required: true, message: '标签索引不能为空', trigger: 'blur' }
          ],
        tagName:
          [
            { required: true, message: '标签名不能为空', trigger: 'blur' }
          ],
        typeIndex:
          [
            { required: true, message: '标签索引不能为空', trigger: 'blur' }
          ],
        typeName:
          [
            { required: true, message: '类型名不能为空', trigger: 'blur' }
          ],
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('sys_tag_type').then(response => {
      this.tagNameOptions = response.data
    })

  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listTags(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.tagsList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {

        ID: undefined,
        tagIndex: undefined,
        tagName: undefined,
        typeIndex: undefined,
        typeName: undefined,
      }
      this.resetForm('form')
    },
    getImgList: function () {
      this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
    },
    fileClose: function () {
      this.fileOpen = false
    },
    tagNameFormat(row) {
      return this.selectDictLabel(this.tagNameOptions, row.tagName)
    },
    // 关系
    // 文件
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加标签'
      this.isEdit = false
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.ID)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      const ID =
        row.ID || this.ids
      getTags(ID).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改标签'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function () {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.ID !== undefined) {
            updateTags(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addTags(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('新增成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      var Ids = (row.ID && [row.ID]) || this.ids

      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function () {
        return delTags({ 'ids': Ids })
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function () {
      })
    }
  }
}
</script>
