module.exports.assertRevert = async function (promise) {
  try {
    await promise
    assert.fail('Expected exception not received')
  } catch (error) {
    const revertFound = error.message.search('revert') >= 0
    assert(revertFound, `Expected "revert", got ${error} instead`)
  }
}

module.exports.assertException = async function (promise, msg) {
  try {
    await promise
    assert.fail('Expected exception not received')
  } catch (error) {
    const revertFound = error.message.search('invalid opcode') >= 0 ||
          error.message.search('assert.fail') >= 0

    assert(revertFound, `Expected "invalid opcode", got ${error} instead. ` + msg || '')
  }
}
