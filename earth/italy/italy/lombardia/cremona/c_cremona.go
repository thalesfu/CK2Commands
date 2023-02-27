package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CremonaCounty interface {
    feud.County
    BCasalmaggiore卡萨尔马焦雷() 	feud.Barony
    BCastelgoffredo卡斯泰尔戈弗雷多() 	feud.Barony
    BCrema克雷马() 	feud.Barony
    BCremona克雷莫纳() 	feud.Barony
    BManerbio马内尔比奥() 	feud.Barony
    BSospiro索斯皮罗() 	feud.Barony
    BVescovato韦斯科瓦托() 	feud.Barony
    BViadana维亚达纳() 	feud.Barony
}

type 克雷莫纳CremonaCounty struct {
	feud.BaseCounty
	卡萨尔马焦雷Casalmaggiore 	feud.Barony
	卡斯泰尔戈弗雷多Castelgoffredo 	feud.Barony
	克雷马Crema 	feud.Barony
	克雷莫纳Cremona 	feud.Barony
	马内尔比奥Manerbio 	feud.Barony
	索斯皮罗Sospiro 	feud.Barony
	韦斯科瓦托Vescovato 	feud.Barony
	维亚达纳Viadana 	feud.Barony
}

func (c *克雷莫纳CremonaCounty) BCasalmaggiore卡萨尔马焦雷() feud.Barony {
	return c.卡萨尔马焦雷Casalmaggiore
}
    
func (c *克雷莫纳CremonaCounty) BCastelgoffredo卡斯泰尔戈弗雷多() feud.Barony {
	return c.卡斯泰尔戈弗雷多Castelgoffredo
}
    
func (c *克雷莫纳CremonaCounty) BCrema克雷马() feud.Barony {
	return c.克雷马Crema
}
    
func (c *克雷莫纳CremonaCounty) BCremona克雷莫纳() feud.Barony {
	return c.克雷莫纳Cremona
}
    
func (c *克雷莫纳CremonaCounty) BManerbio马内尔比奥() feud.Barony {
	return c.马内尔比奥Manerbio
}
    
func (c *克雷莫纳CremonaCounty) BSospiro索斯皮罗() feud.Barony {
	return c.索斯皮罗Sospiro
}
    
func (c *克雷莫纳CremonaCounty) BVescovato韦斯科瓦托() feud.Barony {
	return c.韦斯科瓦托Vescovato
}
    
func (c *克雷莫纳CremonaCounty) BViadana维亚达纳() feud.Barony {
	return c.维亚达纳Viadana
}
    
var CCremona克雷莫纳 CremonaCounty = &克雷莫纳CremonaCounty{}

func init() {
	f := CCremona克雷莫纳.(*克雷莫纳CremonaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "320",
		Title:     "cremona",
		TitleName: "克雷莫纳",
		TitleCode: "c_cremona",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡萨尔马焦雷Casalmaggiore = BCasalmaggiore卡萨尔马焦雷
	f.卡萨尔马焦雷Casalmaggiore.SetParent(f)
	
	f.卡斯泰尔戈弗雷多Castelgoffredo = BCastelgoffredo卡斯泰尔戈弗雷多
	f.卡斯泰尔戈弗雷多Castelgoffredo.SetParent(f)
	
	f.克雷马Crema = BCrema克雷马
	f.克雷马Crema.SetParent(f)
	
	f.克雷莫纳Cremona = BCremona克雷莫纳
	f.克雷莫纳Cremona.SetParent(f)
	
	f.马内尔比奥Manerbio = BManerbio马内尔比奥
	f.马内尔比奥Manerbio.SetParent(f)
	
	f.索斯皮罗Sospiro = BSospiro索斯皮罗
	f.索斯皮罗Sospiro.SetParent(f)
	
	f.韦斯科瓦托Vescovato = BVescovato韦斯科瓦托
	f.韦斯科瓦托Vescovato.SetParent(f)
	
	f.维亚达纳Viadana = BViadana维亚达纳
	f.维亚达纳Viadana.SetParent(f)
	
}
