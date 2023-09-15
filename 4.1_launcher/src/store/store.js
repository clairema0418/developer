import { configureStore } from '@reduxjs/toolkit';
import authSlice from './authSlice';
import gamesSlice from './gamesSlice';

const store = configureStore({
  reducer: {
    auth: authSlice,
    games: gamesSlice,
  },
});

export default store;