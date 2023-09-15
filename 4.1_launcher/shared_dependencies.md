the app is: Launcher App using Electron and React with Login, Game Lobby, and Game Introduction pages, SCSS styling, multiple device support, and RTK Query for API connections

the files we have decided to generate are: src/App.jsx, src/index.js, src/pages/Login.jsx, src/pages/GameLobby.jsx, src/pages/GameIntroduction.jsx, src/components/Logo.jsx, src/components/GameCard.jsx, src/api/launcherAPI.js, src/store.js, src/styles/main.scss, src/styles/_variables.scss, src/styles/_mixins.scss, src/styles/_login.scss, src/styles/_gameLobby.scss, src/styles/_gameIntroduction.scss

Shared dependencies:
1. Exported variables:
   - launcherAPI (from src/api/launcherAPI.js)

2. Data schemas:
   - User: { key: string }
   - Game: { id: string, name: string, imageUrl: string, description: string, version: string }

3. DOM element id names:
   - login-input (Login page input field)
   - confirm-button (Login page Confirm button)
   - game-list (Game Lobby page game list)
   - game-card-{id} (Game Card component with unique game id)
   - download-button (Game Introduction page download button)
   - launch-button (Game Introduction page launch button)

4. Message names:
   - LOGIN_SUCCESS
   - LOGIN_FAILURE
   - FETCH_GAMES_SUCCESS
   - FETCH_GAMES_FAILURE
   - FETCH_GAME_INTRODUCTION_SUCCESS
   - FETCH_GAME_INTRODUCTION_FAILURE
   - DOWNLOAD_GAME_SUCCESS
   - DOWNLOAD_GAME_FAILURE

5. Function names:
   - validateKey (Login page)
   - fetchGames (Game Lobby page)
   - fetchGameIntroduction (Game Introduction page)
   - downloadGame (Game Introduction page)
   - launchGame (Game Introduction page)