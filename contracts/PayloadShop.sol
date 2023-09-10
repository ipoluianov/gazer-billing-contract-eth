// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PayloadShop {
    struct Record {
        bool isRegistered;
        address owner;
        bytes32 payload;
    }

    struct RecordsOfUser {
        uint256[] recordIDs;
    }

    mapping(uint256 => Record) public records;
    mapping(address => RecordsOfUser) private recordsOfUser;
    uint256 public recordsCount = 0;
    address public owner;

    event Registered(
        address indexed _owner,
        bytes32 _payload
    );

    constructor() {
        owner = msg.sender;
    }

    function getBalance() public view returns(uint) {
        return address(this).balance;
    }

    function getRecordsOfUser(
        address addr
    ) public view returns (Record[] memory) {
        Record[] memory result = new Record[](
            recordsOfUser[addr].recordIDs.length
        );
        for (uint i = 0; i < recordsOfUser[addr].recordIDs.length; i++) {
            Record storage rec = records[recordsOfUser[addr].recordIDs[i]];
            result[i] = rec;
        }
        return result;
    }

    function buy(bytes32 _payload) public payable returns (uint256) {
        uint256 recordId = recordsCount + 1;
        records[recordId] = Record(
            true,
            msg.sender,
            _payload
        );
        recordsOfUser[msg.sender].recordIDs.push(recordId);
        recordsCount++;

        emit Registered(msg.sender, _payload);

        return recordId;
    }

    function withdrawAll(address payable _to) public {
        require(owner == _to);
        _to.transfer(address(this).balance);
    }
}
