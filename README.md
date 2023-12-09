# HTMX Exploration

A simple web application exploring the usage of htmx in a text-based role-playing game (RPG) setting. Players can create characters, explore a town, venture into a forest, and engage in combat encounters.

The theme draws inspiration from the classic online text-based game [Legend of the Green Dragon](https://lotgd.net/).

## Technologies Used

- **Go (Golang):** The backend server is built with Go, utilizing the Echo web framework.
  - **Echo:** Webserver
  - **Air:** Live reloading
- **HTML & Tailwind CSS:** The frontend is designed with HTML and styled using Tailwind CSS.
- **htmx:** htmx is employed to enable dynamic and interactive features without the need for heavy JavaScript.

## Project Structure

The project is organized into distinct modules:

- **town:** Handles interactions within the town, such as character creation.
- **user:** Manages user-related functionality, including character attributes and state.
- **combat:** Implements combat mechanics, including user and enemy attacks.
- **forest:** Governs forest exploration and encounters with enemies.
- **enemy:** Defines enemy-related logic and services.

## How to Run

1. Clone the repository:

```bash
git clone https://github.com/JackTheReckless/htmx-exploration.git
cd htmx-exploration
```

2. Install dependencies

```bash
go mod tidy
```

3. Run the application

```bash
go run main.go
```

4. Open your browser and navigate to <http://localhost:8080>

5. Game Master is behind the names gm, admin, jon, or jack
