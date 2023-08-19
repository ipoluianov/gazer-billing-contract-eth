<script>
import Buy from "./components/Buy.vue"
import List from "./components/List.vue";
import Wallet from "./components/Wallet.vue"
import Header from "./components/Header.vue"

export default {
  data() {
    return {
      connected: false,
      updateCounter: 0,
      abi: "",
      domains: [],
      dialog: '',
    }
  },
  computed: {
  },
  components: {
    Buy,
    List,
    Wallet,
    Header,
  },

  mounted() {
  },
  methods: {
    async resetState() {
      this.connected = false;
      this.domains = [];
      this.dialog = '';
    },
    registerDomain() {
      this.dialog = 'reg'
    },
    mainForm() {
      this.dialog = ''
    },
    registered() {
      this.updateDomains();
      this.dialog = ''
    },
    onConnected(signer, contract) {
      this.connected = true;
      this.signer = signer;
      this.contract = contract;
      this.updateDomains();
    },
    onDisconnected() {
      this.connected = false;
      this.resetState();
    },
    updateDomains() {
      if (this.$refs.itemsList == null || this.$refs.itemsList === undefined) {
        return;
      }
      this.$refs.itemsList.updateDomains();
    },
    registerSubdomain(parentDomain) {
    },
  }

}
</script>

<template>
  <div class="wrapper">
    <Header></Header>
    <Wallet @connected="this.onConnected" @disconnected="this.onDisconnected"/>
    <div v-if="this.dialog == 'reg'">
      <Buy @on-register="this.registered" @cancel="this.mainForm" :parent-domain=this.parentDomainForReg />
    </div>
    <div class="content" v-if="this.connected && this.dialog == ''">
      <div class="buttons">
        <button class="add-button" @click="this.registerDomain">Register new domain</button>
        <button class="refresh-button" @click="this.updateDomains">Refresh</button>
      </div>
      <List ref="itemsList"/>
    </div>
  </div>
</template>

<style scoped>
.add-button {
  font-size: 14pt;
  background-color: #2C873A;
  color: #FFFFFF;
  border: 0px solid #FFFFFF;
  border-radius: 5px;
  padding: 5px;
  margin: 5px;
  cursor: pointer;
}

.refresh-button {
  font-size: 14pt;
  background-color: #555555;
  color: #FFFFFF;
  border: 0px solid #FFFFFF;
  border-radius: 5px;
  padding: 5px;
  margin: 5px;
  cursor: pointer;
}

.hr {
    min-height: 5px;
    border-top: 1px solid #888888;
}

.buttons {
  background-color: #111111;
  color: #AAA;
  display: flex;
  flex-direction: row;
  align-items: center;
  flex-wrap: nowrap;
  justify-content: space-between;
}

</style>
