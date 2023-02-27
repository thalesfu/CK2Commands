package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RummahCounty interface {
    feud.County
    BAlqalt阿喀拉特() 	feud.Barony
    BAltheebiyah阿特比杨() 	feud.Barony
    BArraqai阿尔雷喀() 	feud.Barony
    BAssuayerah阿斯苏叶雷() 	feud.Barony
    BAssufayri苏费里() 	feud.Barony
    BHafaralbatin哈夫阿巴廷() 	feud.Barony
    BQaisumah喀斯苏曼() 	feud.Barony
    BSamoudah萨穆达() 	feud.Barony
}

type 艾因赛德RummahCounty struct {
	feud.BaseCounty
	阿喀拉特Alqalt 	feud.Barony
	阿特比杨Altheebiyah 	feud.Barony
	阿尔雷喀Arraqai 	feud.Barony
	阿斯苏叶雷Assuayerah 	feud.Barony
	苏费里Assufayri 	feud.Barony
	哈夫阿巴廷Hafaralbatin 	feud.Barony
	喀斯苏曼Qaisumah 	feud.Barony
	萨穆达Samoudah 	feud.Barony
}

func (c *艾因赛德RummahCounty) BAlqalt阿喀拉特() feud.Barony {
	return c.阿喀拉特Alqalt
}
    
func (c *艾因赛德RummahCounty) BAltheebiyah阿特比杨() feud.Barony {
	return c.阿特比杨Altheebiyah
}
    
func (c *艾因赛德RummahCounty) BArraqai阿尔雷喀() feud.Barony {
	return c.阿尔雷喀Arraqai
}
    
func (c *艾因赛德RummahCounty) BAssuayerah阿斯苏叶雷() feud.Barony {
	return c.阿斯苏叶雷Assuayerah
}
    
func (c *艾因赛德RummahCounty) BAssufayri苏费里() feud.Barony {
	return c.苏费里Assufayri
}
    
func (c *艾因赛德RummahCounty) BHafaralbatin哈夫阿巴廷() feud.Barony {
	return c.哈夫阿巴廷Hafaralbatin
}
    
func (c *艾因赛德RummahCounty) BQaisumah喀斯苏曼() feud.Barony {
	return c.喀斯苏曼Qaisumah
}
    
func (c *艾因赛德RummahCounty) BSamoudah萨穆达() feud.Barony {
	return c.萨穆达Samoudah
}
    
var CRummah艾因赛德 RummahCounty = &艾因赛德RummahCounty{}

func init() {
	f := CRummah艾因赛德.(*艾因赛德RummahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "654",
		Title:     "rummah",
		TitleName: "艾因赛德",
		TitleCode: "c_rummah",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿喀拉特Alqalt = BAlqalt阿喀拉特
	f.阿喀拉特Alqalt.SetParent(f)
	
	f.阿特比杨Altheebiyah = BAltheebiyah阿特比杨
	f.阿特比杨Altheebiyah.SetParent(f)
	
	f.阿尔雷喀Arraqai = BArraqai阿尔雷喀
	f.阿尔雷喀Arraqai.SetParent(f)
	
	f.阿斯苏叶雷Assuayerah = BAssuayerah阿斯苏叶雷
	f.阿斯苏叶雷Assuayerah.SetParent(f)
	
	f.苏费里Assufayri = BAssufayri苏费里
	f.苏费里Assufayri.SetParent(f)
	
	f.哈夫阿巴廷Hafaralbatin = BHafaralbatin哈夫阿巴廷
	f.哈夫阿巴廷Hafaralbatin.SetParent(f)
	
	f.喀斯苏曼Qaisumah = BQaisumah喀斯苏曼
	f.喀斯苏曼Qaisumah.SetParent(f)
	
	f.萨穆达Samoudah = BSamoudah萨穆达
	f.萨穆达Samoudah.SetParent(f)
	
}
