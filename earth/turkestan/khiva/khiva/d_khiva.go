package khiva

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khiva/dashhowuz"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khiva/gurganj"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khiva/khiva"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khiva/urgench"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhivaDuke interface {
	feud.Duke
	CDashhowuz达什霍武兹() dashhowuz.DashhowuzCounty
	CGurganj玉龙杰赤() gurganj.GurganjCounty
	CKhiva希瓦() khiva.KhivaCounty
	CUrgench马德米尼亚() urgench.UrgenchCounty
}

type 希瓦KhivaDuke struct {
	feud.BaseDuke
	达什霍武兹Dashhowuz dashhowuz.DashhowuzCounty
	玉龙杰赤Gurganj    gurganj.GurganjCounty
	希瓦Khiva        khiva.KhivaCounty
	马德米尼亚Urgench   urgench.UrgenchCounty
}

func (d *希瓦KhivaDuke) CDashhowuz达什霍武兹() dashhowuz.DashhowuzCounty {
	return d.达什霍武兹Dashhowuz
}

func (d *希瓦KhivaDuke) CGurganj玉龙杰赤() gurganj.GurganjCounty {
	return d.玉龙杰赤Gurganj
}

func (d *希瓦KhivaDuke) CKhiva希瓦() khiva.KhivaCounty {
	return d.希瓦Khiva
}

func (d *希瓦KhivaDuke) CUrgench马德米尼亚() urgench.UrgenchCounty {
	return d.马德米尼亚Urgench
}

var DKhiva希瓦 KhivaDuke = &希瓦KhivaDuke{}

func init() {
	f := DKhiva希瓦.(*希瓦KhivaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khiva",
		TitleName: "希瓦",
		TitleCode: "d_khiva",
		Counties:  map[string]feud.County{},
	}

	f.达什霍武兹Dashhowuz = dashhowuz.CDashhowuz达什霍武兹
	f.达什霍武兹Dashhowuz.SetParent(f)

	f.玉龙杰赤Gurganj = gurganj.CGurganj玉龙杰赤
	f.玉龙杰赤Gurganj.SetParent(f)

	f.希瓦Khiva = khiva.CKhiva希瓦
	f.希瓦Khiva.SetParent(f)

	f.马德米尼亚Urgench = urgench.CUrgench马德米尼亚
	f.马德米尼亚Urgench.SetParent(f)

}
