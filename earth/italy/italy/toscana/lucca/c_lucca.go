package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LuccaCounty interface {
    feud.County
    BAltopascio阿尔托帕肖() 	feud.Barony
    BCalcinaia卡尔奇纳亚() 	feud.Barony
    BCapannori卡潘诺里() 	feud.Barony
    BCascina卡希纳() 	feud.Barony
    BLucca卢卡() 	feud.Barony
    BPistoia皮斯托亚() 	feud.Barony
    BSeravezza塞拉韦扎() 	feud.Barony
    BViareggio维亚雷焦() 	feud.Barony
}

type 卢卡LuccaCounty struct {
	feud.BaseCounty
	阿尔托帕肖Altopascio 	feud.Barony
	卡尔奇纳亚Calcinaia 	feud.Barony
	卡潘诺里Capannori 	feud.Barony
	卡希纳Cascina 	feud.Barony
	卢卡Lucca 	feud.Barony
	皮斯托亚Pistoia 	feud.Barony
	塞拉韦扎Seravezza 	feud.Barony
	维亚雷焦Viareggio 	feud.Barony
}

func (c *卢卡LuccaCounty) BAltopascio阿尔托帕肖() feud.Barony {
	return c.阿尔托帕肖Altopascio
}
    
func (c *卢卡LuccaCounty) BCalcinaia卡尔奇纳亚() feud.Barony {
	return c.卡尔奇纳亚Calcinaia
}
    
func (c *卢卡LuccaCounty) BCapannori卡潘诺里() feud.Barony {
	return c.卡潘诺里Capannori
}
    
func (c *卢卡LuccaCounty) BCascina卡希纳() feud.Barony {
	return c.卡希纳Cascina
}
    
func (c *卢卡LuccaCounty) BLucca卢卡() feud.Barony {
	return c.卢卡Lucca
}
    
func (c *卢卡LuccaCounty) BPistoia皮斯托亚() feud.Barony {
	return c.皮斯托亚Pistoia
}
    
func (c *卢卡LuccaCounty) BSeravezza塞拉韦扎() feud.Barony {
	return c.塞拉韦扎Seravezza
}
    
func (c *卢卡LuccaCounty) BViareggio维亚雷焦() feud.Barony {
	return c.维亚雷焦Viareggio
}
    
var CLucca卢卡 LuccaCounty = &卢卡LuccaCounty{}

func init() {
	f := CLucca卢卡.(*卢卡LuccaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "323",
		Title:     "lucca",
		TitleName: "卢卡",
		TitleCode: "c_lucca",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔托帕肖Altopascio = BAltopascio阿尔托帕肖
	f.阿尔托帕肖Altopascio.SetParent(f)
	
	f.卡尔奇纳亚Calcinaia = BCalcinaia卡尔奇纳亚
	f.卡尔奇纳亚Calcinaia.SetParent(f)
	
	f.卡潘诺里Capannori = BCapannori卡潘诺里
	f.卡潘诺里Capannori.SetParent(f)
	
	f.卡希纳Cascina = BCascina卡希纳
	f.卡希纳Cascina.SetParent(f)
	
	f.卢卡Lucca = BLucca卢卡
	f.卢卡Lucca.SetParent(f)
	
	f.皮斯托亚Pistoia = BPistoia皮斯托亚
	f.皮斯托亚Pistoia.SetParent(f)
	
	f.塞拉韦扎Seravezza = BSeravezza塞拉韦扎
	f.塞拉韦扎Seravezza.SetParent(f)
	
	f.维亚雷焦Viareggio = BViareggio维亚雷焦
	f.维亚雷焦Viareggio.SetParent(f)
	
}
