module.exports = {
    ci: {
      collect: {
        url: [
          'http://localhost:3000/',
          'http://localhost:3000/about',
          'http://localhost:3000/login',
          'http://localhost:3000/register',
        ],
        startServerCommand: 'npm run build && npm run serve',
      },
      assert: {
        assertions: {
          "categories:performance": ["error", {"minScore": 0.9}],
          "categories:accessibility": ["error", {"minScore": 0.9}],
          "categories:best-practices": ["error", {"minScore": 0.9}],
          "categories:seo": ["error", {"minScore": 0.9}],
        },
      },
      upload: {
        target: 'temporary-public-storage',
      },
    },
  };