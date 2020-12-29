
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="id" prop="ID"><el-input v-model="queryParams.ID" placeholder="请输入id" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="昵称" prop="nickName"><el-input v-model="queryParams.nickName" placeholder="请输入昵称" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="真实姓名" prop="realName"><el-input v-model="queryParams.realName" placeholder="请输入真实姓名" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="电话号" prop="phoneNum"><el-input v-model="queryParams.phoneNum" placeholder="请输入电话号" clearable
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
                                v-permisaction="['admin:userinfo:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:userinfo:edit']"
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
                                v-permisaction="['admin:userinfo:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="userinfoList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="id" align="center" prop="ID"
                                                 :show-overflow-tooltip="true"/><el-table-column label="创建时期" align="center" prop="createdAt"
                                                 :show-overflow-tooltip="true"/><el-table-column label="昵称" align="center" prop="nickName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="性别" align="center" prop="sex"
                                                 :show-overflow-tooltip="true"/><el-table-column label="真实姓名" align="center" prop="realName"
                                                 :show-overflow-tooltip="true"/><el-table-column label="电话号" align="center" prop="phoneNum"
                                                 :show-overflow-tooltip="true"/><el-table-column label="头像" align="center" prop="actor"
                                                 :show-overflow-tooltip="true"/>
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                    v-permisaction="['admin:userinfo:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['admin:userinfo:remove']"
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
                        
                                    <el-form-item label="昵称" prop="nickName">
                                        <el-input v-model="form.nickName" placeholder="昵称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="生日" prop="birthDate">
                                        <el-date-picker
                                                    v-model="form.birthDate"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                                    <el-form-item label="性别" prop="sex">
                                        <el-input v-model="form.sex" placeholder="性别"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="autograph">
                                        <el-input v-model="form.autograph" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="真实姓名" prop="realName">
                                        <el-input v-model="form.realName" placeholder="真实姓名"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="江湖号" prop="jhNum">
                                        <el-input v-model="form.jhNum" placeholder="江湖号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="unionNum">
                                        <el-input v-model="form.unionNum" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="电话号" prop="phoneNum">
                                        <el-input v-model="form.phoneNum" placeholder="电话号"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="password">
                                        <el-input v-model="form.password" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="头像" prop="actor">
                                        <el-input v-model="form.actor" placeholder="头像"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="email">
                                        <el-input v-model="form.email" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="qqOpenId">
                                        <el-input v-model="form.qqOpenId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="wxOpenId">
                                        <el-input v-model="form.wxOpenId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="distId">
                                        <el-input v-model="form.distId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="deviceThemeId">
                                        <el-input v-model="form.deviceThemeId" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="domainName">
                                        <el-input v-model="form.domainName" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="staffNum">
                                        <el-input v-model="form.staffNum" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="province">
                                        <el-input v-model="form.province" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="city">
                                        <el-input v-model="form.city" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="category">
                                        <el-input v-model="form.category" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="payPassword">
                                        <el-input v-model="form.payPassword" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="isfollowWx">
                                        <el-input v-model="form.isfollowWx" placeholder=""
                                                      />
                                    </el-form-item>
                                    <el-form-item label="" prop="firstRecharge">
                                        <el-input v-model="form.firstRecharge" placeholder=""
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
    import {addUserInfo, delUserInfo, getUserInfo, listUserInfo, updateUserInfo} from '@/api/userinfo'
    import FileChoose from '@/components/FileChoose'

    export default {
        name: 'UserInfo',
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
                userinfoList: [],
                
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    ID:undefined,
                    nickName:undefined,
                    realName:undefined,
                    phoneNum:undefined,
                    
                },
                // 表单参数
                form: {
                },
                // 表单校验
                rules: {ID:
                [
                    {required: true, message: 'id不能为空', trigger: 'blur'}
                ],
                nickName:
                [
                    {required: true, message: '昵称不能为空', trigger: 'blur'}
                ],
                realName:
                [
                    {required: true, message: '真实姓名不能为空', trigger: 'blur'}
                ],
                phoneNum:
                [
                    {required: true, message: '电话号不能为空', trigger: 'blur'}
                ],
                }
        }
        },
        created() {
            this.getList()
            },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listUserInfo(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.userinfoList = response.data.list
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
                nickName: undefined,
                birthDate: undefined,
                sex: undefined,
                autograph: undefined,
                realName: undefined,
                jhNum: undefined,
                unionNum: undefined,
                phoneNum: undefined,
                password: undefined,
                actor: undefined,
                email: undefined,
                qqOpenId: undefined,
                wxOpenId: undefined,
                distId: undefined,
                deviceThemeId: undefined,
                domainName: undefined,
                staffNum: undefined,
                province: undefined,
                city: undefined,
                category: undefined,
                payPassword: undefined,
                isfollowWx: undefined,
                firstRecharge: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
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
                this.title = '添加UserInfo'
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
                getUserInfo(ID).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改UserInfo'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.ID !== undefined) {
                            updateUserInfo(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addUserInfo(this.form).then(response => {
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
                      return delUserInfo( { 'ids': Ids })
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>
