package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Galich_merskyCounty interface {
    feud.County
    BBuy布伊() 	feud.Barony
    BChistyebory奇斯特耶博雷() 	feud.Barony
    BGalichmersky加利奇梅尔斯基() 	feud.Barony
    BGradmersky格拉德梅尔斯基() 	feud.Barony
    BIsaevo伊萨耶沃() 	feud.Barony
    BKadyy卡德() 	feud.Barony
    BLevkovo列夫科沃() 	feud.Barony
    BSusanino苏萨尼诺() 	feud.Barony
}

type 加利奇梅尔斯基Galich_merskyCounty struct {
	feud.BaseCounty
	布伊Buy 	feud.Barony
	奇斯特耶博雷Chistyebory 	feud.Barony
	加利奇梅尔斯基Galichmersky 	feud.Barony
	格拉德梅尔斯基Gradmersky 	feud.Barony
	伊萨耶沃Isaevo 	feud.Barony
	卡德Kadyy 	feud.Barony
	列夫科沃Levkovo 	feud.Barony
	苏萨尼诺Susanino 	feud.Barony
}

func (c *加利奇梅尔斯基Galich_merskyCounty) BBuy布伊() feud.Barony {
	return c.布伊Buy
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BChistyebory奇斯特耶博雷() feud.Barony {
	return c.奇斯特耶博雷Chistyebory
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BGalichmersky加利奇梅尔斯基() feud.Barony {
	return c.加利奇梅尔斯基Galichmersky
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BGradmersky格拉德梅尔斯基() feud.Barony {
	return c.格拉德梅尔斯基Gradmersky
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BIsaevo伊萨耶沃() feud.Barony {
	return c.伊萨耶沃Isaevo
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BKadyy卡德() feud.Barony {
	return c.卡德Kadyy
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BLevkovo列夫科沃() feud.Barony {
	return c.列夫科沃Levkovo
}
    
func (c *加利奇梅尔斯基Galich_merskyCounty) BSusanino苏萨尼诺() feud.Barony {
	return c.苏萨尼诺Susanino
}
    
var CGalich_mersky加利奇梅尔斯基 Galich_merskyCounty = &加利奇梅尔斯基Galich_merskyCounty{}

func init() {
	f := CGalich_mersky加利奇梅尔斯基.(*加利奇梅尔斯基Galich_merskyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "586",
		Title:     "galich_mersky",
		TitleName: "加利奇梅尔斯基",
		TitleCode: "c_galich_mersky",
		Baronies:  map[string]feud.Barony{},
	}

	f.布伊Buy = BBuy布伊
	f.布伊Buy.SetParent(f)
	
	f.奇斯特耶博雷Chistyebory = BChistyebory奇斯特耶博雷
	f.奇斯特耶博雷Chistyebory.SetParent(f)
	
	f.加利奇梅尔斯基Galichmersky = BGalichmersky加利奇梅尔斯基
	f.加利奇梅尔斯基Galichmersky.SetParent(f)
	
	f.格拉德梅尔斯基Gradmersky = BGradmersky格拉德梅尔斯基
	f.格拉德梅尔斯基Gradmersky.SetParent(f)
	
	f.伊萨耶沃Isaevo = BIsaevo伊萨耶沃
	f.伊萨耶沃Isaevo.SetParent(f)
	
	f.卡德Kadyy = BKadyy卡德
	f.卡德Kadyy.SetParent(f)
	
	f.列夫科沃Levkovo = BLevkovo列夫科沃
	f.列夫科沃Levkovo.SetParent(f)
	
	f.苏萨尼诺Susanino = BSusanino苏萨尼诺
	f.苏萨尼诺Susanino.SetParent(f)
	
}
