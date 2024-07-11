<template>
  <div class="app-container">
    <div class="search">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" @click="submitForm" :icon="Select">
            {{ $t("common.save") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-upload
            v-model:file-list="fileList"
            :http-request="handleImport"
            :show-file-list="false"
            accept=".yaml"
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
      <el-form ref="enableFormRef" :model="enableForm" inline>
        <el-tooltip :content="$t('hysteria.config.enable')" placement="bottom">
          <el-form-item prop="enable">
            <el-switch
              v-model="enableForm.enable"
              active-value="1"
              inactive-value="0"
              @change="handleChangeEnable"
              :active-text="$t('hysteria.enable')"
              :inactive-text="$t('hysteria.disable')"
              style="height: 32px"
            />
          </el-form-item>
        </el-tooltip>
        <el-form-item>
          <el-dropdown @command="handleDropdownClick">
            <el-button
              type="primary"
              :icon="CirclePlusFilled"
              style="height: 32px"
            >
              {{ $t("hysteria.addConfigItem") }}
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item
                  :key="item"
                  v-for="item in tabs"
                  :command="item.name"
                  >{{ item.desc }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-form-item>
        <el-form-item>
          <div style="display: flex; align-items: center">
            <el-select
              v-model="hysteria2Version"
              :placeholder="$t('hysteria.hysteria2Version')"
              style="height: 32px; width: 160px"
            >
              <el-option
                v-for="item in hysteria2Versions"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
            <el-button
              type="primary"
              @click="handleHysteria2ChangeVersion"
              style="height: 32px"
              >{{ $t("hysteria.hysteria2ChangeVersion") }}
            </el-button>
          </div>
        </el-form-item>
        <el-form-item>
          <el-tag style="height: 32px">
            {{ $t("hysteria.hysteria2Version") }}:
            {{ hysteria2Monitor.version ? hysteria2Monitor.version : "-" }}
          </el-tag>
        </el-form-item>
        <el-form-item>
          <el-tag
            :type="hysteria2Monitor.running ? 'success' : 'danger'"
            style="height: 32px"
          >
            {{ $t("hysteria.hysteria2Running") }}:
            {{
              hysteria2Monitor.running
                ? $t("monitor.hysteria2RunningTrue")
                : $t("monitor.hysteria2RunningFalse")
            }}
          </el-tag>
        </el-form-item>
      </el-form>

      <el-form
        ref="dataFormRef"
        :rules="dataFormRules"
        label-position="top"
        :model="dataForm"
      >
        <el-tabs
          v-model="activeName"
          class="tabs"
          closable
          @tab-remove="closeTabPane"
        >
          <el-tab-pane :label="$t('hysteria.extension')" name="extension">
            <el-tooltip
              :content="$t('hysteria.config.remark')"
              placement="bottom"
            >
              <el-form-item label="remark" prop="remark">
                <el-input v-model="configForm.remark" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.portHopping')"
              placement="bottom"
            >
              <el-form-item label="portHopping" prop="portHopping">
                <el-input v-model="configForm.portHopping" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.listen')" name="listen">
            <el-tooltip
              :content="$t('hysteria.config.listen')"
              placement="bottom"
            >
              <el-form-item label="listen" prop="listen">
                <el-input v-model="dataForm.listen" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.tls')" name="tls">
            <el-tooltip
              :content="$t('hysteria.config.tlsType')"
              placement="bottom"
            >
              <el-form-item label="tls.type" prop="tls.type">
                <el-select v-model="tlsType" style="width: 100%">
                  <el-option
                    v-for="item in tlsTypes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'tls'"
              :content="$t('hysteria.config.tls.cert')"
              placement="bottom"
            >
              <el-form-item label="cert" prop="tls.cert">
                <el-input v-model="dataForm.tls.cert" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'tls'"
              :content="$t('hysteria.config.tls.key')"
              placement="bottom"
            >
              <el-form-item label="key" prop="tls.key">
                <el-input v-model="dataForm.tls.key" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.domains')"
              placement="bottom"
            >
              <el-form-item label="acme.domains" prop="acme.domains">
                <imputMultiple :tags="dataForm.acme.domains" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.email')"
              placement="bottom"
            >
              <el-form-item label="acme.email" prop="acme.email">
                <el-input v-model="dataForm.acme.email" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.ca')"
              placement="bottom"
            >
              <el-form-item label="acme.ca" prop="acme.ca">
                <el-select v-model="dataForm.acme.ca" style="width: 100%">
                  <el-option
                    v-for="item in acmeCas"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.listenHost')"
              placement="bottom"
            >
              <el-form-item label="acme.listenHost" prop="acme.listenHost">
                <el-input v-model="dataForm.acme.listenHost" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.dir')"
              placement="bottom"
            >
              <el-form-item label="acme.dir" prop="acme.dir">
                <el-input v-model="dataForm.acme.dir" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.type')"
              placement="bottom"
            >
              <el-form-item label="acme.type" prop="acme.type">
                <el-select
                  v-model="dataForm.acme.type"
                  style="width: 100%"
                  clearable
                >
                  <el-option
                    v-for="item in acmeTypes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && dataForm.acme.type === 'http'"
              :content="$t('hysteria.config.acme.http.altPort')"
              placement="bottom"
            >
              <el-form-item label="acme.http.altPort" prop="acme.http.altPort">
                <el-input v-model="dataForm.acme.http.altPort" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && dataForm.acme.type === 'tls'"
              :content="$t('hysteria.config.acme.tls.altPort')"
              placement="bottom"
            >
              <el-form-item label="acme.tls.altPort" prop="acme.tls.altPort">
                <el-input v-model="dataForm.acme.tls.altPort" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && dataForm.acme.type === 'dns'"
              :content="$t('hysteria.config.acme.dns.name')"
              placement="bottom"
            >
              <el-form-item label="acme.dns.name" prop="acme.dns.name">
                <el-select v-model="dataForm.acme.dns.name" style="width: 100%">
                  <el-option
                    v-for="item in dnsNames"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && dataForm.acme.type === 'dns'"
              :content="$t('hysteria.config.acme.dns.config')"
              placement="bottom"
            >
              <el-form-item label="acme.dns.config" prop="acme.dns.config">
                <map-add :map-object="dataForm.acme.dns.config" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && !dataForm.acme.type"
              :content="$t('hysteria.config.acme.disableHTTP')"
              placement="bottom"
            >
              <el-form-item label="acme.disableHTTP" prop="acme.disableHTTP">
                <el-switch v-model="dataForm.acme.disableHTTP" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && !dataForm.acme.type"
              :content="$t('hysteria.config.acme.disableTLSALPN')"
              placement="bottom"
            >
              <el-form-item
                label="acme.disableTLSALPN"
                prop="acme.disableTLSALPN"
              >
                <el-switch v-model="dataForm.acme.disableTLSALPN" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && !dataForm.acme.type"
              :content="$t('hysteria.config.acme.altHTTPPort')"
              placement="bottom"
            >
              <el-form-item label="acme.altHTTPPort" prop="acme.altHTTPPort">
                <el-input v-model="dataForm.acme.altHTTPPort" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme' && !dataForm.acme.type"
              :content="$t('hysteria.config.acme.altTLSALPNPort')"
              placement="bottom"
            >
              <el-form-item
                label="acme.altTLSALPNPort"
                prop="acme.altTLSALPNPort"
              >
                <el-input v-model="dataForm.acme.altTLSALPNPort" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.obfs')" name="obfs" v-if="obfs">
            <el-tooltip
              :content="$t('hysteria.config.obfs.type')"
              placement="bottom"
            >
              <el-form-item label="obfs.type" prop="obfs.type">
                <el-select
                  v-model="dataForm.obfs.type"
                  style="width: 100%"
                  clearable
                >
                  <el-option
                    v-for="item in obfsTypes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.obfs.type === 'salamander'"
              :content="$t('hysteria.config.obfs.salamander.password')"
              placement="bottom"
            >
              <el-form-item
                label="obfs.salamander.password"
                prop="obfs.salamander.password"
              >
                <el-input
                  v-model="dataForm.obfs.salamander.password"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.quic')" name="quic" v-if="quic">
            <el-tooltip
              :content="$t('hysteria.config.quic.initStreamReceiveWindow')"
              placement="bottom"
            >
              <el-form-item
                label="quic.initStreamReceiveWindow"
                prop="quic.initStreamReceiveWindow"
              >
                <el-input
                  v-model="dataForm.quic.initStreamReceiveWindow"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.quic.maxStreamReceiveWindow')"
              placement="bottom"
            >
              <el-form-item
                label="quic.maxStreamReceiveWindow"
                prop="quic.maxStreamReceiveWindow"
              >
                <el-input
                  v-model="dataForm.quic.maxStreamReceiveWindow"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.quic.initConnReceiveWindow')"
              placement="bottom"
            >
              <el-form-item
                label="quic.initConnReceiveWindow"
                prop="quic.initConnReceiveWindow"
              >
                <el-input
                  v-model="dataForm.quic.initConnReceiveWindow"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.quic.maxConnReceiveWindow')"
              placement="bottom"
            >
              <el-form-item
                label="quic.maxConnReceiveWindow"
                prop="quic.maxConnReceiveWindow"
              >
                <el-input
                  v-model="dataForm.quic.maxConnReceiveWindow"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.quic.maxIdleTimeout')"
              placement="bottom"
            >
              <el-form-item
                label="quic.maxIdleTimeout"
                prop="quic.maxIdleTimeout"
              >
                <el-input v-model="dataForm.quic.maxIdleTimeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.quic.maxIncomingStreams')"
              placement="bottom"
            >
              <el-form-item
                label="quic.maxIncomingStreams"
                prop="quic.maxIncomingStreams"
              >
                <el-input
                  v-model="dataForm.quic.maxIncomingStreams"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.quic.disablePathMTUDiscovery')"
              placement="bottom"
            >
              <el-form-item
                label="quic.disablePathMTUDiscovery"
                prop="quic.disablePathMTUDiscovery"
              >
                <el-switch v-model="dataForm.quic.disablePathMTUDiscovery" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane
            :label="$t('hysteria.bandwidth')"
            name="bandwidth"
            v-if="bandwidth"
          >
            <el-tooltip
              :content="$t('hysteria.config.bandwidth.up')"
              placement="bottom"
              :trigger-keys="[]"
            >
              <el-form-item label="bandwidth.up" prop="bandwidth.up">
                <el-input v-model="dataForm.bandwidth.up" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.bandwidth.down')"
              placement="bottom"
              :trigger-keys="[]"
            >
              <el-form-item label="bandwidth.down" prop="bandwidth.down">
                <el-input v-model="dataForm.bandwidth.down" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.ignoreClientBandwidth')"
              placement="bottom"
            >
              <el-form-item
                label="ignoreClientBandwidth"
                prop="ignoreClientBandwidth"
              >
                <el-switch v-model="dataForm.ignoreClientBandwidth" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane
            :label="$t('hysteria.speedTest')"
            name="speedTest"
            v-if="speedTest"
          >
            <el-tooltip
              :content="$t('hysteria.config.speedTest')"
              placement="bottom"
            >
              <el-form-item label="speedTest" prop="speedTest">
                <el-switch v-model="dataForm.speedTest" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.udp')" name="udp" v-if="udp">
            <el-tooltip
              :content="$t('hysteria.config.disableUDP')"
              placement="bottom"
            >
              <el-form-item label="disableUDP" prop="disableUDP">
                <el-switch v-model="dataForm.disableUDP" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.udpIdleTimeout')"
              placement="bottom"
            >
              <el-form-item label="udpIdleTimeout" prop="udpIdleTimeout">
                <el-input v-model="dataForm.udpIdleTimeout" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane
            :label="$t('hysteria.resolver')"
            name="resolver"
            v-if="resolver"
          >
            <el-tooltip
              :content="$t('hysteria.config.resolver.type')"
              placement="bottom"
            >
              <el-form-item label="resolver.type" prop="resolver.type">
                <el-select v-model="dataForm.resolver.type" style="width: 100%">
                  <el-option
                    v-for="item in resolverTypes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'tcp'"
              :content="$t('hysteria.config.resolver.tcp.addr')"
              placement="bottom"
            >
              <el-form-item label="resolver.tcp.addr" prop="resolver.tcp.addr">
                <el-input v-model="dataForm.resolver.tcp.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'tcp'"
              :content="$t('hysteria.config.resolver.tcp.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.tcp.timeout"
                prop="resolver.tcp.timeout"
              >
                <el-input v-model="dataForm.resolver.tcp.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'udp'"
              :content="$t('hysteria.config.resolver.udp.addr')"
              placement="bottom"
            >
              <el-form-item label="resolver.udp.addr" prop="resolver.udp.addr">
                <el-input v-model="dataForm.resolver.udp.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'udp'"
              :content="$t('hysteria.config.resolver.udp.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.udp.timeout"
                prop="resolver.udp.timeout"
              >
                <el-input v-model="dataForm.resolver.udp.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.addr')"
              placement="bottom"
            >
              <el-form-item label="resolver.tls.addr" prop="resolver.tls.addr">
                <el-input v-model="dataForm.resolver.tls.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.tls.timeout"
                prop="resolver.tls.timeout"
              >
                <el-input v-model="dataForm.resolver.tls.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.sni')"
              placement="bottom"
            >
              <el-form-item label="resolver.tls.sni" prop="resolver.tls.sni">
                <el-input v-model="dataForm.resolver.tls.sni" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.insecure')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.tls.insecure"
                prop="resolver.tls.insecure"
              >
                <el-switch v-model="dataForm.resolver.tls.insecure" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.addr')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.addr"
                prop="resolver.https.addr"
              >
                <el-input v-model="dataForm.resolver.https.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.timeout"
                prop="resolver.https.timeout"
              >
                <el-input v-model="dataForm.resolver.https.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.sni')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.sni"
                prop="resolver.https.sni"
              >
                <el-input v-model="dataForm.resolver.https.sni" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.insecure')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.insecure"
                prop="resolver.https.insecure"
              >
                <el-switch v-model="dataForm.resolver.https.insecure" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.sniff')" name="sniff" v-if="sniff">
            <el-tooltip
              :content="$t('hysteria.config.sniff.enable')"
              placement="bottom"
            >
              <el-form-item label="sniff.enable" prop="sniff.enable">
                <el-switch v-model="dataForm.sniff.enable" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.sniff.timeout')"
              placement="bottom"
            >
              <el-form-item label="sniff.timeout" prop="sniff.timeout">
                <el-input v-model="dataForm.sniff.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.sniff.rewriteDomain')"
              placement="bottom"
            >
              <el-form-item
                label="sniff.rewriteDomain"
                prop="sniff.rewriteDomain"
              >
                <el-switch v-model="dataForm.sniff.rewriteDomain" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.sniff.tcpPorts')"
              placement="bottom"
            >
              <el-form-item label="sniff.tcpPorts" prop="sniff.tcpPorts">
                <el-input v-model="dataForm.sniff.tcpPorts" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.sniff.udpPorts')"
              placement="bottom"
            >
              <el-form-item label="sniff.udpPorts" prop="sniff.udpPorts">
                <el-input v-model="dataForm.sniff.udpPorts" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.acl')" name="acl" v-if="acl">
            <el-tooltip
              :content="$t('hysteria.config.aclType')"
              placement="bottom"
            >
              <el-form-item label="acl.type" prop="acl.type">
                <el-select v-model="aclType" style="width: 100%" clearable>
                  <el-option
                    v-for="item in aclTypes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="aclType === 'file'"
              :content="$t('hysteria.config.acl.file')"
              placement="bottom"
            >
              <el-form-item label="acl.file" prop="acl.file">
                <el-input v-model="dataForm.acl.file" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="aclType === 'inline'"
              :content="$t('hysteria.config.acl.inline')"
              placement="bottom"
            >
              <el-form-item label="acl.inline" prop="acl.inline">
                <imputMultiple :tags="dataForm.acl.inline" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.acl.geoip')"
              placement="bottom"
            >
              <el-form-item label="acl.geoip" prop="acl.geoip">
                <el-input v-model="dataForm.acl.geoip" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.acl.geosite')"
              placement="bottom"
            >
              <el-form-item label="acl.geosite" prop="acl.geosite">
                <el-input v-model="dataForm.acl.geosite" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.acl.geoUpdateInterval')"
              placement="bottom"
            >
              <el-form-item
                label="acl.geoUpdateInterval"
                prop="acl.geoUpdateInterval"
              >
                <el-input v-model="dataForm.acl.geoUpdateInterval" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane
            :label="$t('hysteria.outbounds')"
            name="outbounds"
            v-if="outbounds"
          >
            <Outbounds :outbounds="dataForm.outbounds" />
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.http')" name="http">
            <el-tooltip
              :content="$t('hysteria.config.trafficStats.listen')"
              placement="bottom"
            >
              <el-form-item
                label="trafficStats.listen"
                prop="trafficStats.listen"
              >
                <el-input v-model="dataForm.trafficStats.listen" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane
            :label="$t('hysteria.masquerade')"
            name="masquerade"
            v-if="masquerade"
          >
            <el-tooltip
              :content="$t('hysteria.config.masquerade.type')"
              placement="bottom"
            >
              <el-form-item label="masquerade.type" prop="masquerade.type">
                <el-select
                  v-model="dataForm.masquerade.type"
                  style="width: 100%"
                  clearable
                >
                  <el-option
                    v-for="item in masqueradeTypes"
                    :key="item"
                    :label="item"
                    :value="item"
                  />
                </el-select>
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.masquerade.type === 'file'"
              :content="$t('hysteria.config.masquerade.file.dir')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.file.dir"
                prop="masquerade.file.dir"
              >
                <el-input v-model="dataForm.masquerade.file.dir" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.masquerade.type === 'proxy'"
              :content="$t('hysteria.config.masquerade.proxy.url')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.proxy.url"
                prop="masquerade.proxy.url"
              >
                <el-input v-model="dataForm.masquerade.proxy.url" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.masquerade.type === 'proxy'"
              :content="$t('hysteria.config.masquerade.proxy.rewriteHost')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.proxy.rewriteHost"
                prop="masquerade.proxy.rewriteHost"
              >
                <el-switch v-model="dataForm.masquerade.proxy.rewriteHost" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.masquerade.type === 'string'"
              :content="$t('hysteria.config.masquerade.string.content')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.string.content"
                prop="masquerade.string.content"
              >
                <el-input
                  v-model="dataForm.masquerade.string.content"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.masquerade.type === 'string'"
              :content="$t('hysteria.config.masquerade.string.headers')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.string.headers"
                prop="masquerade.string.headers"
              >
                <map-add :map-object="dataForm.masquerade.string.headers" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="dataForm.masquerade.type === 'string'"
              :content="$t('hysteria.config.masquerade.string.statusCode')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.string.statusCode"
                prop="masquerade.string.statusCode"
              >
                <el-switch v-model="dataForm.masquerade.string.statusCode" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.masquerade.listenHTTP')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.listenHTTP"
                prop="masquerade.listenHTTP"
              >
                <el-input v-model="dataForm.masquerade.listenHTTP" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.masquerade.listenHTTPS')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.listenHTTPS"
                prop="masquerade.listenHTTPS"
              >
                <el-input v-model="dataForm.masquerade.listenHTTPS" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.masquerade.forceHTTPS')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.forceHTTPS"
                prop="masquerade.forceHTTPS"
              >
                <el-switch v-model="dataForm.masquerade.forceHTTPS" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </el-card>
  </div>
</template>

<script lang="ts">
export default {
  name: "index",
};
</script>

<script setup lang="ts">
import {
  ConfigsUpdateDto,
  defaultHysteria2ServerConfig,
  Hysteria2ServerConfig,
  Tab,
} from "@/api/config/types";
import Outbounds from "./components/Outbounds/index.vue";
import { CirclePlusFilled, Select } from "@element-plus/icons-vue";
import {
  exportHysteria2ConfigApi,
  getHysteria2ConfigApi,
  importHysteria2ConfigApi,
  listConfigApi,
  updateConfigsApi,
  updateHysteria2ConfigApi,
} from "@/api/config";
import { useI18n } from "vue-i18n";
import { assignWith, deepCopy } from "@/utils/copy";
import {
  UploadFile,
  UploadRawFile,
  UploadRequestOptions,
} from "element-plus/lib/components";
import { hysteria2ChangeVersionApi, listReleaseApi } from "@/api/hysteria2";
import { monitorHysteria2Api } from "@/api/monitor";

const { t } = useI18n();

const hysteria2EnableKey = "HYSTERIA2_ENABLE";
const hysteria2Remark = "HYSTERIA2_CONFIG_REMARK";
const hysteria2ConfigPortHopping = "HYSTERIA2_CONFIG_PORT_HOPPING";
const dataFormRef = ref(ElForm);

const dataFormRules = {
  listen: [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
  "trafficStats.listen": [
    {
      required: true,
      message: "Required",
      trigger: ["change", "blur"],
    },
  ],
};

const tlsTypes = ref<string[]>(["tls", "acme"]);
const aclTypes = ref<string[]>(["file", "inline"]);
const acmeCas = ref<string[]>(["zerossl", "letsencrypt"]);
const acmeTypes = ref<string[]>(["http", "tls", "dns"]);
const dnsNames = ref<string[]>([
  "cloudflare",
  "duckdns",
  "gandi",
  "godaddy",
  "namedotcom",
  "vultr",
]);
const obfsTypes = ref<string[]>(["salamander"]);
const resolverTypes = ref<string[]>(["tcp", "udp", "tls", "https"]);
const masqueradeTypes = ref<string[]>(["file", "proxy", "string"]);

const state = reactive({
  enableForm: {
    enable: "0",
  },
  configForm: {
    remark: "",
    portHopping: "",
  },
  dataForm: { ...defaultHysteria2ServerConfig } as Hysteria2ServerConfig,
  activeName: "extension",
  tlsType: "acme",
  aclType: "inline",
  obfs: false,
  quic: false,
  bandwidth: false,
  speedTest: false,
  udp: false,
  resolver: false,
  sniff: false,
  acl: false,
  outbounds: false,
  masquerade: false,
  fileList: [] as UploadFile[],
  hysteria2Version: "",
  hysteria2Versions: [] as string[],
  hysteria2Monitor: {
    version: "",
    running: false,
  },
});

const {
  enableForm,
  configForm,
  activeName,
  dataForm,
  tlsType,
  aclType,
  obfs,
  quic,
  bandwidth,
  speedTest,
  udp,
  resolver,
  sniff,
  acl,
  outbounds,
  masquerade,
  fileList,
  hysteria2Version,
  hysteria2Versions,
  hysteria2Monitor,
} = toRefs(state);

const tabs = computed(() => {
  let tabs: Tab[] = [];
  if (!state.obfs) {
    tabs.push({
      name: "obfs",
      desc: t("hysteria.obfs"),
    });
  }
  if (!state.quic) {
    tabs.push({
      name: "quic",
      desc: t("hysteria.quic"),
    });
  }
  if (!state.bandwidth) {
    tabs.push({
      name: "bandwidth",
      desc: t("hysteria.bandwidth"),
    });
  }
  if (!state.speedTest) {
    tabs.push({
      name: "speedTest",
      desc: t("hysteria.speedTest"),
    });
  }
  if (!state.udp) {
    tabs.push({
      name: "udp",
      desc: t("hysteria.udp"),
    });
  }
  if (!state.resolver) {
    tabs.push({
      name: "resolver",
      desc: t("hysteria.resolver"),
    });
  }
  if (!state.sniff) {
    tabs.push({
      name: "sniff",
      desc: t("hysteria.sniff"),
    });
  }
  if (!state.acl) {
    tabs.push({
      name: "acl",
      desc: t("hysteria.acl"),
    });
  }
  if (!state.outbounds) {
    tabs.push({
      name: "outbounds",
      desc: t("hysteria.outbounds"),
    });
  }
  if (!state.masquerade) {
    tabs.push({
      name: "masquerade",
      desc: t("hysteria.masquerade"),
    });
  }
  return tabs;
});

const handleImport = (params: UploadRequestOptions) => {
  if (state.fileList.length > 0) {
    let formData = new FormData();
    formData.append("file", params.file);
    importHysteria2ConfigApi(formData).then(() => {
      ElMessage.success(t("common.success"));
    });
    state.fileList = [];
  }
};

const beforeImport = (file: UploadRawFile) => {
  if (!file.name.endsWith(".yaml")) {
    ElMessage.error("file format not supported");
    return false;
  }
  if (file.size / 1024 / 1024 > 2) {
    ElMessage.error("the file is too big, less than 2 MB");
    return false;
  }
};

const handleExport = async () => {
  let response = await exportHysteria2ConfigApi();
  try {
    const blob = new Blob([response.data], {
      type: "application/octet-stream",
    });
    let url = window.URL.createObjectURL(blob);
    let a = document.createElement("a");
    document.body.appendChild(a);
    a.href = url;
    let dis = response.headers["content-disposition"];
    a.download = dis.split("attachment; filename=")[1];
    a.click();
    window.URL.revokeObjectURL(url);
    ElMessage.success(t("common.success"));
  } catch (e) {
    /* empty */
  }
};

const handleChangeEnable = async () => {
  const enable = state.enableForm.enable === "1" ? "0" : "1";
  try {
    await updateConfigsApi({
      configUpdateDtos: [
        { key: hysteria2EnableKey, value: state.enableForm.enable },
      ],
    });
    if (state.enableForm.enable === "1") {
      ElMessage.success(t("common.enableSuccess"));
    } else {
      ElMessage.error(t("common.disableSuccess"));
    }
  } catch (e) {
    state.enableForm.enable = enable;
  }
  await setHysteria2Monitor();
};

const submitForm = () => {
  dataFormRef.value.validate((valid: any) => {
    if (valid) {
      let configs: ConfigsUpdateDto[] = [
        {
          key: hysteria2Remark,
          value: state.configForm.remark,
        },
        {
          key: hysteria2ConfigPortHopping,
          value: state.configForm.portHopping,
        },
      ];
      updateConfigsApi({ configUpdateDtos: configs });

      if (state.tlsType == "tls") {
        state.dataForm.acme = undefined;
      } else if (state.tlsType == "acme") {
        state.dataForm.tls = undefined;
        if (state.dataForm.acme?.type == "http") {
          state.dataForm.acme.tls = undefined;
          state.dataForm.acme.dns = undefined;
        } else if (state.dataForm.acme?.type == "tls") {
          state.dataForm.acme.http = undefined;
          state.dataForm.acme.dns = undefined;
        } else if (state.dataForm.acme?.type == "dns") {
          state.dataForm.acme.http = undefined;
          state.dataForm.acme.tls = undefined;
        } else if (state.dataForm.acme?.type == "") {
          state.dataForm.acme.http = undefined;
          state.dataForm.acme.tls = undefined;
          state.dataForm.acme.dns = undefined;
        }
      }

      if (state.dataForm.acl) {
        if (state.aclType == "inline") {
          state.dataForm.acl.file = undefined;
        } else if (state.aclType == "file") {
          state.dataForm.acl.inline = undefined;
        }
      }

      if (state.dataForm.resolver?.type == "tcp") {
        state.dataForm.resolver.udp = undefined;
        state.dataForm.resolver.tls = undefined;
        state.dataForm.resolver.https = undefined;
      } else if (state.dataForm.resolver?.type == "udp") {
        state.dataForm.resolver.tcp = undefined;
        state.dataForm.resolver.tls = undefined;
        state.dataForm.resolver.https = undefined;
      } else if (state.dataForm.resolver?.type == "tls") {
        state.dataForm.resolver.tcp = undefined;
        state.dataForm.resolver.udp = undefined;
        state.dataForm.resolver.https = undefined;
      } else if (state.dataForm.resolver?.type == "https") {
        state.dataForm.resolver.tcp = undefined;
        state.dataForm.resolver.udp = undefined;
        state.dataForm.resolver.tls = undefined;
      }

      if (state.dataForm.masquerade?.type == "file") {
        state.dataForm.masquerade.proxy = undefined;
        state.dataForm.masquerade.string = undefined;
      } else if (state.dataForm.masquerade?.type == "proxy") {
        state.dataForm.masquerade.file = undefined;
        state.dataForm.masquerade.string = undefined;
      } else if (state.dataForm.masquerade?.type == "string") {
        state.dataForm.masquerade.file = undefined;
        state.dataForm.masquerade.proxy = undefined;
      }

      if (!state.obfs) {
        state.dataForm.obfs = undefined;
      }
      if (!state.quic) {
        state.dataForm.quic = undefined;
      }
      if (!state.bandwidth) {
        state.dataForm.bandwidth = undefined;
        state.dataForm.ignoreClientBandwidth = undefined;
      }
      if (!state.speedTest) {
        state.dataForm.speedTest = undefined;
      }
      if (!state.udp) {
        state.dataForm.disableUDP = undefined;
        state.dataForm.udpIdleTimeout = undefined;
      }
      if (!state.resolver) {
        state.dataForm.resolver = undefined;
      }
      if (!state.sniff) {
        state.dataForm.sniff = undefined;
      }
      if (!state.acl) {
        state.dataForm.acl = undefined;
      }
      if (!state.outbounds) {
        state.dataForm.outbounds = undefined;
      }
      if (!state.masquerade) {
        state.dataForm.masquerade = undefined;
      }

      if (state.dataForm?.udpIdleTimeout === "") {
        state.dataForm!.udpIdleTimeout = "60s";
      }
      if (state.dataForm?.quic?.maxIdleTimeout === "") {
        state.dataForm!.quic!.maxIdleTimeout = "30s";
      }
      if (state.dataForm?.resolver?.tcp?.timeout === "") {
        state.dataForm!.resolver!.tcp!.timeout = "4s";
      }
      if (state.dataForm?.resolver?.udp?.timeout === "") {
        state.dataForm!.resolver!.udp!.timeout = "4s";
      }
      if (state.dataForm?.resolver?.tls?.timeout === "") {
        state.dataForm!.resolver!.tls!.timeout = "10s";
      }
      if (state.dataForm?.resolver?.https?.timeout === "") {
        state.dataForm!.resolver!.https!.timeout = "10s";
      }
      if (state.dataForm?.sniff?.timeout === "") {
        state.dataForm!.sniff!.timeout = "2s";
      }
      if (state.dataForm?.acl?.geoUpdateInterval === "") {
        state.dataForm!.acl!.geoUpdateInterval = "168h";
      }

      const params = { ...state.dataForm };
      updateHysteria2ConfigApi(params).then(() => {
        ElMessage.success(t("common.success"));
        setConfig();
      });
    }
  });
};

const setConfig = () => {
  listConfigApi({
    keys: [hysteria2EnableKey, hysteria2Remark, hysteria2ConfigPortHopping],
  }).then((response) => {
    const data = response.data;
    data.forEach((configVo) => {
      if (configVo.key === hysteria2EnableKey) {
        state.enableForm.enable = configVo.value;
      } else if (configVo.key === hysteria2Remark) {
        state.configForm.remark = configVo.value;
      } else if (configVo.key === hysteria2ConfigPortHopping) {
        state.configForm.portHopping = configVo.value;
      }
    });
  });

  getHysteria2ConfigApi().then((response) => {
    const data = response.data;
    if (data) {
      state.tlsType = data?.tls?.cert && data?.tls?.key ? "tls" : "acme";
      state.aclType = data?.acl?.inline ? "inline" : "file";

      state.obfs = !!data?.obfs;
      state.quic = !!data?.quic;
      state.bandwidth = !!data?.bandwidth;
      state.speedTest = !!data?.speedTest;
      state.udp = !!data?.disableUDP || !!data?.udpIdleTimeout;
      state.resolver = !!data?.resolver;
      state.sniff = !!data?.sniff;
      state.acl = !!data?.acl;
      state.outbounds = !!data?.outbounds;
      state.masquerade = !!data?.masquerade;

      state.dataForm = deepCopy(defaultHysteria2ServerConfig);
      assignWith(state.dataForm, data);
    }
  });
};

const closeTabPane = (tabPaneName: string) => {
  if (
    tabPaneName === "extension" ||
    tabPaneName === "listen" ||
    tabPaneName === "tls" ||
    tabPaneName === "http"
  ) {
    ElMessage.error("This tab is required");
    return;
  }

  if (tabPaneName === "obfs") {
    state.obfs = false;
  } else if (tabPaneName === "quic") {
    state.quic = false;
  } else if (tabPaneName === "bandwidth") {
    state.bandwidth = false;
  } else if (tabPaneName === "speedTest") {
    state.speedTest = false;
  } else if (tabPaneName === "udp") {
    state.udp = false;
  } else if (tabPaneName === "resolver") {
    state.resolver = false;
  } else if (tabPaneName === "sniff") {
    state.sniff = false;
  } else if (tabPaneName === "acl") {
    state.acl = false;
  } else if (tabPaneName === "outbounds") {
    state.outbounds = false;
  } else if (tabPaneName === "masquerade") {
    state.masquerade = false;
  }
  state.activeName = "listen";
};

const handleDropdownClick = (command: string) => {
  if (command === "obfs") {
    state.obfs = true;
  } else if (command === "quic") {
    state.quic = true;
  } else if (command === "bandwidth") {
    state.bandwidth = true;
  } else if (command === "speedTest") {
    state.speedTest = true;
  } else if (command === "udp") {
    state.udp = true;
  } else if (command === "resolver") {
    state.resolver = true;
  } else if (command === "sniff") {
    state.sniff = true;
  } else if (command === "acl") {
    state.acl = true;
  } else if (command === "outbounds") {
    state.outbounds = true;
  } else if (command === "masquerade") {
    state.masquerade = true;
  }
  state.activeName = command;
};

const setHysteria2Versions = async () => {
  const { data } = await listReleaseApi();
  state.hysteria2Versions = data;
};

const setHysteria2Monitor = async () => {
  const { data } = await monitorHysteria2Api();
  Object.assign(state.hysteria2Monitor, data);
};

const handleHysteria2ChangeVersion = async () => {
  if (!state.hysteria2Version) {
    ElMessage.error("version number is required");
    return;
  }
  ElMessage.info(t("common.wait"));
  await hysteria2ChangeVersionApi({ version: state.hysteria2Version });
  ElMessage.success(t("common.success"));
  await setHysteria2Monitor();
};

onMounted(() => {
  setConfig();
  setHysteria2Versions();
  setHysteria2Monitor();
});
</script>

<style lang="scss" scoped>
.el-card .el-form {
  max-width: 1000px;
  margin: 0 auto;
}
</style>
