package lothian

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/lothian/dunbar"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/lothian/lothian"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/lothian/teviotdale"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LothianDuke interface {
	feud.Duke
	CDunbar邓巴() dunbar.DunbarCounty
	CLothian洛锡安() lothian.LothianCounty
	CTeviotdale蒂维厄特代尔() teviotdale.TeviotdaleCounty
}

type 洛锡安LothianDuke struct {
	feud.BaseDuke
	邓巴Dunbar         dunbar.DunbarCounty
	洛锡安Lothian       lothian.LothianCounty
	蒂维厄特代尔Teviotdale teviotdale.TeviotdaleCounty
}

func (d *洛锡安LothianDuke) CDunbar邓巴() dunbar.DunbarCounty {
	return d.邓巴Dunbar
}

func (d *洛锡安LothianDuke) CLothian洛锡安() lothian.LothianCounty {
	return d.洛锡安Lothian
}

func (d *洛锡安LothianDuke) CTeviotdale蒂维厄特代尔() teviotdale.TeviotdaleCounty {
	return d.蒂维厄特代尔Teviotdale
}

var DLothian洛锡安 LothianDuke = &洛锡安LothianDuke{}

func init() {
	f := DLothian洛锡安.(*洛锡安LothianDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lothian",
		TitleName: "洛锡安",
		TitleCode: "d_lothian",
		Counties:  map[string]feud.County{},
	}

	f.邓巴Dunbar = dunbar.CDunbar邓巴
	f.邓巴Dunbar.SetParent(f)

	f.洛锡安Lothian = lothian.CLothian洛锡安
	f.洛锡安Lothian.SetParent(f)

	f.蒂维厄特代尔Teviotdale = teviotdale.CTeviotdale蒂维厄特代尔
	f.蒂维厄特代尔Teviotdale.SetParent(f)

}
