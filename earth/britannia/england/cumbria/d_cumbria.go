package cumbria

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/cumbria/cumberland"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/cumbria/westmorland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CumbriaDuke interface {
	feud.Duke
	CCumberland坎伯兰() cumberland.CumberlandCounty
	CWestmorland威斯特摩兰() westmorland.WestmorlandCounty
}

type 坎布里亚CumbriaDuke struct {
	feud.BaseDuke
	坎伯兰Cumberland    cumberland.CumberlandCounty
	威斯特摩兰Westmorland westmorland.WestmorlandCounty
}

func (d *坎布里亚CumbriaDuke) CCumberland坎伯兰() cumberland.CumberlandCounty {
	return d.坎伯兰Cumberland
}

func (d *坎布里亚CumbriaDuke) CWestmorland威斯特摩兰() westmorland.WestmorlandCounty {
	return d.威斯特摩兰Westmorland
}

var DCumbria坎布里亚 CumbriaDuke = &坎布里亚CumbriaDuke{}

func init() {
	f := DCumbria坎布里亚.(*坎布里亚CumbriaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cumbria",
		TitleName: "坎布里亚",
		TitleCode: "d_cumbria",
		Counties:  map[string]feud.County{},
	}

	f.坎伯兰Cumberland = cumberland.CCumberland坎伯兰
	f.坎伯兰Cumberland.SetParent(f)

	f.威斯特摩兰Westmorland = westmorland.CWestmorland威斯特摩兰
	f.威斯特摩兰Westmorland.SetParent(f)

}
