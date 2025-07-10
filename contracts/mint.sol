// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract vericredNFT is ERC721URIStorage {
    address public owner;
    uint256 private _tokenIds;
    mapping(address => bool) public verifiedOrgs;

    constructor() ERC721("document", "VCRD") {
        owner = msg.sender;
    }

    function mintDoc(
        address user,
        string memory tokenURI
    ) public returns (uint256) {
        require(verifiedOrgs[msg.sender], "University is not verified/listed");

        _tokenIds += 1;
        uint256 newItemId = _tokenIds;
        _mint(user, newItemId);
        _setTokenURI(newItemId, tokenURI);
        return newItemId;
    }

    function newOrg(address _org) public {
        require(msg.sender == owner, "Only the owner can push new orgs.");
        verifiedOrgs[_org] = true;
    }
}
