package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VesoulCounty interface {
    feud.County
    BBeaucourt博库尔() 	feud.Barony
    BBelfort贝尔福() 	feud.Barony
    BChamplitte尚普利特() 	feud.Barony
    BJussey瑞塞() 	feud.Barony
    BLure吕尔() 	feud.Barony
    BLuxeuil吕克瑟伊() 	feud.Barony
    BVesoul沃苏勒() 	feud.Barony
}

type 波图瓦VesoulCounty struct {
	feud.BaseCounty
	博库尔Beaucourt 	feud.Barony
	贝尔福Belfort 	feud.Barony
	尚普利特Champlitte 	feud.Barony
	瑞塞Jussey 	feud.Barony
	吕尔Lure 	feud.Barony
	吕克瑟伊Luxeuil 	feud.Barony
	沃苏勒Vesoul 	feud.Barony
}

func (c *波图瓦VesoulCounty) BBeaucourt博库尔() feud.Barony {
	return c.博库尔Beaucourt
}
    
func (c *波图瓦VesoulCounty) BBelfort贝尔福() feud.Barony {
	return c.贝尔福Belfort
}
    
func (c *波图瓦VesoulCounty) BChamplitte尚普利特() feud.Barony {
	return c.尚普利特Champlitte
}
    
func (c *波图瓦VesoulCounty) BJussey瑞塞() feud.Barony {
	return c.瑞塞Jussey
}
    
func (c *波图瓦VesoulCounty) BLure吕尔() feud.Barony {
	return c.吕尔Lure
}
    
func (c *波图瓦VesoulCounty) BLuxeuil吕克瑟伊() feud.Barony {
	return c.吕克瑟伊Luxeuil
}
    
func (c *波图瓦VesoulCounty) BVesoul沃苏勒() feud.Barony {
	return c.沃苏勒Vesoul
}
    
var CVesoul波图瓦 VesoulCounty = &波图瓦VesoulCounty{}

func init() {
	f := CVesoul波图瓦.(*波图瓦VesoulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1609",
		Title:     "vesoul",
		TitleName: "波图瓦",
		TitleCode: "c_vesoul",
		Baronies:  map[string]feud.Barony{},
	}

	f.博库尔Beaucourt = BBeaucourt博库尔
	f.博库尔Beaucourt.SetParent(f)
	
	f.贝尔福Belfort = BBelfort贝尔福
	f.贝尔福Belfort.SetParent(f)
	
	f.尚普利特Champlitte = BChamplitte尚普利特
	f.尚普利特Champlitte.SetParent(f)
	
	f.瑞塞Jussey = BJussey瑞塞
	f.瑞塞Jussey.SetParent(f)
	
	f.吕尔Lure = BLure吕尔
	f.吕尔Lure.SetParent(f)
	
	f.吕克瑟伊Luxeuil = BLuxeuil吕克瑟伊
	f.吕克瑟伊Luxeuil.SetParent(f)
	
	f.沃苏勒Vesoul = BVesoul沃苏勒
	f.沃苏勒Vesoul.SetParent(f)
	
}
