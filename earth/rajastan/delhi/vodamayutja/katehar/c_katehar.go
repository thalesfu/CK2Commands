package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KateharCounty interface {
    feud.County
    BBijnor毗杰诺尔() 	feud.Barony
    BDusajabad菟娑阇跋陀() 	feud.Barony
    BGangadvara殑伽河门() 	feud.Barony
    BJageshwar阇祇湿伐罗() 	feud.Barony
    BMandawar曼荼婆罗() 	feud.Barony
    BNehtaur内多尔() 	feud.Barony
    BRishikesh颉利史计舍() 	feud.Barony
}

type 揭谛诃罗KateharCounty struct {
	feud.BaseCounty
	毗杰诺尔Bijnor 	feud.Barony
	菟娑阇跋陀Dusajabad 	feud.Barony
	殑伽河门Gangadvara 	feud.Barony
	阇祇湿伐罗Jageshwar 	feud.Barony
	曼荼婆罗Mandawar 	feud.Barony
	内多尔Nehtaur 	feud.Barony
	颉利史计舍Rishikesh 	feud.Barony
}

func (c *揭谛诃罗KateharCounty) BBijnor毗杰诺尔() feud.Barony {
	return c.毗杰诺尔Bijnor
}
    
func (c *揭谛诃罗KateharCounty) BDusajabad菟娑阇跋陀() feud.Barony {
	return c.菟娑阇跋陀Dusajabad
}
    
func (c *揭谛诃罗KateharCounty) BGangadvara殑伽河门() feud.Barony {
	return c.殑伽河门Gangadvara
}
    
func (c *揭谛诃罗KateharCounty) BJageshwar阇祇湿伐罗() feud.Barony {
	return c.阇祇湿伐罗Jageshwar
}
    
func (c *揭谛诃罗KateharCounty) BMandawar曼荼婆罗() feud.Barony {
	return c.曼荼婆罗Mandawar
}
    
func (c *揭谛诃罗KateharCounty) BNehtaur内多尔() feud.Barony {
	return c.内多尔Nehtaur
}
    
func (c *揭谛诃罗KateharCounty) BRishikesh颉利史计舍() feud.Barony {
	return c.颉利史计舍Rishikesh
}
    
var CKatehar揭谛诃罗 KateharCounty = &揭谛诃罗KateharCounty{}

func init() {
	f := CKatehar揭谛诃罗.(*揭谛诃罗KateharCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1168",
		Title:     "katehar",
		TitleName: "揭谛诃罗",
		TitleCode: "c_katehar",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗杰诺尔Bijnor = BBijnor毗杰诺尔
	f.毗杰诺尔Bijnor.SetParent(f)
	
	f.菟娑阇跋陀Dusajabad = BDusajabad菟娑阇跋陀
	f.菟娑阇跋陀Dusajabad.SetParent(f)
	
	f.殑伽河门Gangadvara = BGangadvara殑伽河门
	f.殑伽河门Gangadvara.SetParent(f)
	
	f.阇祇湿伐罗Jageshwar = BJageshwar阇祇湿伐罗
	f.阇祇湿伐罗Jageshwar.SetParent(f)
	
	f.曼荼婆罗Mandawar = BMandawar曼荼婆罗
	f.曼荼婆罗Mandawar.SetParent(f)
	
	f.内多尔Nehtaur = BNehtaur内多尔
	f.内多尔Nehtaur.SetParent(f)
	
	f.颉利史计舍Rishikesh = BRishikesh颉利史计舍
	f.颉利史计舍Rishikesh.SetParent(f)
	
}
