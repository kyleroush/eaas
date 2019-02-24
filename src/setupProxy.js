const proxy = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(proxy('/Prod', { target: 'https://7siim5bbqh.execute-api.us-east-1.amazonaws.com/Prod' }));
};