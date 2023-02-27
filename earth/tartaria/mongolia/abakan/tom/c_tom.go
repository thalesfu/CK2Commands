package tom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TomCounty interface {
    feud.County
    BIlinka伊利因卡() 	feud.Barony
    BShorokhovo绍罗霍沃() 	feud.Barony
    BSlavino斯拉维诺() 	feud.Barony
    BTerekhino捷列希诺() 	feud.Barony
    BTom托木() 	feud.Barony
    BUval乌瓦尔() 	feud.Barony
    BZhernovo热尔诺沃() 	feud.Barony
}

type 托木TomCounty struct {
	feud.BaseCounty
	伊利因卡Ilinka 	feud.Barony
	绍罗霍沃Shorokhovo 	feud.Barony
	斯拉维诺Slavino 	feud.Barony
	捷列希诺Terekhino 	feud.Barony
	托木Tom 	feud.Barony
	乌瓦尔Uval 	feud.Barony
	热尔诺沃Zhernovo 	feud.Barony
}

func (c *托木TomCounty) BIlinka伊利因卡() feud.Barony {
	return c.伊利因卡Ilinka
}
    
func (c *托木TomCounty) BShorokhovo绍罗霍沃() feud.Barony {
	return c.绍罗霍沃Shorokhovo
}
    
func (c *托木TomCounty) BSlavino斯拉维诺() feud.Barony {
	return c.斯拉维诺Slavino
}
    
func (c *托木TomCounty) BTerekhino捷列希诺() feud.Barony {
	return c.捷列希诺Terekhino
}
    
func (c *托木TomCounty) BTom托木() feud.Barony {
	return c.托木Tom
}
    
func (c *托木TomCounty) BUval乌瓦尔() feud.Barony {
	return c.乌瓦尔Uval
}
    
func (c *托木TomCounty) BZhernovo热尔诺沃() feud.Barony {
	return c.热尔诺沃Zhernovo
}
    
var CTom托木 TomCounty = &托木TomCounty{}

func init() {
	f := CTom托木.(*托木TomCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1900",
		Title:     "tom",
		TitleName: "托木",
		TitleCode: "c_tom",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊利因卡Ilinka = BIlinka伊利因卡
	f.伊利因卡Ilinka.SetParent(f)
	
	f.绍罗霍沃Shorokhovo = BShorokhovo绍罗霍沃
	f.绍罗霍沃Shorokhovo.SetParent(f)
	
	f.斯拉维诺Slavino = BSlavino斯拉维诺
	f.斯拉维诺Slavino.SetParent(f)
	
	f.捷列希诺Terekhino = BTerekhino捷列希诺
	f.捷列希诺Terekhino.SetParent(f)
	
	f.托木Tom = BTom托木
	f.托木Tom.SetParent(f)
	
	f.乌瓦尔Uval = BUval乌瓦尔
	f.乌瓦尔Uval.SetParent(f)
	
	f.热尔诺沃Zhernovo = BZhernovo热尔诺沃
	f.热尔诺沃Zhernovo.SetParent(f)
	
}
