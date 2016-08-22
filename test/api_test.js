var assert = require('assert');
var api = require("../lib/api");

describe('api', function () {

  var apiObj = api("http://localhost:5000/");

  describe('apiUrlRoot property', function () {
    it('The passed in api root parameter should be available on the returned object', function () {
      assert.equal(apiObj.apiRootUrl, "http://localhost:5000/");
    });
  });

  describe('keysForLfId', function() {
    it('Calling keysForLfId with an lfId returns an object with keys', function() {
      var keys = apiObj.getKeysForLfId("jwhaley")
      assert.equal(keys.keyId.length, 20, "keyId length should be 20");
      assert.equal(keys.secret.length, 40, "secret length should be 40");
    });
  });

});