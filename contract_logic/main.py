class ERC721:
    mintMap = {}
    ipfsMap = {}
    
    def _mint(self, address, itemId):
        self.mintMap[itemId] = address

    def _setTokenURI(self, itemId, tokenURI):
        self.ipfsMap[itemId] = tokenURI
    
    
class veriCredNFT(ERC721):
    tokenId = 0
    verifiedOrg = {}
    
    def __init__(self):
        self.collectionName = "document"
        self.symbol = "VCRD"
        self.msg_sender = "0x416c696365000000000000000000000000000000000000000000000000000000"

    def mintDoc(self, address, tokenURI):
        
        try:
            if self.verifiedOrg[self.msg_sender]:            
                tokenId += 1
                self._mint(address, tokenURI)
                self._setTokenURI(tokenId, address)
                return tokenId
            else:
                return "University not verified."
            
        except Exception as e:
            return "University is not listed."

    def newOrg(self, orgAddress):
        
        # if msg_sender == owner:
        self.verifiedOrg[orgAddress] = True
        return