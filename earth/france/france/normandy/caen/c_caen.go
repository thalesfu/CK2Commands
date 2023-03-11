package caen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CaenCounty interface {
    feud.County
    BBayeux巴约() 	feud.Barony
    BCaen卡昂() 	feud.Barony
    BClecy克莱西() 	feud.Barony
    BDomfront栋夫龙() 	feud.Barony
    BJurques瑞尔克() 	feud.Barony
    BVillers_bocage维莱博卡日() 	feud.Barony
    BVire维尔() 	feud.Barony
}

type 卡昂CaenCounty struct {
	feud.BaseCounty
	巴约Bayeux 	feud.Barony
	卡昂Caen 	feud.Barony
	克莱西Clecy 	feud.Barony
	栋夫龙Domfront 	feud.Barony
	瑞尔克Jurques 	feud.Barony
	维莱博卡日Villers_bocage 	feud.Barony
	维尔Vire 	feud.Barony
}

func (c *卡昂CaenCounty) BBayeux巴约() feud.Barony {
	return c.巴约Bayeux
}
    
func (c *卡昂CaenCounty) BCaen卡昂() feud.Barony {
	return c.卡昂Caen
}
    
func (c *卡昂CaenCounty) BClecy克莱西() feud.Barony {
	return c.克莱西Clecy
}
    
func (c *卡昂CaenCounty) BDomfront栋夫龙() feud.Barony {
	return c.栋夫龙Domfront
}
    
func (c *卡昂CaenCounty) BJurques瑞尔克() feud.Barony {
	return c.瑞尔克Jurques
}
    
func (c *卡昂CaenCounty) BVillers_bocage维莱博卡日() feud.Barony {
	return c.维莱博卡日Villers_bocage
}
    
func (c *卡昂CaenCounty) BVire维尔() feud.Barony {
	return c.维尔Vire
}
    
var CCaen卡昂 CaenCounty = &卡昂CaenCounty{}

func init() {
	f := CCaen卡昂.(*卡昂CaenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1956",
		Title:     "caen",
		TitleName: "卡昂",
		TitleCode: "c_caen",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴约Bayeux = BBayeux巴约
	f.巴约Bayeux.SetParent(f)
	
	f.卡昂Caen = BCaen卡昂
	f.卡昂Caen.SetParent(f)
	
	f.克莱西Clecy = BClecy克莱西
	f.克莱西Clecy.SetParent(f)
	
	f.栋夫龙Domfront = BDomfront栋夫龙
	f.栋夫龙Domfront.SetParent(f)
	
	f.瑞尔克Jurques = BJurques瑞尔克
	f.瑞尔克Jurques.SetParent(f)
	
	f.维莱博卡日Villers_bocage = BVillers_bocage维莱博卡日
	f.维莱博卡日Villers_bocage.SetParent(f)
	
	f.维尔Vire = BVire维尔
	f.维尔Vire.SetParent(f)
	
}
