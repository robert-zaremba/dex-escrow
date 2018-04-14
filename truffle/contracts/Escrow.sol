pragma solidity ^0.4.19;


contract Escrow {
    address public sellerSrv;
    address public seller;
    address public buyer;
    bool public sellerOK;
    bool public buyerOK;
    uint256 public toSell;  // amount we want to sell

    event Cancel();

    function Escrow(address _seller, uint256 _toSell) public {
        require(_seller != address(0));
        seller = _seller;
        sellerSrv = msg.sender;
        toSell = _toSell;
    }

    // eth payment handler
    function () public payable {
        // protect from random payments
        require(msg.sender == seller);
        require(msg.value == toSell);
        sellerOK = true;
    }

    function confirmBuyerPayment(address _buyer) public {
        require(msg.sender == sellerSrv);
        require(_buyer != address(0));
        buyerOK = true;
        buyer = _buyer;
    }

    function cancel() public {
        require(!sellerOK);
        buyerOK = false;
        buyer.transfer(this.balance);
        Cancel();
    }

    function withdraw() public {
        require(msg.sender == buyer);
        require(buyerOK);
        require(sellerOK);
        selfdestruct(seller);
    }
}
