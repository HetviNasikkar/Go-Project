<template>
  <div class="market-trends">
    <!-- Market Trends Section -->
    <h2>Market Trends</h2>

    <!-- Show loading message while fetching data -->
    <div v-if="loading" class="loading">Loading...</div>

    <!-- Show error message if data fetching fails -->
    <div v-else-if="error" class="error">{{ error }}</div>

    <!-- Display market trends when data is successfully fetched -->
    <div v-else class="trends-grid">
      <!-- Loop through marketTrends and create a card for each trend -->
      <div v-for="(trend, index) in marketTrends" :key="index" class="trend-card">
        <div class="trend-icon">{{ trend.icon }}</div> <!-- Trend icon (could be an emoji or symbol) -->
        <div class="trend-details">
          <h3>{{ trend.trend }}</h3> <!-- Trend title -->
          <p><strong>Top Movers:</strong> {{ trend.top_movers }}</p> <!-- List of top-moving assets -->
          <p>{{ trend.description }}</p> <!-- Short description of the trend -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";

export default {
  setup() {
    const marketTrends = ref([]); // Stores market trends data
    const loading = ref(true); // Indicates loading state
    const error = ref(null); // Stores error message if fetching fails

    // Function to fetch market trends data from API
    const fetchMarketTrends = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/market-trends"); // Fetch data from local API
        if (!res.ok) throw new Error("Failed to fetch market trends"); // Handle error if response is not OK
        marketTrends.value = await res.json(); // Store fetched data
      } catch (err) {
        error.value = "Error fetching market trends: " + err.message; // Store error message
      } finally {
        loading.value = false; // Disable loading state after fetching
      }
    };

    onMounted(fetchMarketTrends); // Fetch data when the component is mounted

    return { marketTrends, loading, error }; // Return state variables
  }
};
</script>

<style scoped>
.market-trends {
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

.trends-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr); /* 2 columns layout */
  gap: 20px; /* Space between trend cards */
}

.trend-card {
  background-color: white;
  border-radius: 8px; /* Rounded corners */
  padding: 16px; /* Padding inside each card */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow effect */
  display: flex; /* Arrange icon and details side by side */
  align-items: center; /* Align icon and text vertically */
}

.trend-icon {
  font-size: 32px; /* Large icon size */
  margin-right: 16px; /* Space between icon and text */
}

.trend-details {
  flex: 1; /* Take remaining space */
}

.trend-details h3 {
  margin: 0 0 8px; /* Add spacing below title */
  color: #333; /* Dark gray text */
}

.trend-details p {
  margin: 4px 0; /* Add spacing between paragraphs */
  color: #666; /* Medium gray text */
}

/* Responsive layout for smaller screens */
@media (max-width: 768px) {
  .trends-grid {
    grid-template-columns: 1fr; /* Switch to single-column layout */
  }
}
</style>
