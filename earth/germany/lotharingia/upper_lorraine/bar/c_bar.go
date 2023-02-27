package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarCounty interface {
    feud.County
    BBarleduc巴勒迪克() 	feud.Barony
    BCommercy科梅尔西() 	feud.Barony
    BJoinville茹安维尔() 	feud.Barony
    BSaintdizier圣迪济耶() 	feud.Barony
    BSaintmartinengourgandelle古尔冈德勒的圣马尔坦() 	feud.Barony
    BSaintmihel圣米耶勒() 	feud.Barony
    BVaucouleurs沃库勒尔() 	feud.Barony
}

type 巴鲁瓦BarCounty struct {
	feud.BaseCounty
	巴勒迪克Barleduc 	feud.Barony
	科梅尔西Commercy 	feud.Barony
	茹安维尔Joinville 	feud.Barony
	圣迪济耶Saintdizier 	feud.Barony
	古尔冈德勒的圣马尔坦Saintmartinengourgandelle 	feud.Barony
	圣米耶勒Saintmihel 	feud.Barony
	沃库勒尔Vaucouleurs 	feud.Barony
}

func (c *巴鲁瓦BarCounty) BBarleduc巴勒迪克() feud.Barony {
	return c.巴勒迪克Barleduc
}
    
func (c *巴鲁瓦BarCounty) BCommercy科梅尔西() feud.Barony {
	return c.科梅尔西Commercy
}
    
func (c *巴鲁瓦BarCounty) BJoinville茹安维尔() feud.Barony {
	return c.茹安维尔Joinville
}
    
func (c *巴鲁瓦BarCounty) BSaintdizier圣迪济耶() feud.Barony {
	return c.圣迪济耶Saintdizier
}
    
func (c *巴鲁瓦BarCounty) BSaintmartinengourgandelle古尔冈德勒的圣马尔坦() feud.Barony {
	return c.古尔冈德勒的圣马尔坦Saintmartinengourgandelle
}
    
func (c *巴鲁瓦BarCounty) BSaintmihel圣米耶勒() feud.Barony {
	return c.圣米耶勒Saintmihel
}
    
func (c *巴鲁瓦BarCounty) BVaucouleurs沃库勒尔() feud.Barony {
	return c.沃库勒尔Vaucouleurs
}
    
var CBar巴鲁瓦 BarCounty = &巴鲁瓦BarCounty{}

func init() {
	f := CBar巴鲁瓦.(*巴鲁瓦BarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "930",
		Title:     "bar",
		TitleName: "巴鲁瓦",
		TitleCode: "c_bar",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴勒迪克Barleduc = BBarleduc巴勒迪克
	f.巴勒迪克Barleduc.SetParent(f)
	
	f.科梅尔西Commercy = BCommercy科梅尔西
	f.科梅尔西Commercy.SetParent(f)
	
	f.茹安维尔Joinville = BJoinville茹安维尔
	f.茹安维尔Joinville.SetParent(f)
	
	f.圣迪济耶Saintdizier = BSaintdizier圣迪济耶
	f.圣迪济耶Saintdizier.SetParent(f)
	
	f.古尔冈德勒的圣马尔坦Saintmartinengourgandelle = BSaintmartinengourgandelle古尔冈德勒的圣马尔坦
	f.古尔冈德勒的圣马尔坦Saintmartinengourgandelle.SetParent(f)
	
	f.圣米耶勒Saintmihel = BSaintmihel圣米耶勒
	f.圣米耶勒Saintmihel.SetParent(f)
	
	f.沃库勒尔Vaucouleurs = BVaucouleurs沃库勒尔
	f.沃库勒尔Vaucouleurs.SetParent(f)
	
}
