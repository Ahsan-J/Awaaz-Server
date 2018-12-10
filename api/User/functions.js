let { db } = require('../../config/database')

module.exports = {

  index: (req, res) => {
    let sql = 'SELECT * FROM user';
    db.query(sql, (error, results, fields) => {
      if (error) throw error
      return res.json(results)
    })
  }

}