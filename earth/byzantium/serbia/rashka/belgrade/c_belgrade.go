package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BelgradeCounty interface {
    feud.County
    BBelgrade贝尔格莱德() 	feud.Barony
    BBranicevo布拉尼切沃() 	feud.Barony
    BKragujevac克拉古耶瓦茨() 	feud.Barony
    BLipovic利波维茨() 	feud.Barony
    BPozarevac波扎雷瓦茨() 	feud.Barony
    BRudnik鲁德尼克() 	feud.Barony
    BSmederevo斯梅代雷沃() 	feud.Barony
    BZemun泽姆林() 	feud.Barony
}

type 贝尔格莱德BelgradeCounty struct {
	feud.BaseCounty
	贝尔格莱德Belgrade 	feud.Barony
	布拉尼切沃Branicevo 	feud.Barony
	克拉古耶瓦茨Kragujevac 	feud.Barony
	利波维茨Lipovic 	feud.Barony
	波扎雷瓦茨Pozarevac 	feud.Barony
	鲁德尼克Rudnik 	feud.Barony
	斯梅代雷沃Smederevo 	feud.Barony
	泽姆林Zemun 	feud.Barony
}

func (c *贝尔格莱德BelgradeCounty) BBelgrade贝尔格莱德() feud.Barony {
	return c.贝尔格莱德Belgrade
}
    
func (c *贝尔格莱德BelgradeCounty) BBranicevo布拉尼切沃() feud.Barony {
	return c.布拉尼切沃Branicevo
}
    
func (c *贝尔格莱德BelgradeCounty) BKragujevac克拉古耶瓦茨() feud.Barony {
	return c.克拉古耶瓦茨Kragujevac
}
    
func (c *贝尔格莱德BelgradeCounty) BLipovic利波维茨() feud.Barony {
	return c.利波维茨Lipovic
}
    
func (c *贝尔格莱德BelgradeCounty) BPozarevac波扎雷瓦茨() feud.Barony {
	return c.波扎雷瓦茨Pozarevac
}
    
func (c *贝尔格莱德BelgradeCounty) BRudnik鲁德尼克() feud.Barony {
	return c.鲁德尼克Rudnik
}
    
func (c *贝尔格莱德BelgradeCounty) BSmederevo斯梅代雷沃() feud.Barony {
	return c.斯梅代雷沃Smederevo
}
    
func (c *贝尔格莱德BelgradeCounty) BZemun泽姆林() feud.Barony {
	return c.泽姆林Zemun
}
    
var CBelgrade贝尔格莱德 BelgradeCounty = &贝尔格莱德BelgradeCounty{}

func init() {
	f := CBelgrade贝尔格莱德.(*贝尔格莱德BelgradeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "505",
		Title:     "belgrade",
		TitleName: "贝尔格莱德",
		TitleCode: "c_belgrade",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔格莱德Belgrade = BBelgrade贝尔格莱德
	f.贝尔格莱德Belgrade.SetParent(f)
	
	f.布拉尼切沃Branicevo = BBranicevo布拉尼切沃
	f.布拉尼切沃Branicevo.SetParent(f)
	
	f.克拉古耶瓦茨Kragujevac = BKragujevac克拉古耶瓦茨
	f.克拉古耶瓦茨Kragujevac.SetParent(f)
	
	f.利波维茨Lipovic = BLipovic利波维茨
	f.利波维茨Lipovic.SetParent(f)
	
	f.波扎雷瓦茨Pozarevac = BPozarevac波扎雷瓦茨
	f.波扎雷瓦茨Pozarevac.SetParent(f)
	
	f.鲁德尼克Rudnik = BRudnik鲁德尼克
	f.鲁德尼克Rudnik.SetParent(f)
	
	f.斯梅代雷沃Smederevo = BSmederevo斯梅代雷沃
	f.斯梅代雷沃Smederevo.SetParent(f)
	
	f.泽姆林Zemun = BZemun泽姆林
	f.泽姆林Zemun.SetParent(f)
	
}
