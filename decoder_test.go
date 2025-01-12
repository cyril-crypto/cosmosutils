package cosmosutils_test

import (
    "encoding/hex"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"

    "github.com/calvinlauyh/cosmosutils"
)

func TestDecodeInvalidBase64Tx(t *testing.T) {
    decoder := cosmosutils.NewDecoder()
    _, err := decoder.DecodeBase64("!@#$%^&*()_+=-`")

    assert.EqualError(t, err, "illegal base64 data at input byte 0")
}

func TestDecodeInvalidTxBytes(t *testing.T) {
    decoder := cosmosutils.NewDecoder()
    _, err := decoder.Decode([]byte("INVALID"))

    assert.EqualError(t, err, "expected 2 wire type, got 1: tx parse error")
}

func TestDecodeTxBytes(t *testing.T) {
    anyTxBytes, _ := hex.DecodeString("0A94010A91010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412710A2B7463726F31666D70726D30736A79366C7A396C6C7637726C746E307632617A7A7763777A766B326C73796E122B7463726F31373832676E39687A7161766563756B64617171636C76736E70636B346D747A3376777A70786C1A150A08626173657463726F120931303030303030303012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210359A154BA210C489DA46626A4D631C6F8A471A2FBDA342164DD5FC4A158901F2612040A02087F180412150A100A08626173657463726F12043130303010904E1A40E812FBA1D50115CDDD534C7675B838E6CE7169CDD5AD312182696CABB76F05B82B9557EE1A2842A5579B363F0A5F2E487F7228A81FBF81E125E58F264A29DA07")
    expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\",\"to_address\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\",\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"100000000\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_LEGACY_AMINO_JSON\"}},\"sequence\":\"4\"}],\"fee\":{\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"1000\"}],\"gas_limit\":\"10000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"6BL7odUBFc3dU0x2dbg45s5xac3VrTEhgmlsq7dvBbgrlVfuGihCpVebNj8KXy5If3IoqB+/geEl5Y8mSinaBw==\"]}"
    decoder := cosmosutils.DefaultDecoder

    tx, err := decoder.Decode(anyTxBytes)
    assert.NoError(t, err)
    fmt.Println(tx)

    actual, err := tx.MarshalToJSON()
    assert.NoError(t, err)
    assert.Equal(t, expected, string(actual))
}

func TestDecodeAminoJSONSignedTx(t *testing.T) {
    anyAminoSignModeMsgSendTx := "CpQBCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKK3Rjcm8xZm1wcm0wc2p5Nmx6OWxsdjdybHRuMHYyYXp6d2N3enZrMmxzeW4SK3Rjcm8xNzgyZ245aHpxYXZlY3VrZGFxcWNsdnNucGNrNG10ejN2d3pweGwaFQoIYmFzZXRjcm8SCTEwMDAwMDAwMBJpClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDWaFUuiEMSJ2kZiak1jHG+KRxovvaNCFk3V/EoViQHyYSBAoCCH8YBBIVChAKCGJhc2V0Y3JvEgQxMDAwEJBOGkDoEvuh1QEVzd1TTHZ1uDjmznFpzdWtMSGCaWyrt28FuCuVV+4aKEKlV5s2PwpfLkh/ciioH7+B4SXljyZKKdoH"
    expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\",\"to_address\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\",\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"100000000\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_LEGACY_AMINO_JSON\"}},\"sequence\":\"4\"}],\"fee\":{\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"1000\"}],\"gas_limit\":\"10000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"6BL7odUBFc3dU0x2dbg45s5xac3VrTEhgmlsq7dvBbgrlVfuGihCpVebNj8KXy5If3IoqB+/geEl5Y8mSinaBw==\"]}"
    decoder := cosmosutils.DefaultDecoder

    tx, err := decoder.DecodeBase64(anyAminoSignModeMsgSendTx)
    assert.NoError(t, err)

    actual, err := tx.MarshalToJSON()
    assert.NoError(t, err)
    assert.Equal(t, expected, string(actual))
}

func TestDecodeDirectSignedTx(t *testing.T) {
    anyDirectSignedTx := "CpQBCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKK3Rjcm8xZm1wcm0wc2p5Nmx6OWxsdjdybHRuMHYyYXp6d2N3enZrMmxzeW4SK3Rjcm8xNzgyZ245aHpxYXZlY3VrZGFxcWNsdnNucGNrNG10ejN2d3pweGwaFQoIYmFzZXRjcm8SCTEwMDAwMDAwMBJpClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDWaFUuiEMSJ2kZiak1jHG+KRxovvaNCFk3V/EoViQHyYSBAoCCAEYBBIVChAKCGJhc2V0Y3JvEgQxMDAwEJBOGkAKlWtAxcXT06uPxtdgT2OthBz3WaW4Kz0fJ3XVq3VLxyIqL12OkpbTmKJUsK8avHzu4pz1CXYBlaUoNiPZvqUr"
    expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\",\"to_address\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\",\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"100000000\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_DIRECT\"}},\"sequence\":\"4\"}],\"fee\":{\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"1000\"}],\"gas_limit\":\"10000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"CpVrQMXF09Orj8bXYE9jrYQc91mluCs9Hyd11at1S8ciKi9djpKW05iiVLCvGrx87uKc9Ql2AZWlKDYj2b6lKw==\"]}"
    decoder := cosmosutils.DefaultDecoder

    tx, err := decoder.DecodeBase64(anyDirectSignedTx)
    assert.NoError(t, err)

    actual, err := tx.MarshalToJSON()
    assert.NoError(t, err)
    assert.Equal(t, expected, string(actual))
}

func TestDecodeDirectSignedEthermintTx(t *testing.T) {
    anyDirectSignedTx := "CokBCoYBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmYKKmV0aDFxNm1meHIydGgwMDQ1NjY5ZXZ1a3k3enR6azAzMzlzcDA3cjY5eBIqZXRoMXE2bWZ4cjJ0aDAwNDU2NjlldnVreTd6dHprMDMzOXNwMDdyNjl4GgwKBXN0YWtlEgMxMDASrwEKVwpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQP1JPeLIYCqftFicq5d01/wsBRfBX/7zr5J2FLFzBUxNxIECgIIARJUCk4KRGliYy82QjVBNjY0QkYwQUY0RjcxQjJGMEJBQTMzMTQxRTJGMTMyMTI0MkZCRDVEMTk3NjJGNTQxRUM5NzFBQ0IwODY1EgYyMDAwMDAQwJoMGkF06dsyVQtVOyciSHdCUTJydzTRhn7r9jaiyFHmDSzm9yHmExrJ9vfmSFJN0owQOHhBlwrXFdCKO+pNXYduWzq3AQ=="
    expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"eth1q6mfxr2th0045669evuky7ztzk0339sp07r69x\",\"to_address\":\"eth1q6mfxr2th0045669evuky7ztzk0339sp07r69x\",\"amount\":[{\"denom\":\"stake\",\"amount\":\"100\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/ethermint.crypto.v1.ethsecp256k1.PubKey\",\"key\":\"A/Uk94shgKp+0WJyrl3TX/CwFF8Ff/vOvknYUsXMFTE3\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_DIRECT\"}},\"sequence\":\"0\"}],\"fee\":{\"amount\":[{\"denom\":\"ibc/6B5A664BF0AF4F71B2F0BAA33141E2F1321242FBD5D19762F541EC971ACB0865\",\"amount\":\"200000\"}],\"gas_limit\":\"200000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"dOnbMlULVTsnIkh3QlEycnc00YZ+6/Y2oshR5g0s5vch5hMayfb35khSTdKMEDh4QZcK1xXQijvqTV2Hbls6twE=\"]}"
    decoder := cosmosutils.DefaultDecoder

    tx, err := decoder.DecodeBase64(anyDirectSignedTx)
    assert.NoError(t, err)

    actual, err := tx.MarshalToJSON()
    assert.NoError(t, err)
    assert.Equal(t, expected, string(actual))
}

func TestDecodeIBCBase64Tx(t *testing.T) {
    decoder := cosmosutils.DefaultDecoder
    //https://www.mintscan.io/cosmos/txs/E3944CB1AF60EB1649B2DC9EBB6C67FFCCE17C6E44259417F4C4A657D9206B6E
    _, err := decoder.DecodeBase64("CqABCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKLWNvc21vczF6ZHNmZTN0c3dza2h4eWh1amNuNGFuODdxcGZ1OWdmNG5renJlbhItY29zbW9zMWhwMmc1M3N0cWs3cnBsd3N4ZzA2MHMwNXRmdHpqNGRmZDNzYTB5GhEKBXVhdG9tEggyMzc5MTgxORIKMTU5NjQyNjI0NRJlCk4KRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDbihAoG4eoUY0R8iAGJ1uQEJKb8XwAatZR0f+HzrnFlUSBAoCCH8SEwoNCgV1YXRvbRIEMjAwMBCA8QQaQJ0OaKsjTwhwtJVwXemWhJR2etmR4e/z+o+g2PMfo1J/J/9R/qgLnxVjbRfAPjMf6XfLQcxdlflR7DP+CR+5yA4=")
    assert.NoError(t, err)
}

func TestDecodeNFTBase64Tx(t *testing.T) {
    decoder := cosmosutils.DefaultDecoder
    _, err := decoder.DecodeBase64("Cm0KawofL2NoYWlubWFpbi5uZnQudjEuTXNnSXNzdWVEZW5vbRJICgdkZW5vbWlkEglEZW5vbU5hbWUaBlNjaGVtYSIqY3JvMW5rNHJxM3E0Nmx0Z2pnaHh6ODBoeTM4NXA5dWowdGY1OGFwa2NkEmgKTgpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQIiwYZ9iW+Ai1OXWKSgdyxw6YyBZXa7EFg+i1gx8o6VjBIECgIIARIWChAKB2Jhc2Vjcm8SBTUwMDAwEMCaDBpAVYD6pk8lSj7eCfNJ62EikuQgOejcZuE4XSvZyGyXyOlthWjXisZEgIgHYZSh6R8iRGmiKh/bmCwRvaJN5GxDOQ==")
    assert.NoError(t, err)
}
