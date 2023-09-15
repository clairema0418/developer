import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

const baseURL = 'https://api.example.com';

export const launcherAPI = createApi({
  reducerPath: 'launcherAPI',
  baseQuery: fetchBaseQuery({ baseUrl: baseURL }),
  endpoints: (builder) => ({
    validateKey: builder.query({
      query: (key) => `validate?key=${key}`,
    }),
    fetchGames: builder.query({
      query: (key) => `games?key=${key}`,
    }),
    fetchGameIntroduction: builder.query({
      query: (gameId) => `game/${gameId}/introduction`,
    }),
    fetchGameVersion: builder.query({
      query: (key, gameId) => `game/${gameId}/version?key=${key}`,
    }),
  }),
});

export const {
  useValidateKeyQuery,
  useFetchGamesQuery,
  useFetchGameIntroductionQuery,
  useFetchGameVersionQuery,
} = launcherAPI;