<template>
  <div>
    <h2>Market Trends</h2>
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>
    <div v-else>
      <p><strong>Trend:</strong> {{ marketTrend.trend }}</p>
      <p><strong>Top Movers:</strong> {{ marketTrend.top_movers }}</p>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";

export default {
  setup() {
    const marketTrend = ref({});
    const loading = ref(true);
    const error = ref(null);

    const fetchMarketTrends = async () => {
      try {
        const res = await fetch("/api/market-trends");
        if (!res.ok) throw new Error("Failed to fetch market trends");
        marketTrend.value = await res.json();
      } catch (err) {
        error.value = "Error fetching market trends: " + err.message;
      } finally {
        loading.value = false;
      }
    };

    onMounted(fetchMarketTrends);

    return { marketTrend, loading, error };
  }
};
</script>
