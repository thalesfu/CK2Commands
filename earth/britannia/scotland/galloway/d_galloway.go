package galloway

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/galloway/carrick"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/galloway/clydesdale"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/galloway/galloway"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GallowayDuke interface {
	feud.Duke
	CCarrick卡里克() carrick.CarrickCounty
	CClydesdale克莱兹代尔() clydesdale.ClydesdaleCounty
	CGalloway加洛韦() galloway.GallowayCounty
}

type 加洛韦GallowayDuke struct {
	feud.BaseDuke
	卡里克Carrick      carrick.CarrickCounty
	克莱兹代尔Clydesdale clydesdale.ClydesdaleCounty
	加洛韦Galloway     galloway.GallowayCounty
}

func (d *加洛韦GallowayDuke) CCarrick卡里克() carrick.CarrickCounty {
	return d.卡里克Carrick
}

func (d *加洛韦GallowayDuke) CClydesdale克莱兹代尔() clydesdale.ClydesdaleCounty {
	return d.克莱兹代尔Clydesdale
}

func (d *加洛韦GallowayDuke) CGalloway加洛韦() galloway.GallowayCounty {
	return d.加洛韦Galloway
}

var DGalloway加洛韦 GallowayDuke = &加洛韦GallowayDuke{}

func init() {
	f := DGalloway加洛韦.(*加洛韦GallowayDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "galloway",
		TitleName: "加洛韦",
		TitleCode: "d_galloway",
		Counties:  map[string]feud.County{},
	}

	f.卡里克Carrick = carrick.CCarrick卡里克
	f.卡里克Carrick.SetParent(f)

	f.克莱兹代尔Clydesdale = clydesdale.CClydesdale克莱兹代尔
	f.克莱兹代尔Clydesdale.SetParent(f)

	f.加洛韦Galloway = galloway.CGalloway加洛韦
	f.加洛韦Galloway.SetParent(f)

}
