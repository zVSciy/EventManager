const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/app_review/api',
    createProxyMiddleware({
      target: 'http://review_api:8083',
      changeOrigin: true,
      pathRewrite: {
        '^/app_review/api': '', // Remove /api prefix when forwarding to the backend
      },
    })
  );
};