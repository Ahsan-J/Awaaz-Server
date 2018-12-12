let _ = require('lodash');

module.exports = {
  responseObjectMiddleware: function (req, res, next) {
    res.sendApi = function (response,data) {
      if(_.isUndefined(response.status) || _.isUndefined(response.data) || _.isUndefined(response.message)) {
        return 
      }
      res.send(res)
    }
    next();
  },

  allowCrossDomain: function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE');
    res.header('Access-Control-Allow-Headers', '*');
    next();
  }
}