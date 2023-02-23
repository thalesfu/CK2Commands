package ladakh

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ladakh/diskit"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ladakh/leh"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ladakh/pangong"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LadakhDuke interface {
	feud.Duke
	CDiskit德吉() diskit.DiskitCounty
	CLeh列城() leh.LehCounty
	CPangong班公() pangong.PangongCounty
}

type 拉达克LadakhDuke struct {
	feud.BaseDuke
	德吉Diskit  diskit.DiskitCounty
	列城Leh     leh.LehCounty
	班公Pangong pangong.PangongCounty
}

func (d *拉达克LadakhDuke) CDiskit德吉() diskit.DiskitCounty {
	return d.德吉Diskit
}

func (d *拉达克LadakhDuke) CLeh列城() leh.LehCounty {
	return d.列城Leh
}

func (d *拉达克LadakhDuke) CPangong班公() pangong.PangongCounty {
	return d.班公Pangong
}

var DLadakh拉达克 LadakhDuke = &拉达克LadakhDuke{}

func init() {
	f := DLadakh拉达克.(*拉达克LadakhDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ladakh",
		TitleName: "拉达克",
		TitleCode: "d_ladakh",
		Counties:  map[string]feud.County{},
	}

	f.德吉Diskit = diskit.CDiskit德吉
	f.德吉Diskit.SetParent(f)

	f.列城Leh = leh.CLeh列城
	f.列城Leh.SetParent(f)

	f.班公Pangong = pangong.CPangong班公
	f.班公Pangong.SetParent(f)

}
