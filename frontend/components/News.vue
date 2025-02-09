<template>
  <div>
    <h2>Crypto News</h2>

    <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>
    <ul v-else>
      <li v-for="newsItem in news" :key="newsItem.title">
        <strong>{{ newsItem.title }}</strong> - 
        <small>{{ newsItem.source }}</small>
        <br />
        <a :href="newsItem.url" target="_blank">Read more</a>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";

export default {
  setup() {
    const news = ref([]);
    const loading = ref(true);
    const error = ref(null);

    const fetchNews = async () => {
      try {
        const response = await fetch("/api/news"); // Use relative API path
        if (!response.ok) throw new Error("Failed to fetch news");
        news.value = await response.json();
      } catch (err) {
        error.value = "Error fetching news: " + err.message;
      } finally {
        loading.value = false;
      }
    };

    onMounted(fetchNews);

    return { news, loading, error };
  }
};
</script>

<style>
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin-bottom: 10px;
}
a {
  color: blue;
  text-decoration: none;
}
a:hover {
  text-decoration: underline;
}
</style>
