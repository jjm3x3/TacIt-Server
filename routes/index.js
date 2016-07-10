var express = require('express');
var router = express.Router();

var pgp = require('pg-promise')();
var db = pgp('postgress://tacIt:@localhost:5432/tacItDb');

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.post('/item', function(req, res) {
    console.log(req.body);


    db.none('insert into items (thing) values($1);', [req.body.text])
        .catch( function(error) {
        console.log(error);
    });

    res.redirect(301, 'http://localhost:3000/item');
});

router.get('/item/new', function(req, res) {
    res.send('<p> Enter a note <p>' +
             '<form method=POST action=/item>' + 
             '<p><textarea name=text rows="10" cols="80"></textarea> </p>' +
             '<p><input type=submit value=\"TacIt\"></p>' + 
             '</form>');
});

router.get('/item', function(req, res) {
    db.any('select * from items;').then(function(results) {
        res.setHeader('Content-Type', 'application/json');
        json = JSON.stringify(results);
        res.send('{\"items\":' + json + '}');
    });



    // for web (HTML)
    // var result = '<ul>';
    // db.any('select * from items;').then(function(results) {
    //     for ( i = 0; i < results.length; ++i){
    //         result += "<li>" + results[i].thing + "</li><br>";
    //     }
    //     result += "</ul>";
    //     result += "<a href=\"http://localhost:3000/item/new\">New Note</a>"
    //     console.log(result);
    //     res.send(result);
    // });
});

module.exports = router;
