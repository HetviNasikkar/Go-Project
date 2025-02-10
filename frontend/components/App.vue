<template>
  <div class="dashboard"> 
    <!-- Main dashboard container -->
    <h1>Crypto Dashboard</h1> 
    <!-- Title of the dashboard -->
    <CryptoPrices /> 
    <!-- Component displaying cryptocurrency prices -->
    <MarketTrends /> 
    <!-- Component displaying market trends -->
    <News /> 
    <!-- Component displaying latest crypto news -->
  </div>
</template>

<script>
import CryptoPrices from "./CryptoPrices.vue"; // Importing CryptoPrices component
import MarketTrends from "./MarketTrends.vue"; // Importing MarketTrends component
import News from "./News.vue"; // Importing News component

export default {
  components: {
    CryptoPrices, // Registering CryptoPrices component
    MarketTrends, // Registering MarketTrends component
    News, // Registering News component
  },
};
</script>

<style scoped>
.dashboard {
  font-family: Arial, sans-serif; /* Setting font family */
  padding: 20px; /* Adding padding around the dashboard */
  max-width: 1200px; /* Setting max width for layout consistency */
  margin: 0 auto; /* Centering the dashboard */
  background: linear-gradient(135deg, #f5f7fa, #c3cfe2); /* Light gradient background */
  min-height: 100vh; /* Ensuring full viewport height */
}

h1 {
  text-align: center; /* Center aligning the heading */
  color: #333; /* Setting text color */
  font-size: 2.5rem; /* Increasing font size */
  margin-bottom: 20px; /* Adding spacing below heading */
  padding: 10px; /* Adding padding inside heading */
  background: rgba(255, 255, 255, 0.8); /* Semi-transparent white background */
  border-radius: 8px; /* Rounding corners */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Adding a soft shadow effect */
}
</style>
