
<script>
import Item from "./Item.vue"
import wallet from "./lib/wallet"
import PayloadShop from "../../../artifacts/contracts/PayloadShop.sol/PayloadShop.json"
import ContractAddress from "./contract/ContractAddress.json"

export default {
  mounted() {
    this.updateDomains();
  },
  components: {
    Item,
  },
  data() {
    return {
      domains: [],
    };
  },
  methods: {
    async updateDomains() {
      [this.signer, this.contract] = await wallet.connect(ContractAddress.address, PayloadShop.abi, () => {
        this.resetState();
      });

      this.domains = [];
      const items = await this.contract.getRecordsOfUser(this.signer.getAddress());
      console.log(items);
      console.log("Count:", items.length);
      for (let i = 0; i < items.length; i++) {
        const domain = items[i];
        if (domain === undefined) {
          return "0x00";
        }
        //console.log("domain", i, " = ", this.bytes32ToString(domain.fullName.toString()));
        this.domains.push(domain);
      }

    },
  }
}
</script>

<template>
  <div class="main" v-for="domain in domains">
    <div>
      <Item :domain=domain></Item>
    </div>
  </div>
</template>

<style scoped>
.main {
  padding: 6px;
}
</style>
