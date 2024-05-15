const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
  app.use(
    '/tmf-api',
    createProxyMiddleware({
      target: 'https://localhost:3030',
      changeOrigin: true,
      secure: false,
      pathRewrite: {
        '^/tmf-api': '/tmf-api',
      },
    })
  );

};