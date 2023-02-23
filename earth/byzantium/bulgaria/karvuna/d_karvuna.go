package karvuna

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/karvuna/constantia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/karvuna/karvuna"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria/karvuna/mesembria"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarvunaDuke interface {
	feud.Duke
	CConstantia康斯坦察() constantia.ConstantiaCounty
	CKarvuna卡尔武纳() karvuna.KarvunaCounty
	CMesembria墨森布里亚() mesembria.MesembriaCounty
}

type 卡尔武纳KarvunaDuke struct {
	feud.BaseDuke
	康斯坦察Constantia constantia.ConstantiaCounty
	卡尔武纳Karvuna    karvuna.KarvunaCounty
	墨森布里亚Mesembria mesembria.MesembriaCounty
}

func (d *卡尔武纳KarvunaDuke) CConstantia康斯坦察() constantia.ConstantiaCounty {
	return d.康斯坦察Constantia
}

func (d *卡尔武纳KarvunaDuke) CKarvuna卡尔武纳() karvuna.KarvunaCounty {
	return d.卡尔武纳Karvuna
}

func (d *卡尔武纳KarvunaDuke) CMesembria墨森布里亚() mesembria.MesembriaCounty {
	return d.墨森布里亚Mesembria
}

var DKarvuna卡尔武纳 KarvunaDuke = &卡尔武纳KarvunaDuke{}

func init() {
	f := DKarvuna卡尔武纳.(*卡尔武纳KarvunaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "karvuna",
		TitleName: "卡尔武纳",
		TitleCode: "d_karvuna",
		Counties:  map[string]feud.County{},
	}

	f.康斯坦察Constantia = constantia.CConstantia康斯坦察
	f.康斯坦察Constantia.SetParent(f)

	f.卡尔武纳Karvuna = karvuna.CKarvuna卡尔武纳
	f.卡尔武纳Karvuna.SetParent(f)

	f.墨森布里亚Mesembria = mesembria.CMesembria墨森布里亚
	f.墨森布里亚Mesembria.SetParent(f)

}
