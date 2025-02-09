Crypto Tracking App
This is a simple crypto tracking app built with Vue 3 for the frontend and Go for the backend. The app provides real-time cryptocurrency prices, market trends, and news updates.

Features
Displays live cryptocurrency prices with percentage changes.
Shows current market trends and top movers.
Fetches the latest crypto-related news.
Installation and Setup
Clone the repository.
Navigate to the frontend directory and install dependencies using:
bash
Copy
Edit
npm install
Start the frontend server with:
bash
Copy
Edit
npm run dev
Navigate to the backend directory and run the Go server:
bash
Copy
Edit
go run main.go
The frontend will be available at http://localhost:5173 and the backend at http://localhost:8080.
API Endpoints
GET /api/crypto - Fetches live cryptocurrency prices.
GET /api/market-trends - Provides market trend data.
GET /api/news - Retrieves crypto-related news.
Technologies Used
Frontend: Vue 3, Vite
Backend: Go (net/http)
Database: Not applicable (fetching data from external APIs)
Notes
Ensure that both frontend and backend servers are running simultaneously. If the backend is running on a different port, update API URLs in the Vue components accordingly.