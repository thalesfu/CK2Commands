package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MetzCounty interface {
    feud.County
    BAudunleroman欧丹勒罗芒() 	feud.Barony
    BBouzonville布宗维勒() 	feud.Barony
    BBriey布里埃() 	feud.Barony
    BJoudreville茹德勒维尔() 	feud.Barony
    BMarslatour马尔拉图() 	feud.Barony
    BMetz梅斯() 	feud.Barony
    BSaintjulien圣于连() 	feud.Barony
    BThionville蒂永维勒() 	feud.Barony
}

type 梅斯MetzCounty struct {
	feud.BaseCounty
	欧丹勒罗芒Audunleroman 	feud.Barony
	布宗维勒Bouzonville 	feud.Barony
	布里埃Briey 	feud.Barony
	茹德勒维尔Joudreville 	feud.Barony
	马尔拉图Marslatour 	feud.Barony
	梅斯Metz 	feud.Barony
	圣于连Saintjulien 	feud.Barony
	蒂永维勒Thionville 	feud.Barony
}

func (c *梅斯MetzCounty) BAudunleroman欧丹勒罗芒() feud.Barony {
	return c.欧丹勒罗芒Audunleroman
}
    
func (c *梅斯MetzCounty) BBouzonville布宗维勒() feud.Barony {
	return c.布宗维勒Bouzonville
}
    
func (c *梅斯MetzCounty) BBriey布里埃() feud.Barony {
	return c.布里埃Briey
}
    
func (c *梅斯MetzCounty) BJoudreville茹德勒维尔() feud.Barony {
	return c.茹德勒维尔Joudreville
}
    
func (c *梅斯MetzCounty) BMarslatour马尔拉图() feud.Barony {
	return c.马尔拉图Marslatour
}
    
func (c *梅斯MetzCounty) BMetz梅斯() feud.Barony {
	return c.梅斯Metz
}
    
func (c *梅斯MetzCounty) BSaintjulien圣于连() feud.Barony {
	return c.圣于连Saintjulien
}
    
func (c *梅斯MetzCounty) BThionville蒂永维勒() feud.Barony {
	return c.蒂永维勒Thionville
}
    
var CMetz梅斯 MetzCounty = &梅斯MetzCounty{}

func init() {
	f := CMetz梅斯.(*梅斯MetzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "128",
		Title:     "metz",
		TitleName: "梅斯",
		TitleCode: "c_metz",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧丹勒罗芒Audunleroman = BAudunleroman欧丹勒罗芒
	f.欧丹勒罗芒Audunleroman.SetParent(f)
	
	f.布宗维勒Bouzonville = BBouzonville布宗维勒
	f.布宗维勒Bouzonville.SetParent(f)
	
	f.布里埃Briey = BBriey布里埃
	f.布里埃Briey.SetParent(f)
	
	f.茹德勒维尔Joudreville = BJoudreville茹德勒维尔
	f.茹德勒维尔Joudreville.SetParent(f)
	
	f.马尔拉图Marslatour = BMarslatour马尔拉图
	f.马尔拉图Marslatour.SetParent(f)
	
	f.梅斯Metz = BMetz梅斯
	f.梅斯Metz.SetParent(f)
	
	f.圣于连Saintjulien = BSaintjulien圣于连
	f.圣于连Saintjulien.SetParent(f)
	
	f.蒂永维勒Thionville = BThionville蒂永维勒
	f.蒂永维勒Thionville.SetParent(f)
	
}
