package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EgitaneaCounty interface {
    feud.County
    BAlcains阿尔凯因斯() 	feud.Barony
    BBelver贝韦尔() 	feud.Barony
    BCastelobranco布朗库堡() 	feud.Barony
    BCovilha科维良() 	feud.Barony
    BIdanha伊达尼亚() 	feud.Barony
    BPenha佩尼亚() 	feud.Barony
    BRodao罗当() 	feud.Barony
}

type c_egitaneaEgitaneaCounty struct {
	feud.BaseCounty
	阿尔凯因斯Alcains 	feud.Barony
	贝韦尔Belver 	feud.Barony
	布朗库堡Castelobranco 	feud.Barony
	科维良Covilha 	feud.Barony
	伊达尼亚Idanha 	feud.Barony
	佩尼亚Penha 	feud.Barony
	罗当Rodao 	feud.Barony
}

func (c *c_egitaneaEgitaneaCounty) BAlcains阿尔凯因斯() feud.Barony {
	return c.阿尔凯因斯Alcains
}
    
func (c *c_egitaneaEgitaneaCounty) BBelver贝韦尔() feud.Barony {
	return c.贝韦尔Belver
}
    
func (c *c_egitaneaEgitaneaCounty) BCastelobranco布朗库堡() feud.Barony {
	return c.布朗库堡Castelobranco
}
    
func (c *c_egitaneaEgitaneaCounty) BCovilha科维良() feud.Barony {
	return c.科维良Covilha
}
    
func (c *c_egitaneaEgitaneaCounty) BIdanha伊达尼亚() feud.Barony {
	return c.伊达尼亚Idanha
}
    
func (c *c_egitaneaEgitaneaCounty) BPenha佩尼亚() feud.Barony {
	return c.佩尼亚Penha
}
    
func (c *c_egitaneaEgitaneaCounty) BRodao罗当() feud.Barony {
	return c.罗当Rodao
}
    
var CEgitaneac_egitanea EgitaneaCounty = &c_egitaneaEgitaneaCounty{}

func init() {
	f := CEgitaneac_egitanea.(*c_egitaneaEgitaneaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2012",
		Title:     "egitanea",
		TitleName: "c_egitanea",
		TitleCode: "c_egitanea",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔凯因斯Alcains = BAlcains阿尔凯因斯
	f.阿尔凯因斯Alcains.SetParent(f)
	
	f.贝韦尔Belver = BBelver贝韦尔
	f.贝韦尔Belver.SetParent(f)
	
	f.布朗库堡Castelobranco = BCastelobranco布朗库堡
	f.布朗库堡Castelobranco.SetParent(f)
	
	f.科维良Covilha = BCovilha科维良
	f.科维良Covilha.SetParent(f)
	
	f.伊达尼亚Idanha = BIdanha伊达尼亚
	f.伊达尼亚Idanha.SetParent(f)
	
	f.佩尼亚Penha = BPenha佩尼亚
	f.佩尼亚Penha.SetParent(f)
	
	f.罗当Rodao = BRodao罗当
	f.罗当Rodao.SetParent(f)
	
}
