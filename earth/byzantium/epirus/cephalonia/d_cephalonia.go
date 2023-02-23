package cephalonia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/cephalonia/cephalonia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/cephalonia/corfu"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/cephalonia/zakynthos"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CephaloniaDuke interface {
	feud.Duke
	CCephalonia凯法洛尼亚() cephalonia.CephaloniaCounty
	CCorfu克尔基拉() corfu.CorfuCounty
	CZakynthos扎金索斯() zakynthos.ZakynthosCounty
}

type 凯法洛尼亚CephaloniaDuke struct {
	feud.BaseDuke
	凯法洛尼亚Cephalonia cephalonia.CephaloniaCounty
	克尔基拉Corfu       corfu.CorfuCounty
	扎金索斯Zakynthos   zakynthos.ZakynthosCounty
}

func (d *凯法洛尼亚CephaloniaDuke) CCephalonia凯法洛尼亚() cephalonia.CephaloniaCounty {
	return d.凯法洛尼亚Cephalonia
}

func (d *凯法洛尼亚CephaloniaDuke) CCorfu克尔基拉() corfu.CorfuCounty {
	return d.克尔基拉Corfu
}

func (d *凯法洛尼亚CephaloniaDuke) CZakynthos扎金索斯() zakynthos.ZakynthosCounty {
	return d.扎金索斯Zakynthos
}

var DCephalonia凯法洛尼亚 CephaloniaDuke = &凯法洛尼亚CephaloniaDuke{}

func init() {
	f := DCephalonia凯法洛尼亚.(*凯法洛尼亚CephaloniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cephalonia",
		TitleName: "凯法洛尼亚",
		TitleCode: "d_cephalonia",
		Counties:  map[string]feud.County{},
	}

	f.凯法洛尼亚Cephalonia = cephalonia.CCephalonia凯法洛尼亚
	f.凯法洛尼亚Cephalonia.SetParent(f)

	f.克尔基拉Corfu = corfu.CCorfu克尔基拉
	f.克尔基拉Corfu.SetParent(f)

	f.扎金索斯Zakynthos = zakynthos.CZakynthos扎金索斯
	f.扎金索斯Zakynthos.SetParent(f)

}
