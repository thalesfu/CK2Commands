package esztergom

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/esztergom/esztergom"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/esztergom/pressburg"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/esztergom/sopron"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EsztergomDuke interface {
	feud.Duke
	CEsztergom埃斯泰尔戈姆() esztergom.EsztergomCounty
	CPressburg普雷斯堡() pressburg.PressburgCounty
	CSopron肖普朗() sopron.SopronCounty
}

type 埃斯泰尔戈姆EsztergomDuke struct {
	feud.BaseDuke
	埃斯泰尔戈姆Esztergom esztergom.EsztergomCounty
	普雷斯堡Pressburg   pressburg.PressburgCounty
	肖普朗Sopron       sopron.SopronCounty
}

func (d *埃斯泰尔戈姆EsztergomDuke) CEsztergom埃斯泰尔戈姆() esztergom.EsztergomCounty {
	return d.埃斯泰尔戈姆Esztergom
}

func (d *埃斯泰尔戈姆EsztergomDuke) CPressburg普雷斯堡() pressburg.PressburgCounty {
	return d.普雷斯堡Pressburg
}

func (d *埃斯泰尔戈姆EsztergomDuke) CSopron肖普朗() sopron.SopronCounty {
	return d.肖普朗Sopron
}

var DEsztergom埃斯泰尔戈姆 EsztergomDuke = &埃斯泰尔戈姆EsztergomDuke{}

func init() {
	f := DEsztergom埃斯泰尔戈姆.(*埃斯泰尔戈姆EsztergomDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "esztergom",
		TitleName: "埃斯泰尔戈姆",
		TitleCode: "d_esztergom",
		Counties:  map[string]feud.County{},
	}

	f.埃斯泰尔戈姆Esztergom = esztergom.CEsztergom埃斯泰尔戈姆
	f.埃斯泰尔戈姆Esztergom.SetParent(f)

	f.普雷斯堡Pressburg = pressburg.CPressburg普雷斯堡
	f.普雷斯堡Pressburg.SetParent(f)

	f.肖普朗Sopron = sopron.CSopron肖普朗
	f.肖普朗Sopron.SetParent(f)

}
