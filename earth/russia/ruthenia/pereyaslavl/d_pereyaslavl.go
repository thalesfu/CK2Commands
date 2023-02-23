package pereyaslavl

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/pereyaslavl/hradyzk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/pereyaslavl/pereyaslavl"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/pereyaslavl/priluk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/pereyaslavl/voin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PereyaslavlDuke interface {
	feud.Duke
	CHradyzk格拉季济克() hradyzk.HradyzkCounty
	CPereyaslavl佩列亚斯拉夫尔() pereyaslavl.PereyaslavlCounty
	CPriluk普里卢基() priluk.PrilukCounty
	CVoin沃因() voin.VoinCounty
}

type 佩列亚斯拉夫尔PereyaslavlDuke struct {
	feud.BaseDuke
	格拉季济克Hradyzk       hradyzk.HradyzkCounty
	佩列亚斯拉夫尔Pereyaslavl pereyaslavl.PereyaslavlCounty
	普里卢基Priluk         priluk.PrilukCounty
	沃因Voin             voin.VoinCounty
}

func (d *佩列亚斯拉夫尔PereyaslavlDuke) CHradyzk格拉季济克() hradyzk.HradyzkCounty {
	return d.格拉季济克Hradyzk
}

func (d *佩列亚斯拉夫尔PereyaslavlDuke) CPereyaslavl佩列亚斯拉夫尔() pereyaslavl.PereyaslavlCounty {
	return d.佩列亚斯拉夫尔Pereyaslavl
}

func (d *佩列亚斯拉夫尔PereyaslavlDuke) CPriluk普里卢基() priluk.PrilukCounty {
	return d.普里卢基Priluk
}

func (d *佩列亚斯拉夫尔PereyaslavlDuke) CVoin沃因() voin.VoinCounty {
	return d.沃因Voin
}

var DPereyaslavl佩列亚斯拉夫尔 PereyaslavlDuke = &佩列亚斯拉夫尔PereyaslavlDuke{}

func init() {
	f := DPereyaslavl佩列亚斯拉夫尔.(*佩列亚斯拉夫尔PereyaslavlDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pereyaslavl",
		TitleName: "佩列亚斯拉夫尔",
		TitleCode: "d_pereyaslavl",
		Counties:  map[string]feud.County{},
	}

	f.格拉季济克Hradyzk = hradyzk.CHradyzk格拉季济克
	f.格拉季济克Hradyzk.SetParent(f)

	f.佩列亚斯拉夫尔Pereyaslavl = pereyaslavl.CPereyaslavl佩列亚斯拉夫尔
	f.佩列亚斯拉夫尔Pereyaslavl.SetParent(f)

	f.普里卢基Priluk = priluk.CPriluk普里卢基
	f.普里卢基Priluk.SetParent(f)

	f.沃因Voin = voin.CVoin沃因
	f.沃因Voin.SetParent(f)

}
