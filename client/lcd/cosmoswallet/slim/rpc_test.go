package slim

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/slim/funcInlocal/bech32local"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/slim/funcInlocal/ed25519local"
	"testing"
)

func TestQOSQueryAccountGet(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addr := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Aout := QOSQueryAccountGet(addr)
	t.Log(Aout)
}

func TestQSCQueryAccountGet(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addr := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	Aout := QSCQueryAccountGet(addr)
	t.Log(Aout)
}

func TestQSCtransferSendStr(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addrto := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	coinstr := "10000qos"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	chainid := "qos-test"
	Tout := QSCtransferSendStr(addrto, coinstr, privkey, chainid)
	t.Log(Tout)

}


func TestTransferTxToQSC(t *testing.T) {
	t.Log("1111111111111111")

	SetBlockchainEntrance("47.105.52.237:26657", "forQmoonAddr")
	t.Log("2222222222222222")

	privkey := "j7wbT1lfG0ZRObMKwyYkVFpwRW0OlzFcgVL6ZwQJjeicPcStxppyWMpD1tfFg2YEcDu5JH+M6g+EhSBmjHq0bg=="
	chainid := "test-chain-xHEkEv"
	Tout := AdvertisersTrue(privkey, "ATOM", "100000", chainid)
	t.Log(Tout)


	result:=BroadcastTransferTxToQSC(Tout,"sync")
	t.Log(result)

}


func TestQueryAccount(t *testing.T) {
	t.Log("1111111111111111")

	SetBlockchainEntrance("47.105.52.237:26657", "forQmoonAddr")
	t.Log("2222222222222222")

	privkey := "j7wbT1lfG0ZRObMKwyYkVFpwRW0OlzFcgVL6ZwQJjeicPcStxppyWMpD1tfFg2YEcDu5JH+M6g+EhSBmjHq0bg=="


	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}

	addrben32,_ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	address, _ := getAddrFromBech32(addrben32)

	fmt.Println("address",string(address),addrben32)

	account,err := RpcQueryAccount(address)

	if err!=nil{
		t.Error(err.Error())
	}else{
		t.Log("111",account)

	}


}
