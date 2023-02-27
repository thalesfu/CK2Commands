package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NajeraCounty interface {
    feud.County
    BAlfara阿尔法拉() 	feud.Barony
    BArnedo阿尔内多() 	feud.Barony
    BCalahorra卡拉奥拉() 	feud.Barony
    BHaro阿罗() 	feud.Barony
    BLogrone洛格罗尼奥() 	feud.Barony
    BNajera纳赫拉() 	feud.Barony
    BSantodomingodelacalzada圣多明各德拉卡尔萨达() 	feud.Barony
    BZizurmayor兹祖拉梅拉() 	feud.Barony
}

type 纳赫拉NajeraCounty struct {
	feud.BaseCounty
	阿尔法拉Alfara 	feud.Barony
	阿尔内多Arnedo 	feud.Barony
	卡拉奥拉Calahorra 	feud.Barony
	阿罗Haro 	feud.Barony
	洛格罗尼奥Logrone 	feud.Barony
	纳赫拉Najera 	feud.Barony
	圣多明各德拉卡尔萨达Santodomingodelacalzada 	feud.Barony
	兹祖拉梅拉Zizurmayor 	feud.Barony
}

func (c *纳赫拉NajeraCounty) BAlfara阿尔法拉() feud.Barony {
	return c.阿尔法拉Alfara
}
    
func (c *纳赫拉NajeraCounty) BArnedo阿尔内多() feud.Barony {
	return c.阿尔内多Arnedo
}
    
func (c *纳赫拉NajeraCounty) BCalahorra卡拉奥拉() feud.Barony {
	return c.卡拉奥拉Calahorra
}
    
func (c *纳赫拉NajeraCounty) BHaro阿罗() feud.Barony {
	return c.阿罗Haro
}
    
func (c *纳赫拉NajeraCounty) BLogrone洛格罗尼奥() feud.Barony {
	return c.洛格罗尼奥Logrone
}
    
func (c *纳赫拉NajeraCounty) BNajera纳赫拉() feud.Barony {
	return c.纳赫拉Najera
}
    
func (c *纳赫拉NajeraCounty) BSantodomingodelacalzada圣多明各德拉卡尔萨达() feud.Barony {
	return c.圣多明各德拉卡尔萨达Santodomingodelacalzada
}
    
func (c *纳赫拉NajeraCounty) BZizurmayor兹祖拉梅拉() feud.Barony {
	return c.兹祖拉梅拉Zizurmayor
}
    
var CNajera纳赫拉 NajeraCounty = &纳赫拉NajeraCounty{}

func init() {
	f := CNajera纳赫拉.(*纳赫拉NajeraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "201",
		Title:     "najera",
		TitleName: "纳赫拉",
		TitleCode: "c_najera",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔法拉Alfara = BAlfara阿尔法拉
	f.阿尔法拉Alfara.SetParent(f)
	
	f.阿尔内多Arnedo = BArnedo阿尔内多
	f.阿尔内多Arnedo.SetParent(f)
	
	f.卡拉奥拉Calahorra = BCalahorra卡拉奥拉
	f.卡拉奥拉Calahorra.SetParent(f)
	
	f.阿罗Haro = BHaro阿罗
	f.阿罗Haro.SetParent(f)
	
	f.洛格罗尼奥Logrone = BLogrone洛格罗尼奥
	f.洛格罗尼奥Logrone.SetParent(f)
	
	f.纳赫拉Najera = BNajera纳赫拉
	f.纳赫拉Najera.SetParent(f)
	
	f.圣多明各德拉卡尔萨达Santodomingodelacalzada = BSantodomingodelacalzada圣多明各德拉卡尔萨达
	f.圣多明各德拉卡尔萨达Santodomingodelacalzada.SetParent(f)
	
	f.兹祖拉梅拉Zizurmayor = BZizurmayor兹祖拉梅拉
	f.兹祖拉梅拉Zizurmayor.SetParent(f)
	
}
