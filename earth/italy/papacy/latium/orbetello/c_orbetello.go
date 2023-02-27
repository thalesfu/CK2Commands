package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrbetelloCounty interface {
    feud.County
    BGrosseto格罗塞托() 	feud.Barony
    BOrbetello奥尔贝泰洛() 	feud.Barony
    BPitigliano皮蒂利亚诺() 	feud.Barony
    BRoselle罗塞莱() 	feud.Barony
    BRusellae鲁塞莱() 	feud.Barony
    BSorano索拉诺() 	feud.Barony
    BSovana索瓦纳() 	feud.Barony
    BVetulonia韦图罗尼亚() 	feud.Barony
}

type 奥尔贝泰洛OrbetelloCounty struct {
	feud.BaseCounty
	格罗塞托Grosseto 	feud.Barony
	奥尔贝泰洛Orbetello 	feud.Barony
	皮蒂利亚诺Pitigliano 	feud.Barony
	罗塞莱Roselle 	feud.Barony
	鲁塞莱Rusellae 	feud.Barony
	索拉诺Sorano 	feud.Barony
	索瓦纳Sovana 	feud.Barony
	韦图罗尼亚Vetulonia 	feud.Barony
}

func (c *奥尔贝泰洛OrbetelloCounty) BGrosseto格罗塞托() feud.Barony {
	return c.格罗塞托Grosseto
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BOrbetello奥尔贝泰洛() feud.Barony {
	return c.奥尔贝泰洛Orbetello
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BPitigliano皮蒂利亚诺() feud.Barony {
	return c.皮蒂利亚诺Pitigliano
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BRoselle罗塞莱() feud.Barony {
	return c.罗塞莱Roselle
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BRusellae鲁塞莱() feud.Barony {
	return c.鲁塞莱Rusellae
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BSorano索拉诺() feud.Barony {
	return c.索拉诺Sorano
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BSovana索瓦纳() feud.Barony {
	return c.索瓦纳Sovana
}
    
func (c *奥尔贝泰洛OrbetelloCounty) BVetulonia韦图罗尼亚() feud.Barony {
	return c.韦图罗尼亚Vetulonia
}
    
var COrbetello奥尔贝泰洛 OrbetelloCounty = &奥尔贝泰洛OrbetelloCounty{}

func init() {
	f := COrbetello奥尔贝泰洛.(*奥尔贝泰洛OrbetelloCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "332",
		Title:     "orbetello",
		TitleName: "奥尔贝泰洛",
		TitleCode: "c_orbetello",
		Baronies:  map[string]feud.Barony{},
	}

	f.格罗塞托Grosseto = BGrosseto格罗塞托
	f.格罗塞托Grosseto.SetParent(f)
	
	f.奥尔贝泰洛Orbetello = BOrbetello奥尔贝泰洛
	f.奥尔贝泰洛Orbetello.SetParent(f)
	
	f.皮蒂利亚诺Pitigliano = BPitigliano皮蒂利亚诺
	f.皮蒂利亚诺Pitigliano.SetParent(f)
	
	f.罗塞莱Roselle = BRoselle罗塞莱
	f.罗塞莱Roselle.SetParent(f)
	
	f.鲁塞莱Rusellae = BRusellae鲁塞莱
	f.鲁塞莱Rusellae.SetParent(f)
	
	f.索拉诺Sorano = BSorano索拉诺
	f.索拉诺Sorano.SetParent(f)
	
	f.索瓦纳Sovana = BSovana索瓦纳
	f.索瓦纳Sovana.SetParent(f)
	
	f.韦图罗尼亚Vetulonia = BVetulonia韦图罗尼亚
	f.韦图罗尼亚Vetulonia.SetParent(f)
	
}
