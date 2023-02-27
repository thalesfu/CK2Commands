package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarpaCounty interface {
    feud.County
    BBachanta拜申泰() 	feud.Barony
    BElst叶尔斯特() 	feud.Barony
    BIkburul伊基布鲁尔() 	feud.Barony
    BKaratchaplak卡拉查普拉克() 	feud.Barony
    BKetchenery克切涅雷() 	feud.Barony
    BYashalta亚沙尔塔() 	feud.Barony
    BYashkul亚什库利() 	feud.Barony
    BYsaganaman察甘阿曼() 	feud.Barony
}

type 萨尔帕SarpaCounty struct {
	feud.BaseCounty
	拜申泰Bachanta 	feud.Barony
	叶尔斯特Elst 	feud.Barony
	伊基布鲁尔Ikburul 	feud.Barony
	卡拉查普拉克Karatchaplak 	feud.Barony
	克切涅雷Ketchenery 	feud.Barony
	亚沙尔塔Yashalta 	feud.Barony
	亚什库利Yashkul 	feud.Barony
	察甘阿曼Ysaganaman 	feud.Barony
}

func (c *萨尔帕SarpaCounty) BBachanta拜申泰() feud.Barony {
	return c.拜申泰Bachanta
}
    
func (c *萨尔帕SarpaCounty) BElst叶尔斯特() feud.Barony {
	return c.叶尔斯特Elst
}
    
func (c *萨尔帕SarpaCounty) BIkburul伊基布鲁尔() feud.Barony {
	return c.伊基布鲁尔Ikburul
}
    
func (c *萨尔帕SarpaCounty) BKaratchaplak卡拉查普拉克() feud.Barony {
	return c.卡拉查普拉克Karatchaplak
}
    
func (c *萨尔帕SarpaCounty) BKetchenery克切涅雷() feud.Barony {
	return c.克切涅雷Ketchenery
}
    
func (c *萨尔帕SarpaCounty) BYashalta亚沙尔塔() feud.Barony {
	return c.亚沙尔塔Yashalta
}
    
func (c *萨尔帕SarpaCounty) BYashkul亚什库利() feud.Barony {
	return c.亚什库利Yashkul
}
    
func (c *萨尔帕SarpaCounty) BYsaganaman察甘阿曼() feud.Barony {
	return c.察甘阿曼Ysaganaman
}
    
var CSarpa萨尔帕 SarpaCounty = &萨尔帕SarpaCounty{}

func init() {
	f := CSarpa萨尔帕.(*萨尔帕SarpaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "607",
		Title:     "sarpa",
		TitleName: "萨尔帕",
		TitleCode: "c_sarpa",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜申泰Bachanta = BBachanta拜申泰
	f.拜申泰Bachanta.SetParent(f)
	
	f.叶尔斯特Elst = BElst叶尔斯特
	f.叶尔斯特Elst.SetParent(f)
	
	f.伊基布鲁尔Ikburul = BIkburul伊基布鲁尔
	f.伊基布鲁尔Ikburul.SetParent(f)
	
	f.卡拉查普拉克Karatchaplak = BKaratchaplak卡拉查普拉克
	f.卡拉查普拉克Karatchaplak.SetParent(f)
	
	f.克切涅雷Ketchenery = BKetchenery克切涅雷
	f.克切涅雷Ketchenery.SetParent(f)
	
	f.亚沙尔塔Yashalta = BYashalta亚沙尔塔
	f.亚沙尔塔Yashalta.SetParent(f)
	
	f.亚什库利Yashkul = BYashkul亚什库利
	f.亚什库利Yashkul.SetParent(f)
	
	f.察甘阿曼Ysaganaman = BYsaganaman察甘阿曼
	f.察甘阿曼Ysaganaman.SetParent(f)
	
}
