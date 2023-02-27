package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OpplandCounty interface {
    feud.County
    BBirid比里德() 	feud.Barony
    BFavang福旺() 	feud.Barony
    BFlesberg弗莱斯贝格() 	feud.Barony
    BLillehammer利勒哈默尔() 	feud.Barony
    BOyer厄耶() 	feud.Barony
    BRedalr雷达尔() 	feud.Barony
    BSlidre斯利德勒() 	feud.Barony
}

type 海德马克OpplandCounty struct {
	feud.BaseCounty
	比里德Birid 	feud.Barony
	福旺Favang 	feud.Barony
	弗莱斯贝格Flesberg 	feud.Barony
	利勒哈默尔Lillehammer 	feud.Barony
	厄耶Oyer 	feud.Barony
	雷达尔Redalr 	feud.Barony
	斯利德勒Slidre 	feud.Barony
}

func (c *海德马克OpplandCounty) BBirid比里德() feud.Barony {
	return c.比里德Birid
}
    
func (c *海德马克OpplandCounty) BFavang福旺() feud.Barony {
	return c.福旺Favang
}
    
func (c *海德马克OpplandCounty) BFlesberg弗莱斯贝格() feud.Barony {
	return c.弗莱斯贝格Flesberg
}
    
func (c *海德马克OpplandCounty) BLillehammer利勒哈默尔() feud.Barony {
	return c.利勒哈默尔Lillehammer
}
    
func (c *海德马克OpplandCounty) BOyer厄耶() feud.Barony {
	return c.厄耶Oyer
}
    
func (c *海德马克OpplandCounty) BRedalr雷达尔() feud.Barony {
	return c.雷达尔Redalr
}
    
func (c *海德马克OpplandCounty) BSlidre斯利德勒() feud.Barony {
	return c.斯利德勒Slidre
}
    
var COppland海德马克 OpplandCounty = &海德马克OpplandCounty{}

func init() {
	f := COppland海德马克.(*海德马克OpplandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "273",
		Title:     "oppland",
		TitleName: "海德马克",
		TitleCode: "c_oppland",
		Baronies:  map[string]feud.Barony{},
	}

	f.比里德Birid = BBirid比里德
	f.比里德Birid.SetParent(f)
	
	f.福旺Favang = BFavang福旺
	f.福旺Favang.SetParent(f)
	
	f.弗莱斯贝格Flesberg = BFlesberg弗莱斯贝格
	f.弗莱斯贝格Flesberg.SetParent(f)
	
	f.利勒哈默尔Lillehammer = BLillehammer利勒哈默尔
	f.利勒哈默尔Lillehammer.SetParent(f)
	
	f.厄耶Oyer = BOyer厄耶
	f.厄耶Oyer.SetParent(f)
	
	f.雷达尔Redalr = BRedalr雷达尔
	f.雷达尔Redalr.SetParent(f)
	
	f.斯利德勒Slidre = BSlidre斯利德勒
	f.斯利德勒Slidre.SetParent(f)
	
}
