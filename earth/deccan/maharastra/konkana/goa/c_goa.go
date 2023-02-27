package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GoaCounty interface {
    feud.County
    BChandrapur旃陀罗补罗() 	feud.Barony
    BChitrakul质多罗俱罗() 	feud.Barony
    BGheria盖里亚() 	feud.Barony
    BGopakapattana瞿波迦钵多那() 	feud.Barony
    BKapardikadvipa迦钵地迦提鞞波() 	feud.Barony
    BKapurgarh迦补罗姞利呬() 	feud.Barony
    BMathagram摩侘伽罗摩() 	feud.Barony
}

type 果阿GoaCounty struct {
	feud.BaseCounty
	旃陀罗补罗Chandrapur 	feud.Barony
	质多罗俱罗Chitrakul 	feud.Barony
	盖里亚Gheria 	feud.Barony
	瞿波迦钵多那Gopakapattana 	feud.Barony
	迦钵地迦提鞞波Kapardikadvipa 	feud.Barony
	迦补罗姞利呬Kapurgarh 	feud.Barony
	摩侘伽罗摩Mathagram 	feud.Barony
}

func (c *果阿GoaCounty) BChandrapur旃陀罗补罗() feud.Barony {
	return c.旃陀罗补罗Chandrapur
}
    
func (c *果阿GoaCounty) BChitrakul质多罗俱罗() feud.Barony {
	return c.质多罗俱罗Chitrakul
}
    
func (c *果阿GoaCounty) BGheria盖里亚() feud.Barony {
	return c.盖里亚Gheria
}
    
func (c *果阿GoaCounty) BGopakapattana瞿波迦钵多那() feud.Barony {
	return c.瞿波迦钵多那Gopakapattana
}
    
func (c *果阿GoaCounty) BKapardikadvipa迦钵地迦提鞞波() feud.Barony {
	return c.迦钵地迦提鞞波Kapardikadvipa
}
    
func (c *果阿GoaCounty) BKapurgarh迦补罗姞利呬() feud.Barony {
	return c.迦补罗姞利呬Kapurgarh
}
    
func (c *果阿GoaCounty) BMathagram摩侘伽罗摩() feud.Barony {
	return c.摩侘伽罗摩Mathagram
}
    
var CGoa果阿 GoaCounty = &果阿GoaCounty{}

func init() {
	f := CGoa果阿.(*果阿GoaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1220",
		Title:     "goa",
		TitleName: "果阿",
		TitleCode: "c_goa",
		Baronies:  map[string]feud.Barony{},
	}

	f.旃陀罗补罗Chandrapur = BChandrapur旃陀罗补罗
	f.旃陀罗补罗Chandrapur.SetParent(f)
	
	f.质多罗俱罗Chitrakul = BChitrakul质多罗俱罗
	f.质多罗俱罗Chitrakul.SetParent(f)
	
	f.盖里亚Gheria = BGheria盖里亚
	f.盖里亚Gheria.SetParent(f)
	
	f.瞿波迦钵多那Gopakapattana = BGopakapattana瞿波迦钵多那
	f.瞿波迦钵多那Gopakapattana.SetParent(f)
	
	f.迦钵地迦提鞞波Kapardikadvipa = BKapardikadvipa迦钵地迦提鞞波
	f.迦钵地迦提鞞波Kapardikadvipa.SetParent(f)
	
	f.迦补罗姞利呬Kapurgarh = BKapurgarh迦补罗姞利呬
	f.迦补罗姞利呬Kapurgarh.SetParent(f)
	
	f.摩侘伽罗摩Mathagram = BMathagram摩侘伽罗摩
	f.摩侘伽罗摩Mathagram.SetParent(f)
	
}
