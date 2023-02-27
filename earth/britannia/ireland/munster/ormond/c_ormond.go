package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrmondCounty interface {
    feud.County
    BCahir凯尔() 	feud.Barony
    BCashel卡舍尔() 	feud.Barony
    BClonmel克朗梅尔() 	feud.Barony
    BFethard费萨德() 	feud.Barony
    BLismore利斯莫尔() 	feud.Barony
    BNenagh尼纳() 	feud.Barony
    BRoscrea罗斯克雷() 	feud.Barony
    BWaterford沃特福德() 	feud.Barony
}

type 奥蒙德OrmondCounty struct {
	feud.BaseCounty
	凯尔Cahir 	feud.Barony
	卡舍尔Cashel 	feud.Barony
	克朗梅尔Clonmel 	feud.Barony
	费萨德Fethard 	feud.Barony
	利斯莫尔Lismore 	feud.Barony
	尼纳Nenagh 	feud.Barony
	罗斯克雷Roscrea 	feud.Barony
	沃特福德Waterford 	feud.Barony
}

func (c *奥蒙德OrmondCounty) BCahir凯尔() feud.Barony {
	return c.凯尔Cahir
}
    
func (c *奥蒙德OrmondCounty) BCashel卡舍尔() feud.Barony {
	return c.卡舍尔Cashel
}
    
func (c *奥蒙德OrmondCounty) BClonmel克朗梅尔() feud.Barony {
	return c.克朗梅尔Clonmel
}
    
func (c *奥蒙德OrmondCounty) BFethard费萨德() feud.Barony {
	return c.费萨德Fethard
}
    
func (c *奥蒙德OrmondCounty) BLismore利斯莫尔() feud.Barony {
	return c.利斯莫尔Lismore
}
    
func (c *奥蒙德OrmondCounty) BNenagh尼纳() feud.Barony {
	return c.尼纳Nenagh
}
    
func (c *奥蒙德OrmondCounty) BRoscrea罗斯克雷() feud.Barony {
	return c.罗斯克雷Roscrea
}
    
func (c *奥蒙德OrmondCounty) BWaterford沃特福德() feud.Barony {
	return c.沃特福德Waterford
}
    
var COrmond奥蒙德 OrmondCounty = &奥蒙德OrmondCounty{}

func init() {
	f := COrmond奥蒙德.(*奥蒙德OrmondCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "15",
		Title:     "ormond",
		TitleName: "奥蒙德",
		TitleCode: "c_ormond",
		Baronies:  map[string]feud.Barony{},
	}

	f.凯尔Cahir = BCahir凯尔
	f.凯尔Cahir.SetParent(f)
	
	f.卡舍尔Cashel = BCashel卡舍尔
	f.卡舍尔Cashel.SetParent(f)
	
	f.克朗梅尔Clonmel = BClonmel克朗梅尔
	f.克朗梅尔Clonmel.SetParent(f)
	
	f.费萨德Fethard = BFethard费萨德
	f.费萨德Fethard.SetParent(f)
	
	f.利斯莫尔Lismore = BLismore利斯莫尔
	f.利斯莫尔Lismore.SetParent(f)
	
	f.尼纳Nenagh = BNenagh尼纳
	f.尼纳Nenagh.SetParent(f)
	
	f.罗斯克雷Roscrea = BRoscrea罗斯克雷
	f.罗斯克雷Roscrea.SetParent(f)
	
	f.沃特福德Waterford = BWaterford沃特福德
	f.沃特福德Waterford.SetParent(f)
	
}
