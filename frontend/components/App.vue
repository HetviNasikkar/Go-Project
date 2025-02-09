<template>
  <div>
    <h1>Crypto Dashboard</h1>
    <CryptoPrices />
    <MarketTrends />
    <News />
  </div>
</template>

<script>
import CryptoPrices from "./CryptoPrices.vue";
import MarketTrends from "./MarketTrends.vue";
import News from "./News.vue";

export default {
  components: {
    CryptoPrices,
    MarketTrends,
    News
  }
};
</script>
