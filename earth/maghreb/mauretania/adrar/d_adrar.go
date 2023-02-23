package adrar

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/adrar/idjil"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/adrar/ouadane"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AdrarDuke interface {
	feud.Duke
	CIdjil比尔乌姆格兰() idjil.IdjilCounty
	COuadane瓦丹() ouadane.OuadaneCounty
}

type 阿德拉尔AdrarDuke struct {
	feud.BaseDuke
	比尔乌姆格兰Idjil idjil.IdjilCounty
	瓦丹Ouadane   ouadane.OuadaneCounty
}

func (d *阿德拉尔AdrarDuke) CIdjil比尔乌姆格兰() idjil.IdjilCounty {
	return d.比尔乌姆格兰Idjil
}

func (d *阿德拉尔AdrarDuke) COuadane瓦丹() ouadane.OuadaneCounty {
	return d.瓦丹Ouadane
}

var DAdrar阿德拉尔 AdrarDuke = &阿德拉尔AdrarDuke{}

func init() {
	f := DAdrar阿德拉尔.(*阿德拉尔AdrarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "adrar",
		TitleName: "阿德拉尔",
		TitleCode: "d_adrar",
		Counties:  map[string]feud.County{},
	}

	f.比尔乌姆格兰Idjil = idjil.CIdjil比尔乌姆格兰
	f.比尔乌姆格兰Idjil.SetParent(f)

	f.瓦丹Ouadane = ouadane.COuadane瓦丹
	f.瓦丹Ouadane.SetParent(f)

}
