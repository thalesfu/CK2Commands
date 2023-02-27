package kolzum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolzumCounty interface {
    feud.County
    BClysma克利斯马() 	feud.Barony
    BKolzum科勒祖姆() 	feud.Barony
    BMount_anthony安东尼山() 	feud.Barony
    BShallufa莎鲁法() 	feud.Barony
    BSokhna索赫纳() 	feud.Barony
    BSuez苏伊士() 	feud.Barony
    BTaufiq陶菲克() 	feud.Barony
}

type 科勒祖姆KolzumCounty struct {
	feud.BaseCounty
	克利斯马Clysma 	feud.Barony
	科勒祖姆Kolzum 	feud.Barony
	安东尼山Mount_anthony 	feud.Barony
	莎鲁法Shallufa 	feud.Barony
	索赫纳Sokhna 	feud.Barony
	苏伊士Suez 	feud.Barony
	陶菲克Taufiq 	feud.Barony
}

func (c *科勒祖姆KolzumCounty) BClysma克利斯马() feud.Barony {
	return c.克利斯马Clysma
}
    
func (c *科勒祖姆KolzumCounty) BKolzum科勒祖姆() feud.Barony {
	return c.科勒祖姆Kolzum
}
    
func (c *科勒祖姆KolzumCounty) BMount_anthony安东尼山() feud.Barony {
	return c.安东尼山Mount_anthony
}
    
func (c *科勒祖姆KolzumCounty) BShallufa莎鲁法() feud.Barony {
	return c.莎鲁法Shallufa
}
    
func (c *科勒祖姆KolzumCounty) BSokhna索赫纳() feud.Barony {
	return c.索赫纳Sokhna
}
    
func (c *科勒祖姆KolzumCounty) BSuez苏伊士() feud.Barony {
	return c.苏伊士Suez
}
    
func (c *科勒祖姆KolzumCounty) BTaufiq陶菲克() feud.Barony {
	return c.陶菲克Taufiq
}
    
var CKolzum科勒祖姆 KolzumCounty = &科勒祖姆KolzumCounty{}

func init() {
	f := CKolzum科勒祖姆.(*科勒祖姆KolzumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2006",
		Title:     "kolzum",
		TitleName: "科勒祖姆",
		TitleCode: "c_kolzum",
		Baronies:  map[string]feud.Barony{},
	}

	f.克利斯马Clysma = BClysma克利斯马
	f.克利斯马Clysma.SetParent(f)
	
	f.科勒祖姆Kolzum = BKolzum科勒祖姆
	f.科勒祖姆Kolzum.SetParent(f)
	
	f.安东尼山Mount_anthony = BMount_anthony安东尼山
	f.安东尼山Mount_anthony.SetParent(f)
	
	f.莎鲁法Shallufa = BShallufa莎鲁法
	f.莎鲁法Shallufa.SetParent(f)
	
	f.索赫纳Sokhna = BSokhna索赫纳
	f.索赫纳Sokhna.SetParent(f)
	
	f.苏伊士Suez = BSuez苏伊士
	f.苏伊士Suez.SetParent(f)
	
	f.陶菲克Taufiq = BTaufiq陶菲克
	f.陶菲克Taufiq.SetParent(f)
	
}
