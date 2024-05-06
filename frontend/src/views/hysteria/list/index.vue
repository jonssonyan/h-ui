<template>
  <div class="app-container">
    <div class="search">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="success" @click="submitForm" :icon="Select"
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
      <el-form ref="dataFormRef" label-position="top" :model="formData">
        <el-tabs v-model="activeName" class="tabs" @tab-click="handleClick">
          <el-tab-pane label="base" name="base">
            <el-form-item label="enable" prop="enable">
              <el-switch
                v-model="enableVal"
                active-value="1"
                inactive-value="0"
                @change="handleChange"
              />
            </el-form-item>
            <el-form-item label="listen" prop="listen">
              <el-input
                v-model="formData.listen"
                placeholder="listen"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="trafficStats.listen"
              prop="trafficStats.listen"
            >
              <el-input
                v-model="formData.trafficStats.listen"
                placeholder="trafficStats.listen"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="obfs" name="obfs">
            <el-form-item label="obfs.type" prop="obfs.type">
              <el-input
                v-model="formData.obfs.type"
                placeholder="obfs.type"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="obfs.salamander.password"
              prop="obfs.salamander.password"
            >
              <el-input
                v-model="formData.obfs.salamander.password"
                placeholder="obfs.salamander.password"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="tls" name="tls">
            <el-form-item label="tls.cert" prop="tls.cert">
              <el-input
                v-model="formData.tls.cert"
                placeholder="tls.cert"
                clearable
              />
            </el-form-item>
            <el-form-item label="tls.key" prop="tls.key">
              <el-input
                v-model="formData.tls.key"
                placeholder="tls.key"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.domains" prop="acme.domains">
              <el-input
                v-model="formData.acme.domains"
                placeholder="acme.domains"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.email" prop="acme.email">
              <el-input
                v-model="formData.acme.email"
                placeholder="acme.email"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.ca" prop="acme.ca">
              <el-input
                v-model="formData.acme.ca"
                placeholder="acme.ca"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.disableHTTP" prop="acme.disableHTTP">
              <el-input
                v-model="formData.acme.disableHTTP"
                placeholder="acme.disableHTTP"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="acme.disableTLSALPN"
              prop="acme.disableTLSALPN"
            >
              <el-input
                v-model="formData.acme.disableTLSALPN"
                placeholder="acme.disableTLSALPN"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.listenHost" prop="acme.listenHost">
              <el-input
                v-model="formData.acme.listenHost"
                placeholder="acme.listenHost"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.altHTTPPort" prop="acme.altHTTPPort">
              <el-input
                v-model="formData.acme.altHTTPPort"
                placeholder="acme.altHTTPPort"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="acme.altTLSALPNPort"
              prop="acme.altTLSALPNPort"
            >
              <el-input
                v-model="formData.acme.altTLSALPNPort"
                placeholder="acme.altTLSALPNPort"
                clearable
              />
            </el-form-item>
            <el-form-item label="acme.dir" prop="acme.dir">
              <el-input
                v-model="formData.acme.dir"
                placeholder="acme.dir"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="quic" name="quic">
            <el-form-item
              label="quic.initStreamReceiveWindow"
              prop="quic.initStreamReceiveWindow"
            >
              <el-input
                v-model="formData.quic.initStreamReceiveWindow"
                placeholder="quic.initStreamReceiveWindow"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="quic.maxStreamReceiveWindow"
              prop="quic.maxStreamReceiveWindow"
            >
              <el-input
                v-model="formData.quic.maxStreamReceiveWindow"
                placeholder="quic.maxStreamReceiveWindow"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="quic.initConnectionReceiveWindow"
              prop="quic.initConnectionReceiveWindow"
            >
              <el-input
                v-model="formData.quic.initConnectionReceiveWindow"
                placeholder="quic.initConnectionReceiveWindow"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="quic.maxConnectionReceiveWindow"
              prop="quic.maxConnectionReceiveWindow"
            >
              <el-input
                v-model="formData.quic.maxConnectionReceiveWindow"
                placeholder="quic.maxConnectionReceiveWindow"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="quic.maxIdleTimeout"
              prop="quic.maxIdleTimeout"
            >
              <el-input
                v-model="formData.quic.maxIdleTimeout"
                placeholder="quic.maxIdleTimeout"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="quic.maxIncomingStreams"
              prop="quic.maxIncomingStreams"
            >
              <el-input
                v-model="formData.quic.maxIncomingStreams"
                placeholder="quic.maxIncomingStreams"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="quic.disablePathMTUDiscovery"
              prop="quic.disablePathMTUDiscovery"
            >
              <el-input
                v-model="formData.quic.disablePathMTUDiscovery"
                placeholder="quic.disablePathMTUDiscovery"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="bandwidth" name="bandwidth">
            <el-form-item label="bandwidth.up" prop="bandwidth.up">
              <el-input
                v-model="formData.bandwidth.up"
                placeholder="bandwidth.up"
                clearable
              />
            </el-form-item>
            <el-form-item label="bandwidth.down" prop="bandwidth.down">
              <el-input
                v-model="formData.bandwidth.down"
                placeholder="bandwidth.down"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="ignoreClientBandwidth"
              prop="ignoreClientBandwidth"
            >
              <el-input
                v-model="formData.ignoreClientBandwidth"
                placeholder="ignoreClientBandwidth"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="speedTest" name="speedTest">
            <el-form-item label="speedTest" prop="speedTest">
              <el-input
                v-model="formData.speedTest"
                placeholder="speedTest"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="udp" name="udp">
            <el-form-item label="disableUDP" prop="disableUDP">
              <el-input
                v-model="formData.disableUDP"
                placeholder="disableUDP"
                clearable
              />
            </el-form-item>
            <el-form-item label="udpIdleTimeout" prop="udpIdleTimeout">
              <el-input
                v-model="formData.udpIdleTimeout"
                placeholder="udpIdleTimeout"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="resolver" name="resolver">
            <el-form-item label="resolver.type" prop="resolver.type">
              <el-input
                v-model="formData.resolver.type"
                placeholder="resolver.type"
                clearable
              />
            </el-form-item>
            <el-form-item label="resolver.tcp.addr" prop="resolver.tcp.addr">
              <el-input
                v-model="formData.resolver.tcp.addr"
                placeholder="resolver.tcp.addr"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.tcp.timeout"
              prop="resolver.tcp.timeout"
            >
              <el-input
                v-model="formData.resolver.tcp.timeout"
                placeholder="resolver.tcp.timeout"
                clearable
              />
            </el-form-item>
            <el-form-item label="resolver.udp.addr" prop="resolver.udp.addr">
              <el-input
                v-model="formData.resolver.udp.addr"
                placeholder="resolver.udp.addr"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.udp.timeout"
              prop="resolver.udp.timeout"
            >
              <el-input
                v-model="formData.resolver.udp.timeout"
                placeholder="resolver.udp.timeout"
                clearable
              />
            </el-form-item>
            <el-form-item label="resolver.tls.addr" prop="resolver.tls.addr">
              <el-input
                v-model="formData.resolver.tls.addr"
                placeholder="resolver.tls.addr"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.tls.timeout"
              prop="resolver.tls.timeout"
            >
              <el-input
                v-model="formData.resolver.tls.timeout"
                placeholder="resolver.tls.timeout"
                clearable
              />
            </el-form-item>
            <el-form-item label="resolver.tls.sni" prop="resolver.tls.sni">
              <el-input
                v-model="formData.resolver.tls.sni"
                placeholder="resolver.tls.sni"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.tls.insecure"
              prop="resolver.tls.insecure"
            >
              <el-input
                v-model="formData.resolver.tls.insecure"
                placeholder="resolver.tls.insecure"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.https.addr"
              prop="resolver.https.addr"
            >
              <el-input
                v-model="formData.resolver.https.addr"
                placeholder="resolver.https.addr"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.https.timeout"
              prop="resolver.https.timeout"
            >
              <el-input
                v-model="formData.resolver.https.timeout"
                placeholder="resolver.https.timeout"
                clearable
              />
            </el-form-item>
            <el-form-item label="resolver.https.sni" prop="resolver.https.sni">
              <el-input
                v-model="formData.resolver.https.sni"
                placeholder="resolver.https.sni"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="resolver.https.insecure"
              prop="resolver.https.insecure"
            >
              <el-input
                v-model="formData.resolver.https.insecure"
                placeholder="resolver.https.insecure"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="acl" name="acl">
            <el-form-item label="acl.file" prop="acl.file">
              <el-input
                v-model="formData.acl.file"
                placeholder="acl.file"
                clearable
              />
            </el-form-item>
            <el-form-item label="acl.inline" prop="acl.inline">
              <el-input
                v-model="formData.acl.inline"
                placeholder="acl.inline"
                clearable
              />
            </el-form-item>
            <el-form-item label="acl.geoip" prop="acl.geoip">
              <el-input
                v-model="formData.acl.geoip"
                placeholder="acl.geoip"
                clearable
              />
            </el-form-item>
            <el-form-item label="acl.geosite" prop="acl.geosite">
              <el-input
                v-model="formData.acl.geosite"
                placeholder="acl.geosite"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="acl.geoUpdateInterval"
              prop="acl.geoUpdateInterval"
            >
              <el-input
                v-model="formData.acl.geoUpdateInterval"
                placeholder="acl.geoUpdateInterval"
                clearable
              />
            </el-form-item>
          </el-tab-pane>
          <el-tab-pane label="outbounds" name="outbounds">
            <!--          <el-col :xs="24" :sm="8">-->
            <!--            <el-form-item label="outbounds.name" prop="outbounds.name">-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.name"-->
            <!--                placeholder="outbounds.name"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item label="outbounds.type" prop="outbounds.type">-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.type"-->
            <!--                placeholder="outbounds.type"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.direct.mode"-->
            <!--              prop="outbounds.direct.mode"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.mode"-->
            <!--                placeholder="outbounds.direct.mode"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.direct.bindIPv4"-->
            <!--              prop="outbounds.direct.bindIPv4"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.bindIPv4"-->
            <!--                placeholder="outbounds.direct.bindIPv4"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.direct.bindIPv6"-->
            <!--              prop="outbounds.direct.bindIPv6"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.bindIPv6"-->
            <!--                placeholder="outbounds.direct.bindIPv6"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.direct.bindDevice"-->
            <!--              prop="outbounds.direct.bindDevice"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.direct.bindDevice"-->
            <!--                placeholder="outbounds.direct.bindDevice"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.socks5.addr"-->
            <!--              prop="outbounds.socks5.addr"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.socks5.addr"-->
            <!--                placeholder="outbounds.socks5.addr"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.socks5.username"-->
            <!--              prop="outbounds.socks5.username"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.socks5.username"-->
            <!--                placeholder="outbounds.socks5.username"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.socks5.password"-->
            <!--              prop="outbounds.socks5.password"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.socks5.password"-->
            <!--                placeholder="outbounds.socks5.password"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item label="outbounds.http.url" prop="outbounds.http.url">-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.http.url"-->
            <!--                placeholder="outbounds.http.url"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--            <el-form-item-->
            <!--              label="outbounds.http.insecure"-->
            <!--              prop="outbounds.http.insecure"-->
            <!--            >-->
            <!--              <el-input-->
            <!--                v-model="formData.outbounds.http.insecure"-->
            <!--                placeholder="outbounds.http.insecure"-->
            <!--                clearable-->
            <!--              />-->
            <!--            </el-form-item>-->
            <!--          </el-col>-->
          </el-tab-pane>
          <el-tab-pane label="masquerade" name="masquerade">
            <el-form-item label="masquerade.type" prop="masquerade.type">
              <el-input
                v-model="formData.masquerade.type"
                placeholder="masquerade.type"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.file.dir"
              prop="masquerade.file.dir"
            >
              <el-input
                v-model="formData.masquerade.file.dir"
                placeholder="masquerade.file.dir"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.proxy.url"
              prop="masquerade.proxy.url"
            >
              <el-input
                v-model="formData.masquerade.proxy.url"
                placeholder="masquerade.proxy.url"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.proxy.rewriteHost"
              prop="masquerade.proxy.rewriteHost"
            >
              <el-input
                v-model="formData.masquerade.proxy.rewriteHost"
                placeholder="masquerade.proxy.rewriteHost"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.string.content"
              prop="masquerade.string.content"
            >
              <el-input
                v-model="formData.masquerade.string.content"
                placeholder="masquerade.string.content"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.string.headers"
              prop="masquerade.string.headers"
            >
              <el-input
                v-model="formData.masquerade.string.headers"
                placeholder="masquerade.string.headers"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.string.statusCode"
              prop="masquerade.string.statusCode"
            >
              <el-input
                v-model="formData.masquerade.string.statusCode"
                placeholder="masquerade.string.statusCode"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.listenHTTP"
              prop="masquerade.listenHTTP"
            >
              <el-input
                v-model="formData.masquerade.listenHTTP"
                placeholder="masquerade.listenHTTP"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.listenHTTPS"
              prop="masquerade.listenHTTPS"
            >
              <el-input
                v-model="formData.masquerade.listenHTTPS"
                placeholder="masquerade.listenHTTPS"
                clearable
              />
            </el-form-item>
            <el-form-item
              label="masquerade.forceHTTPS"
              prop="masquerade.forceHTTPS"
            >
              <el-input
                v-model="formData.masquerade.forceHTTPS"
                placeholder="masquerade.forceHTTPS"
                clearable
              />
            </el-form-item>
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
import { Select } from "@element-plus/icons-vue";
import { getConfig, getHysteria2Config } from "@/api/config";

const state = reactive({
  activeName: "base",
  enableKey: "HYSTERIA2_ENABLE",
  enableVal: "0",
  formData: defaultHysteria2ServerConfig as Hysteria2ServerConfig,
});

const { activeName, enableVal, formData } = toRefs(state);

const handleImport = () => {};
const handleExport = () => {};

const handleChange = () => {};
const submitForm = () => {};

const handleClick = () => {};

onMounted(() => {
  getConfig({ key: state.enableKey }).then((response) => {
    if (response.data) {
      const { value } = response.data;
      state.enableVal = value;
    }
  });
  getHysteria2Config().then((response) => {
    Object.assign(state.formData, response.data);
  });
});
</script>

<style lang="scss" scoped>
.el-card .el-form {
  max-width: 1000px;
  margin: 0 auto;
}

.tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
