package savonia

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/savonia/joensuu"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/savonia/kuopio"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/savonia/laukaa"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/savonia/savolaks"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SavoniaDuke interface {
	feud.Duke
	CJoensuu约恩苏() joensuu.JoensuuCounty
	CKuopio卡拉韦西() kuopio.KuopioCounty
	CLaukaa凯泰莱() laukaa.LaukaaCounty
	CSavolaks萨沃() savolaks.SavolaksCounty
}

type 萨沃尼亚SavoniaDuke struct {
	feud.BaseDuke
	约恩苏Joensuu joensuu.JoensuuCounty
	卡拉韦西Kuopio kuopio.KuopioCounty
	凯泰莱Laukaa  laukaa.LaukaaCounty
	萨沃Savolaks savolaks.SavolaksCounty
}

func (d *萨沃尼亚SavoniaDuke) CJoensuu约恩苏() joensuu.JoensuuCounty {
	return d.约恩苏Joensuu
}

func (d *萨沃尼亚SavoniaDuke) CKuopio卡拉韦西() kuopio.KuopioCounty {
	return d.卡拉韦西Kuopio
}

func (d *萨沃尼亚SavoniaDuke) CLaukaa凯泰莱() laukaa.LaukaaCounty {
	return d.凯泰莱Laukaa
}

func (d *萨沃尼亚SavoniaDuke) CSavolaks萨沃() savolaks.SavolaksCounty {
	return d.萨沃Savolaks
}

var DSavonia萨沃尼亚 SavoniaDuke = &萨沃尼亚SavoniaDuke{}

func init() {
	f := DSavonia萨沃尼亚.(*萨沃尼亚SavoniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "savonia",
		TitleName: "萨沃尼亚",
		TitleCode: "d_savonia",
		Counties:  map[string]feud.County{},
	}

	f.约恩苏Joensuu = joensuu.CJoensuu约恩苏
	f.约恩苏Joensuu.SetParent(f)

	f.卡拉韦西Kuopio = kuopio.CKuopio卡拉韦西
	f.卡拉韦西Kuopio.SetParent(f)

	f.凯泰莱Laukaa = laukaa.CLaukaa凯泰莱
	f.凯泰莱Laukaa.SetParent(f)

	f.萨沃Savolaks = savolaks.CSavolaks萨沃
	f.萨沃Savolaks.SetParent(f)

}
