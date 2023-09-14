<script>
import Buy from "./components/Buy.vue"
import List from "./components/List.vue";
import Wallet from "./components/Wallet.vue"
import Header from "./components/Header.vue"
import Admin from "./components/Admin.vue"

export default {
  // https://gazer.cloud/action/buy-premium/?addr=jkx7pkp3f3gfhecunlubkfgegadmaqtzihro5w5537nergsd
  data() {
    return {
      connected: false,
      updateCounter: 0,
      abi: "",
      domains: [],
      dialog: '',
      isAdmin: false,
    }
  },
  computed: {
  },
  components: {
    Buy,
    List,
    Wallet,
    Header,
    Admin,
  },

  mounted() {
  },
  methods: {
    async resetState() {
      console.log("app resetState");
      this.connected = false;
      this.domains = [];
      this.dialog = '';
      this.isAdmin = false;
    },
    registerDomain() {
      this.dialog = 'reg'
    },
    mainForm() {
      this.dialog = ''
    },
    registered() {
      console.log("app registered")
      this.updateDomains();
      this.dialog = ''
    },
    onConnected(signer, contract) {
      console.log("app onConnected")
      this.connected = true;
      this.signer = signer;
      this.contract = contract;
      this.updateDomains();
      this.updateAdmin();
    },
    onDisconnected() {
      console.log("app onDisconnected")
      this.connected = false;
      this.resetState();
    },
    updateDomains() {
      console.log("app updateDomains")
      if (this.$refs.itemsList == null || this.$refs.itemsList === undefined) {
        return;
      }
      this.$refs.itemsList.updateDomains();

      if (this.$refs.adminPanel == null || this.$refs.adminPanel === undefined) {
        return;
      }
      this.$refs.adminPanel.updateUI();
    },

    async updateAdmin() {
        console.log("wallet updateAdmin");
      console.log("updateAdmin ---------------------------- ");
      this.domains = [];
      let contractOwnerAddress = await this.contract.owner();
      contractOwnerAddress = contractOwnerAddress.toLowerCase();
      let localAddress = await this.signer.getAddress();
      localAddress = localAddress.toLowerCase();
      console.log(
        "Contract owner:" + contractOwnerAddress + " localAddr:" + localAddress
      );
      if (contractOwnerAddress == localAddress) {
        console.log("updateAdmin THIS IS ADMIN");
        this.isAdmin = true;
      }
      console.log("updateAdmin ok");
    },

  }

}
</script>

<template>
  <div class="wrapper">
    <Header></Header>
    <Wallet @connected="this.onConnected" @disconnected="this.onDisconnected"/>
    <div v-if="this.dialog == 'reg'">
      <Buy @on-register="this.registered" @cancel="this.mainForm"/>
    </div>
    <div class="content" v-if="this.connected && this.dialog == ''">
      <div class="buttons">
        <button class="add-button" @click="this.registerDomain">Register Premium Address</button>
        <button class="refresh-button" @click="this.updateDomains">Refresh</button>
      </div>
      <List ref="itemsList"/>
    </div>
    <div class="content" v-if="this.isAdmin && this.dialog == ''">
      <Admin ref="adminPanel"/>
    </div>
  </div>
</template>

<style scoped>
.add-button {
  font-size: 14pt;
  background-color: #25c43a;
  color: #FFFFFF;
  border: 0px solid #FFFFFF;
  border-radius: 5px;
  padding: 12px;
  margin: 5px;
  cursor: pointer;
}

.refresh-button {
  font-size: 14pt;
  background-color: #555555;
  color: #FFFFFF;
  border: 0px solid #FFFFFF;
  border-radius: 5px;
  padding: 12px;
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
