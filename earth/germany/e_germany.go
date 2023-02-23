package germany

import (
	"github.com/thalesfu/CK2Commands/earth/germany/austria"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia"
	"github.com/thalesfu/CK2Commands/earth/germany/germany"
	"github.com/thalesfu/CK2Commands/earth/germany/hansa"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GermanyEmpire interface {
	feud.Empire
	KAustria奥地利() austria.AustriaKingdom
	KBavaria巴伐利亚() bavaria.BavariaKingdom
	KCarinthia卡林西亚() carinthia.CarinthiaKingdom
	KFrisia弗里西亚() frisia.FrisiaKingdom
	KGermany德意志() germany.GermanyKingdom
	KHansa汉萨同盟() hansa.HansaKingdom
	KLotharingia中法兰克() lotharingia.LotharingiaKingdom
	KSaxony萨克森() saxony.SaxonyKingdom
}

type 日耳曼尼亚GermanyEmpire struct {
	feud.BaseEmpire
	奥地利Austria      austria.AustriaKingdom
	巴伐利亚Bavaria     bavaria.BavariaKingdom
	卡林西亚Carinthia   carinthia.CarinthiaKingdom
	弗里西亚Frisia      frisia.FrisiaKingdom
	德意志Germany      germany.GermanyKingdom
	汉萨同盟Hansa       hansa.HansaKingdom
	中法兰克Lotharingia lotharingia.LotharingiaKingdom
	萨克森Saxony       saxony.SaxonyKingdom
}

func (e *日耳曼尼亚GermanyEmpire) KAustria奥地利() austria.AustriaKingdom {
	return e.奥地利Austria
}

func (e *日耳曼尼亚GermanyEmpire) KBavaria巴伐利亚() bavaria.BavariaKingdom {
	return e.巴伐利亚Bavaria
}

func (e *日耳曼尼亚GermanyEmpire) KCarinthia卡林西亚() carinthia.CarinthiaKingdom {
	return e.卡林西亚Carinthia
}

func (e *日耳曼尼亚GermanyEmpire) KFrisia弗里西亚() frisia.FrisiaKingdom {
	return e.弗里西亚Frisia
}

func (e *日耳曼尼亚GermanyEmpire) KGermany德意志() germany.GermanyKingdom {
	return e.德意志Germany
}

func (e *日耳曼尼亚GermanyEmpire) KHansa汉萨同盟() hansa.HansaKingdom {
	return e.汉萨同盟Hansa
}

func (e *日耳曼尼亚GermanyEmpire) KLotharingia中法兰克() lotharingia.LotharingiaKingdom {
	return e.中法兰克Lotharingia
}

func (e *日耳曼尼亚GermanyEmpire) KSaxony萨克森() saxony.SaxonyKingdom {
	return e.萨克森Saxony
}

var EGermany日耳曼尼亚 GermanyEmpire = &日耳曼尼亚GermanyEmpire{}

func init() {
	f := EGermany日耳曼尼亚.(*日耳曼尼亚GermanyEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "germany",
		TitleName: "日耳曼尼亚",
		TitleCode: "e_germany",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.奥地利Austria = austria.KAustria奥地利
	f.奥地利Austria.SetParent(f)
	f.巴伐利亚Bavaria = bavaria.KBavaria巴伐利亚
	f.巴伐利亚Bavaria.SetParent(f)
	f.卡林西亚Carinthia = carinthia.KCarinthia卡林西亚
	f.卡林西亚Carinthia.SetParent(f)
	f.弗里西亚Frisia = frisia.KFrisia弗里西亚
	f.弗里西亚Frisia.SetParent(f)
	f.德意志Germany = germany.KGermany德意志
	f.德意志Germany.SetParent(f)
	f.汉萨同盟Hansa = hansa.KHansa汉萨同盟
	f.汉萨同盟Hansa.SetParent(f)
	f.中法兰克Lotharingia = lotharingia.KLotharingia中法兰克
	f.中法兰克Lotharingia.SetParent(f)
	f.萨克森Saxony = saxony.KSaxony萨克森
	f.萨克森Saxony.SetParent(f)
}
