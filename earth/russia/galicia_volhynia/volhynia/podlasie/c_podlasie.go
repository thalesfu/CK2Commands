package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PodlasieCounty interface {
    feud.County
    BDrohiczyn德罗希琴() 	feud.Barony
    BGoniadz戈尼翁兹() 	feud.Barony
    BKrynki克伦基() 	feud.Barony
    BLapy瓦佩() 	feud.Barony
    BRajgrod拉伊格鲁德() 	feud.Barony
    BSejny塞伊内() 	feud.Barony
    BTykocin蒂科钦() 	feud.Barony
    BZambrow赞布鲁夫() 	feud.Barony
}

type 波德拉谢PodlasieCounty struct {
	feud.BaseCounty
	德罗希琴Drohiczyn 	feud.Barony
	戈尼翁兹Goniadz 	feud.Barony
	克伦基Krynki 	feud.Barony
	瓦佩Lapy 	feud.Barony
	拉伊格鲁德Rajgrod 	feud.Barony
	塞伊内Sejny 	feud.Barony
	蒂科钦Tykocin 	feud.Barony
	赞布鲁夫Zambrow 	feud.Barony
}

func (c *波德拉谢PodlasieCounty) BDrohiczyn德罗希琴() feud.Barony {
	return c.德罗希琴Drohiczyn
}
    
func (c *波德拉谢PodlasieCounty) BGoniadz戈尼翁兹() feud.Barony {
	return c.戈尼翁兹Goniadz
}
    
func (c *波德拉谢PodlasieCounty) BKrynki克伦基() feud.Barony {
	return c.克伦基Krynki
}
    
func (c *波德拉谢PodlasieCounty) BLapy瓦佩() feud.Barony {
	return c.瓦佩Lapy
}
    
func (c *波德拉谢PodlasieCounty) BRajgrod拉伊格鲁德() feud.Barony {
	return c.拉伊格鲁德Rajgrod
}
    
func (c *波德拉谢PodlasieCounty) BSejny塞伊内() feud.Barony {
	return c.塞伊内Sejny
}
    
func (c *波德拉谢PodlasieCounty) BTykocin蒂科钦() feud.Barony {
	return c.蒂科钦Tykocin
}
    
func (c *波德拉谢PodlasieCounty) BZambrow赞布鲁夫() feud.Barony {
	return c.赞布鲁夫Zambrow
}
    
var CPodlasie波德拉谢 PodlasieCounty = &波德拉谢PodlasieCounty{}

func init() {
	f := CPodlasie波德拉谢.(*波德拉谢PodlasieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "425",
		Title:     "podlasie",
		TitleName: "波德拉谢",
		TitleCode: "c_podlasie",
		Baronies:  map[string]feud.Barony{},
	}

	f.德罗希琴Drohiczyn = BDrohiczyn德罗希琴
	f.德罗希琴Drohiczyn.SetParent(f)
	
	f.戈尼翁兹Goniadz = BGoniadz戈尼翁兹
	f.戈尼翁兹Goniadz.SetParent(f)
	
	f.克伦基Krynki = BKrynki克伦基
	f.克伦基Krynki.SetParent(f)
	
	f.瓦佩Lapy = BLapy瓦佩
	f.瓦佩Lapy.SetParent(f)
	
	f.拉伊格鲁德Rajgrod = BRajgrod拉伊格鲁德
	f.拉伊格鲁德Rajgrod.SetParent(f)
	
	f.塞伊内Sejny = BSejny塞伊内
	f.塞伊内Sejny.SetParent(f)
	
	f.蒂科钦Tykocin = BTykocin蒂科钦
	f.蒂科钦Tykocin.SetParent(f)
	
	f.赞布鲁夫Zambrow = BZambrow赞布鲁夫
	f.赞布鲁夫Zambrow.SetParent(f)
	
}
