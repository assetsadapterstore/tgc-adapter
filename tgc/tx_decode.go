/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tgc

import (
	"encoding/hex"
	"fmt"
	"github.com/blocktree/eosio-adapter/eosio"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/eoscanada/eos-go"
	"time"
)

// TransactionDecoder 交易单解析器
type TransactionDecoder struct {
	*eosio.TransactionDecoder
	tgcWM *WalletManager
}

//NewTransactionDecoder 交易单解析器
func NewTransactionDecoder(wm *WalletManager) *TransactionDecoder {
	decoder := TransactionDecoder{tgcWM: wm}
	decoder.TransactionDecoder = eosio.NewTransactionDecoder(wm.WalletManager)
	return &decoder
}

// SubmitRawTransaction 广播交易单
func (decoder *TransactionDecoder) SubmitRawTransaction(wrapper openwallet.WalletDAI, rawTx *openwallet.RawTransaction) (*openwallet.Transaction, error) {

	var stx eos.SignedTransaction
	txHex, err := hex.DecodeString(rawTx.RawHex)
	if err != nil {
		return nil, fmt.Errorf("transaction decode failed, unexpected error: %v", err)
	}
	err = eos.UnmarshalBinary(txHex, &stx)
	if err != nil {
		return nil, fmt.Errorf("transaction decode failed, unexpected error: %v", err)
	}

	packedTx, err := stx.Pack(eos.CompressionNone)
	if err != nil {
		return nil, err
	}

	response, err := decoder.tgcWM.MasterAPI.PushTransaction(packedTx)
	if err != nil {
		return nil, fmt.Errorf("push transaction: %s", err)
	}

	decoder.tgcWM.Log.Infof("Transaction [%s] submitted to the network successfully.", hex.EncodeToString(response.Processed.ID))

	rawTx.TxID = hex.EncodeToString(response.Processed.ID)
	rawTx.IsSubmit = true

	decimals := int32(rawTx.Coin.Contract.Decimals)
	fees := "0"

	//记录一个交易单
	tx := &openwallet.Transaction{
		From:       rawTx.TxFrom,
		To:         rawTx.TxTo,
		Amount:     rawTx.TxAmount,
		Coin:       rawTx.Coin,
		TxID:       rawTx.TxID,
		Decimal:    decimals,
		AccountID:  rawTx.Account.AccountID,
		Fees:       fees,
		SubmitTime: time.Now().Unix(),
		ExtParam:   rawTx.ExtParam,
	}

	tx.WxID = openwallet.GenTransactionWxID(tx)

	return tx, nil
}
