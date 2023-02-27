package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VairagaraCounty interface {
    feud.County
    BBhandugara婆度加罗() 	feud.Barony
    BCikamburika锡甘布利加() 	feud.Barony
    BGondia贡迪亚() 	feud.Barony
    BKas伽斯() 	feud.Barony
    BKauria乔梨() 	feud.Barony
    BKelod艺卢() 	feud.Barony
    BVairagara吠罗伽罗() 	feud.Barony
}

type 吠罗伽罗VairagaraCounty struct {
	feud.BaseCounty
	婆度加罗Bhandugara 	feud.Barony
	锡甘布利加Cikamburika 	feud.Barony
	贡迪亚Gondia 	feud.Barony
	伽斯Kas 	feud.Barony
	乔梨Kauria 	feud.Barony
	艺卢Kelod 	feud.Barony
	吠罗伽罗Vairagara 	feud.Barony
}

func (c *吠罗伽罗VairagaraCounty) BBhandugara婆度加罗() feud.Barony {
	return c.婆度加罗Bhandugara
}
    
func (c *吠罗伽罗VairagaraCounty) BCikamburika锡甘布利加() feud.Barony {
	return c.锡甘布利加Cikamburika
}
    
func (c *吠罗伽罗VairagaraCounty) BGondia贡迪亚() feud.Barony {
	return c.贡迪亚Gondia
}
    
func (c *吠罗伽罗VairagaraCounty) BKas伽斯() feud.Barony {
	return c.伽斯Kas
}
    
func (c *吠罗伽罗VairagaraCounty) BKauria乔梨() feud.Barony {
	return c.乔梨Kauria
}
    
func (c *吠罗伽罗VairagaraCounty) BKelod艺卢() feud.Barony {
	return c.艺卢Kelod
}
    
func (c *吠罗伽罗VairagaraCounty) BVairagara吠罗伽罗() feud.Barony {
	return c.吠罗伽罗Vairagara
}
    
var CVairagara吠罗伽罗 VairagaraCounty = &吠罗伽罗VairagaraCounty{}

func init() {
	f := CVairagara吠罗伽罗.(*吠罗伽罗VairagaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1254",
		Title:     "vairagara",
		TitleName: "吠罗伽罗",
		TitleCode: "c_vairagara",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆度加罗Bhandugara = BBhandugara婆度加罗
	f.婆度加罗Bhandugara.SetParent(f)
	
	f.锡甘布利加Cikamburika = BCikamburika锡甘布利加
	f.锡甘布利加Cikamburika.SetParent(f)
	
	f.贡迪亚Gondia = BGondia贡迪亚
	f.贡迪亚Gondia.SetParent(f)
	
	f.伽斯Kas = BKas伽斯
	f.伽斯Kas.SetParent(f)
	
	f.乔梨Kauria = BKauria乔梨
	f.乔梨Kauria.SetParent(f)
	
	f.艺卢Kelod = BKelod艺卢
	f.艺卢Kelod.SetParent(f)
	
	f.吠罗伽罗Vairagara = BVairagara吠罗伽罗
	f.吠罗伽罗Vairagara.SetParent(f)
	
}
