package germany

import (
	"github.com/thalesfu/CK2Commands/earth/germany/austria"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia"
	"github.com/thalesfu/CK2Commands/earth/germany/eastern_marches"
	"github.com/thalesfu/CK2Commands/earth/germany/franconia_otto"
	"github.com/thalesfu/CK2Commands/earth/germany/frisia"
	"github.com/thalesfu/CK2Commands/earth/germany/germany"
	"github.com/thalesfu/CK2Commands/earth/germany/hansa"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony"
	"github.com/thalesfu/CK2Commands/earth/germany/swabia_otto"
	"github.com/thalesfu/CK2Commands/earth/germany/thuringia_otto"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GermanyEmpire interface {
    feud.Empire
    KAustria奥地利() 	austria.AustriaKingdom
    KBavaria巴伐利亚() 	bavaria.BavariaKingdom
    KCarinthia卡林西亚() 	carinthia.CarinthiaKingdom
    KEastern_marches奥地利() 	eastern_marches.Eastern_marchesKingdom
    KFranconia_otto法兰克尼亚() 	franconia_otto.Franconia_ottoKingdom
    KFrisia弗里西亚() 	frisia.FrisiaKingdom
    KGermany德意志() 	germany.GermanyKingdom
    KHansa汉萨同盟() 	hansa.HansaKingdom
    KLotharingia中法兰克() 	lotharingia.LotharingiaKingdom
    KSaxony萨克森() 	saxony.SaxonyKingdom
    KSwabia_otto施瓦本() 	swabia_otto.Swabia_ottoKingdom
    KThuringia_otto图林根() 	thuringia_otto.Thuringia_ottoKingdom
}

type 日耳曼尼亚GermanyEmpire struct {
	feud.BaseEmpire
	奥地利Austria 	austria.AustriaKingdom
	巴伐利亚Bavaria 	bavaria.BavariaKingdom
	卡林西亚Carinthia 	carinthia.CarinthiaKingdom
	奥地利Eastern_marches 	eastern_marches.Eastern_marchesKingdom
	法兰克尼亚Franconia_otto 	franconia_otto.Franconia_ottoKingdom
	弗里西亚Frisia 	frisia.FrisiaKingdom
	德意志Germany 	germany.GermanyKingdom
	汉萨同盟Hansa 	hansa.HansaKingdom
	中法兰克Lotharingia 	lotharingia.LotharingiaKingdom
	萨克森Saxony 	saxony.SaxonyKingdom
	施瓦本Swabia_otto 	swabia_otto.Swabia_ottoKingdom
	图林根Thuringia_otto 	thuringia_otto.Thuringia_ottoKingdom
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
    
func (e *日耳曼尼亚GermanyEmpire) KEastern_marches奥地利() eastern_marches.Eastern_marchesKingdom {
	return e.奥地利Eastern_marches
}
    
func (e *日耳曼尼亚GermanyEmpire) KFranconia_otto法兰克尼亚() franconia_otto.Franconia_ottoKingdom {
	return e.法兰克尼亚Franconia_otto
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
    
func (e *日耳曼尼亚GermanyEmpire) KSwabia_otto施瓦本() swabia_otto.Swabia_ottoKingdom {
	return e.施瓦本Swabia_otto
}
    
func (e *日耳曼尼亚GermanyEmpire) KThuringia_otto图林根() thuringia_otto.Thuringia_ottoKingdom {
	return e.图林根Thuringia_otto
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
	f.奥地利Eastern_marches = eastern_marches.KEastern_marches奥地利
	f.奥地利Eastern_marches.SetParent(f)
	f.法兰克尼亚Franconia_otto = franconia_otto.KFranconia_otto法兰克尼亚
	f.法兰克尼亚Franconia_otto.SetParent(f)
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
	f.施瓦本Swabia_otto = swabia_otto.KSwabia_otto施瓦本
	f.施瓦本Swabia_otto.SetParent(f)
	f.图林根Thuringia_otto = thuringia_otto.KThuringia_otto图林根
	f.图林根Thuringia_otto.SetParent(f)
}
