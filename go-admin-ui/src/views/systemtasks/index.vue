
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="标签索引" prop="typeIndex"><el-select v-model="queryParams.typeIndex"
                                               placeholder="系统任务标签索引" clearable size="small">
                                        <el-option
                                                v-for="dict in typeIndexOptions"
                                                :key="dict.dictValue"
                                                :label="dict.dictLabel"
                                                :value="dict.dictValue"
                                        />
                                    </el-select>
                            </el-form-item>
                        <el-form-item label="奖励金" prop="reward"><el-input v-model="queryParams.reward" placeholder="请输入奖励金" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="是否需要关联上级" prop="isRelat"><el-input v-model="queryParams.isRelat" placeholder="请输入是否需要关联上级" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:systemtasks:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:systemtasks:edit']"
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
                                v-permisaction="['admin:systemtasks:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="systemtasksList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="id" align="center" prop="ID"
                                                 :show-overflow-tooltip="true"/><el-table-column label="创建时间" align="center" prop="createdAt"
                                                 :show-overflow-tooltip="true"/><el-table-column label="标签索引" align="center" prop="typeIndex"
                                                 :formatter="typeIndexFormat" width="100">
                                    <template slot-scope="scope">
                                        {{ typeIndexFormat(scope.row) }}
                                    </template>
                                </el-table-column><el-table-column label="标签名" align="center" prop="typeName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="任务内容" align="center" prop="content"
                                                 :show-overflow-tooltip="true"/><el-table-column label="跳转地址" align="center" prop="jumpUrl"
                                                 :show-overflow-tooltip="true"/><el-table-column label="静态地址" align="center" prop="staticUrl"
                                                 :show-overflow-tooltip="true"/><el-table-column label="奖励金" align="center" prop="reward"
                                                 :show-overflow-tooltip="true"/><el-table-column label="是否需要关联上级" align="center" prop="isRelat"
                                                 :show-overflow-tooltip="true"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['admin:systemtasks:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['admin:systemtasks:remove']"
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
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <el-dialog :title="title" :visible.sync="open" width="500px">
                    <el-form ref="form" :model="form" :rules="rules" label-width="80px">
                        
                                    <el-form-item label="标签索引" prop="typeIndex">
                                        <el-select v-model="form.typeIndex"
                                                       placeholder="请选择" >
                                                <el-option
                                                        v-for="dict in typeIndexOptions"
                                                        :key="dict.dictValue"
                                                        :label="dict.dictLabel"
                                                        :value="dict.dictValue"
                                                />
                                            </el-select>
                                    </el-form-item>
                                    <el-form-item label="标签名" prop="typeName">
                                        <el-input v-model="form.typeName" placeholder="标签名"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="任务内容" prop="content">
                                        <el-input v-model="form.content" placeholder="任务内容"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="跳转地址" prop="jumpUrl">
                                        <el-input v-model="form.jumpUrl" placeholder="跳转地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="静态地址" prop="staticUrl">
                                        <el-input v-model="form.staticUrl" placeholder="静态地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="tasksUuid">
                                        <el-input v-model="form.tasksUuid" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="upTkUuid">
                                        <el-input v-model="form.upTkUuid" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="奖励金" prop="reward">
                                        <el-input v-model="form.reward" placeholder="奖励金"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="是否需要关联上级" prop="isRelat">
                                        <el-input v-model="form.isRelat" placeholder="是否需要关联上级"
                                                      />
                                    </el-form-item>
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog>
                <FileChoose ref="fileChoose" :dialog-form-visible="fileOpen" @confirm="getImgList" @close="fileClose" />
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addSystemTasks, delSystemTasks, getSystemTasks, listSystemTasks, updateSystemTasks} from '@/api/systemtasks'
    import FileChoose from '@/components/FileChoose'

    export default {
        name: 'SystemTasks',
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
                systemtasksList: [],
                typeIndexOptions: [],
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    typeIndex:undefined,
                    reward:undefined,
                    isRelat:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {typeIndex:
                [
                    {required: true, message: '标签索引不能为空', trigger: 'blur'}
                ],
                reward:
                [
                    {required: true, message: '奖励金不能为空', trigger: 'blur'}
                ],
                isRelat:
                [
                    {required: true, message: '是否需要关联上级不能为空', trigger: 'blur'}
                ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('sys_tag_type').then(response => {
                this.typeIndexOptions = response.data
            })

            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listSystemTasks(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.systemtasksList = response.data.list
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
                typeIndex: undefined,
                typeName: undefined,
                content: undefined,
                jumpUrl: undefined,
                staticUrl: undefined,
                tasksUuid: undefined,
                upTkUuid: undefined,
                reward: undefined,
                isRelat: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            typeIndexFormat(row) {
                return this.selectDictLabel(this.typeIndexOptions, row.typeIndex)
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
                this.title = '添加系统任务'
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
                getSystemTasks(ID).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改系统任务'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.ID !== undefined) {
                            updateSystemTasks(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addSystemTasks(this.form).then(response => {
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
                      return delSystemTasks( { 'ids': Ids })
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>
