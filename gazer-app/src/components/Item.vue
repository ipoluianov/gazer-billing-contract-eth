<script>
import wallet from "./lib/wallet"
import bytes32 from "./lib/bytes32"

export default {
    props: ['domain'],
    emits: ['registerSubdomain'],
    computed: {
        xchgAddressString() {
            return wallet.bytes32ToXchgAddressString(this.domain.payload);
        }
    },
    methods: {
    }
}
</script>

<template>
    <div class="item">
        <div class="subitem">
            <div class="xchg_addr">#{{ this.xchgAddressString }}</div>
        </div>
    </div>
    
</template>

<style>
.item {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    border-bottom: 1px solid #555;
}

.subitem {
    display: flex;
    justify-content: space-between;
    flex-direction: row;
}

.xchg_addr {
    font-size: 11pt;
    color: #ffd700;
    font-family: monospace;
}

.button_register_subdomain {
    cursor: pointer;
}
</style>
