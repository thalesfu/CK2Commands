package kumara_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kumara_mandalaCounty interface {
    feud.County
    BBetmangala苾曼伽罗() 	feud.Barony
    BEkdala夷迦陀罗() 	feud.Barony
    BFathabad法特哈巴德() 	feud.Barony
    BKatasgarh羯怛罗娑姞利呬() 	feud.Barony
    BKumarakhali鸠摩罗佉梨() 	feud.Barony
    BMolvailup牟毗卢() 	feud.Barony
    BNaneli那内梨() 	feud.Barony
}

type 鸠摩罗曼荼罗Kumara_mandalaCounty struct {
	feud.BaseCounty
	苾曼伽罗Betmangala 	feud.Barony
	夷迦陀罗Ekdala 	feud.Barony
	法特哈巴德Fathabad 	feud.Barony
	羯怛罗娑姞利呬Katasgarh 	feud.Barony
	鸠摩罗佉梨Kumarakhali 	feud.Barony
	牟毗卢Molvailup 	feud.Barony
	那内梨Naneli 	feud.Barony
}

func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BBetmangala苾曼伽罗() feud.Barony {
	return c.苾曼伽罗Betmangala
}
    
func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BEkdala夷迦陀罗() feud.Barony {
	return c.夷迦陀罗Ekdala
}
    
func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BFathabad法特哈巴德() feud.Barony {
	return c.法特哈巴德Fathabad
}
    
func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BKatasgarh羯怛罗娑姞利呬() feud.Barony {
	return c.羯怛罗娑姞利呬Katasgarh
}
    
func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BKumarakhali鸠摩罗佉梨() feud.Barony {
	return c.鸠摩罗佉梨Kumarakhali
}
    
func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BMolvailup牟毗卢() feud.Barony {
	return c.牟毗卢Molvailup
}
    
func (c *鸠摩罗曼荼罗Kumara_mandalaCounty) BNaneli那内梨() feud.Barony {
	return c.那内梨Naneli
}
    
var CKumara_mandala鸠摩罗曼荼罗 Kumara_mandalaCounty = &鸠摩罗曼荼罗Kumara_mandalaCounty{}

func init() {
	f := CKumara_mandala鸠摩罗曼荼罗.(*鸠摩罗曼荼罗Kumara_mandalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1240",
		Title:     "kumara_mandala",
		TitleName: "鸠摩罗曼荼罗",
		TitleCode: "c_kumara_mandala",
		Baronies:  map[string]feud.Barony{},
	}

	f.苾曼伽罗Betmangala = BBetmangala苾曼伽罗
	f.苾曼伽罗Betmangala.SetParent(f)
	
	f.夷迦陀罗Ekdala = BEkdala夷迦陀罗
	f.夷迦陀罗Ekdala.SetParent(f)
	
	f.法特哈巴德Fathabad = BFathabad法特哈巴德
	f.法特哈巴德Fathabad.SetParent(f)
	
	f.羯怛罗娑姞利呬Katasgarh = BKatasgarh羯怛罗娑姞利呬
	f.羯怛罗娑姞利呬Katasgarh.SetParent(f)
	
	f.鸠摩罗佉梨Kumarakhali = BKumarakhali鸠摩罗佉梨
	f.鸠摩罗佉梨Kumarakhali.SetParent(f)
	
	f.牟毗卢Molvailup = BMolvailup牟毗卢
	f.牟毗卢Molvailup.SetParent(f)
	
	f.那内梨Naneli = BNaneli那内梨
	f.那内梨Naneli.SetParent(f)
	
}
