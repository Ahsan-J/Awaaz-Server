let router = require('express').Router();
let Controller = require('./functions.js')

router.get('/', Controller['index'])

module.exports = router;