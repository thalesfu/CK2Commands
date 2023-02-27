package besancon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BesanconCounty interface {
    feud.County
    BArlay阿尔莱() 	feud.Barony
    BBellevaux贝勒沃() 	feud.Barony
    BBesancon贝桑松() 	feud.Barony
    BMaiche迈什() 	feud.Barony
    BMarchaux马尔绍() 	feud.Barony
    BMorteau莫尔托() 	feud.Barony
    BPontarlier蓬塔利耶() 	feud.Barony
}

type 贝桑松BesanconCounty struct {
	feud.BaseCounty
	阿尔莱Arlay 	feud.Barony
	贝勒沃Bellevaux 	feud.Barony
	贝桑松Besancon 	feud.Barony
	迈什Maiche 	feud.Barony
	马尔绍Marchaux 	feud.Barony
	莫尔托Morteau 	feud.Barony
	蓬塔利耶Pontarlier 	feud.Barony
}

func (c *贝桑松BesanconCounty) BArlay阿尔莱() feud.Barony {
	return c.阿尔莱Arlay
}
    
func (c *贝桑松BesanconCounty) BBellevaux贝勒沃() feud.Barony {
	return c.贝勒沃Bellevaux
}
    
func (c *贝桑松BesanconCounty) BBesancon贝桑松() feud.Barony {
	return c.贝桑松Besancon
}
    
func (c *贝桑松BesanconCounty) BMaiche迈什() feud.Barony {
	return c.迈什Maiche
}
    
func (c *贝桑松BesanconCounty) BMarchaux马尔绍() feud.Barony {
	return c.马尔绍Marchaux
}
    
func (c *贝桑松BesanconCounty) BMorteau莫尔托() feud.Barony {
	return c.莫尔托Morteau
}
    
func (c *贝桑松BesanconCounty) BPontarlier蓬塔利耶() feud.Barony {
	return c.蓬塔利耶Pontarlier
}
    
var CBesancon贝桑松 BesanconCounty = &贝桑松BesanconCounty{}

func init() {
	f := CBesancon贝桑松.(*贝桑松BesanconCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "135",
		Title:     "besancon",
		TitleName: "贝桑松",
		TitleCode: "c_besancon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔莱Arlay = BArlay阿尔莱
	f.阿尔莱Arlay.SetParent(f)
	
	f.贝勒沃Bellevaux = BBellevaux贝勒沃
	f.贝勒沃Bellevaux.SetParent(f)
	
	f.贝桑松Besancon = BBesancon贝桑松
	f.贝桑松Besancon.SetParent(f)
	
	f.迈什Maiche = BMaiche迈什
	f.迈什Maiche.SetParent(f)
	
	f.马尔绍Marchaux = BMarchaux马尔绍
	f.马尔绍Marchaux.SetParent(f)
	
	f.莫尔托Morteau = BMorteau莫尔托
	f.莫尔托Morteau.SetParent(f)
	
	f.蓬塔利耶Pontarlier = BPontarlier蓬塔利耶
	f.蓬塔利耶Pontarlier.SetParent(f)
	
}
