package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmorionCounty interface {
    feud.County
    BAmorion阿摩里翁() 	feud.Barony
    BDocimium_amorion多基米翁() 	feud.Barony
    BKrasos克拉索斯() 	feud.Barony
    BPhilomelion菲洛梅利翁() 	feud.Barony
    BPolybotos波利波图斯() 	feud.Barony
    BSalutaris萨卢塔里斯() 	feud.Barony
    BSynnada希南达() 	feud.Barony
}

type 阿摩里翁AmorionCounty struct {
	feud.BaseCounty
	阿摩里翁Amorion 	feud.Barony
	多基米翁Docimium_amorion 	feud.Barony
	克拉索斯Krasos 	feud.Barony
	菲洛梅利翁Philomelion 	feud.Barony
	波利波图斯Polybotos 	feud.Barony
	萨卢塔里斯Salutaris 	feud.Barony
	希南达Synnada 	feud.Barony
}

func (c *阿摩里翁AmorionCounty) BAmorion阿摩里翁() feud.Barony {
	return c.阿摩里翁Amorion
}
    
func (c *阿摩里翁AmorionCounty) BDocimium_amorion多基米翁() feud.Barony {
	return c.多基米翁Docimium_amorion
}
    
func (c *阿摩里翁AmorionCounty) BKrasos克拉索斯() feud.Barony {
	return c.克拉索斯Krasos
}
    
func (c *阿摩里翁AmorionCounty) BPhilomelion菲洛梅利翁() feud.Barony {
	return c.菲洛梅利翁Philomelion
}
    
func (c *阿摩里翁AmorionCounty) BPolybotos波利波图斯() feud.Barony {
	return c.波利波图斯Polybotos
}
    
func (c *阿摩里翁AmorionCounty) BSalutaris萨卢塔里斯() feud.Barony {
	return c.萨卢塔里斯Salutaris
}
    
func (c *阿摩里翁AmorionCounty) BSynnada希南达() feud.Barony {
	return c.希南达Synnada
}
    
var CAmorion阿摩里翁 AmorionCounty = &阿摩里翁AmorionCounty{}

func init() {
	f := CAmorion阿摩里翁.(*阿摩里翁AmorionCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1931",
		Title:     "amorion",
		TitleName: "阿摩里翁",
		TitleCode: "c_amorion",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿摩里翁Amorion = BAmorion阿摩里翁
	f.阿摩里翁Amorion.SetParent(f)
	
	f.多基米翁Docimium_amorion = BDocimium_amorion多基米翁
	f.多基米翁Docimium_amorion.SetParent(f)
	
	f.克拉索斯Krasos = BKrasos克拉索斯
	f.克拉索斯Krasos.SetParent(f)
	
	f.菲洛梅利翁Philomelion = BPhilomelion菲洛梅利翁
	f.菲洛梅利翁Philomelion.SetParent(f)
	
	f.波利波图斯Polybotos = BPolybotos波利波图斯
	f.波利波图斯Polybotos.SetParent(f)
	
	f.萨卢塔里斯Salutaris = BSalutaris萨卢塔里斯
	f.萨卢塔里斯Salutaris.SetParent(f)
	
	f.希南达Synnada = BSynnada希南达
	f.希南达Synnada.SetParent(f)
	
}
