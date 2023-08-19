<script>
import connect1 from './lib/wallet'
import ErrorMessage from './ErrorMessage.vue'
import xchg_utils from './lib/xchg_utils'
import bytes32 from './lib/bytes32'

import PayloadShop from "../../../artifacts/contracts/PayloadShop.sol/PayloadShop.json"
import ContractAddress from "./contract/ContractAddress.json"


export default {
    data() {
        return {
            message: "",
            errorText: null,
            currentTransactionHash: null,
            xchgAddress: "#7ctd4bacwqo5xjxypmhblgmffr3t3x5croyetd5n7tsfobnn",
        }
    },
    props: ['parentDomain'],
    emits: ['onRegister', 'cancel'],
    components: {
        ErrorMessage
    },
    methods: {
        async register() {
            this.errorText = "";
            [this.signer, this.contract] = await connect1.connect(ContractAddress.address, PayloadShop.abi);
            console.log("Register", this.domainNameToRegister);
            console.log("Signer", await this.signer.getAddress());
            try {
                let xchgAddr = xchg_utils.xchgAddressToBinary(this.xchgAddress);
                let tx = await this.contract.buy(xchgAddr);
                this.message = "processing ...";
                await tx.wait();
                this.message = "complete";
                this.errorText = null;
                this.$emit('onRegister');
            } catch (err) {
                this.errorText = err;
            }
        },
    },
}
</script>

<template>
    <div class="main">
        <div class="line2">
            <input placeholder="Gazer Address" class="input-address" v-model="this.xchgAddress" />
        </div>
        <div class="line3">
            <button class="button-ok" @click="this.register">Register</button>
            <button class="button-cancel" @click="this.$emit('cancel')">Cancel</button>
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

.input-name {
    background-color: black;
    color: #FFFFFF;
    border: 0px;
    font-size: 14pt;
    flex-grow: 2;
}

.input-address {
    background-color: black;
    color: #FFFFFF;
    border: 0px;
    font-size: 14pt;
    flex-grow: 2;
}

.button-ok {
    font-size: 14pt;
  background-color: #2C873A;
  color: #FFFFFF;
  border: 0px solid #FFFFFF;
  border-radius: 5px;
  padding: 5px;
  margin: 5px;
  cursor: pointer;
}

.button-cancel {
    font-size: 14pt;
  background-color: #555555;
  color: #FFFFFF;
  border: 0px solid #FFFFFF;
  border-radius: 5px;
  padding: 5px;
  margin: 5px;
  cursor: pointer;
}
</style>