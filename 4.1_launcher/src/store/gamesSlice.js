import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  games: [],
  status: 'idle',
  error: null,
};

const gamesSlice = createSlice({
  name: 'games',
  initialState,
  reducers: {
    fetchGamesStart: (state) => {
      state.status = 'loading';
    },
    fetchGamesSuccess: (state, action) => {
      state.status = 'succeeded';
      state.games = action.payload;
      state.error = null;
    },
    fetchGamesFailure: (state, action) => {
      state.status = 'failed';
      state.error = action.payload;
    },
  },
});

export const { fetchGamesStart, fetchGamesSuccess, fetchGamesFailure } = gamesSlice.actions;

export default gamesSlice.reducer;