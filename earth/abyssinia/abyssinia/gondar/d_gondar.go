package gondar

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/gondar/begemder"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/gondar/gondar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GondarDuke interface {
	feud.Duke
	CBegemder贝格姆迪尔() begemder.BegemderCounty
	CGondar贡德尔() gondar.GondarCounty
}

type 贡德尔GondarDuke struct {
	feud.BaseDuke
	贝格姆迪尔Begemder begemder.BegemderCounty
	贡德尔Gondar     gondar.GondarCounty
}

func (d *贡德尔GondarDuke) CBegemder贝格姆迪尔() begemder.BegemderCounty {
	return d.贝格姆迪尔Begemder
}

func (d *贡德尔GondarDuke) CGondar贡德尔() gondar.GondarCounty {
	return d.贡德尔Gondar
}

var DGondar贡德尔 GondarDuke = &贡德尔GondarDuke{}

func init() {
	f := DGondar贡德尔.(*贡德尔GondarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gondar",
		TitleName: "贡德尔",
		TitleCode: "d_gondar",
		Counties:  map[string]feud.County{},
	}

	f.贝格姆迪尔Begemder = begemder.CBegemder贝格姆迪尔
	f.贝格姆迪尔Begemder.SetParent(f)

	f.贡德尔Gondar = gondar.CGondar贡德尔
	f.贡德尔Gondar.SetParent(f)

}
