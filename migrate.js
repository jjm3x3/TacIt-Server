var pgp = require('pg-promise')();
var db = pgp('postgress://tacIt:@localhost:5432/tacItDb');

exports.up = function(){
    createTable("items", '(id serial, thing text)');
    // makeNotNull("items", "id");
    // makeSerial("items", "id");

}

function createTable(name,values){
    db.none('create table ' + name +
            ' ' + values + ';');
}

function makeNotNull(tableName, coulmnName) {
    db.one('alter table ' + tableName + 
            ' alter column ' + coulmnName + ' set not null;')
    .then(function() {
        consoe.log("all is right!");
    }).catch( function(error) {
        console.log(error);
    });

}


function makeSerial(tableName,columnName) {
    db.one('alter table ' + tableName +
           ' alter column ' + columnName + ' set serial');
}
