package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LepielCounty interface {
    feud.County
    BDokshytsy多克希齐() 	feud.Barony
    BHlybokaye格卢博科耶() 	feud.Barony
    BLepiel列佩利() 	feud.Barony
    BLukoml卢科姆利() 	feud.Barony
    BNovalukoml诺瓦卢科姆利() 	feud.Barony
    BSharkawshchyna沙尔科夫希纳() 	feud.Barony
    BUshachy乌沙奇() 	feud.Barony
}

type 卢科姆利LepielCounty struct {
	feud.BaseCounty
	多克希齐Dokshytsy 	feud.Barony
	格卢博科耶Hlybokaye 	feud.Barony
	列佩利Lepiel 	feud.Barony
	卢科姆利Lukoml 	feud.Barony
	诺瓦卢科姆利Novalukoml 	feud.Barony
	沙尔科夫希纳Sharkawshchyna 	feud.Barony
	乌沙奇Ushachy 	feud.Barony
}

func (c *卢科姆利LepielCounty) BDokshytsy多克希齐() feud.Barony {
	return c.多克希齐Dokshytsy
}
    
func (c *卢科姆利LepielCounty) BHlybokaye格卢博科耶() feud.Barony {
	return c.格卢博科耶Hlybokaye
}
    
func (c *卢科姆利LepielCounty) BLepiel列佩利() feud.Barony {
	return c.列佩利Lepiel
}
    
func (c *卢科姆利LepielCounty) BLukoml卢科姆利() feud.Barony {
	return c.卢科姆利Lukoml
}
    
func (c *卢科姆利LepielCounty) BNovalukoml诺瓦卢科姆利() feud.Barony {
	return c.诺瓦卢科姆利Novalukoml
}
    
func (c *卢科姆利LepielCounty) BSharkawshchyna沙尔科夫希纳() feud.Barony {
	return c.沙尔科夫希纳Sharkawshchyna
}
    
func (c *卢科姆利LepielCounty) BUshachy乌沙奇() feud.Barony {
	return c.乌沙奇Ushachy
}
    
var CLepiel卢科姆利 LepielCounty = &卢科姆利LepielCounty{}

func init() {
	f := CLepiel卢科姆利.(*卢科姆利LepielCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "934",
		Title:     "lepiel",
		TitleName: "卢科姆利",
		TitleCode: "c_lepiel",
		Baronies:  map[string]feud.Barony{},
	}

	f.多克希齐Dokshytsy = BDokshytsy多克希齐
	f.多克希齐Dokshytsy.SetParent(f)
	
	f.格卢博科耶Hlybokaye = BHlybokaye格卢博科耶
	f.格卢博科耶Hlybokaye.SetParent(f)
	
	f.列佩利Lepiel = BLepiel列佩利
	f.列佩利Lepiel.SetParent(f)
	
	f.卢科姆利Lukoml = BLukoml卢科姆利
	f.卢科姆利Lukoml.SetParent(f)
	
	f.诺瓦卢科姆利Novalukoml = BNovalukoml诺瓦卢科姆利
	f.诺瓦卢科姆利Novalukoml.SetParent(f)
	
	f.沙尔科夫希纳Sharkawshchyna = BSharkawshchyna沙尔科夫希纳
	f.沙尔科夫希纳Sharkawshchyna.SetParent(f)
	
	f.乌沙奇Ushachy = BUshachy乌沙奇
	f.乌沙奇Ushachy.SetParent(f)
	
}
