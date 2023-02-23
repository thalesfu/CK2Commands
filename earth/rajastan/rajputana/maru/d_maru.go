package maru

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/maru/godwad"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/maru/mandavyapura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/maru/medantaka"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaruDuke interface {
	feud.Duke
	CGodwad乔荼婆荼() godwad.GodwadCounty
	CMandavyapura曼荼毗耶补罗() mandavyapura.MandavyapuraCounty
	CMedantaka迷檀多迦() medantaka.MedantakaCounty
}

type 摩楼MaruDuke struct {
	feud.BaseDuke
	乔荼婆荼Godwad         godwad.GodwadCounty
	曼荼毗耶补罗Mandavyapura mandavyapura.MandavyapuraCounty
	迷檀多迦Medantaka      medantaka.MedantakaCounty
}

func (d *摩楼MaruDuke) CGodwad乔荼婆荼() godwad.GodwadCounty {
	return d.乔荼婆荼Godwad
}

func (d *摩楼MaruDuke) CMandavyapura曼荼毗耶补罗() mandavyapura.MandavyapuraCounty {
	return d.曼荼毗耶补罗Mandavyapura
}

func (d *摩楼MaruDuke) CMedantaka迷檀多迦() medantaka.MedantakaCounty {
	return d.迷檀多迦Medantaka
}

var DMaru摩楼 MaruDuke = &摩楼MaruDuke{}

func init() {
	f := DMaru摩楼.(*摩楼MaruDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "maru",
		TitleName: "摩楼",
		TitleCode: "d_maru",
		Counties:  map[string]feud.County{},
	}

	f.乔荼婆荼Godwad = godwad.CGodwad乔荼婆荼
	f.乔荼婆荼Godwad.SetParent(f)

	f.曼荼毗耶补罗Mandavyapura = mandavyapura.CMandavyapura曼荼毗耶补罗
	f.曼荼毗耶补罗Mandavyapura.SetParent(f)

	f.迷檀多迦Medantaka = medantaka.CMedantaka迷檀多迦
	f.迷檀多迦Medantaka.SetParent(f)

}
