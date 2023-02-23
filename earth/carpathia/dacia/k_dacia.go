package dacia

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/bessarabia"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/moldau"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/oltenia"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/wallachia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DaciaKingdom interface {
	feud.Kingdom
	DBessarabia比萨拉比亚() bessarabia.BessarabiaDuke
	DMoldau摩尔达维亚() moldau.MoldauDuke
	DOltenia奥尔泰尼亚() oltenia.OlteniaDuke
	DWallachia蒙泰尼亚() wallachia.WallachiaDuke
}

type 瓦拉吉亚DaciaKingdom struct {
	feud.BaseKingdom
	比萨拉比亚Bessarabia bessarabia.BessarabiaDuke
	摩尔达维亚Moldau     moldau.MoldauDuke
	奥尔泰尼亚Oltenia    oltenia.OlteniaDuke
	蒙泰尼亚Wallachia   wallachia.WallachiaDuke
}

func (k *瓦拉吉亚DaciaKingdom) DBessarabia比萨拉比亚() bessarabia.BessarabiaDuke {
	return k.比萨拉比亚Bessarabia
}

func (k *瓦拉吉亚DaciaKingdom) DMoldau摩尔达维亚() moldau.MoldauDuke {
	return k.摩尔达维亚Moldau
}

func (k *瓦拉吉亚DaciaKingdom) DOltenia奥尔泰尼亚() oltenia.OlteniaDuke {
	return k.奥尔泰尼亚Oltenia
}

func (k *瓦拉吉亚DaciaKingdom) DWallachia蒙泰尼亚() wallachia.WallachiaDuke {
	return k.蒙泰尼亚Wallachia
}

var KDacia瓦拉吉亚 DaciaKingdom = &瓦拉吉亚DaciaKingdom{}

func init() {
	f := KDacia瓦拉吉亚.(*瓦拉吉亚DaciaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "dacia",
		TitleName: "瓦拉吉亚",
		TitleCode: "k_dacia",
		Dukes:     map[string]feud.Duke{},
	}

	f.比萨拉比亚Bessarabia = bessarabia.DBessarabia比萨拉比亚
	f.比萨拉比亚Bessarabia.SetParent(f)

	f.摩尔达维亚Moldau = moldau.DMoldau摩尔达维亚
	f.摩尔达维亚Moldau.SetParent(f)

	f.奥尔泰尼亚Oltenia = oltenia.DOltenia奥尔泰尼亚
	f.奥尔泰尼亚Oltenia.SetParent(f)

	f.蒙泰尼亚Wallachia = wallachia.DWallachia蒙泰尼亚
	f.蒙泰尼亚Wallachia.SetParent(f)

}
