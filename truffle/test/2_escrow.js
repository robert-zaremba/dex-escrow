const Escrow = artifacts.require('Escrow')
let { assertException, assertRevert } = require('./utils')

/*
let accounts = ['0x989b8dab9711460e099cde907133729ad9980e08', '0x3d20b3d9783bcbc0376fa5fc9e832ea7cdf18204', '0x599fe06e68b8c6a9d787b035d831558e9f4c5e04']
let escrow
Escrow.deployed((x) => escrow = x)
web3.eth.sendTransaction({from: seller, to: buyer, value: 1e18})
*/

let eth = web3.eth

contract('Escrow', function(accounts) {
  let escrow
  let sellerSrv = accounts[0]
  let seller = accounts[1]
  let buyer = accounts[2]

  before(async () => {
    escrow = await Escrow.deployed()
  })

  it('only accept exact payment', async () => {
    let call = () =>
      eth.sendTransaction({ from: seller, to: escrow.address, value: web3.toWei(0.1) })
    assertException(call)
  })

  it('can receive payment', async () => {
    let toSend = web3.toWei(1)
    await eth.sendTransaction({from: seller, to: escrow.address, value: toSend})
    // console.log('>>> ok', ok)
    // console.log('>>> sender after', eth.getTransactionReceipt(ok))

    let a = await eth.getBalance(escrow.address)
    assert.equal(a.toNumber(), toSend)
    assert.equal(await escrow.sellerOK(), true)
  })

  it('can\'t withdraw before confirmation', async () => {
    let call = () =>
      escrow.withdraw({ from: buyer })
    await assertException(call)
  })

  it('only sellerSrv can confirm', async () => {
    await assertException(() => escrow.confirmBuyerPayment(buyer))
    await escrow.confirmBuyerPayment(buyer, {from: sellerSrv})
  })

  it('only buyer can withdraw', async () => {
    let call = () =>
      escrow.withdraw({from: seller})
    await assertException(call)
    await escrow.withdraw({from: buyer})
  })
})
