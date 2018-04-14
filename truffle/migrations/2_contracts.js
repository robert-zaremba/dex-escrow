var Escrow = artifacts.require('./Escrow.sol')

let accounts = [
  '0xfff61542cf8599767bd006f25d698a7e2ec3a316',
  '0x1b42b0422555c4fab7fdbd287ba0fdd6280051b1',
  '0x6b2f7b1789231690d090688732adc97c3efc2bed'
]

module.exports = function (deployer) {
  console.log()
  deployer.deploy(Escrow, accounts[1], web3.toWei(1))
}
