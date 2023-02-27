package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DorpatCounty interface {
    feud.County
    BFellin费林() 	feud.Barony
    BJarva耶尔瓦() 	feud.Barony
    BKirrumpa基伦佩() 	feud.Barony
    BLais莱斯() 	feud.Barony
    BOdenpah奥登佩赫() 	feud.Barony
    BTartu塔尔图() 	feud.Barony
    BVanakastre瓦纳_卡斯特雷() 	feud.Barony
    BVastseliina瓦斯采利纳() 	feud.Barony
}

type 塔尔图DorpatCounty struct {
	feud.BaseCounty
	费林Fellin 	feud.Barony
	耶尔瓦Jarva 	feud.Barony
	基伦佩Kirrumpa 	feud.Barony
	莱斯Lais 	feud.Barony
	奥登佩赫Odenpah 	feud.Barony
	塔尔图Tartu 	feud.Barony
	瓦纳_卡斯特雷Vanakastre 	feud.Barony
	瓦斯采利纳Vastseliina 	feud.Barony
}

func (c *塔尔图DorpatCounty) BFellin费林() feud.Barony {
	return c.费林Fellin
}
    
func (c *塔尔图DorpatCounty) BJarva耶尔瓦() feud.Barony {
	return c.耶尔瓦Jarva
}
    
func (c *塔尔图DorpatCounty) BKirrumpa基伦佩() feud.Barony {
	return c.基伦佩Kirrumpa
}
    
func (c *塔尔图DorpatCounty) BLais莱斯() feud.Barony {
	return c.莱斯Lais
}
    
func (c *塔尔图DorpatCounty) BOdenpah奥登佩赫() feud.Barony {
	return c.奥登佩赫Odenpah
}
    
func (c *塔尔图DorpatCounty) BTartu塔尔图() feud.Barony {
	return c.塔尔图Tartu
}
    
func (c *塔尔图DorpatCounty) BVanakastre瓦纳_卡斯特雷() feud.Barony {
	return c.瓦纳_卡斯特雷Vanakastre
}
    
func (c *塔尔图DorpatCounty) BVastseliina瓦斯采利纳() feud.Barony {
	return c.瓦斯采利纳Vastseliina
}
    
var CDorpat塔尔图 DorpatCounty = &塔尔图DorpatCounty{}

func init() {
	f := CDorpat塔尔图.(*塔尔图DorpatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "379",
		Title:     "dorpat",
		TitleName: "塔尔图",
		TitleCode: "c_dorpat",
		Baronies:  map[string]feud.Barony{},
	}

	f.费林Fellin = BFellin费林
	f.费林Fellin.SetParent(f)
	
	f.耶尔瓦Jarva = BJarva耶尔瓦
	f.耶尔瓦Jarva.SetParent(f)
	
	f.基伦佩Kirrumpa = BKirrumpa基伦佩
	f.基伦佩Kirrumpa.SetParent(f)
	
	f.莱斯Lais = BLais莱斯
	f.莱斯Lais.SetParent(f)
	
	f.奥登佩赫Odenpah = BOdenpah奥登佩赫
	f.奥登佩赫Odenpah.SetParent(f)
	
	f.塔尔图Tartu = BTartu塔尔图
	f.塔尔图Tartu.SetParent(f)
	
	f.瓦纳_卡斯特雷Vanakastre = BVanakastre瓦纳_卡斯特雷
	f.瓦纳_卡斯特雷Vanakastre.SetParent(f)
	
	f.瓦斯采利纳Vastseliina = BVastseliina瓦斯采利纳
	f.瓦斯采利纳Vastseliina.SetParent(f)
	
}
