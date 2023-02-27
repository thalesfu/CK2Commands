package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SripuriCounty interface {
    feud.County
    BBagbahara巴格巴哈拉() 	feud.Barony
    BBasna巴斯纳() 	feud.Barony
    BNuapada努瓦帕达() 	feud.Barony
    BRajauli罗韶梨() 	feud.Barony
    BRajiva_lochana罗什伐卢遮那() 	feud.Barony
    BSanjeli桑祇梨() 	feud.Barony
    BSripuri室利补梨() 	feud.Barony
}

type 室利补梨SripuriCounty struct {
	feud.BaseCounty
	巴格巴哈拉Bagbahara 	feud.Barony
	巴斯纳Basna 	feud.Barony
	努瓦帕达Nuapada 	feud.Barony
	罗韶梨Rajauli 	feud.Barony
	罗什伐卢遮那Rajiva_lochana 	feud.Barony
	桑祇梨Sanjeli 	feud.Barony
	室利补梨Sripuri 	feud.Barony
}

func (c *室利补梨SripuriCounty) BBagbahara巴格巴哈拉() feud.Barony {
	return c.巴格巴哈拉Bagbahara
}
    
func (c *室利补梨SripuriCounty) BBasna巴斯纳() feud.Barony {
	return c.巴斯纳Basna
}
    
func (c *室利补梨SripuriCounty) BNuapada努瓦帕达() feud.Barony {
	return c.努瓦帕达Nuapada
}
    
func (c *室利补梨SripuriCounty) BRajauli罗韶梨() feud.Barony {
	return c.罗韶梨Rajauli
}
    
func (c *室利补梨SripuriCounty) BRajiva_lochana罗什伐卢遮那() feud.Barony {
	return c.罗什伐卢遮那Rajiva_lochana
}
    
func (c *室利补梨SripuriCounty) BSanjeli桑祇梨() feud.Barony {
	return c.桑祇梨Sanjeli
}
    
func (c *室利补梨SripuriCounty) BSripuri室利补梨() feud.Barony {
	return c.室利补梨Sripuri
}
    
var CSripuri室利补梨 SripuriCounty = &室利补梨SripuriCounty{}

func init() {
	f := CSripuri室利补梨.(*室利补梨SripuriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1155",
		Title:     "sripuri",
		TitleName: "室利补梨",
		TitleCode: "c_sripuri",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴格巴哈拉Bagbahara = BBagbahara巴格巴哈拉
	f.巴格巴哈拉Bagbahara.SetParent(f)
	
	f.巴斯纳Basna = BBasna巴斯纳
	f.巴斯纳Basna.SetParent(f)
	
	f.努瓦帕达Nuapada = BNuapada努瓦帕达
	f.努瓦帕达Nuapada.SetParent(f)
	
	f.罗韶梨Rajauli = BRajauli罗韶梨
	f.罗韶梨Rajauli.SetParent(f)
	
	f.罗什伐卢遮那Rajiva_lochana = BRajiva_lochana罗什伐卢遮那
	f.罗什伐卢遮那Rajiva_lochana.SetParent(f)
	
	f.桑祇梨Sanjeli = BSanjeli桑祇梨
	f.桑祇梨Sanjeli.SetParent(f)
	
	f.室利补梨Sripuri = BSripuri室利补梨
	f.室利补梨Sripuri.SetParent(f)
	
}
