# Echo Go React CRUD Project

This project is a CRUD application built with Go (Echo framework) and connected to a React.js frontend created with Vite. The frontend is styled using Tailwind CSS and utilizes React Router for navigation.

## Features

- Fetch and display a list of albums.
- View detailed information about a specific album by clicking on it.
- Serve static files using Echo in Go.

## Technologies Used

- **Backend**: Go, Echo framework
- **Frontend**: React.js, Vite, Tailwind CSS, React Router
- **API Requests**: Axios

## Getting Started

### Prerequisites

- Go 1.16+
- Node.js 14+

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/your-username/echo-go-react-crud.git
   cd echo-go-react-crud
   cd backend
   go mod tidy
   go run main.go
   // frontend
   cd ui
   npm i
   npm run dev
   // npm run prepare for remove ./dist and run npm run build
   ```

### Next

- [x] Improve UI
- [x] Connect PostgreSQL
- [x] Complete API operations
- [x] Login and register
- [x] Fix login, register and validate routes
- [x] Fix middlewares
