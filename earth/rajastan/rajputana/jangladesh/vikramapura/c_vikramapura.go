package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VikramapuraCounty interface {
    feud.County
    BBhandasar畔荼娑罗() 	feud.Barony
    BBhurupal浮卢婆() 	feud.Barony
    BKolayat拘罗耶多() 	feud.Barony
    BOdur乌杜尔() 	feud.Barony
    BPallikondai波梨军荼() 	feud.Barony
    BShekhsar舍迦尔() 	feud.Barony
    BVikramapura毗讫罗摩补罗() 	feud.Barony
}

type 毗讫罗摩补罗VikramapuraCounty struct {
	feud.BaseCounty
	畔荼娑罗Bhandasar 	feud.Barony
	浮卢婆Bhurupal 	feud.Barony
	拘罗耶多Kolayat 	feud.Barony
	乌杜尔Odur 	feud.Barony
	波梨军荼Pallikondai 	feud.Barony
	舍迦尔Shekhsar 	feud.Barony
	毗讫罗摩补罗Vikramapura 	feud.Barony
}

func (c *毗讫罗摩补罗VikramapuraCounty) BBhandasar畔荼娑罗() feud.Barony {
	return c.畔荼娑罗Bhandasar
}
    
func (c *毗讫罗摩补罗VikramapuraCounty) BBhurupal浮卢婆() feud.Barony {
	return c.浮卢婆Bhurupal
}
    
func (c *毗讫罗摩补罗VikramapuraCounty) BKolayat拘罗耶多() feud.Barony {
	return c.拘罗耶多Kolayat
}
    
func (c *毗讫罗摩补罗VikramapuraCounty) BOdur乌杜尔() feud.Barony {
	return c.乌杜尔Odur
}
    
func (c *毗讫罗摩补罗VikramapuraCounty) BPallikondai波梨军荼() feud.Barony {
	return c.波梨军荼Pallikondai
}
    
func (c *毗讫罗摩补罗VikramapuraCounty) BShekhsar舍迦尔() feud.Barony {
	return c.舍迦尔Shekhsar
}
    
func (c *毗讫罗摩补罗VikramapuraCounty) BVikramapura毗讫罗摩补罗() feud.Barony {
	return c.毗讫罗摩补罗Vikramapura
}
    
var CVikramapura毗讫罗摩补罗 VikramapuraCounty = &毗讫罗摩补罗VikramapuraCounty{}

func init() {
	f := CVikramapura毗讫罗摩补罗.(*毗讫罗摩补罗VikramapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1349",
		Title:     "vikramapura",
		TitleName: "毗讫罗摩补罗",
		TitleCode: "c_vikramapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.畔荼娑罗Bhandasar = BBhandasar畔荼娑罗
	f.畔荼娑罗Bhandasar.SetParent(f)
	
	f.浮卢婆Bhurupal = BBhurupal浮卢婆
	f.浮卢婆Bhurupal.SetParent(f)
	
	f.拘罗耶多Kolayat = BKolayat拘罗耶多
	f.拘罗耶多Kolayat.SetParent(f)
	
	f.乌杜尔Odur = BOdur乌杜尔
	f.乌杜尔Odur.SetParent(f)
	
	f.波梨军荼Pallikondai = BPallikondai波梨军荼
	f.波梨军荼Pallikondai.SetParent(f)
	
	f.舍迦尔Shekhsar = BShekhsar舍迦尔
	f.舍迦尔Shekhsar.SetParent(f)
	
	f.毗讫罗摩补罗Vikramapura = BVikramapura毗讫罗摩补罗
	f.毗讫罗摩补罗Vikramapura.SetParent(f)
	
}
