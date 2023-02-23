package poitou

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/poitou/lusignan"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/poitou/poitiers"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/poitou/saintonge"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/poitou/thouars"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PoitouDuke interface {
	feud.Duke
	CLusignan吕西尼昂() lusignan.LusignanCounty
	CPoitiers普瓦捷() poitiers.PoitiersCounty
	CSaintonge圣通日() saintonge.SaintongeCounty
	CThouars图阿尔() thouars.ThouarsCounty
}

type 普瓦图PoitouDuke struct {
	feud.BaseDuke
	吕西尼昂Lusignan lusignan.LusignanCounty
	普瓦捷Poitiers  poitiers.PoitiersCounty
	圣通日Saintonge saintonge.SaintongeCounty
	图阿尔Thouars   thouars.ThouarsCounty
}

func (d *普瓦图PoitouDuke) CLusignan吕西尼昂() lusignan.LusignanCounty {
	return d.吕西尼昂Lusignan
}

func (d *普瓦图PoitouDuke) CPoitiers普瓦捷() poitiers.PoitiersCounty {
	return d.普瓦捷Poitiers
}

func (d *普瓦图PoitouDuke) CSaintonge圣通日() saintonge.SaintongeCounty {
	return d.圣通日Saintonge
}

func (d *普瓦图PoitouDuke) CThouars图阿尔() thouars.ThouarsCounty {
	return d.图阿尔Thouars
}

var DPoitou普瓦图 PoitouDuke = &普瓦图PoitouDuke{}

func init() {
	f := DPoitou普瓦图.(*普瓦图PoitouDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "poitou",
		TitleName: "普瓦图",
		TitleCode: "d_poitou",
		Counties:  map[string]feud.County{},
	}

	f.吕西尼昂Lusignan = lusignan.CLusignan吕西尼昂
	f.吕西尼昂Lusignan.SetParent(f)

	f.普瓦捷Poitiers = poitiers.CPoitiers普瓦捷
	f.普瓦捷Poitiers.SetParent(f)

	f.圣通日Saintonge = saintonge.CSaintonge圣通日
	f.圣通日Saintonge.SetParent(f)

	f.图阿尔Thouars = thouars.CThouars图阿尔
	f.图阿尔Thouars.SetParent(f)

}
