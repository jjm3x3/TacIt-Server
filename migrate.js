var pgp = require('pg-promise')();
var db = pgp('postgress://tacIt:secret@localhost:5432/tacItDb');

exports.up = function(){
    createTable("items", '(id serial, thing text)');
    console.log("migration ran");
    // makeNotNull("items", "id");
    // makeSerial("items", "id");

}

function createTable(name,values){
    result = db.none('create table ' + name +
            ' ' + values + ';')
                .then(function(){
			console.log("it worked")
                 }).catch(function(error){
			console.log(error);
});
    console.log(result)
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
