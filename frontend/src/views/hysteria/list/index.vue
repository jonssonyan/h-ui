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
          <el-button @click="handleImport">
            <template #icon>
              <i-ep-upload />
            </template>
            导入
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
      <el-form ref="enableFormDataRef" :model="enableFormData" inline>
        <el-tooltip :content="$t('hysteria.config.enable')" placement="bottom">
          <el-form-item :label="$t('hysteria.enable')" prop="enable">
            <el-switch
              v-model="enableFormData.enable"
              active-value="1"
              inactive-value="0"
              @change="handleChangeEnable"
            />
          </el-form-item>
        </el-tooltip>
        <el-form-item>
          <el-dropdown @command="handleDropdownClick">
            <el-button type="primary" :icon="CirclePlusFilled">
              {{ $t("hysteria.addConfigItem") }}
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="obfs"
                  >{{ $t("hysteria.obfs") }}
                </el-dropdown-item>
                <el-dropdown-item command="quic"
                  >{{ $t("hysteria.quic") }}
                </el-dropdown-item>
                <el-dropdown-item command="bandwidth"
                  >{{ $t("hysteria.bandwidth") }}
                </el-dropdown-item>
                <el-dropdown-item command="speedTest"
                  >{{ $t("hysteria.speedTest") }}
                </el-dropdown-item>
                <el-dropdown-item command="udp"
                  >{{ $t("hysteria.udp") }}
                </el-dropdown-item>
                <el-dropdown-item command="resolver"
                  >{{ $t("hysteria.resolver") }}
                </el-dropdown-item>
                <el-dropdown-item command="acl"
                  >{{ $t("hysteria.acl") }}
                </el-dropdown-item>
                <el-dropdown-item command="outbounds"
                  >{{ $t("hysteria.outbounds") }}
                </el-dropdown-item>
                <el-dropdown-item command="masquerade"
                  >{{ $t("hysteria.masquerade") }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
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
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.name')"-->
            <!--              prop="outbounds.name"-->
            <!--            >-->
            <!--              <el-input v-model="formData.outbounds.name" clearable />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.type')"-->
            <!--              prop="outbounds.type"-->
            <!--            >-->
            <!--              <el-input v-model="formData.outbounds.type" clearable />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.socks5.addr')"-->
            <!--              prop="outbounds.socks5.addr"-->
            <!--            >-->
            <!--              <el-input v-model="formData.outbounds.socks5.addr" clearable />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.socks5.username')"-->
            <!--              prop="outbounds.socks5.username"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.socks5.username"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.socks5.password')"-->
            <!--              prop="outbounds.socks5.password"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.socks5.password"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.http.url')"-->
            <!--              prop="outbounds.http.url"-->
            <!--            >-->
            <!--              <el-input v-model="formData.outbounds.http.url" clearable />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.http.insecure')"-->
            <!--              prop="outbounds.http.insecure"-->
            <!--            >-->
            <!--              <el-input v-model="formData.outbounds.http.insecure" clearable />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.direct.mode')"-->
            <!--              prop="outbounds.direct.mode"-->
            <!--            >-->
            <!--              <el-input v-model="formData.outbounds.direct.mode" clearable />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.direct.bindIPv4')"-->
            <!--              prop="outbounds.direct.bindIPv4"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.bindIPv4"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.direct.bindIPv6')"-->
            <!--              prop="outbounds.direct.bindIPv6"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.bindIPv6"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              :label="$t('hysteria.config.outbounds.direct.bindDevice')"-->
            <!--              prop="outbounds.direct.bindDevice"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.bindDevice"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
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
} from "@/api/config/types";
import { CirclePlusFilled, Select } from "@element-plus/icons-vue";
import { getConfigApi, getHysteria2ConfigApi } from "@/api/config";

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
} = toRefs(state);

const handleImport = () => {};
const handleExport = () => {};

const handleChangeEnable = () => {};
const submitForm = () => {};

const setConfig = () => {
  getConfigApi({ key: hysteria2EnableKey }).then((response) => {
    if (response.data) {
      const { value } = response.data;
      state.enableFormData.enable = value;
    }
  });
  getHysteria2ConfigApi().then((response) => {
    if (response.data) {
      Object.assign(state.formData, response.data);
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
  if (tabPaneName === "obfs") {
    state.obfs = false;
  } else if (tabPaneName === "quic") {
    state.quic = false;
  } else if (tabPaneName === "bandwidth") {
    state.bandwidth = false;
  } else if (tabPaneName === "speedTest") {
    state.speedTest = true;
  } else if (tabPaneName === "udp") {
    state.udp = false;
  } else if (tabPaneName === "resolver") {
    state.resolver = false;
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
  } else if (command === "acl") {
    state.acl = true;
  } else if (command === "outbounds") {
    state.outbounds = true;
  } else if (command === "masquerade") {
    state.masquerade = true;
  }
  state.activeName = command;
};

onMounted(() => {
  setConfig();
});
</script>

<style lang="scss" scoped>
.el-card .el-form {
  max-width: 1000px;
  margin: 0 auto;
}
</style>
