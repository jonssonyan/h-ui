<template>
  <div class="app-container">
    <!-- 用户数据 -->
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="queryParams.username"
            placeholder="用户名"
            clearable
            style="width: 200px"
            @keyup.enter="handleQuery"
          />
        </el-form-item>

        <el-form-item label="状态" prop="deleted">
          <el-select
            v-model="queryParams.deleted"
            placeholder="全部"
            clearable
            style="width: 200px"
          >
            <el-option label="启用" value="0" />
            <el-option label="禁用" value="1" />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :icon="Search" @click="handleQuery"
            >搜索
          </el-button>
          <el-button :icon="Refresh" @click="resetQuery">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card>
      <template #header>
        <el-form-item style="float: left">
          <el-button type="success" :icon="Plus" @click="handleAdd"
            >新增
          </el-button>
          <el-button
            type="danger"
            :icon="Delete"
            :disabled="ids.length === 0"
            @click="handleDelete"
            >删除
          </el-button>
        </el-form-item>
      </template>

      <el-table
        v-loading="loading"
        :data="records"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="50" align="center" />
        <el-table-column
          key="id"
          label="编号"
          align="center"
          prop="id"
          width="100"
        />
        <el-table-column
          key="username"
          label="用户名"
          align="center"
          prop="username"
        />

        <el-table-column
          label="创建时间"
          align="center"
          prop="createTime"
          width="180"
        ></el-table-column>
        <el-table-column label="操作" align="left" width="200">
          <template #default="scope">
            <el-button type="primary" link @click="handleUpdate(scope.row)"
              >编辑
            </el-button>
            <el-button type="danger" link @click="handleDelete(scope.row)"
              >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination
        v-if="total > 0"
        :total="total"
        v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize"
        @pagination="handleQuery"
      />
    </el-card>

    <!-- 用户表单 -->
    <el-dialog
      :title="dialog.title"
      v-model="dialog.visible"
      width="600px"
      append-to-body
      @close="closeDialog"
    >
      <el-form
        ref="dataFormRef"
        :model="formData"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            :readonly="!!formData.id"
            v-model="formData.username"
            placeholder="请输入用户名"
          />
        </el-form-item>

        <el-form-item label="邮箱" prop="email">
          <el-input
            v-model="formData.email"
            placeholder="请输入邮箱"
            maxlength="50"
          />
        </el-form-item>

        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.deleted">
            <el-radio :label="0">正常</el-radio>
            <el-radio :label="1">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="closeDialog">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
export default {
  name: "index",
};
</script>

<script setup lang="ts">
import {
  AccountForm,
  AccountPageDto,
  AccountUpdateDto,
  AccountVo,
} from "@/api/account/types";
import {
  saveAccountApi,
  deleteAccountApi,
  getAccountApi,
  pageAccountApi,
  updateAccountApi,
} from "@/api/account";
import { Search, Plus, Refresh, Delete } from "@element-plus/icons-vue";

const queryFormRef = ref(ElForm); // 查询表单
const dataFormRef = ref(ElForm); // 用户表单

const state = reactive({
  // 遮罩层
  loading: true,
  // 选中数组
  ids: [] as number[],
  // 总条数
  total: 0,
  records: [] as AccountVo[],
  dialog: {
    visible: false,
  } as DialogType,
  formData: {} as AccountForm,
  queryParams: {
    username: undefined,
    deleted: undefined,
    pageNum: 1,
    pageSize: 10,
  } as AccountPageDto,
});

const { ids, loading, queryParams, records, total, formData, dialog } =
  toRefs(state);

/**
 * 查询
 */
function handleQuery() {
  state.loading = true;
  pageAccountApi(state.queryParams).then(({ data }) => {
    state.records = data.records;
    state.total = data.total;
    state.loading = false;
  });
}

/**
 * 重置
 */
function resetQuery() {
  queryFormRef.value.resetFields();
  handleQuery();
}

/**
 * 行选中
 */
function handleSelectionChange(selection: any) {
  state.ids = selection.map((item: any) => item.id);
}

/**
 * 保存
 **/
async function handleAdd() {
  state.dialog = {
    title: "添加用户",
    visible: true,
  };
}

/**
 * 修改
 **/
async function handleUpdate(row: { [key: string]: any }) {
  dialog.value = {
    title: "修改用户",
    visible: true,
  };

  // const id = row.id || state.ids;
  getAccountApi({ id: row.id }).then(({ data }) => {
    Object.assign(formData.value, data);
  });
}

/**
 * 表单提交
 */
function submitForm() {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      const accountId = state.formData.id;
      let accountUpdateDto: AccountUpdateDto = Object.assign(
        {},
        state.formData
      );
      if (accountId) {
        updateAccountApi(accountUpdateDto).then(() => {
          ElMessage.success("修改用户成功");
          closeDialog();
          handleQuery();
        });
      } else {
        saveAccountApi(accountUpdateDto).then(() => {
          ElMessage.success("新增用户成功");
          closeDialog();
          handleQuery();
        });
      }
    }
  });
}

/**
 * 删除
 */
function handleDelete(row: { [key: string]: any }) {
  const id = row.id;
  const username = row.username;
  ElMessageBox.confirm(
    "是否确认删除用户名为「" + username + "」的数据项?",
    "警告",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    }
  )
    .then(function () {
      deleteAccountApi({ id: id }).then(() => {
        ElMessage.success("删除成功");
        handleQuery();
      });
    })
    .catch(() => ElMessage.info("已取消删除"));
}

/**
 * 关闭用户弹窗
 */
function closeDialog() {
  dialog.value.visible = false;
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();
}

onMounted(() => {
  // 初始化用户列表数据
  handleQuery();
});
</script>

<style scoped></style>
