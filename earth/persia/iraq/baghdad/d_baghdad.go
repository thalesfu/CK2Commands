package baghdad

import (
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/baghdad/baghdad"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/baghdad/ilam"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/baghdad/karbala"
	"github.com/thalesfu/CK2Commands/earth/persia/iraq/baghdad/kermanshah"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaghdadDuke interface {
	feud.Duke
	CBaghdad巴格达() baghdad.BaghdadCounty
	CIlam伊拉姆() ilam.IlamCounty
	CKarbala卡尔巴拉() karbala.KarbalaCounty
	CKermanshah克尔曼沙赫() kermanshah.KermanshahCounty
}

type 巴格达BaghdadDuke struct {
	feud.BaseDuke
	巴格达Baghdad      baghdad.BaghdadCounty
	伊拉姆Ilam         ilam.IlamCounty
	卡尔巴拉Karbala     karbala.KarbalaCounty
	克尔曼沙赫Kermanshah kermanshah.KermanshahCounty
}

func (d *巴格达BaghdadDuke) CBaghdad巴格达() baghdad.BaghdadCounty {
	return d.巴格达Baghdad
}

func (d *巴格达BaghdadDuke) CIlam伊拉姆() ilam.IlamCounty {
	return d.伊拉姆Ilam
}

func (d *巴格达BaghdadDuke) CKarbala卡尔巴拉() karbala.KarbalaCounty {
	return d.卡尔巴拉Karbala
}

func (d *巴格达BaghdadDuke) CKermanshah克尔曼沙赫() kermanshah.KermanshahCounty {
	return d.克尔曼沙赫Kermanshah
}

var DBaghdad巴格达 BaghdadDuke = &巴格达BaghdadDuke{}

func init() {
	f := DBaghdad巴格达.(*巴格达BaghdadDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "baghdad",
		TitleName: "巴格达",
		TitleCode: "d_baghdad",
		Counties:  map[string]feud.County{},
	}

	f.巴格达Baghdad = baghdad.CBaghdad巴格达
	f.巴格达Baghdad.SetParent(f)

	f.伊拉姆Ilam = ilam.CIlam伊拉姆
	f.伊拉姆Ilam.SetParent(f)

	f.卡尔巴拉Karbala = karbala.CKarbala卡尔巴拉
	f.卡尔巴拉Karbala.SetParent(f)

	f.克尔曼沙赫Kermanshah = kermanshah.CKermanshah克尔曼沙赫
	f.克尔曼沙赫Kermanshah.SetParent(f)

}
