package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrielCounty interface {
    feud.County
    BArdee阿迪() 	feud.Barony
    BArmagh阿马() 	feud.Barony
    BClogher克洛赫() 	feud.Barony
    BClones克朗斯() 	feud.Barony
    BDrogheda德罗赫达() 	feud.Barony
    BDundalk邓多克() 	feud.Barony
    BMonaghan莫纳亨() 	feud.Barony
    BOlouth劳斯() 	feud.Barony
}

type 奥里尔OrielCounty struct {
	feud.BaseCounty
	阿迪Ardee 	feud.Barony
	阿马Armagh 	feud.Barony
	克洛赫Clogher 	feud.Barony
	克朗斯Clones 	feud.Barony
	德罗赫达Drogheda 	feud.Barony
	邓多克Dundalk 	feud.Barony
	莫纳亨Monaghan 	feud.Barony
	劳斯Olouth 	feud.Barony
}

func (c *奥里尔OrielCounty) BArdee阿迪() feud.Barony {
	return c.阿迪Ardee
}
    
func (c *奥里尔OrielCounty) BArmagh阿马() feud.Barony {
	return c.阿马Armagh
}
    
func (c *奥里尔OrielCounty) BClogher克洛赫() feud.Barony {
	return c.克洛赫Clogher
}
    
func (c *奥里尔OrielCounty) BClones克朗斯() feud.Barony {
	return c.克朗斯Clones
}
    
func (c *奥里尔OrielCounty) BDrogheda德罗赫达() feud.Barony {
	return c.德罗赫达Drogheda
}
    
func (c *奥里尔OrielCounty) BDundalk邓多克() feud.Barony {
	return c.邓多克Dundalk
}
    
func (c *奥里尔OrielCounty) BMonaghan莫纳亨() feud.Barony {
	return c.莫纳亨Monaghan
}
    
func (c *奥里尔OrielCounty) BOlouth劳斯() feud.Barony {
	return c.劳斯Olouth
}
    
var COriel奥里尔 OrielCounty = &奥里尔OrielCounty{}

func init() {
	f := COriel奥里尔.(*奥里尔OrielCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "7",
		Title:     "oriel",
		TitleName: "奥里尔",
		TitleCode: "c_oriel",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿迪Ardee = BArdee阿迪
	f.阿迪Ardee.SetParent(f)
	
	f.阿马Armagh = BArmagh阿马
	f.阿马Armagh.SetParent(f)
	
	f.克洛赫Clogher = BClogher克洛赫
	f.克洛赫Clogher.SetParent(f)
	
	f.克朗斯Clones = BClones克朗斯
	f.克朗斯Clones.SetParent(f)
	
	f.德罗赫达Drogheda = BDrogheda德罗赫达
	f.德罗赫达Drogheda.SetParent(f)
	
	f.邓多克Dundalk = BDundalk邓多克
	f.邓多克Dundalk.SetParent(f)
	
	f.莫纳亨Monaghan = BMonaghan莫纳亨
	f.莫纳亨Monaghan.SetParent(f)
	
	f.劳斯Olouth = BOlouth劳斯
	f.劳斯Olouth.SetParent(f)
	
}
