/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tgc

import (
	"github.com/blocktree/eosio-adapter/eosio"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/eoscanada/eos-go"
)

const (
	maxAddresNum = 10000
)

type WalletManager struct {
	*eosio.WalletManager
	MasterAPI *eos.API // 连接官方主站的API，需要通过主站广播交易
}

func NewWalletManager(cacheManager openwallet.ICacheManager) *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = eosio.NewWalletManager(cacheManager)
	wm.Config = eosio.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	wm.Decoder = NewAddressDecoder(&wm)
	wm.DecoderV2 = NewAddressDecoder(&wm)
	//wm.TxDecoder = NewTransactionDecoder(&wm)
	return &wm
}
