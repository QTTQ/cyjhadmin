
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="id" prop="ID"><el-input v-model="queryParams.ID" placeholder="请输入id" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="装扮名字" prop="title"><el-input v-model="queryParams.title" placeholder="请输入装扮名字" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="装扮小名" prop="smTitle"><el-input v-model="queryParams.smTitle" placeholder="请输入装扮小名" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="装扮类型索引" prop="typeIndex"><el-select v-model="queryParams.typeIndex"
                                               placeholder="装扮装扮类型索引" clearable size="small">
                                        <el-option
                                                v-for="dict in typeIndexOptions"
                                                :key="dict.dictValue"
                                                :label="dict.dictLabel"
                                                :value="dict.dictValue"
                                        />
                                    </el-select>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:dress:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:dress:edit']"
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
                                v-permisaction="['admin:dress:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="dressList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="id" align="center" prop="ID"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮名字" align="center" prop="title"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮小名" align="center" prop="smTitle"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮类型" align="center" prop="typeName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮描述" align="center" prop="desc"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮icon" align="center" prop="icon"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮图片地址" align="center" prop="url"
                                                 :show-overflow-tooltip="true"/><el-table-column label="备注" align="center" prop="remarks"
                                                 :show-overflow-tooltip="true"/><el-table-column label="标签" align="center" prop="tags"
                                                 :show-overflow-tooltip="true"/><el-table-column label="装扮等级" align="center" prop="level"
                                                 :show-overflow-tooltip="true"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['admin:dress:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['admin:dress:remove']"
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
                        
                                    <el-form-item label="装扮名字" prop="title">
                                        <el-input v-model="form.title" placeholder="装扮名字"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="装扮小名" prop="smTitle">
                                        <el-input v-model="form.smTitle" placeholder="装扮小名"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="装扮类型索引" prop="typeIndex">
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
                                    <el-form-item label="装扮类型" prop="typeName">
                                        <el-input v-model="form.typeName" placeholder="装扮类型"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="装扮描述" prop="desc">
                                        <el-input v-model="form.desc" placeholder="装扮描述"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="装扮icon" prop="icon">
                                        <el-input v-model="form.icon" placeholder="装扮icon"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="装扮图片地址" prop="url">
                                        <el-input v-model="form.url" placeholder="装扮图片地址"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="备注" prop="remarks">
                                        <el-input v-model="form.remarks" placeholder="备注"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="标签" prop="tags">
                                        <el-input v-model="form.tags" placeholder="标签"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="装扮等级" prop="level">
                                        <el-input v-model="form.level" placeholder="装扮等级"
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
    import {addDress, delDress, getDress, listDress, updateDress} from '@/api/dress'
    import FileChoose from '@/components/FileChoose'

    export default {
        name: 'Dress',
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
                dressList: [],
                typeIndexOptions: [],
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    ID:undefined,
                    title:undefined,
                    smTitle:undefined,
                    typeIndex:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {ID:
                [
                    {required: true, message: 'id不能为空', trigger: 'blur'}
                ],
                title:
                [
                    {required: true, message: '装扮名字不能为空', trigger: 'blur'}
                ],
                smTitle:
                [
                    {required: true, message: '装扮小名不能为空', trigger: 'blur'}
                ],
                typeIndex:
                [
                    {required: true, message: '装扮类型索引不能为空', trigger: 'blur'}
                ],
                }
        }
        },
        created() {
            this.getList()
            this.getDicts('sys_dress_type').then(response => {
                this.typeIndexOptions = response.data
            })

            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listDress(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.dressList = response.data.list
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
                title: undefined,
                smTitle: undefined,
                typeIndex: undefined,
                typeName: undefined,
                desc: undefined,
                icon: undefined,
                url: undefined,
                remarks: undefined,
                tags: undefined,
                level: undefined,
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
                this.title = '添加装扮'
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
                getDress(ID).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改装扮'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.ID !== undefined) {
                            updateDress(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addDress(this.form).then(response => {
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
                      return delDress( { 'ids': Ids })
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>
