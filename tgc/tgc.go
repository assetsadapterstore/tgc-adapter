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

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "TGC"
}

//LoadAssetsConfig 加载外部配置
//func (wm *WalletManager) LoadAssetsConfig(c config.Configer) error {
//
//	wm.WalletManager.LoadAssetsConfig(c)
//
//	//设置官方主站的API
//	masterAPI := c.String("masterAPI")
//	wm.MasterAPI = eos.New(masterAPI)
//	wm.MasterAPI.Debug = false
//	wm.Api.Debug = false
//
//	return nil
//}
