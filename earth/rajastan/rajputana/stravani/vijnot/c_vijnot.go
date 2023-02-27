package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VijnotCounty interface {
    feud.County
    BBhatiya般提耶() 	feud.Barony
    BGhotki科德吉() 	feud.Barony
    BKishangarh计霜姞利呬() 	feud.Barony
    BRavel罗婆() 	feud.Barony
    BSaitual室利兜阿() 	feud.Barony
    BSakyang萨姜() 	feud.Barony
    BVijnot毗者努多() 	feud.Barony
}

type 毗者努多VijnotCounty struct {
	feud.BaseCounty
	般提耶Bhatiya 	feud.Barony
	科德吉Ghotki 	feud.Barony
	计霜姞利呬Kishangarh 	feud.Barony
	罗婆Ravel 	feud.Barony
	室利兜阿Saitual 	feud.Barony
	萨姜Sakyang 	feud.Barony
	毗者努多Vijnot 	feud.Barony
}

func (c *毗者努多VijnotCounty) BBhatiya般提耶() feud.Barony {
	return c.般提耶Bhatiya
}
    
func (c *毗者努多VijnotCounty) BGhotki科德吉() feud.Barony {
	return c.科德吉Ghotki
}
    
func (c *毗者努多VijnotCounty) BKishangarh计霜姞利呬() feud.Barony {
	return c.计霜姞利呬Kishangarh
}
    
func (c *毗者努多VijnotCounty) BRavel罗婆() feud.Barony {
	return c.罗婆Ravel
}
    
func (c *毗者努多VijnotCounty) BSaitual室利兜阿() feud.Barony {
	return c.室利兜阿Saitual
}
    
func (c *毗者努多VijnotCounty) BSakyang萨姜() feud.Barony {
	return c.萨姜Sakyang
}
    
func (c *毗者努多VijnotCounty) BVijnot毗者努多() feud.Barony {
	return c.毗者努多Vijnot
}
    
var CVijnot毗者努多 VijnotCounty = &毗者努多VijnotCounty{}

func init() {
	f := CVijnot毗者努多.(*毗者努多VijnotCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1336",
		Title:     "vijnot",
		TitleName: "毗者努多",
		TitleCode: "c_vijnot",
		Baronies:  map[string]feud.Barony{},
	}

	f.般提耶Bhatiya = BBhatiya般提耶
	f.般提耶Bhatiya.SetParent(f)
	
	f.科德吉Ghotki = BGhotki科德吉
	f.科德吉Ghotki.SetParent(f)
	
	f.计霜姞利呬Kishangarh = BKishangarh计霜姞利呬
	f.计霜姞利呬Kishangarh.SetParent(f)
	
	f.罗婆Ravel = BRavel罗婆
	f.罗婆Ravel.SetParent(f)
	
	f.室利兜阿Saitual = BSaitual室利兜阿
	f.室利兜阿Saitual.SetParent(f)
	
	f.萨姜Sakyang = BSakyang萨姜
	f.萨姜Sakyang.SetParent(f)
	
	f.毗者努多Vijnot = BVijnot毗者努多
	f.毗者努多Vijnot.SetParent(f)
	
}
