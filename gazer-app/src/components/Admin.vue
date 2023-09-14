
<script>
import wallet from "./lib/wallet";
import PayloadShop from "../../../artifacts/contracts/PayloadShop.sol/PayloadShop.json";
import ContractAddress from "./contract/ContractAddress.json";
import js_utils from "./lib/js_utils";
import { ethers } from "ethers";
import ErrorMessage from "./ErrorMessage.vue";

export default {
  mounted() {
    this.updateUI();
  },
  components: {
    ErrorMessage,
  },
  data() {
    return {
      message: "",
      errorText: null,
      balance: 0,
      priceToSet: 0,
      currentPrice: 0,

      withdrawAddress: "",
    };
  },
  methods: {
    async updateUI() {
      console.log("admin updateUI");
      [this.signer, this.contract] = await wallet.connect(
        ContractAddress.address,
        PayloadShop.abi,
        () => {
          this.resetState();
        }
      );

      this.domains = [];
      // Get Contract Balance
      let contractBalance = await this.contract.getBalance();
      this.balance = js_utils.truncate(
        ethers.utils.formatEther(contractBalance),
        18
      );
      console.log("contractBalance" + contractBalance);

      // Get Contract Balance
      let currentPrice = await this.contract.currentPrice();
      this.currentPrice = js_utils.truncate(
        ethers.utils.formatEther(currentPrice),
        18
      );
      //this.currentPrice = this.currentPrice;
      console.log("currentPrice: " + this.currentPrice);
    },

    async setPrice() {
      console.log("SET PRICE:" + this.priceToSet);

      this.errorText = "";
      [this.signer, this.contract] = await wallet.connect(
        ContractAddress.address,
        PayloadShop.abi
      );
      console.log("Signer", await this.signer.getAddress());
      try {
        let priceToSet = ethers.BigNumber.from(this.priceToSet * 1000000000);
        priceToSet = priceToSet.mul(1000000000);
        let tx = await this.contract.setPrice(priceToSet);
        this.message = "processing ...";
        await tx.wait();
        this.message = "complete";
        this.errorText = null;
        //this.$emit("onRegister");
        this.updateUI();
      } catch (err) {
        this.errorText = err;
      }
    },

    async withdraw() {
      console.log("WITHDRAW:" + this.withdrawAddress);

      this.errorText = "";
      [this.signer, this.contract] = await wallet.connect(
        ContractAddress.address,
        PayloadShop.abi
      );
      console.log("Signer", await this.signer.getAddress());
      try {
        let tx = await this.contract.withdrawAll(this.withdrawAddress);
        this.message = "processing ...";
        await tx.wait();
        this.message = "complete withdraw";
        this.errorText = null;
        //this.$emit("onRegister");
        this.updateUI();
      } catch (err) {
        this.errorText = err;
      }
    },

    resetState() {
      console.log("admin resetState");
    },
  },
};
</script>

<template>
  <div class="main" style="border-top: 10px solid #DDD; margin-top: 30px; background-color: #555;">
    <h1>Admin</h1>
    <div>
      <div>contract balance: {{ this.balance }} ETH</div>
      <div>contract currentPrice: {{ this.currentPrice }} ETH</div>
    </div>
    <div>
      <h1>SetPrice</h1>
      <input v-model="this.priceToSet" />
      <button @click="this.setPrice">SET PRICE (ETH)</button>
    </div>
    <div>
      <h1>Withdraw</h1>
      <input v-model="this.withdrawAddress" />
      <button @click="this.withdraw">WITHDRAW</button>
    </div>
  </div>
  <div>{{ this.message }}</div>
  <div>
    <ErrorMessage :error="this.errorText" />
  </div>
</template>

<style scoped>
.main {
  padding: 6px;
}
</style>
