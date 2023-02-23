package ostrobothnia

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/ostrobothnia/kajaani"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/ostrobothnia/osterbotten"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/ostrobothnia/oulu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OstrobothniaDuke interface {
	feud.Duke
	CKajaani凯努() kajaani.KajaaniCounty
	COsterbotten东博滕() osterbotten.OsterbottenCounty
	COulu奥卢() oulu.OuluCounty
}

type 东博滕OstrobothniaDuke struct {
	feud.BaseDuke
	凯努Kajaani      kajaani.KajaaniCounty
	东博滕Osterbotten osterbotten.OsterbottenCounty
	奥卢Oulu         oulu.OuluCounty
}

func (d *东博滕OstrobothniaDuke) CKajaani凯努() kajaani.KajaaniCounty {
	return d.凯努Kajaani
}

func (d *东博滕OstrobothniaDuke) COsterbotten东博滕() osterbotten.OsterbottenCounty {
	return d.东博滕Osterbotten
}

func (d *东博滕OstrobothniaDuke) COulu奥卢() oulu.OuluCounty {
	return d.奥卢Oulu
}

var DOstrobothnia东博滕 OstrobothniaDuke = &东博滕OstrobothniaDuke{}

func init() {
	f := DOstrobothnia东博滕.(*东博滕OstrobothniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ostrobothnia",
		TitleName: "东博滕",
		TitleCode: "d_ostrobothnia",
		Counties:  map[string]feud.County{},
	}

	f.凯努Kajaani = kajaani.CKajaani凯努
	f.凯努Kajaani.SetParent(f)

	f.东博滕Osterbotten = osterbotten.COsterbotten东博滕
	f.东博滕Osterbotten.SetParent(f)

	f.奥卢Oulu = oulu.COulu奥卢
	f.奥卢Oulu.SetParent(f)

}
