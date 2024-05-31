<template>
  <div class="app-container">
    <github-corner class="github-corner" />
    <!-- 用户数据 -->
    <div class="search">
      <el-form ref="queryFormRef" :model="queryParams" :inline="true">
        <el-form-item :label="$t('account.username')" prop="username">
          <el-input
            v-model="queryParams.username"
            :placeholder="$t('account.username')"
            clearable
            style="width: 200px"
            @keyup.enter="handleQuery"
          />
        </el-form-item>

        <el-form-item :label="$t('common.deleted')" prop="deleted">
          <el-select
            v-model="queryParams.deleted"
            :placeholder="$t('common.all')"
            clearable
            style="width: 200px"
          >
            <el-option :label="$t('common.enable')" value="0" />
            <el-option :label="$t('common.disable')" value="1" />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :icon="Search" @click="handleQuery"
            >{{ $t("common.search") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button :icon="Refresh" @click="resetQuery"
            >{{ $t("common.reset") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button :icon="Plus" @click="handleAdd"
            >{{ $t("common.add") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleExport">
            <template #icon>
              <i-ep-download />
            </template>
            导出
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-card shadow="never">
      <el-table v-loading="loading" :data="records">
        <el-table-column
          key="id"
          :label="$t('common.id')"
          align="center"
          prop="id"
          width="100"
        />
        <el-table-column
          key="username"
          :label="$t('account.username')"
          align="center"
          prop="username"
        />
        <el-table-column
          key="quota"
          :label="$t('account.quota')"
          align="center"
          prop="quota"
        >
          <template #default="scope">
            {{ formatBytes(scope.row.quota) }}
          </template>
        </el-table-column>
        <el-table-column
          key="download"
          :label="$t('account.download')"
          align="center"
          prop="download"
        >
          <template #default="scope">
            {{ formatBytes(scope.row.download) }}
          </template>
        </el-table-column>
        <el-table-column
          key="upload"
          :label="$t('account.upload')"
          align="center"
          prop="upload"
        >
          <template #default="scope">
            {{ formatBytes(scope.row.download) }}
          </template>
        </el-table-column>
        <el-table-column
          key="expireTime"
          :label="$t('account.expireTime')"
          align="center"
          prop="expireTime"
          width="180"
        >
          <template #default="scope">
            {{ timestampToDateTime(scope.row.expireTime) }}
          </template>
        </el-table-column>
        <el-table-column
          key="kickUtilTime"
          :label="$t('account.kickUtilTime')"
          align="center"
          prop="kickUtilTime"
          width="180"
        >
          <template #default="scope">
            {{ calculateTimeDifference(scope.row.kickUtilTime) }}
          </template>
        </el-table-column>
        <el-table-column
          key="deviceNo"
          :label="$t('account.deviceNo')"
          align="center"
          prop="deviceNo"
          width="180"
        />
        <el-table-column
          key="online"
          :label="$t('account.online')"
          align="center"
          prop="online"
          width="180"
        />
        <el-table-column
          key="device"
          :label="$t('account.device')"
          align="center"
          prop="device"
          width="180"
        />
        <el-table-column
          key="role"
          :label="$t('account.role')"
          align="center"
          prop="role"
        />
        <el-table-column :label="$t('common.deleted')" align="center">
          <template #default="scope">
            <el-tag v-if="scope.row.deleted === 0" type="success"
              >{{ $t("common.enable") }}
            </el-tag>
            <el-tag v-else type="info">{{ $t("common.disable") }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          :label="$t('common.createTime')"
          align="center"
          prop="createTime"
        >
          <template #default="scope">
            {{ timestampToDateTime(scope.row.createTime) }}
          </template>
        </el-table-column>

        <el-table-column :label="$t('common.operate')" align="left" width="200">
          <template #default="scope">
            <el-button type="primary" link @click="handleUpdate(scope.row)"
              >{{ $t("common.edit") }}
            </el-button>
            <el-button type="danger" link @click="handleDelete(scope.row)"
              >{{ $t("common.delete") }}
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
      <el-form ref="dataFormRef" :model="formData" label-width="80px">
        <el-form-item :label="$t('account.username')" prop="username">
          <el-input
            v-model="formData.username"
            :placeholder="$t('account.username')"
            maxlength="50"
            clearable
          />
        </el-form-item>
        <el-form-item :label="$t('account.pass')" prop="pass">
          <el-input
            v-model="formData.pass"
            :placeholder="$t('account.pass')"
            maxlength="50"
            clearable
            type="password"
            show-password
          />
        </el-form-item>
        <el-form-item :label="$t('account.conPass')" prop="conPass">
          <el-input
            v-model="formData.conPass"
            :placeholder="$t('account.conPass')"
            maxlength="50"
            clearable
            type="password"
            show-password
          />
        </el-form-item>
        <el-form-item :label="$t('account.quota')" prop="quota">
          <unit-select
            v-model:quota="formData.quota"
            v-model:quotaTmp="quotaTmp"
          />
        </el-form-item>
        <el-form-item :label="$t('account.expireTime')" prop="expireTime">
          <el-date-picker
            v-model="formData.expireTime"
            type="datetime"
            :placeholder="$t('account.expireTime')"
            value-format="x"
            :shortcuts="shortcuts"
            clearable
          />
        </el-form-item>

        <el-form-item :label="$t('common.deleted')" prop="status">
          <el-radio-group v-model="formData.deleted">
            <el-radio :label="0">{{ $t("common.enable") }}</el-radio>
            <el-radio :label="1">{{ $t("common.disable") }}</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm"
            >{{ $t("common.confirm") }}
          </el-button>
          <el-button @click="closeDialog">{{ $t("common.cancel") }}</el-button>
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
  exportAccountApi,
} from "@/api/account";
import { Search, Plus, Refresh } from "@element-plus/icons-vue";
import {
  timestampToDateTime,
  getMonthLater,
  getWeekLater,
  getYearLater,
  calculateTimeDifference,
} from "@/utils/time";
import { formatBytes } from "@/utils/byte";

const queryFormRef = ref(ElForm); // 查询表单
const dataFormRef = ref(ElForm); // 用户表单

const state = reactive({
  // 遮罩层
  loading: true,
  // 总条数
  total: 0,
  records: [] as AccountVo[],
  dialog: {
    visible: false,
  } as DialogType,
  formData: {
    quota: 0,
    expireTime: getMonthLater(),
    deleted: 0,
  } as AccountForm,
  queryParams: {
    username: undefined,
    deleted: undefined,
    pageNum: 1,
    pageSize: 10,
  } as AccountPageDto,
  quotaTmp: 0,
});

const { loading, total, records, dialog, formData, queryParams, quotaTmp } =
  toRefs(state);

const shortcuts = [
  {
    text: "A week later",
    value: getWeekLater,
  },
  {
    text: "A month later",
    value: getMonthLater,
  },
  {
    text: "A year later",
    value: getYearLater,
  },
];

function resetFormData() {
  Object.assign(state.formData, {
    id: undefined,
    quota: 0,
    expireTime: getMonthLater(),
    deleted: 0,
  });
  quotaTmp.value = 0;
}

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
  const id = row.id;
  getAccountApi({ id: id }).then(({ data }) => {
    Object.assign(formData.value, data);
    quotaTmp.value = data.quota;
  });

  dialog.value = {
    title: "修改用户",
    visible: true,
  };
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

  if (dialog.value.title == "修改用户") {
    resetFormData();
  }
}

/**
 * 导出
 */
function handleExport() {
  exportAccountApi().then((res) => {
    const blob = new Blob([res.data], {
      type: "application/octet-stream",
    });
    let url = window.URL.createObjectURL(blob);
    let a = document.createElement("a");
    document.body.appendChild(a);
    a.href = url;
    let dis = res.headers["content-disposition"];
    a.download = dis.split("attachment; filename=")[1];
    // 模拟点击下载
    a.click();
    window.URL.revokeObjectURL(url);
    ElMessage.success("导出成功");
  });
}

onMounted(() => {
  // 初始化用户列表数据
  handleQuery();
});
</script>

<style lang="scss" scoped>
.github-corner {
  position: absolute;
  top: 0;
  right: 0;
  z-index: 1;
  border: 0;
}
</style>
