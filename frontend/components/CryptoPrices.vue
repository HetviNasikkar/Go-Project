<template>
  <div>
    <h2>Crypto Prices</h2>
    
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>
    <table v-else>
      <thead>
        <tr>
          <th>Coin</th>
          <th>Price</th>
          <th>Change</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="coin in cryptoData" :key="coin.name">
          <td>{{ coin.name }}</td>
          <td>${{ coin.price.toFixed(2) }}</td>
          <td :class="getChangeClass(coin.change)">
            {{ coin.change.toFixed(2) }}%
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";

export default {
  setup() {
    const cryptoData = ref([]);
    const loading = ref(true);
    const error = ref(null);

    const fetchCryptoData = async () => {
      try {
        const response = await fetch("/api/crypto"); // Use relative API path
        if (!response.ok) throw new Error("Failed to fetch crypto data");
        cryptoData.value = await response.json();
      } catch (err) {
        error.value = "Error fetching crypto data: " + err.message;
      } finally {
        loading.value = false;
      }
    };

    const getChangeClass = (change) => {
      return change > 0 ? "positive" : change < 0 ? "negative" : "";
    };

    onMounted(fetchCryptoData);

    return { cryptoData, loading, error, getChangeClass };
  }
};
</script>

<style>
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  border: 1px solid black;
  padding: 8px;
  text-align: left;
}
.positive {
  color: green;
}
.negative {
  color: red;
}
</style>
