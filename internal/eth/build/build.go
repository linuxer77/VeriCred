// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package build

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BuildMetaData contains all meta data concerning the Build contract.
var BuildMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Orgs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allOrgs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_org\",\"type\":\"address\"}],\"name\":\"newOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifiedOrgs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600881526020017f646f63756d656e740000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f5643524400000000000000000000000000000000000000000000000000000000815250815f908161008a919061031f565b50806001908161009a919061031f565b5050503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103ee565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061015d57607f821691505b6020821081036101705761016f610119565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101d27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610197565b6101dc8683610197565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61022061021b610216846101f4565b6101fd565b6101f4565b9050919050565b5f819050919050565b61023983610206565b61024d61024582610227565b8484546101a3565b825550505050565b5f5f905090565b610264610255565b61026f818484610230565b505050565b5b81811015610292576102875f8261025c565b600181019050610275565b5050565b601f8211156102d7576102a881610176565b6102b184610188565b810160208510156102c0578190505b6102d46102cc85610188565b830182610274565b50505b505050565b5f82821c905092915050565b5f6102f75f19846008026102dc565b1980831691505092915050565b5f61030f83836102e8565b9150826002028217905092915050565b610328826100e2565b67ffffffffffffffff811115610341576103406100ec565b5b61034b8254610146565b610356828285610296565b5f60209050601f831160018114610387575f8415610375578287015190505b61037f8582610304565b8655506103e6565b601f19841661039586610176565b5f5b828110156103bc57848901518255600182019150602085019450602081019050610397565b868310156103d957848901516103d5601f8916826102e8565b8355505b6001600288020188555050505b505050505050565b6128b9806103fb5f395ff3fe608060405234801561000f575f5ffd5b5060043610610129575f3560e01c80636352211e116100ab57806395d89b411161006f57806395d89b4114610365578063a22cb46514610383578063b88d4fde1461039f578063c87b56dd146103bb578063e985e9c5146103eb57610129565b80636352211e146102875780636fae464c146102b757806370a08231146102e757806376c77fff146103175780638da5cb5b1461034757610129565b806323b872dd116100f257806323b872dd146101e557806342842e0e1461020157806345050d411461021d5780634ec6e9371461023b5780635a696f371461025757610129565b80629a9b7b1461012d57806301ffc9a71461014b57806306fdde031461017b578063081812fc14610199578063095ea7b3146101c9575b5f5ffd5b61013561041b565b6040516101429190611b5c565b60405180910390f35b61016560048036038101906101609190611bdb565b610424565b6040516101729190611c20565b60405180910390f35b610183610484565b6040516101909190611ca9565b60405180910390f35b6101b360048036038101906101ae9190611cf3565b610513565b6040516101c09190611d5d565b60405180910390f35b6101e360048036038101906101de9190611da0565b61052e565b005b6101ff60048036038101906101fa9190611dde565b610544565b005b61021b60048036038101906102169190611dde565b610643565b005b610225610662565b6040516102329190611ee5565b60405180910390f35b61025560048036038101906102509190611f05565b6106ed565b005b610271600480360381019061026c9190611cf3565b610834565b60405161027e9190611d5d565b60405180910390f35b6102a1600480360381019061029c9190611cf3565b61086f565b6040516102ae9190611d5d565b60405180910390f35b6102d160048036038101906102cc9190611f05565b610880565b6040516102de9190611c20565b60405180910390f35b61030160048036038101906102fc9190611f05565b61089d565b60405161030e9190611b5c565b60405180910390f35b610331600480360381019061032c919061205c565b610953565b60405161033e9190611b5c565b60405180910390f35b61034f610a1a565b60405161035c9190611d5d565b60405180910390f35b61036d610a3f565b60405161037a9190611ca9565b60405180910390f35b61039d600480360381019061039891906120e0565b610acf565b005b6103b960048036038101906103b491906121bc565b610ae5565b005b6103d560048036038101906103d09190611cf3565b610b0a565b6040516103e29190611ca9565b60405180910390f35b6104056004803603810190610400919061223c565b610c15565b6040516104129190611c20565b60405180910390f35b5f600854905090565b5f634906490660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061047d575061047c82610ca3565b5b9050919050565b60605f8054610492906122a7565b80601f01602080910402602001604051908101604052809291908181526020018280546104be906122a7565b80156105095780601f106104e057610100808354040283529160200191610509565b820191905f5260205f20905b8154815290600101906020018083116104ec57829003601f168201915b5050505050905090565b5f61051d82610d84565b5061052782610e0a565b9050919050565b610540828261053b610e43565b610e4a565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105b4575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105ab9190611d5d565b60405180910390fd5b5f6105c783836105c2610e43565b610e5c565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461063d578382826040517f64283d7b000000000000000000000000000000000000000000000000000000008152600401610634939291906122d7565b60405180910390fd5b50505050565b61065d83838360405180602001604052805f815250610ae5565b505050565b6060600a8054806020026020016040519081016040528092919081815260200182805480156106e357602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161069a575b5050505050905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461077c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107739061237c565b60405180910390fd5b600160095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600a81908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600a8181548110610843575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f61087982610d84565b9050919050565b6009602052805f5260405f205f915054906101000a900460ff1681565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361090e575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016109059190611d5d565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166109dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d49061240a565b60405180910390fd5b600160085f8282546109ef9190612455565b925050819055505f6008549050610a068482611067565b610a10818461115a565b8091505092915050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060018054610a4e906122a7565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7a906122a7565b8015610ac55780601f10610a9c57610100808354040283529160200191610ac5565b820191905f5260205f20905b815481529060010190602001808311610aa857829003601f168201915b5050505050905090565b610ae1610ada610e43565b83836111b4565b5050565b610af0848484610544565b610b04610afb610e43565b8585858561131d565b50505050565b6060610b1582610d84565b505f60065f8481526020019081526020015f208054610b33906122a7565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5f906122a7565b8015610baa5780601f10610b8157610100808354040283529160200191610baa565b820191905f5260205f20905b815481529060010190602001808311610b8d57829003601f168201915b505050505090505f610bba6114c9565b90505f815103610bce578192505050610c10565b5f82511115610c02578082604051602001610bea9291906124c2565b60405160208183030381529060405292505050610c10565b610c0b846114df565b925050505b919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610d6d57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610d7d5750610d7c82611545565b5b9050919050565b5f5f610d8f836115ae565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e0157826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610df89190611b5c565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610e5783838360016115e7565b505050565b5f5f610e67846115ae565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610ea857610ea78184866117a6565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610f3357610ee75f855f5f6115e7565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610fb257600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036110d7575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016110ce9190611d5d565b60405180910390fd5b5f6110e383835f610e5c565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611155575f6040517f73c6ac6e00000000000000000000000000000000000000000000000000000000815260040161114c9190611d5d565b60405180910390fd5b505050565b8060065f8481526020019081526020015f2090816111789190612685565b507ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7826040516111a89190611b5c565b60405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361122457816040517f5b08ba1800000000000000000000000000000000000000000000000000000000815260040161121b9190611d5d565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516113109190611c20565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b11156114c2578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b815260040161137b94939291906127a6565b6020604051808303815f875af19250505080156113b657506040513d601f19601f820116820180604052508101906113b39190612804565b60015b611437573d805f81146113e4576040519150601f19603f3d011682016040523d82523d5f602084013e6113e9565b606091505b505f81510361142f57836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114269190611d5d565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146114c057836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114b79190611d5d565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60606114ea82610d84565b505f6114f46114c9565b90505f8151116115125760405180602001604052805f81525061153d565b8061151c84611869565b60405160200161152d9291906124c2565b6040516020818303038152906040525b915050919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061161f57505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611751575f61162e84610d84565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561169857508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156116ab57506116a98184610c15565b155b156116ed57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016116e49190611d5d565b60405180910390fd5b811561174f57838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6117b1838383611933565b611864575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361182557806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161181c9190611b5c565b60405180910390fd5b81816040517f177e802f00000000000000000000000000000000000000000000000000000000815260040161185b92919061282f565b60405180910390fd5b505050565b60605f6001611877846119f3565b0190505f8167ffffffffffffffff81111561189557611894611f38565b5b6040519080825280601f01601f1916602001820160405280156118c75781602001600182028036833780820191505090505b5090505f82602001820190505b600115611928578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a858161191d5761191c612856565b5b0494505f85036118d4575b819350505050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156119ea57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806119ab57506119aa8484610c15565b5b806119e957508273ffffffffffffffffffffffffffffffffffffffff166119d183610e0a565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611a4f577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381611a4557611a44612856565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310611a8c576d04ee2d6d415b85acef81000000008381611a8257611a81612856565b5b0492506020810190505b662386f26fc100008310611abb57662386f26fc100008381611ab157611ab0612856565b5b0492506010810190505b6305f5e1008310611ae4576305f5e1008381611ada57611ad9612856565b5b0492506008810190505b6127108310611b09576127108381611aff57611afe612856565b5b0492506004810190505b60648310611b2c5760648381611b2257611b21612856565b5b0492506002810190505b600a8310611b3b576001810190505b80915050919050565b5f819050919050565b611b5681611b44565b82525050565b5f602082019050611b6f5f830184611b4d565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611bba81611b86565b8114611bc4575f5ffd5b50565b5f81359050611bd581611bb1565b92915050565b5f60208284031215611bf057611bef611b7e565b5b5f611bfd84828501611bc7565b91505092915050565b5f8115159050919050565b611c1a81611c06565b82525050565b5f602082019050611c335f830184611c11565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611c7b82611c39565b611c858185611c43565b9350611c95818560208601611c53565b611c9e81611c61565b840191505092915050565b5f6020820190508181035f830152611cc18184611c71565b905092915050565b611cd281611b44565b8114611cdc575f5ffd5b50565b5f81359050611ced81611cc9565b92915050565b5f60208284031215611d0857611d07611b7e565b5b5f611d1584828501611cdf565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611d4782611d1e565b9050919050565b611d5781611d3d565b82525050565b5f602082019050611d705f830184611d4e565b92915050565b611d7f81611d3d565b8114611d89575f5ffd5b50565b5f81359050611d9a81611d76565b92915050565b5f5f60408385031215611db657611db5611b7e565b5b5f611dc385828601611d8c565b9250506020611dd485828601611cdf565b9150509250929050565b5f5f5f60608486031215611df557611df4611b7e565b5b5f611e0286828701611d8c565b9350506020611e1386828701611d8c565b9250506040611e2486828701611cdf565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611e6081611d3d565b82525050565b5f611e718383611e57565b60208301905092915050565b5f602082019050919050565b5f611e9382611e2e565b611e9d8185611e38565b9350611ea883611e48565b805f5b83811015611ed8578151611ebf8882611e66565b9750611eca83611e7d565b925050600181019050611eab565b5085935050505092915050565b5f6020820190508181035f830152611efd8184611e89565b905092915050565b5f60208284031215611f1a57611f19611b7e565b5b5f611f2784828501611d8c565b91505092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611f6e82611c61565b810181811067ffffffffffffffff82111715611f8d57611f8c611f38565b5b80604052505050565b5f611f9f611b75565b9050611fab8282611f65565b919050565b5f67ffffffffffffffff821115611fca57611fc9611f38565b5b611fd382611c61565b9050602081019050919050565b828183375f83830152505050565b5f612000611ffb84611fb0565b611f96565b90508281526020810184848401111561201c5761201b611f34565b5b612027848285611fe0565b509392505050565b5f82601f83011261204357612042611f30565b5b8135612053848260208601611fee565b91505092915050565b5f5f6040838503121561207257612071611b7e565b5b5f61207f85828601611d8c565b925050602083013567ffffffffffffffff8111156120a05761209f611b82565b5b6120ac8582860161202f565b9150509250929050565b6120bf81611c06565b81146120c9575f5ffd5b50565b5f813590506120da816120b6565b92915050565b5f5f604083850312156120f6576120f5611b7e565b5b5f61210385828601611d8c565b9250506020612114858286016120cc565b9150509250929050565b5f67ffffffffffffffff82111561213857612137611f38565b5b61214182611c61565b9050602081019050919050565b5f61216061215b8461211e565b611f96565b90508281526020810184848401111561217c5761217b611f34565b5b612187848285611fe0565b509392505050565b5f82601f8301126121a3576121a2611f30565b5b81356121b384826020860161214e565b91505092915050565b5f5f5f5f608085870312156121d4576121d3611b7e565b5b5f6121e187828801611d8c565b94505060206121f287828801611d8c565b935050604061220387828801611cdf565b925050606085013567ffffffffffffffff81111561222457612223611b82565b5b6122308782880161218f565b91505092959194509250565b5f5f6040838503121561225257612251611b7e565b5b5f61225f85828601611d8c565b925050602061227085828601611d8c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806122be57607f821691505b6020821081036122d1576122d061227a565b5b50919050565b5f6060820190506122ea5f830186611d4e565b6122f76020830185611b4d565b6123046040830184611d4e565b949350505050565b7f4f6e6c7920746865206f776e65722063616e2070757368206e6577206f7267735f8201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f612366602183611c43565b91506123718261230c565b604082019050919050565b5f6020820190508181035f8301526123938161235a565b9050919050565b7f556e6976657273697479206973206e6f742076657269666965642f6c697374655f8201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b5f6123f4602183611c43565b91506123ff8261239a565b604082019050919050565b5f6020820190508181035f830152612421816123e8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61245f82611b44565b915061246a83611b44565b925082820190508082111561248257612481612428565b5b92915050565b5f81905092915050565b5f61249c82611c39565b6124a68185612488565b93506124b6818560208601611c53565b80840191505092915050565b5f6124cd8285612492565b91506124d98284612492565b91508190509392505050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026125417fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612506565b61254b8683612506565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61258661258161257c84611b44565b612563565b611b44565b9050919050565b5f819050919050565b61259f8361256c565b6125b36125ab8261258d565b848454612512565b825550505050565b5f5f905090565b6125ca6125bb565b6125d5818484612596565b505050565b5b818110156125f8576125ed5f826125c2565b6001810190506125db565b5050565b601f82111561263d5761260e816124e5565b612617846124f7565b81016020851015612626578190505b61263a612632856124f7565b8301826125da565b50505b505050565b5f82821c905092915050565b5f61265d5f1984600802612642565b1980831691505092915050565b5f612675838361264e565b9150826002028217905092915050565b61268e82611c39565b67ffffffffffffffff8111156126a7576126a6611f38565b5b6126b182546122a7565b6126bc8282856125fc565b5f60209050601f8311600181146126ed575f84156126db578287015190505b6126e5858261266a565b86555061274c565b601f1984166126fb866124e5565b5f5b82811015612722578489015182556001820191506020850194506020810190506126fd565b8683101561273f578489015161273b601f89168261264e565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f61277882612754565b612782818561275e565b9350612792818560208601611c53565b61279b81611c61565b840191505092915050565b5f6080820190506127b95f830187611d4e565b6127c66020830186611d4e565b6127d36040830185611b4d565b81810360608301526127e5818461276e565b905095945050505050565b5f815190506127fe81611bb1565b92915050565b5f6020828403121561281957612818611b7e565b5b5f612826848285016127f0565b91505092915050565b5f6040820190506128425f830185611d4e565b61284f6020830184611b4d565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea26469706673582212200d1c940f8bf0c9927a8f72078b391e3dc4b9e4301ebd4fb6de752d02fa0daa8964736f6c634300081e0033",
}

// BuildABI is the input ABI used to generate the binding from.
// Deprecated: Use BuildMetaData.ABI instead.
var BuildABI = BuildMetaData.ABI

// BuildBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuildMetaData.Bin instead.
var BuildBin = BuildMetaData.Bin

// DeployBuild deploys a new Ethereum contract, binding an instance of Build to it.
func DeployBuild(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Build, error) {
	parsed, err := BuildMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuildBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Build{BuildCaller: BuildCaller{contract: contract}, BuildTransactor: BuildTransactor{contract: contract}, BuildFilterer: BuildFilterer{contract: contract}}, nil
}

// Build is an auto generated Go binding around an Ethereum contract.
type Build struct {
	BuildCaller     // Read-only binding to the contract
	BuildTransactor // Write-only binding to the contract
	BuildFilterer   // Log filterer for contract events
}

// BuildCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuildCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuildTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuildTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuildFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuildFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuildSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuildSession struct {
	Contract     *Build            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuildCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuildCallerSession struct {
	Contract *BuildCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuildTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuildTransactorSession struct {
	Contract     *BuildTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuildRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuildRaw struct {
	Contract *Build // Generic contract binding to access the raw methods on
}

// BuildCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuildCallerRaw struct {
	Contract *BuildCaller // Generic read-only contract binding to access the raw methods on
}

// BuildTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuildTransactorRaw struct {
	Contract *BuildTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuild creates a new instance of Build, bound to a specific deployed contract.
func NewBuild(address common.Address, backend bind.ContractBackend) (*Build, error) {
	contract, err := bindBuild(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Build{BuildCaller: BuildCaller{contract: contract}, BuildTransactor: BuildTransactor{contract: contract}, BuildFilterer: BuildFilterer{contract: contract}}, nil
}

// NewBuildCaller creates a new read-only instance of Build, bound to a specific deployed contract.
func NewBuildCaller(address common.Address, caller bind.ContractCaller) (*BuildCaller, error) {
	contract, err := bindBuild(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuildCaller{contract: contract}, nil
}

// NewBuildTransactor creates a new write-only instance of Build, bound to a specific deployed contract.
func NewBuildTransactor(address common.Address, transactor bind.ContractTransactor) (*BuildTransactor, error) {
	contract, err := bindBuild(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuildTransactor{contract: contract}, nil
}

// NewBuildFilterer creates a new log filterer instance of Build, bound to a specific deployed contract.
func NewBuildFilterer(address common.Address, filterer bind.ContractFilterer) (*BuildFilterer, error) {
	contract, err := bindBuild(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuildFilterer{contract: contract}, nil
}

// bindBuild binds a generic wrapper to an already deployed contract.
func bindBuild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuildMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Build *BuildRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Build.Contract.BuildCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Build *BuildRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Build.Contract.BuildTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Build *BuildRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Build.Contract.BuildTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Build *BuildCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Build.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Build *BuildTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Build.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Build *BuildTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Build.Contract.contract.Transact(opts, method, params...)
}

// Orgs is a free data retrieval call binding the contract method 0x5a696f37.
//
// Solidity: function Orgs(uint256 ) view returns(address)
func (_Build *BuildCaller) Orgs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "Orgs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Orgs is a free data retrieval call binding the contract method 0x5a696f37.
//
// Solidity: function Orgs(uint256 ) view returns(address)
func (_Build *BuildSession) Orgs(arg0 *big.Int) (common.Address, error) {
	return _Build.Contract.Orgs(&_Build.CallOpts, arg0)
}

// Orgs is a free data retrieval call binding the contract method 0x5a696f37.
//
// Solidity: function Orgs(uint256 ) view returns(address)
func (_Build *BuildCallerSession) Orgs(arg0 *big.Int) (common.Address, error) {
	return _Build.Contract.Orgs(&_Build.CallOpts, arg0)
}

// AllOrgs is a free data retrieval call binding the contract method 0x45050d41.
//
// Solidity: function allOrgs() view returns(address[])
func (_Build *BuildCaller) AllOrgs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "allOrgs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllOrgs is a free data retrieval call binding the contract method 0x45050d41.
//
// Solidity: function allOrgs() view returns(address[])
func (_Build *BuildSession) AllOrgs() ([]common.Address, error) {
	return _Build.Contract.AllOrgs(&_Build.CallOpts)
}

// AllOrgs is a free data retrieval call binding the contract method 0x45050d41.
//
// Solidity: function allOrgs() view returns(address[])
func (_Build *BuildCallerSession) AllOrgs() ([]common.Address, error) {
	return _Build.Contract.AllOrgs(&_Build.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Build *BuildCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Build *BuildSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Build.Contract.BalanceOf(&_Build.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Build *BuildCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Build.Contract.BalanceOf(&_Build.CallOpts, owner)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Build *BuildCaller) CurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Build *BuildSession) CurrentTokenId() (*big.Int, error) {
	return _Build.Contract.CurrentTokenId(&_Build.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Build *BuildCallerSession) CurrentTokenId() (*big.Int, error) {
	return _Build.Contract.CurrentTokenId(&_Build.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Build *BuildCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Build *BuildSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.GetApproved(&_Build.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Build *BuildCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.GetApproved(&_Build.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Build *BuildCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Build *BuildSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Build.Contract.IsApprovedForAll(&_Build.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Build *BuildCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Build.Contract.IsApprovedForAll(&_Build.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Build *BuildCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Build *BuildSession) Name() (string, error) {
	return _Build.Contract.Name(&_Build.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Build *BuildCallerSession) Name() (string, error) {
	return _Build.Contract.Name(&_Build.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Build *BuildCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Build *BuildSession) Owner() (common.Address, error) {
	return _Build.Contract.Owner(&_Build.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Build *BuildCallerSession) Owner() (common.Address, error) {
	return _Build.Contract.Owner(&_Build.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Build *BuildCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Build *BuildSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.OwnerOf(&_Build.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Build *BuildCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.OwnerOf(&_Build.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Build *BuildCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Build *BuildSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Build.Contract.SupportsInterface(&_Build.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Build *BuildCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Build.Contract.SupportsInterface(&_Build.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Build *BuildCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Build *BuildSession) Symbol() (string, error) {
	return _Build.Contract.Symbol(&_Build.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Build *BuildCallerSession) Symbol() (string, error) {
	return _Build.Contract.Symbol(&_Build.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Build *BuildCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Build *BuildSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Build.Contract.TokenURI(&_Build.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Build *BuildCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Build.Contract.TokenURI(&_Build.CallOpts, tokenId)
}

// VerifiedOrgs is a free data retrieval call binding the contract method 0x6fae464c.
//
// Solidity: function verifiedOrgs(address ) view returns(bool)
func (_Build *BuildCaller) VerifiedOrgs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "verifiedOrgs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifiedOrgs is a free data retrieval call binding the contract method 0x6fae464c.
//
// Solidity: function verifiedOrgs(address ) view returns(bool)
func (_Build *BuildSession) VerifiedOrgs(arg0 common.Address) (bool, error) {
	return _Build.Contract.VerifiedOrgs(&_Build.CallOpts, arg0)
}

// VerifiedOrgs is a free data retrieval call binding the contract method 0x6fae464c.
//
// Solidity: function verifiedOrgs(address ) view returns(bool)
func (_Build *BuildCallerSession) VerifiedOrgs(arg0 common.Address) (bool, error) {
	return _Build.Contract.VerifiedOrgs(&_Build.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Build *BuildTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Build *BuildSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.Approve(&_Build.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Build *BuildTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.Approve(&_Build.TransactOpts, to, tokenId)
}

// MintDoc is a paid mutator transaction binding the contract method 0x76c77fff.
//
// Solidity: function mintDoc(address user, string tokenURI) returns(uint256)
func (_Build *BuildTransactor) MintDoc(opts *bind.TransactOpts, user common.Address, tokenURI string) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "mintDoc", user, tokenURI)
}

// MintDoc is a paid mutator transaction binding the contract method 0x76c77fff.
//
// Solidity: function mintDoc(address user, string tokenURI) returns(uint256)
func (_Build *BuildSession) MintDoc(user common.Address, tokenURI string) (*types.Transaction, error) {
	return _Build.Contract.MintDoc(&_Build.TransactOpts, user, tokenURI)
}

// MintDoc is a paid mutator transaction binding the contract method 0x76c77fff.
//
// Solidity: function mintDoc(address user, string tokenURI) returns(uint256)
func (_Build *BuildTransactorSession) MintDoc(user common.Address, tokenURI string) (*types.Transaction, error) {
	return _Build.Contract.MintDoc(&_Build.TransactOpts, user, tokenURI)
}

// NewOrg is a paid mutator transaction binding the contract method 0x4ec6e937.
//
// Solidity: function newOrg(address _org) returns()
func (_Build *BuildTransactor) NewOrg(opts *bind.TransactOpts, _org common.Address) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "newOrg", _org)
}

// NewOrg is a paid mutator transaction binding the contract method 0x4ec6e937.
//
// Solidity: function newOrg(address _org) returns()
func (_Build *BuildSession) NewOrg(_org common.Address) (*types.Transaction, error) {
	return _Build.Contract.NewOrg(&_Build.TransactOpts, _org)
}

// NewOrg is a paid mutator transaction binding the contract method 0x4ec6e937.
//
// Solidity: function newOrg(address _org) returns()
func (_Build *BuildTransactorSession) NewOrg(_org common.Address) (*types.Transaction, error) {
	return _Build.Contract.NewOrg(&_Build.TransactOpts, _org)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Build *BuildTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Build *BuildSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom0(&_Build.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Build *BuildTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom0(&_Build.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Build *BuildTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Build *BuildSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Build.Contract.SetApprovalForAll(&_Build.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Build *BuildTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Build.Contract.SetApprovalForAll(&_Build.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.TransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.TransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// BuildApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Build contract.
type BuildApprovalIterator struct {
	Event *BuildApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BuildApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BuildApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BuildApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildApproval represents a Approval event raised by the Build contract.
type BuildApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Build *BuildFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BuildApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BuildApprovalIterator{contract: _Build.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Build *BuildFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BuildApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildApproval)
				if err := _Build.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Build *BuildFilterer) ParseApproval(log types.Log) (*BuildApproval, error) {
	event := new(BuildApproval)
	if err := _Build.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Build contract.
type BuildApprovalForAllIterator struct {
	Event *BuildApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BuildApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BuildApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BuildApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildApprovalForAll represents a ApprovalForAll event raised by the Build contract.
type BuildApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Build *BuildFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BuildApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Build.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BuildApprovalForAllIterator{contract: _Build.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Build *BuildFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BuildApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Build.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildApprovalForAll)
				if err := _Build.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Build *BuildFilterer) ParseApprovalForAll(log types.Log) (*BuildApprovalForAll, error) {
	event := new(BuildApprovalForAll)
	if err := _Build.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Build contract.
type BuildBatchMetadataUpdateIterator struct {
	Event *BuildBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BuildBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BuildBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BuildBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Build contract.
type BuildBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Build *BuildFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*BuildBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Build.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &BuildBatchMetadataUpdateIterator{contract: _Build.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Build *BuildFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *BuildBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Build.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildBatchMetadataUpdate)
				if err := _Build.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Build *BuildFilterer) ParseBatchMetadataUpdate(log types.Log) (*BuildBatchMetadataUpdate, error) {
	event := new(BuildBatchMetadataUpdate)
	if err := _Build.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Build contract.
type BuildMetadataUpdateIterator struct {
	Event *BuildMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BuildMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BuildMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BuildMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildMetadataUpdate represents a MetadataUpdate event raised by the Build contract.
type BuildMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Build *BuildFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*BuildMetadataUpdateIterator, error) {

	logs, sub, err := _Build.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &BuildMetadataUpdateIterator{contract: _Build.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Build *BuildFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *BuildMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Build.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildMetadataUpdate)
				if err := _Build.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Build *BuildFilterer) ParseMetadataUpdate(log types.Log) (*BuildMetadataUpdate, error) {
	event := new(BuildMetadataUpdate)
	if err := _Build.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Build contract.
type BuildTransferIterator struct {
	Event *BuildTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BuildTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BuildTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BuildTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildTransfer represents a Transfer event raised by the Build contract.
type BuildTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Build *BuildFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BuildTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BuildTransferIterator{contract: _Build.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Build *BuildFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BuildTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildTransfer)
				if err := _Build.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Build *BuildFilterer) ParseTransfer(log types.Log) (*BuildTransfer, error) {
	event := new(BuildTransfer)
	if err := _Build.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
