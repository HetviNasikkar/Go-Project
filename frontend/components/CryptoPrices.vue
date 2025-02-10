<template>
  <div class="crypto-prices">
    <!-- Crypto Prices Section -->
    <h2>Crypto Prices</h2>

    <!-- Show loading message while fetching data -->
    <div v-if="loading" class="loading">Loading...</div>

    <!-- Show error message if data fetching fails -->
    <div v-else-if="error" class="error">{{ error }}</div>

    <!-- Show crypto prices table if data is successfully fetched -->
    <table v-else class="crypto-table">
      <thead>
        <tr>
          <th>Coin</th>
          <th>Price</th>
          <th>Change</th>
        </tr>
      </thead>
      <tbody>
        <!-- Loop through cryptoData and display each coin's details -->
        <tr v-for="(value, key) in cryptoData" :key="key">
          <td>{{ key }}</td>
          <td>${{ Number(value.price).toFixed(2) }}</td>
          <!-- Apply dynamic class based on price change -->
          <td :class="getChangeClass(Number(value.change))">
            {{ Number(value.change).toFixed(2) }}%
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
    const cryptoData = ref({}); // Stores crypto price data
    const loading = ref(true); // Indicates loading state
    const error = ref(null); // Stores error message if fetching fails

    // Function to fetch crypto price data from API
    const fetchCryptoData = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/crypto"); // Fetch data from local API
        if (!response.ok) throw new Error("Failed to fetch crypto data"); // Handle error if response is not OK
        const data = await response.json();
        cryptoData.value = data; // Store fetched data
      } catch (err) {
        error.value = "Error fetching crypto data: " + err.message; // Store error message
      } finally {
        loading.value = false; // Disable loading state after fetching
      }
    };

    // Function to apply CSS class based on price change (positive or negative)
    const getChangeClass = (change) => {
      return change > 0 ? "positive" : change < 0 ? "negative" : "";
    };

    onMounted(fetchCryptoData); // Fetch data when the component is mounted

    return { cryptoData, loading, error, getChangeClass }; // Return state variables and functions
  }
};
</script>

<style scoped>
.crypto-prices {
  padding: 20px;
  background: rgba(255, 255, 255, 0.8); /* Semi-transparent white background */
  border-radius: 8px; /* Rounded corners */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow for depth */
  margin-bottom: 20px; /* Spacing between sections */
}

h2 {
  text-align: center; /* Center align heading */
  color: #333; /* Dark gray text color */
  font-size: 2rem; /* Large font size */
  margin-bottom: 16px; /* Space below heading */
}

.loading, .error {
  text-align: center; /* Center align loading and error messages */
  font-size: 18px; /* Readable font size */
  color: #666; /* Gray text */
}

.crypto-table {
  width: 100%; /* Make table take full width */
  border-collapse: collapse; /* Remove extra spacing between table borders */
  background-color: white; /* White background */
  border-radius: 8px; /* Rounded corners */
  overflow: hidden; /* Hide overflowing content */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow effect */
}

th, td {
  border: 1px solid #ddd; /* Light gray borders */
  padding: 12px; /* Add padding for better spacing */
  text-align: left; /* Align text to the left */
}

th {
  background-color: #f4f4f4; /* Light gray background for table headers */
  font-weight: bold; /* Make header text bold */
}

.positive {
  color: green; /* Green color for positive price change */
}

.negative {
  color: red; /* Red color for negative price change */
}
</style>
