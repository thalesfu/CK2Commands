package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalinganagarCounty interface {
    feud.County
    BKalinganagara羯陵伽那揭罗() 	feud.Barony
    BKalingapattanam羯陵伽波多那() 	feud.Barony
    BKotra拘多罗() 	feud.Barony
    BManappakkam摩那波甘() 	feud.Barony
    BMandasa曼陀娑() 	feud.Barony
    BSrikakulam室利迦俱蓝() 	feud.Barony
    BSrikurmam室利俱利蒙() 	feud.Barony
}

type 羯陵伽那揭罗KalinganagarCounty struct {
	feud.BaseCounty
	羯陵伽那揭罗Kalinganagara 	feud.Barony
	羯陵伽波多那Kalingapattanam 	feud.Barony
	拘多罗Kotra 	feud.Barony
	摩那波甘Manappakkam 	feud.Barony
	曼陀娑Mandasa 	feud.Barony
	室利迦俱蓝Srikakulam 	feud.Barony
	室利俱利蒙Srikurmam 	feud.Barony
}

func (c *羯陵伽那揭罗KalinganagarCounty) BKalinganagara羯陵伽那揭罗() feud.Barony {
	return c.羯陵伽那揭罗Kalinganagara
}
    
func (c *羯陵伽那揭罗KalinganagarCounty) BKalingapattanam羯陵伽波多那() feud.Barony {
	return c.羯陵伽波多那Kalingapattanam
}
    
func (c *羯陵伽那揭罗KalinganagarCounty) BKotra拘多罗() feud.Barony {
	return c.拘多罗Kotra
}
    
func (c *羯陵伽那揭罗KalinganagarCounty) BManappakkam摩那波甘() feud.Barony {
	return c.摩那波甘Manappakkam
}
    
func (c *羯陵伽那揭罗KalinganagarCounty) BMandasa曼陀娑() feud.Barony {
	return c.曼陀娑Mandasa
}
    
func (c *羯陵伽那揭罗KalinganagarCounty) BSrikakulam室利迦俱蓝() feud.Barony {
	return c.室利迦俱蓝Srikakulam
}
    
func (c *羯陵伽那揭罗KalinganagarCounty) BSrikurmam室利俱利蒙() feud.Barony {
	return c.室利俱利蒙Srikurmam
}
    
var CKalinganagar羯陵伽那揭罗 KalinganagarCounty = &羯陵伽那揭罗KalinganagarCounty{}

func init() {
	f := CKalinganagar羯陵伽那揭罗.(*羯陵伽那揭罗KalinganagarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1224",
		Title:     "kalinganagar",
		TitleName: "羯陵伽那揭罗",
		TitleCode: "c_kalinganagar",
		Baronies:  map[string]feud.Barony{},
	}

	f.羯陵伽那揭罗Kalinganagara = BKalinganagara羯陵伽那揭罗
	f.羯陵伽那揭罗Kalinganagara.SetParent(f)
	
	f.羯陵伽波多那Kalingapattanam = BKalingapattanam羯陵伽波多那
	f.羯陵伽波多那Kalingapattanam.SetParent(f)
	
	f.拘多罗Kotra = BKotra拘多罗
	f.拘多罗Kotra.SetParent(f)
	
	f.摩那波甘Manappakkam = BManappakkam摩那波甘
	f.摩那波甘Manappakkam.SetParent(f)
	
	f.曼陀娑Mandasa = BMandasa曼陀娑
	f.曼陀娑Mandasa.SetParent(f)
	
	f.室利迦俱蓝Srikakulam = BSrikakulam室利迦俱蓝
	f.室利迦俱蓝Srikakulam.SetParent(f)
	
	f.室利俱利蒙Srikurmam = BSrikurmam室利俱利蒙
	f.室利俱利蒙Srikurmam.SetParent(f)
	
}
