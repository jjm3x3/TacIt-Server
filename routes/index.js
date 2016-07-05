var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.post('/item', function(req, res) {
    console.log(req.body);

    var pgp = require('pg-promise')();
    var db = pgp('postgress://tacIt:@localhost:5432/tacItDb');

    db.none('insert into items (thing) values(\'' + req.body.text + '\');');

    res.send("thing recived");
});

router.get('/item', function(req, res) {
    res.send("here is a thing");
});

module.exports = router;
