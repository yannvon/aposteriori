pragma solidity >=0.4.17;

contract Cascade {

    string public message;

    function readMessage(string memory initialMessage) public {
        message = initialMessage;
    }

    function setMessage(string memory newMessage) public {
        message = newMessage;
    }
}