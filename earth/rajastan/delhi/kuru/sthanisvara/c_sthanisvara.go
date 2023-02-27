package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SthanisvaraCounty interface {
    feud.County
    BKarnal羯罗拏罗() 	feud.Barony
    BPanipat波尼跋() 	feud.Barony
    BSaharanpur娑诃兰那城() 	feud.Barony
    BSamana娑摩那() 	feud.Barony
    BShakumbhri_devi舍甘婆利提毗居处() 	feud.Barony
    BSirhind西尔欣德() 	feud.Barony
    BSthanisvara萨他泥湿伐罗() 	feud.Barony
}

type 萨他泥湿伐罗SthanisvaraCounty struct {
	feud.BaseCounty
	羯罗拏罗Karnal 	feud.Barony
	波尼跋Panipat 	feud.Barony
	娑诃兰那城Saharanpur 	feud.Barony
	娑摩那Samana 	feud.Barony
	舍甘婆利提毗居处Shakumbhri_devi 	feud.Barony
	西尔欣德Sirhind 	feud.Barony
	萨他泥湿伐罗Sthanisvara 	feud.Barony
}

func (c *萨他泥湿伐罗SthanisvaraCounty) BKarnal羯罗拏罗() feud.Barony {
	return c.羯罗拏罗Karnal
}
    
func (c *萨他泥湿伐罗SthanisvaraCounty) BPanipat波尼跋() feud.Barony {
	return c.波尼跋Panipat
}
    
func (c *萨他泥湿伐罗SthanisvaraCounty) BSaharanpur娑诃兰那城() feud.Barony {
	return c.娑诃兰那城Saharanpur
}
    
func (c *萨他泥湿伐罗SthanisvaraCounty) BSamana娑摩那() feud.Barony {
	return c.娑摩那Samana
}
    
func (c *萨他泥湿伐罗SthanisvaraCounty) BShakumbhri_devi舍甘婆利提毗居处() feud.Barony {
	return c.舍甘婆利提毗居处Shakumbhri_devi
}
    
func (c *萨他泥湿伐罗SthanisvaraCounty) BSirhind西尔欣德() feud.Barony {
	return c.西尔欣德Sirhind
}
    
func (c *萨他泥湿伐罗SthanisvaraCounty) BSthanisvara萨他泥湿伐罗() feud.Barony {
	return c.萨他泥湿伐罗Sthanisvara
}
    
var CSthanisvara萨他泥湿伐罗 SthanisvaraCounty = &萨他泥湿伐罗SthanisvaraCounty{}

func init() {
	f := CSthanisvara萨他泥湿伐罗.(*萨他泥湿伐罗SthanisvaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1367",
		Title:     "sthanisvara",
		TitleName: "萨他泥湿伐罗",
		TitleCode: "c_sthanisvara",
		Baronies:  map[string]feud.Barony{},
	}

	f.羯罗拏罗Karnal = BKarnal羯罗拏罗
	f.羯罗拏罗Karnal.SetParent(f)
	
	f.波尼跋Panipat = BPanipat波尼跋
	f.波尼跋Panipat.SetParent(f)
	
	f.娑诃兰那城Saharanpur = BSaharanpur娑诃兰那城
	f.娑诃兰那城Saharanpur.SetParent(f)
	
	f.娑摩那Samana = BSamana娑摩那
	f.娑摩那Samana.SetParent(f)
	
	f.舍甘婆利提毗居处Shakumbhri_devi = BShakumbhri_devi舍甘婆利提毗居处
	f.舍甘婆利提毗居处Shakumbhri_devi.SetParent(f)
	
	f.西尔欣德Sirhind = BSirhind西尔欣德
	f.西尔欣德Sirhind.SetParent(f)
	
	f.萨他泥湿伐罗Sthanisvara = BSthanisvara萨他泥湿伐罗
	f.萨他泥湿伐罗Sthanisvara.SetParent(f)
	
}
