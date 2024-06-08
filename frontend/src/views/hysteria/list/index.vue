<template>
  <div class="app-container">
    <div class="search">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" @click="submitForm" :icon="Select"
            >保存
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
              导入
            </el-button>
          </el-upload>
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
      <el-form ref="enableFormDataRef" :model="enableFormData" inline>
        <el-tooltip :content="$t('hysteria.config.enable')" placement="bottom">
          <el-form-item prop="enable">
            <el-switch
              v-model="enableFormData.enable"
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
          <el-select
            v-model="hysteria2Version"
            :placeholder="t('hysteria.hysteria2Version')"
            style="height: 32px"
          >
            <el-option
              v-for="item in hysteria2Versions"
              :key="item.tagName"
              :label="item.tagName"
              :value="item.browserDownloadURL"
            />
          </el-select>
          <el-button
            type="primary"
            @click="handleChangeHysteria2Version"
            style="height: 32px"
            >{{ t("hysteria.changeHysteria2Version") }}
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-tag style="height: 32px">
            Hysteria2 版本: {{ hysteria2Monitor.version }}
          </el-tag>
        </el-form-item>
        <el-form-item>
          <el-tag
            :type="hysteria2Monitor.running ? 'success' : 'danger'"
            style="height: 32px"
          >
            Hysteria2 状态:
            {{
              hysteria2Monitor.running
                ? $t("monitor.hysteria2RunningTrue")
                : $t("monitor.hysteria2RunningFalse")
            }}
          </el-tag>
        </el-form-item>
      </el-form>

      <el-form ref="formDataRef" label-position="top" :model="formData">
        <el-tabs
          v-model="activeName"
          class="tabs"
          closable
          @tab-remove="closeTabPane"
        >
          <el-tab-pane :label="$t('hysteria.listen')" name="listen">
            <el-tooltip
              :content="$t('hysteria.config.listen')"
              placement="bottom"
            >
              <el-form-item label="listen" prop="listen">
                <el-input v-model="formData.listen" clearable />
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
                <el-input v-model="formData.tls.cert" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'tls'"
              :content="$t('hysteria.config.tls.key')"
              placement="bottom"
            >
              <el-form-item label="key" prop="tls.key">
                <el-input v-model="formData.tls.key" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.domains')"
              placement="bottom"
            >
              <el-form-item label="acme.domains" prop="acme.domains">
                <imputMultiple :tags="formData.acme.domains" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.email')"
              placement="bottom"
            >
              <el-form-item label="acme.email" prop="acme.email">
                <el-input v-model="formData.acme.email" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.ca')"
              placement="bottom"
            >
              <el-form-item label="acme.ca" prop="acme.ca">
                <el-select v-model="formData.acme.ca" style="width: 100%">
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
              :content="$t('hysteria.config.acme.disableHTTP')"
              placement="bottom"
            >
              <el-form-item label="acme.disableHTTP" prop="acme.disableHTTP">
                <el-switch v-model="formData.acme.disableHTTP" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.disableTLSALPN')"
              placement="bottom"
            >
              <el-form-item
                label="acme.disableTLSALPN"
                prop="acme.disableTLSALPN"
              >
                <el-switch v-model="formData.acme.disableTLSALPN" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.altHTTPPort')"
              placement="bottom"
            >
              <el-form-item label="acme.altHTTPPort" prop="acme.altHTTPPort">
                <el-input v-model="formData.acme.altHTTPPort" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.altTLSALPNPort')"
              placement="bottom"
            >
              <el-form-item
                label="acme.altTLSALPNPort"
                prop="acme.altTLSALPNPort"
              >
                <el-input v-model="formData.acme.altTLSALPNPort" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.dir')"
              placement="bottom"
            >
              <el-form-item label="acme.dir" prop="acme.dir">
                <el-input v-model="formData.acme.dir" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="tlsType === 'acme'"
              :content="$t('hysteria.config.acme.listenHost')"
              placement="bottom"
            >
              <el-form-item label="acme.listenHost" prop="acme.listenHost">
                <el-input v-model="formData.acme.listenHost" clearable />
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
                  v-model="formData.obfs.type"
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
              v-if="formData.obfs.type === 'salamander'"
              :content="$t('hysteria.config.obfs.salamander.password')"
              placement="bottom"
            >
              <el-form-item
                label="obfs.salamander.password"
                prop="obfs.salamander.password"
              >
                <el-input
                  v-model="formData.obfs.salamander.password"
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
                  v-model="formData.quic.initStreamReceiveWindow"
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
                  v-model="formData.quic.maxStreamReceiveWindow"
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
                  v-model="formData.quic.initConnReceiveWindow"
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
                  v-model="formData.quic.maxConnReceiveWindow"
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
                <el-input v-model="formData.quic.maxIdleTimeout" clearable />
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
                  v-model="formData.quic.maxIncomingStreams"
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
                <el-switch v-model="formData.quic.disablePathMTUDiscovery" />
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
            >
              <el-form-item label="bandwidth.up" prop="bandwidth.up">
                <el-input v-model="formData.bandwidth.up" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.bandwidth.down')"
              placement="bottom"
            >
              <el-form-item label="bandwidth.down" prop="bandwidth.down">
                <el-input v-model="formData.bandwidth.down" clearable />
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
                <el-switch v-model="formData.ignoreClientBandwidth" />
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
                <el-switch v-model="formData.speedTest" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.udp')" name="udp" v-if="udp">
            <el-tooltip
              :content="$t('hysteria.config.disableUDP')"
              placement="bottom"
            >
              <el-form-item label="disableUDP" prop="disableUDP">
                <el-switch v-model="formData.disableUDP" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.udpIdleTimeout')"
              placement="bottom"
            >
              <el-form-item label="udpIdleTimeout" prop="udpIdleTimeout">
                <el-input v-model="formData.udpIdleTimeout" clearable />
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
                <el-select v-model="formData.resolver.type" style="width: 100%">
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
              v-if="formData.resolver.type === 'tcp'"
              :content="$t('hysteria.config.resolver.tcp.addr')"
              placement="bottom"
            >
              <el-form-item label="resolver.tcp.addr" prop="resolver.tcp.addr">
                <el-input v-model="formData.resolver.tcp.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'tcp'"
              :content="$t('hysteria.config.resolver.tcp.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.tcp.timeout"
                prop="resolver.tcp.timeout"
              >
                <el-input v-model="formData.resolver.tcp.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'udp'"
              :content="$t('hysteria.config.resolver.udp.addr')"
              placement="bottom"
            >
              <el-form-item label="resolver.udp.addr" prop="resolver.udp.addr">
                <el-input v-model="formData.resolver.udp.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'udp'"
              :content="$t('hysteria.config.resolver.udp.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.udp.timeout"
                prop="resolver.udp.timeout"
              >
                <el-input v-model="formData.resolver.udp.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.addr')"
              placement="bottom"
            >
              <el-form-item label="resolver.tls.addr" prop="resolver.tls.addr">
                <el-input v-model="formData.resolver.tls.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.tls.timeout"
                prop="resolver.tls.timeout"
              >
                <el-input v-model="formData.resolver.tls.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.sni')"
              placement="bottom"
            >
              <el-form-item label="resolver.tls.sni" prop="resolver.tls.sni">
                <el-input v-model="formData.resolver.tls.sni" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'tls'"
              :content="$t('hysteria.config.resolver.tls.insecure')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.tls.insecure"
                prop="resolver.tls.insecure"
              >
                <el-switch v-model="formData.resolver.tls.insecure" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.addr')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.addr"
                prop="resolver.https.addr"
              >
                <el-input v-model="formData.resolver.https.addr" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.timeout')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.timeout"
                prop="resolver.https.timeout"
              >
                <el-input v-model="formData.resolver.https.timeout" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.sni')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.sni"
                prop="resolver.https.sni"
              >
                <el-input v-model="formData.resolver.https.sni" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.resolver.type === 'https'"
              :content="$t('hysteria.config.resolver.https.insecure')"
              placement="bottom"
            >
              <el-form-item
                label="resolver.https.insecure"
                prop="resolver.https.insecure"
              >
                <el-switch v-model="formData.resolver.https.insecure" />
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
              v-if="aclTypes === 'file'"
              :content="$t('hysteria.config.acl.file')"
              placement="bottom"
            >
              <el-form-item label="acl.file" prop="acl.file">
                <el-input v-model="formData.acl.file" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="aclTypes === 'inline'"
              :content="$t('hysteria.config.acl.inline')"
              placement="bottom"
            >
              <el-form-item label="acl.inline" prop="acl.inline">
                <imputMultiple :tags="formData.acl.inline" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.acl.geoip')"
              placement="bottom"
            >
              <el-form-item label="acl.geoip" prop="acl.geoip">
                <el-input v-model="formData.acl.geoip" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              :content="$t('hysteria.config.acl.geosite')"
              placement="bottom"
            >
              <el-form-item label="acl.geosite" prop="acl.geosite">
                <el-input v-model="formData.acl.geosite" clearable />
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
                <el-input v-model="formData.acl.geoUpdateInterval" clearable />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
          <el-tab-pane
            :label="$t('hysteria.outbounds')"
            name="outbounds"
            v-if="outbounds"
          >
            <Outbounds :outbounds="formData.outbounds" />
          </el-tab-pane>
          <el-tab-pane :label="$t('hysteria.http')" name="http">
            <el-tooltip
              :content="$t('hysteria.config.trafficStats.listen')"
              placement="bottom"
            >
              <el-form-item
                :label="$t('hysteria.config.trafficStats.listen')"
                prop="trafficStats.listen"
              >
                <el-input v-model="formData.trafficStats.listen" clearable />
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
                  v-model="formData.masquerade.type"
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
              v-if="formData.masquerade.type === 'file'"
              :content="$t('hysteria.config.masquerade.file.dir')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.file.dir"
                prop="masquerade.file.dir"
              >
                <el-input v-model="formData.masquerade.file.dir" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.masquerade.type === 'proxy'"
              :content="$t('hysteria.config.masquerade.proxy.url')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.proxy.url"
                prop="masquerade.proxy.url"
              >
                <el-input v-model="formData.masquerade.proxy.url" clearable />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.masquerade.type === 'proxy'"
              :content="$t('hysteria.config.masquerade.proxy.rewriteHost')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.proxy.rewriteHost"
                prop="masquerade.proxy.rewriteHost"
              >
                <el-switch v-model="formData.masquerade.proxy.rewriteHost" />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.masquerade.type === 'string'"
              :content="$t('hysteria.config.masquerade.string.content')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.string.content"
                prop="masquerade.string.content"
              >
                <el-input
                  v-model="formData.masquerade.string.content"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.masquerade.type === 'string'"
              :content="$t('hysteria.config.masquerade.string.headers')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.string.headers"
                prop="masquerade.string.headers"
              >
                <el-input
                  v-model="formData.masquerade.string.headers"
                  clearable
                />
              </el-form-item>
            </el-tooltip>
            <el-tooltip
              v-if="formData.masquerade.type === 'string'"
              :content="$t('hysteria.config.masquerade.string.statusCode')"
              placement="bottom"
            >
              <el-form-item
                label="masquerade.string.statusCode"
                prop="masquerade.string.statusCode"
              >
                <el-switch v-model="formData.masquerade.string.statusCode" />
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
                <el-input v-model="formData.masquerade.listenHTTP" clearable />
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
                <el-input v-model="formData.masquerade.listenHTTPS" clearable />
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
                <el-switch v-model="formData.masquerade.forceHTTPS" />
              </el-form-item>
            </el-tooltip>
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import {
  defaultHysteria2ServerConfig,
  Hysteria2ServerConfig,
  Tab,
} from "@/api/config/types";
import Outbounds from "./components/Outbounds.vue";
import { CirclePlusFilled, Select } from "@element-plus/icons-vue";
import {
  exportHysteria2ConfigApi,
  getConfigApi,
  getHysteria2ConfigApi,
  importHysteria2ConfigApi,
  updateHysteria2ConfigApi,
} from "@/api/config";
import { useI18n } from "vue-i18n";
import { assignIgnoringNull } from "@/utils/object";
import {
  UploadFile,
  UploadRawFile,
  UploadRequestOptions,
} from "element-plus/lib/components";
import { listReleaseApi } from "@/api/hysteria2";
import { Hysteria2ReleaseVo } from "@/api/hysteria2/types";
import { monitorHysteria2Api } from "@/api/monitor";

const { t } = useI18n();

const hysteria2EnableKey = "HYSTERIA2_ENABLE";

const tlsTypes = ref<string[]>(["tls", "acme"]);
const aclTypes = ref<string[]>(["file", "inline"]);
const acmeCas = ref<string[]>(["zerossl", "letsencrypt"]);
const obfsTypes = ref<string[]>(["salamander"]);
const resolverTypes = ref<string[]>(["tcp", "udp", "tls", "https"]);
const masqueradeTypes = ref<string[]>(["file", "proxy", "string"]);

const state = reactive({
  enableFormData: {
    enable: "0",
  },
  formData: defaultHysteria2ServerConfig as Hysteria2ServerConfig,
  activeName: "listen",
  tlsType: "acme",
  aclType: "inline",
  obfs: false,
  quic: false,
  bandwidth: false,
  speedTest: false,
  udp: false,
  resolver: false,
  acl: false,
  outbounds: false,
  masquerade: false,
  fileList: [] as UploadFile[],
  hysteria2Version: "",
  hysteria2Versions: [] as Hysteria2ReleaseVo[],
  hysteria2Monitor: {
    version: "",
    running: false,
  },
});

const {
  activeName,
  formData,
  enableFormData,
  tlsType,
  aclType,
  obfs,
  quic,
  bandwidth,
  speedTest,
  udp,
  resolver,
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
      ElMessage.success("导入成功");
    });
    state.fileList = [];
  }
};

function beforeImport(file: UploadRawFile) {
  if (!file.name.endsWith(".yaml")) {
    ElMessage.error("file format not supported");
    return false;
  }
  if (file.size / 1024 / 1024 > 2) {
    ElMessage.error("the file is too big, less than 2 MB");
    return false;
  }
}

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
    // 模拟点击下载
    a.click();
    window.URL.revokeObjectURL(url);
    ElMessage.success("导出成功");
  } catch (err) {
    ElMessage.error("导出失败");
  }
};

const handleChangeEnable = () => {};
const submitForm = () => {
  const params = { ...state.formData };
  updateHysteria2ConfigApi(params).then(() => {
    ElMessage.success("更新成功");
    setConfig();
  });
};

const setConfig = () => {
  getConfigApi({ key: hysteria2EnableKey }).then((response) => {
    if (response.data) {
      const { value } = response.data;
      state.enableFormData.enable = value;
    }
  });
  getHysteria2ConfigApi().then((response) => {
    if (response.data) {
      assignIgnoringNull(state.formData, response.data);
      if (state.formData?.tls?.cert && state.formData?.tls?.key) {
        state.tlsType = "tls";
      } else {
        state.tlsType = "acme";
      }
      if (state.formData?.acl?.inline) {
        state.aclType = "inline";
      } else {
        state.aclType = "file";
      }
      if (state.formData?.obfs) {
        state.obfs = true;
      }
      if (state.formData?.quic) {
        state.quic = true;
      }
      if (state.formData?.bandwidth) {
        state.bandwidth = true;
      }
      if (state.formData?.speedTest) {
        state.speedTest = true;
      }
      if (state.formData?.disableUDP || state.formData?.udpIdleTimeout) {
        state.udp = true;
      }
      if (state.formData?.resolver) {
        state.resolver = true;
      }
      if (state.formData?.acl) {
        state.acl = true;
      }
      if (state.formData?.outbounds) {
        state.outbounds = true;
      }
      if (state.formData?.masquerade) {
        state.masquerade = true;
      }
    }
  });
};

const closeTabPane = (tabPaneName: string) => {
  if (
    tabPaneName === "listen" ||
    tabPaneName === "tls" ||
    tabPaneName === "http"
  ) {
    ElMessage.error("This tab is required");
    return;
  }

  if (tabPaneName === "obfs") {
    state.formData.obfs = undefined;
    state.obfs = false;
  } else if (tabPaneName === "quic") {
    state.formData.quic = undefined;
    state.quic = false;
  } else if (tabPaneName === "bandwidth") {
    state.formData.bandwidth = undefined;
    state.formData.ignoreClientBandwidth = undefined;
    state.bandwidth = false;
  } else if (tabPaneName === "speedTest") {
    state.formData.speedTest = undefined;
    state.speedTest = true;
  } else if (tabPaneName === "udp") {
    state.formData.disableUDP = undefined;
    state.formData.udpIdleTimeout = undefined;
    state.udp = false;
  } else if (tabPaneName === "resolver") {
    state.formData.resolver = undefined;
    state.resolver = false;
  } else if (tabPaneName === "acl") {
    state.formData.acl = undefined;
    state.acl = false;
  } else if (tabPaneName === "outbounds") {
    state.formData.outbounds = undefined;
    state.outbounds = false;
  } else if (tabPaneName === "masquerade") {
    state.formData.masquerade = undefined;
    state.masquerade = false;
  }
  state.activeName = "listen";
};

const handleDropdownClick = (command: string) => {
  if (command === "obfs") {
    state.formData.obfs = {
      type: "salamander",
      salamander: {
        password: "cry_me_a_r1ver",
      },
    };
    state.obfs = true;
  } else if (command === "quic") {
    state.formData.quic = {
      initStreamReceiveWindow: 8388608,
      maxStreamReceiveWindow: 8388608,
      initConnReceiveWindow: 20971520,
      maxConnReceiveWindow: 20971520,
      maxIdleTimeout: "30s",
      maxIncomingStreams: 1024,
      disablePathMTUDiscovery: false,
    };
    state.quic = true;
  } else if (command === "bandwidth") {
    state.formData.bandwidth = {
      up: "1 gbps",
      down: "1 gbps",
    };
    state.formData.ignoreClientBandwidth = false;
    state.bandwidth = true;
  } else if (command === "speedTest") {
    state.formData.speedTest = false;
    state.speedTest = true;
  } else if (command === "udp") {
    state.formData.disableUDP = false;
    state.formData.udpIdleTimeout = "60s";
    state.udp = true;
  } else if (command === "resolver") {
    state.formData.resolver = {
      type: "",
      tcp: {
        addr: "8.8.8.8:53",
        timeout: "4s",
      },
      udp: {
        addr: "8.8.4.4:53",
        timeout: "4s",
      },
      tls: {
        addr: "1.1.1.1:853",
        timeout: "10s",
        sni: "cloudflare-dns.com",
        insecure: false,
      },
      https: {
        addr: "1.1.1.1:443",
        timeout: "10s",
        sni: "cloudflare-dns.com",
        insecure: false,
      },
    };
    state.resolver = true;
  } else if (command === "acl") {
    state.formData.acl = {
      file: undefined,
      inline: undefined,
      geoip: undefined,
      geosite: undefined,
      geoUpdateInterval: undefined,
    };
    state.acl = true;
  } else if (command === "outbounds") {
    state.formData.outbounds = [];
    state.outbounds = true;
  } else if (command === "masquerade") {
    state.formData.masquerade = {
      type: "",
      file: {
        dir: "",
      },
      proxy: {
        url: "",
        rewriteHost: true,
      },
      string: {
        content: "hello stupid world",
        headers: undefined,
        statusCode: 200,
      },
      listenHTTP: ":80",
      listenHTTPS: ":443",
      forceHTTPS: true,
    };
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

const handleChangeHysteria2Version = async () => {};

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
