/*
 * Copyright 2019 The openwallet Authors
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
	"github.com/astaxie/beego/config"
	"github.com/blocktree/eosio-adapter/eosio"
	"github.com/blocktree/openwallet/log"
	"path/filepath"
	"testing"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	cache := eosio.NewCacheManager()
	wm := NewWalletManager(&cache)
	//读取配置
	absFile := filepath.Join("conf", "conf.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		panic(err)
	}
	wm.LoadAssetsConfig(c)
	return wm
}


func TestWalletManager_GetInfo(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetInfo()
	if err != nil {
		log.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetAccount(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetAccount("bitbanktest11")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetBlock(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetBlockByNum(10000)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetTransaction(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetTransaction("042308511b749ef48db1c76e0301e5a80df911c034dba97b1fcc89b4542e0b40")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetABI(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetABI("evsio.token")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetCurrencyBalance(t *testing.T) {
	wm := testNewWalletManager()
	r, err := wm.Api.GetCurrencyBalance("bitbanktest11", "TGC", "evsio.token")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	log.Infof("%+v", r)
}

func TestWalletManager_GetBlockHeight(t *testing.T) {
	wm := testNewWalletManager()
	header, err := wm.Blockscanner.GetCurrentBlockHeader()
	if err != nil {
		t.Errorf("GetCurrentBlockHeader error: %v", err)
		return
	}
	log.Infof("header: %+v", header)
}