package alsace

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/alsace/nordgau"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/alsace/sundgau"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlsaceDuke interface {
	feud.Duke
	CNordgau诺德高() nordgau.NordgauCounty
	CSundgau松德高() sundgau.SundgauCounty
}

type 阿尔萨斯AlsaceDuke struct {
	feud.BaseDuke
	诺德高Nordgau nordgau.NordgauCounty
	松德高Sundgau sundgau.SundgauCounty
}

func (d *阿尔萨斯AlsaceDuke) CNordgau诺德高() nordgau.NordgauCounty {
	return d.诺德高Nordgau
}

func (d *阿尔萨斯AlsaceDuke) CSundgau松德高() sundgau.SundgauCounty {
	return d.松德高Sundgau
}

var DAlsace阿尔萨斯 AlsaceDuke = &阿尔萨斯AlsaceDuke{}

func init() {
	f := DAlsace阿尔萨斯.(*阿尔萨斯AlsaceDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "alsace",
		TitleName: "阿尔萨斯",
		TitleCode: "d_alsace",
		Counties:  map[string]feud.County{},
	}

	f.诺德高Nordgau = nordgau.CNordgau诺德高
	f.诺德高Nordgau.SetParent(f)

	f.松德高Sundgau = sundgau.CSundgau松德高
	f.松德高Sundgau.SetParent(f)

}
