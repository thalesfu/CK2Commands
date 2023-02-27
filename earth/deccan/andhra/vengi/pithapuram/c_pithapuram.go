package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PithapuramCounty interface {
    feud.County
    BDraksharama达洛叉罗摩() 	feud.Barony
    BKakinada卡基纳达() 	feud.Barony
    BMandapeta曼达佩塔() 	feud.Barony
    BPithapuram毗多补罗() 	feud.Barony
    BSamalkota萨马尔科塔() 	feud.Barony
    BSomarama苏马拉玛() 	feud.Barony
    BWardha跋罗陀() 	feud.Barony
}

type 毗多补罗PithapuramCounty struct {
	feud.BaseCounty
	达洛叉罗摩Draksharama 	feud.Barony
	卡基纳达Kakinada 	feud.Barony
	曼达佩塔Mandapeta 	feud.Barony
	毗多补罗Pithapuram 	feud.Barony
	萨马尔科塔Samalkota 	feud.Barony
	苏马拉玛Somarama 	feud.Barony
	跋罗陀Wardha 	feud.Barony
}

func (c *毗多补罗PithapuramCounty) BDraksharama达洛叉罗摩() feud.Barony {
	return c.达洛叉罗摩Draksharama
}
    
func (c *毗多补罗PithapuramCounty) BKakinada卡基纳达() feud.Barony {
	return c.卡基纳达Kakinada
}
    
func (c *毗多补罗PithapuramCounty) BMandapeta曼达佩塔() feud.Barony {
	return c.曼达佩塔Mandapeta
}
    
func (c *毗多补罗PithapuramCounty) BPithapuram毗多补罗() feud.Barony {
	return c.毗多补罗Pithapuram
}
    
func (c *毗多补罗PithapuramCounty) BSamalkota萨马尔科塔() feud.Barony {
	return c.萨马尔科塔Samalkota
}
    
func (c *毗多补罗PithapuramCounty) BSomarama苏马拉玛() feud.Barony {
	return c.苏马拉玛Somarama
}
    
func (c *毗多补罗PithapuramCounty) BWardha跋罗陀() feud.Barony {
	return c.跋罗陀Wardha
}
    
var CPithapuram毗多补罗 PithapuramCounty = &毗多补罗PithapuramCounty{}

func init() {
	f := CPithapuram毗多补罗.(*毗多补罗PithapuramCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1415",
		Title:     "pithapuram",
		TitleName: "毗多补罗",
		TitleCode: "c_pithapuram",
		Baronies:  map[string]feud.Barony{},
	}

	f.达洛叉罗摩Draksharama = BDraksharama达洛叉罗摩
	f.达洛叉罗摩Draksharama.SetParent(f)
	
	f.卡基纳达Kakinada = BKakinada卡基纳达
	f.卡基纳达Kakinada.SetParent(f)
	
	f.曼达佩塔Mandapeta = BMandapeta曼达佩塔
	f.曼达佩塔Mandapeta.SetParent(f)
	
	f.毗多补罗Pithapuram = BPithapuram毗多补罗
	f.毗多补罗Pithapuram.SetParent(f)
	
	f.萨马尔科塔Samalkota = BSamalkota萨马尔科塔
	f.萨马尔科塔Samalkota.SetParent(f)
	
	f.苏马拉玛Somarama = BSomarama苏马拉玛
	f.苏马拉玛Somarama.SetParent(f)
	
	f.跋罗陀Wardha = BWardha跋罗陀
	f.跋罗陀Wardha.SetParent(f)
	
}
