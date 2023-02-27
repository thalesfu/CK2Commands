package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZeelandCounty interface {
    feud.County
    BBorsele博尔瑟莱() 	feud.Barony
    BCadzand卡德赞德() 	feud.Barony
    BMiddelburg米德尔堡() 	feud.Barony
    BRenesse雷讷瑟() 	feud.Barony
    BTholen托伦() 	feud.Barony
    BVeere费勒() 	feud.Barony
    BVlissingen弗利辛亨() 	feud.Barony
    BZierikzee济里克泽() 	feud.Barony
}

type 泽兰ZeelandCounty struct {
	feud.BaseCounty
	博尔瑟莱Borsele 	feud.Barony
	卡德赞德Cadzand 	feud.Barony
	米德尔堡Middelburg 	feud.Barony
	雷讷瑟Renesse 	feud.Barony
	托伦Tholen 	feud.Barony
	费勒Veere 	feud.Barony
	弗利辛亨Vlissingen 	feud.Barony
	济里克泽Zierikzee 	feud.Barony
}

func (c *泽兰ZeelandCounty) BBorsele博尔瑟莱() feud.Barony {
	return c.博尔瑟莱Borsele
}
    
func (c *泽兰ZeelandCounty) BCadzand卡德赞德() feud.Barony {
	return c.卡德赞德Cadzand
}
    
func (c *泽兰ZeelandCounty) BMiddelburg米德尔堡() feud.Barony {
	return c.米德尔堡Middelburg
}
    
func (c *泽兰ZeelandCounty) BRenesse雷讷瑟() feud.Barony {
	return c.雷讷瑟Renesse
}
    
func (c *泽兰ZeelandCounty) BTholen托伦() feud.Barony {
	return c.托伦Tholen
}
    
func (c *泽兰ZeelandCounty) BVeere费勒() feud.Barony {
	return c.费勒Veere
}
    
func (c *泽兰ZeelandCounty) BVlissingen弗利辛亨() feud.Barony {
	return c.弗利辛亨Vlissingen
}
    
func (c *泽兰ZeelandCounty) BZierikzee济里克泽() feud.Barony {
	return c.济里克泽Zierikzee
}
    
var CZeeland泽兰 ZeelandCounty = &泽兰ZeelandCounty{}

func init() {
	f := CZeeland泽兰.(*泽兰ZeelandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "79",
		Title:     "zeeland",
		TitleName: "泽兰",
		TitleCode: "c_zeeland",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔瑟莱Borsele = BBorsele博尔瑟莱
	f.博尔瑟莱Borsele.SetParent(f)
	
	f.卡德赞德Cadzand = BCadzand卡德赞德
	f.卡德赞德Cadzand.SetParent(f)
	
	f.米德尔堡Middelburg = BMiddelburg米德尔堡
	f.米德尔堡Middelburg.SetParent(f)
	
	f.雷讷瑟Renesse = BRenesse雷讷瑟
	f.雷讷瑟Renesse.SetParent(f)
	
	f.托伦Tholen = BTholen托伦
	f.托伦Tholen.SetParent(f)
	
	f.费勒Veere = BVeere费勒
	f.费勒Veere.SetParent(f)
	
	f.弗利辛亨Vlissingen = BVlissingen弗利辛亨
	f.弗利辛亨Vlissingen.SetParent(f)
	
	f.济里克泽Zierikzee = BZierikzee济里克泽
	f.济里克泽Zierikzee.SetParent(f)
	
}
