var express = require('express');

var app = express();

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

app.listen(8080);
console.log('Listening on port 8080...');