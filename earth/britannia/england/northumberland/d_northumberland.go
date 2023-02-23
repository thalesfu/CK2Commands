package northumberland

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/northumberland/durham"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/northumberland/lindisfarne"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/northumberland/northumberland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorthumberlandDuke interface {
	feud.Duke
	CDurham达勒姆() durham.DurhamCounty
	CLindisfarne林迪斯法恩() lindisfarne.LindisfarneCounty
	CNorthumberland诺森伯兰() northumberland.NorthumberlandCounty
}

type 诺森伯兰NorthumberlandDuke struct {
	feud.BaseDuke
	达勒姆Durham          durham.DurhamCounty
	林迪斯法恩Lindisfarne   lindisfarne.LindisfarneCounty
	诺森伯兰Northumberland northumberland.NorthumberlandCounty
}

func (d *诺森伯兰NorthumberlandDuke) CDurham达勒姆() durham.DurhamCounty {
	return d.达勒姆Durham
}

func (d *诺森伯兰NorthumberlandDuke) CLindisfarne林迪斯法恩() lindisfarne.LindisfarneCounty {
	return d.林迪斯法恩Lindisfarne
}

func (d *诺森伯兰NorthumberlandDuke) CNorthumberland诺森伯兰() northumberland.NorthumberlandCounty {
	return d.诺森伯兰Northumberland
}

var DNorthumberland诺森伯兰 NorthumberlandDuke = &诺森伯兰NorthumberlandDuke{}

func init() {
	f := DNorthumberland诺森伯兰.(*诺森伯兰NorthumberlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "northumberland",
		TitleName: "诺森伯兰",
		TitleCode: "d_northumberland",
		Counties:  map[string]feud.County{},
	}

	f.达勒姆Durham = durham.CDurham达勒姆
	f.达勒姆Durham.SetParent(f)

	f.林迪斯法恩Lindisfarne = lindisfarne.CLindisfarne林迪斯法恩
	f.林迪斯法恩Lindisfarne.SetParent(f)

	f.诺森伯兰Northumberland = northumberland.CNorthumberland诺森伯兰
	f.诺森伯兰Northumberland.SetParent(f)

}
