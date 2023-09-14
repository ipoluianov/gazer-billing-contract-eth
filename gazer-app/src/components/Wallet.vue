<template>
  <div class="header">
    <div v-if="this.connected" class="widget">{{ this.shortAddress }}</div>
    <div v-if="!this.connected" class="widget">Use MetaMask Wallet</div>
    <div v-if="!this.connected" class="widget">
      <button class="connect-button" @click="this.connectWallet">
        CONNECT WALLET
      </button>
    </div>
    <button v-if="this.connected" class="disconnect-button" @click="this.disconnectWallet">disconnect</button>
  </div>
  <div class="header" v-if="this.connected">
    <div class="widget">
      <div v-if="this.connected" class="widget">My Balance: {{ this.balance }} {{ this.balanceCurrency }}</div>
    </div>
  </div>
  <div class="header" v-if="this.connected">
    <div class="widget">
      <div v-if="this.connected" class="widget">ChainID: {{ this.chainName }}</div>
      <div v-if="this.chainId != 1 && this.chainId != 31337" style="color: #FF0000; font-weight: bold; font-size: 24pt;">
        Wrong ChainID
      </div>
    </div>
  </div>
</template>

<script>
import wallet from "./lib/wallet";
import { ethers } from "ethers";
import js_utils from "./lib/js_utils";

import PayloadShop from "../../../artifacts/contracts/PayloadShop.sol/PayloadShop.json";
import ContractAddress from "./contract/ContractAddress.json";

export default {
  data() {
    return {
      chainId: 0,
      connected: false,
      balance: 0,
      balanceCurrency: "",
      selectedAddress: null,
      chainName: "",
    };
  },
  computed: {
    shortAddress() {
      return this.selectedAddress;
      return wallet.shortAddress(this.selectedAddress);
    },
  },
  emits: ["connected", "disconnected"],
  methods: {
    async connectWallet() {
        console.log("wallet connectWallet");
      [this.signer, this.contract] = await wallet.connect(
        ContractAddress.address,
        PayloadShop.abi,
        () => {
          this.resetState();
        }
      );
      this.selectedAddress = await this.signer.getAddress();
      this.updateBalance();
      this.$emit("connected", this.signer, this.contract);
    },
    disconnectWallet() {
        console.log("wallet disconnectWallet");
      this.resetState();
      this.$emit("disconnected");
    },
    async resetState() {
        console.log("wallet resetState");
      this.connected = false;
      this.selectedAddress = null;
      this.balance = 0;
      this.errorText = "";
      this.domains = [];
      this.dialog = "";
      this.$emit("disconnected");
    },

    async updateBalance() {
        console.log("wallet updateBalance");
      if (this.signer == null) {
        console.log("no signer");
        this.resetState();
        return;
      }
      console.log("get-balance");
      this.balance = await this.signer.getBalance();
      this.balance = js_utils.truncate(
        ethers.utils.formatEther(this.balance),
        5
      );
      this.connected = true;
      this.signer.provider.getNetwork().then((network) => {
        this.chainId = network.chainId;
        console.log("------------- CHAIN:" + this.chainId);
        let currencyCode = wallet.getCurrencyCode(this.chainId);
        this.chainName = wallet.getChainName(this.chainId);
        this.balanceCurrency = "ETH";
      });
    },
  },
};
</script>

<style scoped>

.header {
  margin: 10px;
}
.connect-button {
  background-color: #009900;
  border: 1px solid #00bb00;
  padding: 5px;
  margin: 5px;
  cursor: pointer;
  border-radius: 3px;
  color: #ffffff;
  font-family: "Courier New", Courier, monospace;
  font-size: 14pt;
  font-weight: 600;
}

.disconnect-button {
  background-color: #00000000;
  border: 0px solid #888888;
  padding: 5px;
  margin: 5px;
  cursor: pointer;
  border-radius: 3px;
  color: #008899;
  text-decoration: underline;
  
  font-size: 10pt;
  font-weight: 600;
  
}

.hr {
  min-height: 5px;
  border-top: 1px solid #888888;
}
</style>