package vladimir

import (
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/gorodez"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/suzdal"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/vladimir"
	"github.com/thalesfu/CK2Commands/earth/russia/vladimir/vladimir/yuryev"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VladimirDuke interface {
	feud.Duke
	CGorodez戈罗杰茨() gorodez.GorodezCounty
	CSuzdal苏兹达尔() suzdal.SuzdalCounty
	CVladimir弗拉基米尔() vladimir.VladimirCounty
	CYuryev尤里耶夫() yuryev.YuryevCounty
}

type 弗拉基米尔VladimirDuke struct {
	feud.BaseDuke
	戈罗杰茨Gorodez   gorodez.GorodezCounty
	苏兹达尔Suzdal    suzdal.SuzdalCounty
	弗拉基米尔Vladimir vladimir.VladimirCounty
	尤里耶夫Yuryev    yuryev.YuryevCounty
}

func (d *弗拉基米尔VladimirDuke) CGorodez戈罗杰茨() gorodez.GorodezCounty {
	return d.戈罗杰茨Gorodez
}

func (d *弗拉基米尔VladimirDuke) CSuzdal苏兹达尔() suzdal.SuzdalCounty {
	return d.苏兹达尔Suzdal
}

func (d *弗拉基米尔VladimirDuke) CVladimir弗拉基米尔() vladimir.VladimirCounty {
	return d.弗拉基米尔Vladimir
}

func (d *弗拉基米尔VladimirDuke) CYuryev尤里耶夫() yuryev.YuryevCounty {
	return d.尤里耶夫Yuryev
}

var DVladimir弗拉基米尔 VladimirDuke = &弗拉基米尔VladimirDuke{}

func init() {
	f := DVladimir弗拉基米尔.(*弗拉基米尔VladimirDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vladimir",
		TitleName: "弗拉基米尔",
		TitleCode: "d_vladimir",
		Counties:  map[string]feud.County{},
	}

	f.戈罗杰茨Gorodez = gorodez.CGorodez戈罗杰茨
	f.戈罗杰茨Gorodez.SetParent(f)

	f.苏兹达尔Suzdal = suzdal.CSuzdal苏兹达尔
	f.苏兹达尔Suzdal.SetParent(f)

	f.弗拉基米尔Vladimir = vladimir.CVladimir弗拉基米尔
	f.弗拉基米尔Vladimir.SetParent(f)

	f.尤里耶夫Yuryev = yuryev.CYuryev尤里耶夫
	f.尤里耶夫Yuryev.SetParent(f)

}
