<template>
  <div class="crypto-news">
    <!-- Section title -->
    <h2>Crypto News</h2>

    <!-- Show loading message while data is being fetched -->
    <div v-if="loading" class="loading">Loading...</div>

    <!-- Show error message if fetching data fails -->
    <div v-else-if="error" class="error">{{ error }}</div>

    <!-- Display the news list when data is successfully fetched -->
    <ul v-else class="news-list">
      <!-- Loop through the news array and create a list item for each news article -->
      <li v-for="newsItem in news" :key="newsItem.title" class="news-item">
        <strong>{{ newsItem.title }}</strong> <!-- News title -->
        - <small>{{ newsItem.source }}</small> <!-- Source of the news -->
        <br />
        <a :href="newsItem.url" target="_blank" class="read-more">Read more</a> <!-- Link to the full article -->
      </li>
    </ul>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";

export default {
  setup() {
    const news = ref([]); // Stores fetched news articles
    const loading = ref(true); // Indicates if data is loading
    const error = ref(null); // Stores error message if fetching fails

    // Function to fetch news from API
    const fetchNews = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/news"); // Fetch data from local API
        if (!response.ok) throw new Error("Failed to fetch news"); // Handle non-successful responses
        news.value = await response.json(); // Store fetched news data
      } catch (err) {
        error.value = "Error fetching news: " + err.message; // Store error message
      } finally {
        loading.value = false; // Disable loading state once data fetching is complete
      }
    };

    onMounted(fetchNews); // Fetch news when the component is mounted

    return { news, loading, error }; // Return state variables
  }
};
</script>

<style scoped>
.crypto-news {
  padding: 20px;
  background: rgba(255, 255, 255, 0.8); /* Semi-transparent white background */
  border-radius: 8px; /* Rounded corners */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow for depth */
  margin-bottom: 20px; /* Space below the section */
}

h2 {
  text-align: center; /* Center align heading */
  color: #333; /* Dark gray text */
  font-size: 2rem; /* Large font size */
  margin-bottom: 16px; /* Space below heading */
}

.loading, .error {
  text-align: center; /* Center align loading and error messages */
  font-size: 18px; /* Readable font size */
  color: #666; /* Medium gray text */
}

.news-list {
  list-style-type: none; /* Remove default list styles */
  padding: 0; /* Remove default padding */
}

.news-item {
  margin-bottom: 16px; /* Space between news items */
  padding: 12px; /* Padding inside each news item */
  background-color: white; /* White background for contrast */
  border-radius: 8px; /* Rounded corners */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow */
}

.read-more {
  color: #007bff; /* Blue color for links */
  text-decoration: none; /* Remove underline from links */
}

.read-more:hover {
  text-decoration: underline; /* Underline on hover for better UX */
}
</style>
