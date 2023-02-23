package uppland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/uppland/aland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/uppland/gastrikland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/uppland/sodermanland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/uppland/uppland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UpplandDuke interface {
	feud.Duke
	CAland奥兰() aland.AlandCounty
	CGastrikland耶斯特里克兰() gastrikland.GastriklandCounty
	CSodermanland南曼兰() sodermanland.SodermanlandCounty
	CUppland乌普兰() uppland.UpplandCounty
}

type 乌普兰UpplandDuke struct {
	feud.BaseDuke
	奥兰Aland           aland.AlandCounty
	耶斯特里克兰Gastrikland gastrikland.GastriklandCounty
	南曼兰Sodermanland   sodermanland.SodermanlandCounty
	乌普兰Uppland        uppland.UpplandCounty
}

func (d *乌普兰UpplandDuke) CAland奥兰() aland.AlandCounty {
	return d.奥兰Aland
}

func (d *乌普兰UpplandDuke) CGastrikland耶斯特里克兰() gastrikland.GastriklandCounty {
	return d.耶斯特里克兰Gastrikland
}

func (d *乌普兰UpplandDuke) CSodermanland南曼兰() sodermanland.SodermanlandCounty {
	return d.南曼兰Sodermanland
}

func (d *乌普兰UpplandDuke) CUppland乌普兰() uppland.UpplandCounty {
	return d.乌普兰Uppland
}

var DUppland乌普兰 UpplandDuke = &乌普兰UpplandDuke{}

func init() {
	f := DUppland乌普兰.(*乌普兰UpplandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "uppland",
		TitleName: "乌普兰",
		TitleCode: "d_uppland",
		Counties:  map[string]feud.County{},
	}

	f.奥兰Aland = aland.CAland奥兰
	f.奥兰Aland.SetParent(f)

	f.耶斯特里克兰Gastrikland = gastrikland.CGastrikland耶斯特里克兰
	f.耶斯特里克兰Gastrikland.SetParent(f)

	f.南曼兰Sodermanland = sodermanland.CSodermanland南曼兰
	f.南曼兰Sodermanland.SetParent(f)

	f.乌普兰Uppland = uppland.CUppland乌普兰
	f.乌普兰Uppland.SetParent(f)

}
