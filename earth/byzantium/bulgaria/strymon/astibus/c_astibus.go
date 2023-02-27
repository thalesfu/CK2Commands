package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AstibusCounty interface {
    feud.County
    BAstibus阿斯提博斯() 	feud.Barony
    BDukena杜凯纳() 	feud.Barony
    BPresevo普雷舍沃() 	feud.Barony
    BSkopje斯科普里() 	feud.Barony
    BStobi斯托比() 	feud.Barony
    BVeles韦莱斯() 	feud.Barony
    BZegligovo热格利戈沃() 	feud.Barony
}

type 阿斯提博斯AstibusCounty struct {
	feud.BaseCounty
	阿斯提博斯Astibus 	feud.Barony
	杜凯纳Dukena 	feud.Barony
	普雷舍沃Presevo 	feud.Barony
	斯科普里Skopje 	feud.Barony
	斯托比Stobi 	feud.Barony
	韦莱斯Veles 	feud.Barony
	热格利戈沃Zegligovo 	feud.Barony
}

func (c *阿斯提博斯AstibusCounty) BAstibus阿斯提博斯() feud.Barony {
	return c.阿斯提博斯Astibus
}
    
func (c *阿斯提博斯AstibusCounty) BDukena杜凯纳() feud.Barony {
	return c.杜凯纳Dukena
}
    
func (c *阿斯提博斯AstibusCounty) BPresevo普雷舍沃() feud.Barony {
	return c.普雷舍沃Presevo
}
    
func (c *阿斯提博斯AstibusCounty) BSkopje斯科普里() feud.Barony {
	return c.斯科普里Skopje
}
    
func (c *阿斯提博斯AstibusCounty) BStobi斯托比() feud.Barony {
	return c.斯托比Stobi
}
    
func (c *阿斯提博斯AstibusCounty) BVeles韦莱斯() feud.Barony {
	return c.韦莱斯Veles
}
    
func (c *阿斯提博斯AstibusCounty) BZegligovo热格利戈沃() feud.Barony {
	return c.热格利戈沃Zegligovo
}
    
var CAstibus阿斯提博斯 AstibusCounty = &阿斯提博斯AstibusCounty{}

func init() {
	f := CAstibus阿斯提博斯.(*阿斯提博斯AstibusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1888",
		Title:     "astibus",
		TitleName: "阿斯提博斯",
		TitleCode: "c_astibus",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯提博斯Astibus = BAstibus阿斯提博斯
	f.阿斯提博斯Astibus.SetParent(f)
	
	f.杜凯纳Dukena = BDukena杜凯纳
	f.杜凯纳Dukena.SetParent(f)
	
	f.普雷舍沃Presevo = BPresevo普雷舍沃
	f.普雷舍沃Presevo.SetParent(f)
	
	f.斯科普里Skopje = BSkopje斯科普里
	f.斯科普里Skopje.SetParent(f)
	
	f.斯托比Stobi = BStobi斯托比
	f.斯托比Stobi.SetParent(f)
	
	f.韦莱斯Veles = BVeles韦莱斯
	f.韦莱斯Veles.SetParent(f)
	
	f.热格利戈沃Zegligovo = BZegligovo热格利戈沃
	f.热格利戈沃Zegligovo.SetParent(f)
	
}
