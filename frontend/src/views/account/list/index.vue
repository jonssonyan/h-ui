<template>
  <div class="app-container">
    <github-corner class="github-corner" />
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
          <el-upload
            v-model:file-list="fileList"
            :http-request="handleImport"
            :show-file-list="false"
            accept=".json"
            :limit="1"
            :before-upload="beforeImport"
          >
            <el-button>
              <template #icon>
                <i-ep-upload />
              </template>
              {{ $t("common.import") }}
            </el-button>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button @click="handleExport">
            <template #icon>
              <i-ep-download />
            </template>
            {{ $t("common.export") }}
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
        />
        <el-table-column
          key="username"
          :label="$t('account.username')"
          align="center"
          prop="username"
        />
        <el-table-column
          key="role"
          :label="$t('account.role')"
          align="center"
          prop="role"
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
            {{ formatBytes(scope.row.upload) }}
          </template>
        </el-table-column>
        <el-table-column
          key="online"
          :label="$t('account.onlineStatus')"
          align="center"
          prop="online"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.online" type="success"
              >{{ $t("account.online") }}
            </el-tag>
            <el-tag v-else type="info">{{ $t("account.offline") }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          key="device"
          :label="$t('account.device')"
          align="center"
          prop="device"
        />
        <el-table-column
          key="deviceNo"
          :label="$t('account.deviceNo')"
          align="center"
          prop="deviceNo"
        />
        <el-table-column
          key="kickUtilTime"
          :label="$t('account.kickUtilTimeLast')"
          align="center"
          prop="kickUtilTime"
        >
          <template #default="scope">
            {{ calculateTimeDifference(scope.row.kickUtilTime) }}
          </template>
        </el-table-column>
        <el-table-column
          key="expireTime"
          :label="$t('account.expireTime')"
          align="center"
          prop="expireTime"
        >
          <template #default="scope">
            {{ timestampToDateTime(scope.row.expireTime) }}
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
        <el-table-column
          key="deleted"
          :label="$t('common.deleted')"
          align="center"
          prop="deleted"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.deleted === 0" type="success"
              >{{ $t("common.enable") }}
            </el-tag>
            <el-tag v-else type="danger">{{ $t("common.disable") }}</el-tag>
          </template>
        </el-table-column>

        <el-table-column
          :label="$t('common.operate')"
          align="center"
          width="300"
        >
          <template #default="scope">
            <el-button type="primary" link @click="handleSubscribe(scope.row)"
              >{{ $t("common.subscribe") }}
            </el-button>
            <el-button type="primary" link @click="handleNodeUrl(scope.row)"
              >{{ $t("common.nodeUrl") }}
            </el-button>
            <el-button type="primary" link @click="handleQrCode(scope.row)">
              {{ $t("common.nodeQrCode") }}
            </el-button>
            <el-popconfirm
              title="Are you sure to reset traffic?"
              @confirm="resetTraffic(scope.row)"
            >
              <template #reference>
                <el-button type="primary" link
                  >{{ $t("common.resetTraffic") }}
                </el-button>
              </template>
            </el-popconfirm>
            <el-button type="primary" link @click="handleUpdate(scope.row)"
              >{{ $t("common.edit") }}
            </el-button>
            <el-button type="danger" link @click="handleDelete(scope.row)"
              >{{ $t("common.delete") }}
            </el-button>
            <el-button type="danger" link @click="handleKick(scope.row)"
              >{{ $t("account.kick") }}
            </el-button>
            <el-popconfirm
              :title="$t('account.releaseKickTip')"
              @confirm="confirmReleaseKick(scope.row)"
              v-if="calculateTimeDifference(scope.row.kickUtilTime) !== '-'"
            >
              <template #reference>
                <el-button type="danger" link
                  >{{ $t("account.releaseKick") }}
                </el-button>
              </template>
            </el-popconfirm>
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

    <el-dialog
      :title="dialog.title"
      v-model="dialog.visible"
      width="620px"
      append-to-body
      @close="closeDialog"
    >
      <el-form
        ref="dataFormRef"
        :rules="
          dialog.title === t('common.add')
            ? dataFormAddRules
            : dataFormUpdateRules
        "
        :model="dataForm"
        label-width="100px"
      >
        <el-form-item :label="$t('account.username')" prop="username">
          <el-input
            v-model="dataForm.username"
            :placeholder="$t('account.username')"
            maxlength="50"
            clearable
          />
        </el-form-item>
        <el-form-item :label="$t('account.pass')" prop="pass">
          <el-input
            v-model="dataForm.pass"
            :placeholder="$t('account.pass')"
            maxlength="50"
            clearable
            type="password"
            show-password
          />
        </el-form-item>
        <el-form-item :label="$t('account.conPass')" prop="conPass">
          <el-input
            v-model="dataForm.conPass"
            :placeholder="$t('account.conPass')"
            maxlength="50"
            clearable
            type="password"
            show-password
          />
        </el-form-item>
        <el-form-item :label="$t('account.quota')" prop="quota">
          <unit-select :setValue="setQuota" :valueTmp="quotaTmp" />
        </el-form-item>
        <el-form-item :label="$t('account.deviceNo')" prop="deviceNo">
          <el-input-number
            v-model="dataForm.deviceNo"
            :placeholder="$t('account.deviceNo')"
            :min="1"
            :controls="false"
            :precision="0"
            clearable
            style="width: 220px"
          />
        </el-form-item>
        <el-form-item :label="$t('account.expireTime')" prop="expireTime">
          <el-date-picker
            v-model="dataForm.expireTime"
            type="datetime"
            :placeholder="$t('account.expireTime')"
            value-format="x"
            :shortcuts="shortcuts"
            clearable
          />
        </el-form-item>
        <el-form-item :label="$t('common.deleted')" prop="deleted">
          <el-radio-group v-model="dataForm.deleted">
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

    <el-dialog
      :title="dialogKick.title"
      v-model="dialogKick.visible"
      width="600px"
      append-to-body
      @close="closeDialogKick"
    >
      <el-form ref="kickFormRef" :model="kickForm" label-width="100px">
        <el-form-item :label="$t('account.kickUtilTime')" prop="kickUtilTime">
          <el-date-picker
            v-model="kickForm.kickUtilTime"
            type="datetime"
            :placeholder="$t('account.kickUtilTime')"
            value-format="x"
            :shortcuts="shortcutsKick"
            clearable
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitKickForm"
            >{{ $t("common.confirm") }}
          </el-button>
          <el-button @click="closeDialogKick"
            >{{ $t("common.cancel") }}
          </el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      :title="qrCodeDialog.title"
      v-model="qrCodeDialog.visible"
      width="600px"
      append-to-body
      @close="qrCodeDialog.visible = false"
    >
      <el-form style="text-align: center">
        <el-image
          style="width: 300px; height: 300px"
          :src="qrCodeSrc"
        ></el-image>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="qrCodeDialog.visible = false"
            >{{ $t("common.confirm") }}
          </el-button>
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
  KickAccountForm,
} from "@/api/account/types";
import {
  saveAccountApi,
  deleteAccountApi,
  getAccountApi,
  pageAccountApi,
  updateAccountApi,
  exportAccountApi,
  releaseKickAccountApi,
  importAccountApi,
  resetTrafficApi,
} from "@/api/account";
import { Search, Plus, Refresh } from "@element-plus/icons-vue";
import {
  timestampToDateTime,
  getMonthLater,
  getWeekLater,
  getYearLater,
  calculateTimeDifference,
  getHourLater,
  getDayLater,
} from "@/utils/time";
import { formatBytes } from "@/utils/byte";

import {
  hysteria2KickApi,
  hysteria2SubscribeUrlApi,
  hysteria2UrlApi,
} from "@/api/hysteria2";
import {
  UploadFile,
  UploadRawFile,
  UploadRequestOptions,
} from "element-plus/lib/components";
import { useI18n } from "vue-i18n";
import {
  Hysteria2SubscribeUrlDto,
  Hysteria2UrlDto,
} from "@/api/hysteria2/types";
import copy from "copy-to-clipboard";

const { t } = useI18n();

const queryFormRef = ref(ElForm); // 查询表单
const dataFormRef = ref(ElForm); // 用户表单
const kickFormRef = ref(ElForm); // 下线表单

const dataFormAddRules = {
  username: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
    {
      pattern: /^[a-zA-Z0-9!@#$%^&*()_+-=]{6,32}$/,
      message: "Username format is incorrect",
      trigger: ["change", "blur"],
    },
  ],
  pass: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
    {
      pattern: /^[a-zA-Z0-9!@#$%^&*()_+-=]{6,32}$/,
      message: "Pass format is incorrect",
      trigger: ["change", "blur"],
    },
  ],
  conPass: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
    {
      pattern: /^[a-zA-Z0-9!@#$%^&*()_+-=]{6,32}$/,
      message: "ConPass format is incorrect",
      trigger: ["change", "blur"],
    },
  ],
  expireTime: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
  deviceNo: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
  deleted: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
};

const dataFormUpdateRules = {
  username: [
    {
      pattern: /^[a-zA-Z0-9!@#$%^&*()_+-=]{6,32}$/,
      message: "Username format is incorrect",
      trigger: ["change", "blur"],
    },
  ],
  pass: [
    {
      pattern: /^[a-zA-Z0-9!@#$%^&*()_+-=]{6,32}$/,
      message: "Pass format is incorrect",
      trigger: ["change", "blur"],
    },
  ],
  conPass: [
    {
      pattern: /^[a-zA-Z0-9!@#$%^&*()_+-=]{6,32}$/,
      message: "Pass format is incorrect",
      trigger: ["change", "blur"],
    },
  ],
};

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

const shortcutsKick = [
  {
    text: "A hour later",
    value: getHourLater,
  },
  {
    text: "A day later",
    value: getDayLater,
  },
  {
    text: "A week later",
    value: getWeekLater,
  },
  {
    text: "A month later",
    value: getMonthLater,
  },
];

const state = reactive({
  loading: true,
  total: 0,
  records: [] as AccountVo[],
  dialog: {
    visible: false,
  } as DialogType,
  dialogKick: {
    visible: false,
  } as DialogType,
  dataForm: {
    quota: 0,
    expireTime: getMonthLater(),
    deviceNo: 6,
    deleted: 0,
  } as AccountForm,
  kickForm: {
    kickUtilTime: getHourLater(),
  } as KickAccountForm,
  queryParams: {
    username: undefined,
    deleted: undefined,
    pageNum: 1,
    pageSize: 10,
  } as AccountPageDto,
  quotaTmp: 0,
  fileList: [] as UploadFile[],
  qrCodeDialog: {
    title: "QR Code",
    visible: false,
  } as DialogType,
  qrCodeSrc: "",
});

const {
  loading,
  total,
  records,
  dialog,
  dialogKick,
  dataForm,
  kickForm,
  queryParams,
  quotaTmp,
  fileList,
  qrCodeDialog,
  qrCodeSrc,
} = toRefs(state);

const resetDataForm = () => {
  Object.assign(state.dataForm, {
    id: undefined,
    quota: 0,
    expireTime: getMonthLater(),
    deleted: 0,
  });
  quotaTmp.value = 0;
};

/**
 * 查询
 */
const handleQuery = async () => {
  state.loading = true;
  try {
    const { data } = await pageAccountApi(state.queryParams);
    state.records = data.records;
    state.total = data.total;
  } finally {
    state.loading = false;
  }
};

/**
 * 重置
 */
const resetQuery = () => {
  queryFormRef.value.resetFields();
  handleQuery();
};

/**
 * 保存
 **/
const handleAdd = () => {
  state.dialog = {
    title: t("common.add"),
    visible: true,
  };
};

/**
 * 修改
 **/
const handleUpdate = async (row: { [key: string]: any }) => {
  const id = row.id;
  const { data } = await getAccountApi({ id: id });
  Object.assign(state.dataForm, data);
  quotaTmp.value = data.quota;
  dialog.value = {
    title: t("common.update"),
    visible: true,
  };
};

const setQuota = (newQuota: number) => {
  state.dataForm.quota = newQuota;
};

/**
 * 表单提交
 */
const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      const accountId = state.dataForm.id;
      let accountUpdateDto: AccountUpdateDto = { ...state.dataForm };
      if (accountId) {
        updateAccountApi(accountUpdateDto).then(() => {
          ElMessage.success(t("common.success"));
          closeDialog();
          handleQuery();
        });
      } else {
        saveAccountApi(accountUpdateDto).then(() => {
          ElMessage.success(t("common.success"));
          closeDialog();
          handleQuery();
        });
      }
    }
  });
};

/**
 * 下线表单提交
 */
const submitKickForm = () => {
  kickFormRef.value.validate((valid: any) => {
    if (valid) {
      const params = { ...state.kickForm };
      hysteria2KickApi(params).then(() => {
        ElMessage.success(t("common.success"));
        closeDialogKick();
        handleQuery();
      });
    }
  });
};

/**
 * 删除
 */
const handleDelete = (row: { [key: string]: any }) => {
  const id = row.id;
  const username = row.username;
  ElMessageBox.confirm(
    "Are you sure to delete the data item with the username「" +
      username +
      "」?",
    "Warning",
    {
      confirmButtonText: t("common.confirm"),
      cancelButtonText: t("common.cancel"),
      type: "warning",
    }
  )
    .then(() => {
      deleteAccountApi({ id: id }).then(() => {
        ElMessage.success(t("common.success"));
        handleQuery();
      });
    })
    .catch(() => ElMessage.info(t("common.cancel")));
};

/**
 * 强制用户下线
 * @param row
 */
const handleKick = (row: { [key: string]: any }) => {
  state.kickForm.ids = [row.id];
  dialogKick.value = {
    title: t("account.kickTip"),
    visible: true,
  };
};

/**
 * 解除下线状态
 * @param row
 */
const confirmReleaseKick = (row: { [key: string]: any }) => {
  releaseKickAccountApi({ id: row.id }).then(() => {
    ElMessage.success(t("account.releaseSuccess"));
    handleQuery();
  });
};

/**
 * 关闭用户弹窗
 */
const closeDialog = () => {
  dialog.value.visible = false;
  dataFormRef.value.resetFields();
  dataFormRef.value.clearValidate();

  if (dialog.value.title == t("common.update")) {
    resetDataForm();
  }
};

/**
 * 关闭下线弹窗
 */
const closeDialogKick = () => {
  dialogKick.value.visible = false;
  kickFormRef.value.resetFields();
  kickFormRef.value.clearValidate();
};

/**
 * 导入
 */
const handleImport = (params: UploadRequestOptions) => {
  if (state.fileList.length > 0) {
    let formData = new FormData();
    formData.append("file", params.file);
    importAccountApi(formData).then(() => {
      ElMessage.success(t("common.success"));
    });
    state.fileList = [];
  }
};

const beforeImport = (file: UploadRawFile) => {
  if (!file.name.endsWith(".json")) {
    ElMessage.error("file format not supported");
    return false;
  }
  if (file.size / 1024 / 1024 > 2) {
    ElMessage.error("the file is too big, less than 2 MB");
    return false;
  }
};

/**
 * 导出
 */
const handleExport = () => {
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
    ElMessage.success(t("common.success"));
  });
};

const handleSubscribe = async (row: { [key: string]: any }) => {
  try {
    const dto: Hysteria2SubscribeUrlDto = {
      accountId: row.id,
      protocol: window.location.protocol,
      host: window.location.host,
    };
    const { data } = await hysteria2SubscribeUrlApi(dto);
    copy(data.url);
    ElMessage.success(t("common.copySuccess"));
  } catch (e) {
    /* empty */
  }
};

const handleNodeUrl = async (row: { [key: string]: any }) => {
  try {
    const dto: Hysteria2UrlDto = {
      accountId: row.id,
      hostname: window.location.hostname,
    };
    const { data } = await hysteria2UrlApi(dto);
    copy(data.url);
    ElMessage.success(t("common.copySuccess"));
  } catch (e) {
    /* empty */
  }
};

const handleQrCode = async (row: { [key: string]: any }) => {
  try {
    const dto: Hysteria2UrlDto = {
      accountId: row.id,
      hostname: window.location.hostname,
    };
    const { data } = await hysteria2UrlApi(dto);
    state.qrCodeSrc = "data:image/png;base64," + data.qrCode;
    state.qrCodeDialog.visible = true;
  } catch (e) {
    /* empty */
  }
};

const resetTraffic = async (row: { [key: string]: any }) => {
  try {
    await resetTrafficApi({ id: row.id });
    ElMessage.success(t("common.success"));
    await handleQuery();
  } catch (e) {
    /* empty */
  }
};

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
