package alger

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/alger/orania"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlgerDuke interface {
	feud.Duke
	COrania奥兰() orania.OraniaCounty
}

type 阿尔及尔AlgerDuke struct {
	feud.BaseDuke
	奥兰Orania orania.OraniaCounty
}

func (d *阿尔及尔AlgerDuke) COrania奥兰() orania.OraniaCounty {
	return d.奥兰Orania
}

var DAlger阿尔及尔 AlgerDuke = &阿尔及尔AlgerDuke{}

func init() {
	f := DAlger阿尔及尔.(*阿尔及尔AlgerDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "alger",
		TitleName: "阿尔及尔",
		TitleCode: "d_alger",
		Counties:  map[string]feud.County{},
	}

	f.奥兰Orania = orania.COrania奥兰
	f.奥兰Orania.SetParent(f)

}
