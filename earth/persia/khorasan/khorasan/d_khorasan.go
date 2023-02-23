package khorasan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/khorasan/birjand"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/khorasan/nishapur"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/khorasan/qohistan"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/khorasan/tus"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhorasanDuke interface {
	feud.Duke
	CBirjand比尔詹德() birjand.BirjandCounty
	CNishapur内沙布尔() nishapur.NishapurCounty
	CQohistan忽希思丹() qohistan.QohistanCounty
	CTus图斯() tus.TusCounty
}

type 呼罗珊KhorasanDuke struct {
	feud.BaseDuke
	比尔詹德Birjand  birjand.BirjandCounty
	内沙布尔Nishapur nishapur.NishapurCounty
	忽希思丹Qohistan qohistan.QohistanCounty
	图斯Tus        tus.TusCounty
}

func (d *呼罗珊KhorasanDuke) CBirjand比尔詹德() birjand.BirjandCounty {
	return d.比尔詹德Birjand
}

func (d *呼罗珊KhorasanDuke) CNishapur内沙布尔() nishapur.NishapurCounty {
	return d.内沙布尔Nishapur
}

func (d *呼罗珊KhorasanDuke) CQohistan忽希思丹() qohistan.QohistanCounty {
	return d.忽希思丹Qohistan
}

func (d *呼罗珊KhorasanDuke) CTus图斯() tus.TusCounty {
	return d.图斯Tus
}

var DKhorasan呼罗珊 KhorasanDuke = &呼罗珊KhorasanDuke{}

func init() {
	f := DKhorasan呼罗珊.(*呼罗珊KhorasanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khorasan",
		TitleName: "呼罗珊",
		TitleCode: "d_khorasan",
		Counties:  map[string]feud.County{},
	}

	f.比尔詹德Birjand = birjand.CBirjand比尔詹德
	f.比尔詹德Birjand.SetParent(f)

	f.内沙布尔Nishapur = nishapur.CNishapur内沙布尔
	f.内沙布尔Nishapur.SetParent(f)

	f.忽希思丹Qohistan = qohistan.CQohistan忽希思丹
	f.忽希思丹Qohistan.SetParent(f)

	f.图斯Tus = tus.CTus图斯
	f.图斯Tus.SetParent(f)

}
