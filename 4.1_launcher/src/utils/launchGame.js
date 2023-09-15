const { shell } = require('electron');

const launchGame = (gamePath) => {
  shell.openPath(gamePath)
    .then(() => {
      console.log('Game launched successfully');
    })
    .catch((error) => {
      console.error('Failed to launch game:', error);
    });
};

module.exports = launchGame;