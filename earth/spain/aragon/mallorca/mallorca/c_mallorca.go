package mallorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MallorcaCounty interface {
    feud.County
    BAlcudia阿尔库迪亚() 	feud.Barony
    BAlgaida阿尔盖达() 	feud.Barony
    BEivissa伊维萨() 	feud.Barony
    BFelanitx费拉尼奇() 	feud.Barony
    BFormentera福门特拉() 	feud.Barony
    BManacor马纳科尔() 	feud.Barony
    BPalma帕尔马() 	feud.Barony
    BSantaponsa圣蓬萨() 	feud.Barony
}

type 马略卡MallorcaCounty struct {
	feud.BaseCounty
	阿尔库迪亚Alcudia 	feud.Barony
	阿尔盖达Algaida 	feud.Barony
	伊维萨Eivissa 	feud.Barony
	费拉尼奇Felanitx 	feud.Barony
	福门特拉Formentera 	feud.Barony
	马纳科尔Manacor 	feud.Barony
	帕尔马Palma 	feud.Barony
	圣蓬萨Santaponsa 	feud.Barony
}

func (c *马略卡MallorcaCounty) BAlcudia阿尔库迪亚() feud.Barony {
	return c.阿尔库迪亚Alcudia
}
    
func (c *马略卡MallorcaCounty) BAlgaida阿尔盖达() feud.Barony {
	return c.阿尔盖达Algaida
}
    
func (c *马略卡MallorcaCounty) BEivissa伊维萨() feud.Barony {
	return c.伊维萨Eivissa
}
    
func (c *马略卡MallorcaCounty) BFelanitx费拉尼奇() feud.Barony {
	return c.费拉尼奇Felanitx
}
    
func (c *马略卡MallorcaCounty) BFormentera福门特拉() feud.Barony {
	return c.福门特拉Formentera
}
    
func (c *马略卡MallorcaCounty) BManacor马纳科尔() feud.Barony {
	return c.马纳科尔Manacor
}
    
func (c *马略卡MallorcaCounty) BPalma帕尔马() feud.Barony {
	return c.帕尔马Palma
}
    
func (c *马略卡MallorcaCounty) BSantaponsa圣蓬萨() feud.Barony {
	return c.圣蓬萨Santaponsa
}
    
var CMallorca马略卡 MallorcaCounty = &马略卡MallorcaCounty{}

func init() {
	f := CMallorca马略卡.(*马略卡MallorcaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "827",
		Title:     "mallorca",
		TitleName: "马略卡",
		TitleCode: "c_mallorca",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔库迪亚Alcudia = BAlcudia阿尔库迪亚
	f.阿尔库迪亚Alcudia.SetParent(f)
	
	f.阿尔盖达Algaida = BAlgaida阿尔盖达
	f.阿尔盖达Algaida.SetParent(f)
	
	f.伊维萨Eivissa = BEivissa伊维萨
	f.伊维萨Eivissa.SetParent(f)
	
	f.费拉尼奇Felanitx = BFelanitx费拉尼奇
	f.费拉尼奇Felanitx.SetParent(f)
	
	f.福门特拉Formentera = BFormentera福门特拉
	f.福门特拉Formentera.SetParent(f)
	
	f.马纳科尔Manacor = BManacor马纳科尔
	f.马纳科尔Manacor.SetParent(f)
	
	f.帕尔马Palma = BPalma帕尔马
	f.帕尔马Palma.SetParent(f)
	
	f.圣蓬萨Santaponsa = BSantaponsa圣蓬萨
	f.圣蓬萨Santaponsa.SetParent(f)
	
}
