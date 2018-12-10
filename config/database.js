var mysql = require('mysql');

module.exports = {

  /* *Local Configuration* */
  // db: mysql.createConnection({
  //   host: 'localhost',
  //   user: 'root',
  //   password: '',
  //   database: 'awaaz_case_store'
  // }),

  /* *DB4FREE configuration* */
  db: mysql.createConnection({
    host: 'db4free.net',
    port: 3306,
    user: 'awaaz_admin',
    password: 'qwerty12345',
    database: 'awaaz_case_store'
  })
}