let express           = require('express');
let socket            = require('socket.io');
var createError       = require('http-errors');
var path              = require('path');
var cookieParser      = require('cookie-parser');
var logger            = require('morgan');
let bodyParser        = require('body-parser');
let debug             = require('debug')('awaaz-express:server');
let customMiddlewares = require('./config/middlewareFunctons.js')

// Init the Server
let app           = express();
let port          = 9000;

app.use(customMiddlewares.allowCrossDomain);
app.use(customMiddlewares.responseObjectMiddleware);
app.use(logger('dev'));
app.use(cookieParser());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(express.static(path.join(__dirname, 'uploads')));

// catch 404 and forward to error handler
// app.use(function (req, res, next) {
//   next(createError(404));
// });

// Using the imported Routes
app.use('/user',require('./api/User/routes.js'));


const server = app.listen(port)

serverWithIO = socket(server)
app.on('error', onError);
app.on('listening', onListening);

module.exports =  serverWithIO;

// Server Index path
app.get('/', function (request, response) {
  response.send('<h1>Server Running</h1>')
})

// error handler
app.use(function (err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.send('error');
});

// Listening Function
function onListening() {
  var addr = server.address();
  debug('Listening on ' + port);
}

// on Error
function onError(error) {
  if (error.syscall !== 'listen') {
    throw error;
  }

  var bind = typeof port === 'string'
    ? 'Pipe ' + port
    : 'Port ' + port;

  // handle specific listen errors with friendly messages
  switch (error.code) {
    case 'EACCES':
      console.error(bind + ' requires elevated privileges');
      process.exit(1);
      break;
    case 'EADDRINUSE':
      console.error(bind + ' is already in use');
      process.exit(1);
      break;
    default:
      throw error;
  }
}