<script>
export default {
    data() {
        return {
            showFull: false,
        };
    },
    props: ['error'],
    computed: {

        hasErrors() {
            if (this.error === undefined || this.error == null) {
                return false;
            }
            let text = this.error.toString();
            if (text.length == 0)
                return false; 
            return true;
        },

        shortText() {
            if (this.error === undefined || this.error == null) {
                return "";
            }
            let text = this.error.toString();
            let resultText = "";
            if (text.toLowerCase().includes("reject")) {
                resultText = "Action rejected";
            }
            if (text.toUpperCase().includes("CERR:ACCESS_DENIED")) {
                resultText = "Access Denied";
            }
            if (text.toUpperCase().includes("CERR:WRONG_PRICE")) {
                resultText = "Wrong Price";
            }
            return resultText
        }
    }
}
</script>


<template>
    <div class="shortErrorText" v-if="this.shortText != ''">{{ this.shortText }}</div>
    <div v-if="this.hasErrors && !this.showFull"><button @click="this.showFull = true">SHOW FULL TEXT</button></div>
    <div v-if="this.hasErrors && this.showFull"><button @click="this.showFull = false">HIDE FULL TEXT</button></div>
    <div v-if="this.showFull">{{ this.error }}</div>
</template>

<style>
.shortErrorText {
    border: 1px solid #F00;
    padding: 12px;
    margin: 12px;
    color: red;
    font-size: 24pt;
}
</style>