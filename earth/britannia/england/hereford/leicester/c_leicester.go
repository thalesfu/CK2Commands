package leicester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeicesterCounty interface {
    feud.County
    BHucknall哈克纳尔() 	feud.Barony
    BLeicester莱斯特() 	feud.Barony
    BNewark纽瓦克() 	feud.Barony
    BNewstead纽斯特德() 	feud.Barony
    BNottingham诺丁汉() 	feud.Barony
    BSouthwell索斯韦尔() 	feud.Barony
    BTickhill蒂克希尔() 	feud.Barony
    BWorksop沃克索普() 	feud.Barony
}

type 莱斯特LeicesterCounty struct {
	feud.BaseCounty
	哈克纳尔Hucknall 	feud.Barony
	莱斯特Leicester 	feud.Barony
	纽瓦克Newark 	feud.Barony
	纽斯特德Newstead 	feud.Barony
	诺丁汉Nottingham 	feud.Barony
	索斯韦尔Southwell 	feud.Barony
	蒂克希尔Tickhill 	feud.Barony
	沃克索普Worksop 	feud.Barony
}

func (c *莱斯特LeicesterCounty) BHucknall哈克纳尔() feud.Barony {
	return c.哈克纳尔Hucknall
}
    
func (c *莱斯特LeicesterCounty) BLeicester莱斯特() feud.Barony {
	return c.莱斯特Leicester
}
    
func (c *莱斯特LeicesterCounty) BNewark纽瓦克() feud.Barony {
	return c.纽瓦克Newark
}
    
func (c *莱斯特LeicesterCounty) BNewstead纽斯特德() feud.Barony {
	return c.纽斯特德Newstead
}
    
func (c *莱斯特LeicesterCounty) BNottingham诺丁汉() feud.Barony {
	return c.诺丁汉Nottingham
}
    
func (c *莱斯特LeicesterCounty) BSouthwell索斯韦尔() feud.Barony {
	return c.索斯韦尔Southwell
}
    
func (c *莱斯特LeicesterCounty) BTickhill蒂克希尔() feud.Barony {
	return c.蒂克希尔Tickhill
}
    
func (c *莱斯特LeicesterCounty) BWorksop沃克索普() feud.Barony {
	return c.沃克索普Worksop
}
    
var CLeicester莱斯特 LeicesterCounty = &莱斯特LeicesterCounty{}

func init() {
	f := CLeicester莱斯特.(*莱斯特LeicesterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "62",
		Title:     "leicester",
		TitleName: "莱斯特",
		TitleCode: "c_leicester",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈克纳尔Hucknall = BHucknall哈克纳尔
	f.哈克纳尔Hucknall.SetParent(f)
	
	f.莱斯特Leicester = BLeicester莱斯特
	f.莱斯特Leicester.SetParent(f)
	
	f.纽瓦克Newark = BNewark纽瓦克
	f.纽瓦克Newark.SetParent(f)
	
	f.纽斯特德Newstead = BNewstead纽斯特德
	f.纽斯特德Newstead.SetParent(f)
	
	f.诺丁汉Nottingham = BNottingham诺丁汉
	f.诺丁汉Nottingham.SetParent(f)
	
	f.索斯韦尔Southwell = BSouthwell索斯韦尔
	f.索斯韦尔Southwell.SetParent(f)
	
	f.蒂克希尔Tickhill = BTickhill蒂克希尔
	f.蒂克希尔Tickhill.SetParent(f)
	
	f.沃克索普Worksop = BWorksop沃克索普
	f.沃克索普Worksop.SetParent(f)
	
}
