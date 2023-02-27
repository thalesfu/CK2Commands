package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DimapurCounty interface {
    feud.County
    BAskot阿尸拘吒() 	feud.Barony
    BBagchara薄伽遮罗() 	feud.Barony
    BDimapur提摩补罗() 	feud.Barony
    BDindori陈豆利() 	feud.Barony
    BHerombial诃夷卢摩毗阿罗() 	feud.Barony
    BMaibong摩夷菩那迦() 	feud.Barony
}

type 提摩补罗DimapurCounty struct {
	feud.BaseCounty
	阿尸拘吒Askot 	feud.Barony
	薄伽遮罗Bagchara 	feud.Barony
	提摩补罗Dimapur 	feud.Barony
	陈豆利Dindori 	feud.Barony
	诃夷卢摩毗阿罗Herombial 	feud.Barony
	摩夷菩那迦Maibong 	feud.Barony
}

func (c *提摩补罗DimapurCounty) BAskot阿尸拘吒() feud.Barony {
	return c.阿尸拘吒Askot
}
    
func (c *提摩补罗DimapurCounty) BBagchara薄伽遮罗() feud.Barony {
	return c.薄伽遮罗Bagchara
}
    
func (c *提摩补罗DimapurCounty) BDimapur提摩补罗() feud.Barony {
	return c.提摩补罗Dimapur
}
    
func (c *提摩补罗DimapurCounty) BDindori陈豆利() feud.Barony {
	return c.陈豆利Dindori
}
    
func (c *提摩补罗DimapurCounty) BHerombial诃夷卢摩毗阿罗() feud.Barony {
	return c.诃夷卢摩毗阿罗Herombial
}
    
func (c *提摩补罗DimapurCounty) BMaibong摩夷菩那迦() feud.Barony {
	return c.摩夷菩那迦Maibong
}
    
var CDimapur提摩补罗 DimapurCounty = &提摩补罗DimapurCounty{}

func init() {
	f := CDimapur提摩补罗.(*提摩补罗DimapurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1296",
		Title:     "dimapur",
		TitleName: "提摩补罗",
		TitleCode: "c_dimapur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尸拘吒Askot = BAskot阿尸拘吒
	f.阿尸拘吒Askot.SetParent(f)
	
	f.薄伽遮罗Bagchara = BBagchara薄伽遮罗
	f.薄伽遮罗Bagchara.SetParent(f)
	
	f.提摩补罗Dimapur = BDimapur提摩补罗
	f.提摩补罗Dimapur.SetParent(f)
	
	f.陈豆利Dindori = BDindori陈豆利
	f.陈豆利Dindori.SetParent(f)
	
	f.诃夷卢摩毗阿罗Herombial = BHerombial诃夷卢摩毗阿罗
	f.诃夷卢摩毗阿罗Herombial.SetParent(f)
	
	f.摩夷菩那迦Maibong = BMaibong摩夷菩那迦
	f.摩夷菩那迦Maibong.SetParent(f)
	
}
