class ERC721:
    mintMap = {}
    ipfsMap = {}
    
    def _mint(self, address, itemId):
        self.mintMap[itemId] = address

    def _setTokenURI(self, itemId, tokenURI):
        self.ipfsMap[itemId] = tokenURI
    
    
class veriCredNFT(ERC721):
    tokenId = 0
    
    def __init__(self):
        self.collectionName = "document"
        self.symbol = "VCRD"

    def mintDoc(self, address, tokenURI):
        tokenId += 1
        self._mint(address, tokenURI)
        self._setTokenURI(itemId, address)
        return tokenId
