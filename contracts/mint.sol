// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract vericredNFT is ERC721URIStorage {
    uint256 private _tokenIds;

    constructor() ERC721("certificate", "VCRD") {}

    function mintDoc(
        address user,
        string memory tokenURI
    ) public returns (uint256) {
        _tokenIds += 1;
        uint256 newItemId = _tokenIds;
        _mint(user, newItemId);
        _setTokenURI(newItemId, tokenURI);
        return newItemId;
    }
}
