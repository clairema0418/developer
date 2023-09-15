import { launcherAPI } from "../api/launcherAPI";

async function fetchGameVersion(key, gameId) {
  try {
    const response = await launcherAPI.get(`/games/${gameId}/version`, {
      headers: { Authorization: `Bearer ${key}` },
    });
    return response.data.version;
  } catch (error) {
    console.error("Error fetching game version:", error);
    throw error;
  }
}

export default fetchGameVersion;