var fs = require("fs")
var express = require('express');
var http = require('http');
var https = require('https');
var privateKey = fs.readFileSync('sslcert/notifai.key', 'utf8');
var certificate = fs.readFileSync('sslcert/1_notifai.us_bundle.crt', 'utf8');

var app = express();

var credentials = {key: privateKey, cert: certificate}
var httpServer = http.createServer(app);
var httpsServer = https.createServer(credentials, app)

app.get('/api/notification', function(req, res) {
  var body = "";
  req.on('data', function (chunk) {
    body += chunk;
  });
  req.on('end', function () {
    console.log('body: ' + body);
    var jsonObj = JSON.parse(body);
  })
  res.end('Hello, World!');
});

httpServer.listen(8080);
httpsServer.listen(8443);
console.log('HTTP: Listening on port 8080...');
console.log('HTTPS: Listening on port 8080...');