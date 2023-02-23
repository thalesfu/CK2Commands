package gotland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/gotland/gotland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GotlandDuke interface {
	feud.Duke
	CGotland哥特兰() gotland.GotlandCounty
}

type 哥特兰GotlandDuke struct {
	feud.BaseDuke
	哥特兰Gotland gotland.GotlandCounty
}

func (d *哥特兰GotlandDuke) CGotland哥特兰() gotland.GotlandCounty {
	return d.哥特兰Gotland
}

var DGotland哥特兰 GotlandDuke = &哥特兰GotlandDuke{}

func init() {
	f := DGotland哥特兰.(*哥特兰GotlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gotland",
		TitleName: "哥特兰",
		TitleCode: "d_gotland",
		Counties:  map[string]feud.County{},
	}

	f.哥特兰Gotland = gotland.CGotland哥特兰
	f.哥特兰Gotland.SetParent(f)

}
