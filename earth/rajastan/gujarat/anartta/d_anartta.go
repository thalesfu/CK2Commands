package anartta

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/anartta/dhamalpur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/anartta/dvaraka"
	"github.com/thalesfu/CK2Commands/earth/rajastan/gujarat/anartta/kutch"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnarttaDuke interface {
	feud.Duke
	CDhamalpur陀摩罗补罗() dhamalpur.DhamalpurCounty
	CDvaraka堕罗迦() dvaraka.DvarakaCounty
	CKutch契吒() kutch.KutchCounty
}

type 阿捺多AnarttaDuke struct {
	feud.BaseDuke
	陀摩罗补罗Dhamalpur dhamalpur.DhamalpurCounty
	堕罗迦Dvaraka     dvaraka.DvarakaCounty
	契吒Kutch        kutch.KutchCounty
}

func (d *阿捺多AnarttaDuke) CDhamalpur陀摩罗补罗() dhamalpur.DhamalpurCounty {
	return d.陀摩罗补罗Dhamalpur
}

func (d *阿捺多AnarttaDuke) CDvaraka堕罗迦() dvaraka.DvarakaCounty {
	return d.堕罗迦Dvaraka
}

func (d *阿捺多AnarttaDuke) CKutch契吒() kutch.KutchCounty {
	return d.契吒Kutch
}

var DAnartta阿捺多 AnarttaDuke = &阿捺多AnarttaDuke{}

func init() {
	f := DAnartta阿捺多.(*阿捺多AnarttaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "anartta",
		TitleName: "阿捺多",
		TitleCode: "d_anartta",
		Counties:  map[string]feud.County{},
	}

	f.陀摩罗补罗Dhamalpur = dhamalpur.CDhamalpur陀摩罗补罗
	f.陀摩罗补罗Dhamalpur.SetParent(f)

	f.堕罗迦Dvaraka = dvaraka.CDvaraka堕罗迦
	f.堕罗迦Dvaraka.SetParent(f)

	f.契吒Kutch = kutch.CKutch契吒
	f.契吒Kutch.SetParent(f)

}
