// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DoSimpleArb

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
)

// DoSimpleArbMetaData contains all meta data concerning the DoSimpleArb contract.
var DoSimpleArbMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"flashAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"flashToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paybackToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"printMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611bc68061007e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806313f62e1d14610838578063715018a61461084d5780638da5cb5b14610855578063f2fde38b14610881575b60008061005c3660048184611a43565b81019061006991906115a4565b9350505091506000806000806000808680602001905181019061008c91906117cb565b955095509550955095509550816000815181106100d2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610173576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600160248201527f310000000000000000000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff881630146101f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600160248201527f3200000000000000000000000000000000000000000000000000000000000000604482015260640161016a565b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83600181518110610248577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151886040518363ffffffff1660e01b815260040161028f92919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b602060405180830381600087803b1580156102a957600080fd5b505af11580156102bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e19190611675565b508160018151811061031c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663022c0d9f83600181518110610379577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b1580156103c657600080fd5b505afa1580156103da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103fe9190611588565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610437576000610439565b825b84600181518110610473577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b1580156104c057600080fd5b505afa1580156104d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f89190611588565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610531576000610533565b835b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815260048101929092526024820152306044820152608060648201526000608482015260a401600060405180830381600087803b15801561059d57600080fd5b505af11580156105b1573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8360008151811061060b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151876040518363ffffffff1660e01b815260040161065292919073ffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b602060405180830381600087803b15801561066c57600080fd5b505af1158015610680573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a49190611675565b508273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6106e060005473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260019073ffffffffffffffffffffffffffffffffffffffff8816906370a082319060240160206040518083038186803b15801561074857600080fd5b505afa15801561075c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078091906116e3565b61078a9190611af9565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401602060405180830381600087803b1580156107f557600080fd5b505af1158015610809573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082d9190611675565b505050505050505050005b61084b6108463660046116fb565b610894565b005b61084b61126d565b6000546040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61084b61088f366004611565565b6112fa565b600080826000815181106108d1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561091e57600080fd5b505afa158015610932573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109569190611695565b50915091506000610b788460008151811061099a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b1580156109e757600080fd5b505afa1580156109fb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1f9190611588565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610a575782610a59565b835b6dffffffffffffffffffffffffffff1685600081518110610aa3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610af057600080fd5b505afa158015610b04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b289190611588565b73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610b605784610b62565b835b6dffffffffffffffffffffffffffff168961142a565b90508684600081518110610bb5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610c0257600080fd5b505afa158015610c16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3a9190611588565b73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610c725782610c74565b835b6dffffffffffffffffffffffffffff161015610c9257505050611267565b60008085600181518110610ccf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015610d1c57600080fd5b505afa158015610d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d549190611695565b50915091506000610f7687600181518110610d98577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610de557600080fd5b505afa158015610df9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1d9190611588565b73ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1614610e555782610e57565b835b6dffffffffffffffffffffffffffff1688600181518110610ea1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b158015610eee57600080fd5b505afa158015610f02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f269190611588565b73ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff1614610f5e5784610f60565b835b6dffffffffffffffffffffffffffff168c61147e565b90508381111561126057600087600081518110610fbc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b15801561100957600080fd5b505afa15801561101d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110419190611588565b73ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff161461107a57600061107c565b8a5b90506000886000815181106110ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561110757600080fd5b505afa15801561111b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113f9190611588565b73ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff161461117857600061117a565b8b5b905060008c878d8d8d88604051602001611199969594939291906118ac565b6040516020818303038152906040529050896000815181106111e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663022c0d9f848430856040518563ffffffff1660e01b815260040161122a9493929190611933565b600060405180830381600087803b15801561124457600080fd5b505af1158015611258573d6000803e3d6000fd5b505050505050505b5050505050505b50505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161016a565b6112f860006114cd565b565b60005473ffffffffffffffffffffffffffffffffffffffff16331461137b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161016a565b73ffffffffffffffffffffffffffffffffffffffff811661141e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161016a565b611427816114cd565b50565b6000806114378386611abc565b611443906103e8611abc565b905060006114518486611af9565b61145d906103e5611abc565b90506114698183611a83565b611474906001611a6b565b9695505050505050565b60008061148d836103e5611abc565b9050600061149b8583611abc565b90506000826114ac886103e8611abc565b6114b69190611a6b565b90506114c28183611a83565b979650505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516dffffffffffffffffffffffffffff8116811461156057600080fd5b919050565b600060208284031215611576578081fd5b813561158181611b6e565b9392505050565b600060208284031215611599578081fd5b815161158181611b6e565b600080600080608085870312156115b9578283fd5b84356115c481611b6e565b9350602085810135935060408601359250606086013567ffffffffffffffff808211156115ef578384fd5b818801915088601f830112611602578384fd5b81358181111561161457611614611b3f565b611644847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016119d0565b91508082528984828501011115611659578485fd5b8084840185840137810190920192909252939692955090935050565b600060208284031215611686578081fd5b81518015158114611581578182fd5b6000806000606084860312156116a9578283fd5b6116b284611542565b92506116c060208501611542565b9150604084015163ffffffff811681146116d8578182fd5b809150509250925092565b6000602082840312156116f4578081fd5b5051919050565b60008060008060808587031215611710578384fd5b8435935060208086013561172381611b6e565b9350604086013561173381611b6e565b9250606086013567ffffffffffffffff81111561174e578283fd5b8601601f8101881361175e578283fd5b803561177161176c82611a1f565b6119d0565b8082825284820191508484018b868560051b8701011115611790578687fd5b8694505b838510156117bb5780356117a781611b6e565b835260019490940193918501918501611794565b50979a9699509497505050505050565b60008060008060008060c087890312156117e3578182fd5b86519550602080880151955060408801516117fd81611b6e565b606089015190955061180e81611b6e565b608089015190945067ffffffffffffffff81111561182a578384fd5b8801601f81018a1361183a578384fd5b805161184861176c82611a1f565b8082825284820191508484018d868560051b8701011115611867578788fd5b8794505b8385101561189257805161187e81611b6e565b83526001949094019391850191850161186b565b50809650505050505060a087015190509295509295509295565b600060c082018883526020888185015273ffffffffffffffffffffffffffffffffffffffff8089166040860152808816606086015260c0608086015282875180855260e0870191508389019450855b818110156119195785518416835294840194918401916001016118fb565b5050809450505050508260a0830152979650505050505050565b84815260006020858184015273ffffffffffffffffffffffffffffffffffffffff85166040840152608060608401528351806080850152825b818110156119885785810183015185820160a00152820161196c565b81811115611999578360a083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160a0019695505050505050565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611a1757611a17611b3f565b604052919050565b600067ffffffffffffffff821115611a3957611a39611b3f565b5060051b60200190565b60008085851115611a52578182fd5b83861115611a5e578182fd5b5050820193919092039150565b60008219821115611a7e57611a7e611b10565b500190565b600082611ab7577f4e487b710000000000000000000000000000000000000000000000000000000081526012600452602481fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611af457611af4611b10565b500290565b600082821015611b0b57611b0b611b10565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461142757600080fdfea2646970667358221220bb331e3bcf95e75e88d8c9cc11e655b82fc8415cb85e9d5ea1c58e165f65917264736f6c63430008040033",
}

// DoSimpleArbABI is the input ABI used to generate the binding from.
// Deprecated: Use DoSimpleArbMetaData.ABI instead.
var DoSimpleArbABI = DoSimpleArbMetaData.ABI

// DoSimpleArbBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DoSimpleArbMetaData.Bin instead.
var DoSimpleArbBin = DoSimpleArbMetaData.Bin

// DeployDoSimpleArb deploys a new Ethereum contract, binding an instance of DoSimpleArb to it.
func DeployDoSimpleArb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DoSimpleArb, error) {
	parsed, err := DoSimpleArbMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DoSimpleArbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DoSimpleArb{DoSimpleArbCaller: DoSimpleArbCaller{contract: contract}, DoSimpleArbTransactor: DoSimpleArbTransactor{contract: contract}, DoSimpleArbFilterer: DoSimpleArbFilterer{contract: contract}}, nil
}

// DoSimpleArb is an auto generated Go binding around an Ethereum contract.
type DoSimpleArb struct {
	DoSimpleArbCaller     // Read-only binding to the contract
	DoSimpleArbTransactor // Write-only binding to the contract
	DoSimpleArbFilterer   // Log filterer for contract events
}

// DoSimpleArbCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoSimpleArbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoSimpleArbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoSimpleArbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoSimpleArbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoSimpleArbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoSimpleArbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoSimpleArbSession struct {
	Contract     *DoSimpleArb      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoSimpleArbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoSimpleArbCallerSession struct {
	Contract *DoSimpleArbCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DoSimpleArbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoSimpleArbTransactorSession struct {
	Contract     *DoSimpleArbTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DoSimpleArbRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoSimpleArbRaw struct {
	Contract *DoSimpleArb // Generic contract binding to access the raw methods on
}

// DoSimpleArbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoSimpleArbCallerRaw struct {
	Contract *DoSimpleArbCaller // Generic read-only contract binding to access the raw methods on
}

// DoSimpleArbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoSimpleArbTransactorRaw struct {
	Contract *DoSimpleArbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoSimpleArb creates a new instance of DoSimpleArb, bound to a specific deployed contract.
func NewDoSimpleArb(address common.Address, backend bind.ContractBackend) (*DoSimpleArb, error) {
	contract, err := bindDoSimpleArb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoSimpleArb{DoSimpleArbCaller: DoSimpleArbCaller{contract: contract}, DoSimpleArbTransactor: DoSimpleArbTransactor{contract: contract}, DoSimpleArbFilterer: DoSimpleArbFilterer{contract: contract}}, nil
}

// NewDoSimpleArbCaller creates a new read-only instance of DoSimpleArb, bound to a specific deployed contract.
func NewDoSimpleArbCaller(address common.Address, caller bind.ContractCaller) (*DoSimpleArbCaller, error) {
	contract, err := bindDoSimpleArb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoSimpleArbCaller{contract: contract}, nil
}

// NewDoSimpleArbTransactor creates a new write-only instance of DoSimpleArb, bound to a specific deployed contract.
func NewDoSimpleArbTransactor(address common.Address, transactor bind.ContractTransactor) (*DoSimpleArbTransactor, error) {
	contract, err := bindDoSimpleArb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoSimpleArbTransactor{contract: contract}, nil
}

// NewDoSimpleArbFilterer creates a new log filterer instance of DoSimpleArb, bound to a specific deployed contract.
func NewDoSimpleArbFilterer(address common.Address, filterer bind.ContractFilterer) (*DoSimpleArbFilterer, error) {
	contract, err := bindDoSimpleArb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoSimpleArbFilterer{contract: contract}, nil
}

// bindDoSimpleArb binds a generic wrapper to an already deployed contract.
func bindDoSimpleArb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DoSimpleArbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoSimpleArb *DoSimpleArbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoSimpleArb.Contract.DoSimpleArbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoSimpleArb *DoSimpleArbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.DoSimpleArbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoSimpleArb *DoSimpleArbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.DoSimpleArbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoSimpleArb *DoSimpleArbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoSimpleArb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoSimpleArb *DoSimpleArbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoSimpleArb *DoSimpleArbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoSimpleArb *DoSimpleArbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DoSimpleArb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoSimpleArb *DoSimpleArbSession) Owner() (common.Address, error) {
	return _DoSimpleArb.Contract.Owner(&_DoSimpleArb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoSimpleArb *DoSimpleArbCallerSession) Owner() (common.Address, error) {
	return _DoSimpleArb.Contract.Owner(&_DoSimpleArb.CallOpts)
}

// PrintMoney is a paid mutator transaction binding the contract method 0x13f62e1d.
//
// Solidity: function printMoney(uint256 flashAmount, address flashToken, address paybackToken, address[] path) returns()
func (_DoSimpleArb *DoSimpleArbTransactor) PrintMoney(opts *bind.TransactOpts, flashAmount *big.Int, flashToken common.Address, paybackToken common.Address, path []common.Address) (*types.Transaction, error) {
	return _DoSimpleArb.contract.Transact(opts, "printMoney", flashAmount, flashToken, paybackToken, path)
}

// PrintMoney is a paid mutator transaction binding the contract method 0x13f62e1d.
//
// Solidity: function printMoney(uint256 flashAmount, address flashToken, address paybackToken, address[] path) returns()
func (_DoSimpleArb *DoSimpleArbSession) PrintMoney(flashAmount *big.Int, flashToken common.Address, paybackToken common.Address, path []common.Address) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.PrintMoney(&_DoSimpleArb.TransactOpts, flashAmount, flashToken, paybackToken, path)
}

// PrintMoney is a paid mutator transaction binding the contract method 0x13f62e1d.
//
// Solidity: function printMoney(uint256 flashAmount, address flashToken, address paybackToken, address[] path) returns()
func (_DoSimpleArb *DoSimpleArbTransactorSession) PrintMoney(flashAmount *big.Int, flashToken common.Address, paybackToken common.Address, path []common.Address) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.PrintMoney(&_DoSimpleArb.TransactOpts, flashAmount, flashToken, paybackToken, path)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoSimpleArb *DoSimpleArbTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoSimpleArb.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoSimpleArb *DoSimpleArbSession) RenounceOwnership() (*types.Transaction, error) {
	return _DoSimpleArb.Contract.RenounceOwnership(&_DoSimpleArb.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoSimpleArb *DoSimpleArbTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DoSimpleArb.Contract.RenounceOwnership(&_DoSimpleArb.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoSimpleArb *DoSimpleArbTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DoSimpleArb.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoSimpleArb *DoSimpleArbSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.TransferOwnership(&_DoSimpleArb.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoSimpleArb *DoSimpleArbTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.TransferOwnership(&_DoSimpleArb.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DoSimpleArb *DoSimpleArbTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DoSimpleArb.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DoSimpleArb *DoSimpleArbSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.Fallback(&_DoSimpleArb.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DoSimpleArb *DoSimpleArbTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DoSimpleArb.Contract.Fallback(&_DoSimpleArb.TransactOpts, calldata)
}

// DoSimpleArbOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DoSimpleArb contract.
type DoSimpleArbOwnershipTransferredIterator struct {
	Event *DoSimpleArbOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DoSimpleArbOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoSimpleArbOwnershipTransferred)
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
		it.Event = new(DoSimpleArbOwnershipTransferred)
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
func (it *DoSimpleArbOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoSimpleArbOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoSimpleArbOwnershipTransferred represents a OwnershipTransferred event raised by the DoSimpleArb contract.
type DoSimpleArbOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoSimpleArb *DoSimpleArbFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DoSimpleArbOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DoSimpleArb.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DoSimpleArbOwnershipTransferredIterator{contract: _DoSimpleArb.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoSimpleArb *DoSimpleArbFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DoSimpleArbOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DoSimpleArb.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoSimpleArbOwnershipTransferred)
				if err := _DoSimpleArb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoSimpleArb *DoSimpleArbFilterer) ParseOwnershipTransferred(log types.Log) (*DoSimpleArbOwnershipTransferred, error) {
	event := new(DoSimpleArbOwnershipTransferred)
	if err := _DoSimpleArb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
