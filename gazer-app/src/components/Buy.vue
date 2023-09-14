<script>
import wallet from "./lib/wallet";
import ErrorMessage from "./ErrorMessage.vue";
import xchg_utils from "./lib/xchg_utils";
import bytes32 from "./lib/bytes32";

import js_utils from "./lib/js_utils";
import { ethers } from "ethers";

import PayloadShop from "../../../artifacts/contracts/PayloadShop.sol/PayloadShop.json";
import ContractAddress from "./contract/ContractAddress.json";

export default {
  mounted() {
    this.updateCurrentPrice();
    const search = window.location.search;
    console.log("Parameter:" + search);
    let i = search.indexOf("addr=");
    if (i > 0) {
      this.xchgAddress = search.substring(i + 5);
    }
  },
  data() {
    return {
      chainId: 0,
      message: "",
      errorText: null,
      currentTransactionHash: null,
      xchgAddress: "#7ctd4bacwqo5xjxypmhblgmffr3t3x5croyetd5n7tsfobnn",

      currentPriceETH: 0,
      currentPrice: ethers.BigNumber.from(0),
    };
  },
  props: ["parentDomain"],
  emits: ["onRegister", "cancel"],
  components: {
    ErrorMessage,
  },
  methods: {
    async register() {
      this.errorText = "";
      [this.signer, this.contract] = await wallet.connect(
        ContractAddress.address,
        PayloadShop.abi
      );
      console.log("Register", this.domainNameToRegister);
      console.log("Signer", await this.signer.getAddress());
      try {
        if (this.chainId != 1 && this.chainId != 31337) {
          throw "Wrong Chain";
        }
        let xchgAddr = xchg_utils.xchgAddressToBinary(this.xchgAddress);

        const options = { value: this.currentPrice };
        //const options = {value: 0}
        let tx = await this.contract.buy(xchgAddr, options);
        this.message = "processing ...";
        await tx.wait();
        this.message = "complete";
        this.errorText = null;
        this.$emit("onRegister");
      } catch (err) {
        this.errorText = err;
      }
    },
    async updateCurrentPrice() {
      [this.signer, this.contract] = await wallet.connect(
        ContractAddress.address,
        PayloadShop.abi,
        () => {
          this.resetState();
        }
      );

      // Get Contract Balance
      let currentPrice = await this.contract.currentPrice();
      this.currentPriceETH = js_utils.truncate(
        ethers.utils.formatEther(currentPrice),
        18
      );
      this.currentPrice = currentPrice;
      console.log("currentPrice: " + this.currentPrice);

      this.signer.provider.getNetwork().then((network) => {
        this.chainId = network.chainId;
      });
    },
  },
};
</script>

<template>
  <div class="main">
    <h2 style="color: #ffd700">Register a Premium Address</h2>
    <div>
      <div style="color: #4169e1; font-size: 14pt">
        Price: {{ this.currentPriceETH }} ETH
      </div>
    </div>
    <div>Address:</div>
    <div class="line2">
      <input
        placeholder="Gazer Address"
        class="input-address"
        v-model="this.xchgAddress"
      />
    </div>
    <div class="line3">
      <button class="button-ok" @click="this.register">Register</button>
      <button class="button-cancel" @click="this.$emit('cancel')">
        Cancel
      </button>
    </div>
  </div>
  <div>{{ this.message }}</div>
  <div>
    <ErrorMessage :error="this.errorText" />
  </div>
</template>

<style scoped>
.main {
  display: flex;
  flex-direction: column;
}

.line1 {
  display: flex;
  flex-direction: row;
  padding: 6px;
}

.line2 {
  display: flex;
  flex-direction: row;
  padding: 6px;
}

.line3 {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: 6px;
}

.input-address {
  background-color: black;
  color: #ffffff;
  border: 0px;
  font-size: 14pt;
  flex-grow: 2;
  height: 50px;
}

.button-ok {
  font-size: 14pt;
  background-color: #25c43a;
  color: #ffffff;
  border: 0px solid #ffffff;
  border-radius: 5px;
  padding: 12px;
  margin: 5px;
  width: 150px;
  cursor: pointer;
}

.button-cancel {
  font-size: 14pt;
  background-color: #555555;
  color: #ffffff;
  border: 0px solid #ffffff;
  border-radius: 5px;
  padding: 12px;
  margin: 5px;
  cursor: pointer;
}
</style>