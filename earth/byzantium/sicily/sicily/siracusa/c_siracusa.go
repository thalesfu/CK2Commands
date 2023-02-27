package siracusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SiracusaCounty interface {
    feud.County
    BAugusta奥古斯塔() 	feud.Barony
    BCaltagirone卡尔塔吉罗内() 	feud.Barony
    BCatania卡塔尼亚() 	feud.Barony
    BCenturipe琴图里佩() 	feud.Barony
    BLentini伦蒂尼() 	feud.Barony
    BNoto诺托() 	feud.Barony
    BPaterno帕泰尔诺() 	feud.Barony
    BSyracusa锡拉库萨() 	feud.Barony
}

type 锡拉库萨SiracusaCounty struct {
	feud.BaseCounty
	奥古斯塔Augusta 	feud.Barony
	卡尔塔吉罗内Caltagirone 	feud.Barony
	卡塔尼亚Catania 	feud.Barony
	琴图里佩Centuripe 	feud.Barony
	伦蒂尼Lentini 	feud.Barony
	诺托Noto 	feud.Barony
	帕泰尔诺Paterno 	feud.Barony
	锡拉库萨Syracusa 	feud.Barony
}

func (c *锡拉库萨SiracusaCounty) BAugusta奥古斯塔() feud.Barony {
	return c.奥古斯塔Augusta
}
    
func (c *锡拉库萨SiracusaCounty) BCaltagirone卡尔塔吉罗内() feud.Barony {
	return c.卡尔塔吉罗内Caltagirone
}
    
func (c *锡拉库萨SiracusaCounty) BCatania卡塔尼亚() feud.Barony {
	return c.卡塔尼亚Catania
}
    
func (c *锡拉库萨SiracusaCounty) BCenturipe琴图里佩() feud.Barony {
	return c.琴图里佩Centuripe
}
    
func (c *锡拉库萨SiracusaCounty) BLentini伦蒂尼() feud.Barony {
	return c.伦蒂尼Lentini
}
    
func (c *锡拉库萨SiracusaCounty) BNoto诺托() feud.Barony {
	return c.诺托Noto
}
    
func (c *锡拉库萨SiracusaCounty) BPaterno帕泰尔诺() feud.Barony {
	return c.帕泰尔诺Paterno
}
    
func (c *锡拉库萨SiracusaCounty) BSyracusa锡拉库萨() feud.Barony {
	return c.锡拉库萨Syracusa
}
    
var CSiracusa锡拉库萨 SiracusaCounty = &锡拉库萨SiracusaCounty{}

func init() {
	f := CSiracusa锡拉库萨.(*锡拉库萨SiracusaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "343",
		Title:     "siracusa",
		TitleName: "锡拉库萨",
		TitleCode: "c_siracusa",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥古斯塔Augusta = BAugusta奥古斯塔
	f.奥古斯塔Augusta.SetParent(f)
	
	f.卡尔塔吉罗内Caltagirone = BCaltagirone卡尔塔吉罗内
	f.卡尔塔吉罗内Caltagirone.SetParent(f)
	
	f.卡塔尼亚Catania = BCatania卡塔尼亚
	f.卡塔尼亚Catania.SetParent(f)
	
	f.琴图里佩Centuripe = BCenturipe琴图里佩
	f.琴图里佩Centuripe.SetParent(f)
	
	f.伦蒂尼Lentini = BLentini伦蒂尼
	f.伦蒂尼Lentini.SetParent(f)
	
	f.诺托Noto = BNoto诺托
	f.诺托Noto.SetParent(f)
	
	f.帕泰尔诺Paterno = BPaterno帕泰尔诺
	f.帕泰尔诺Paterno.SetParent(f)
	
	f.锡拉库萨Syracusa = BSyracusa锡拉库萨
	f.锡拉库萨Syracusa.SetParent(f)
	
}
