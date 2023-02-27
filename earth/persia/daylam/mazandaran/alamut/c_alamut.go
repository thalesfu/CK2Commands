package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlamutCounty interface {
    feud.County
    BAlamut阿拉穆特() 	feud.Barony
    BChaloos恰卢斯() 	feud.Barony
    BDamavand达马万德() 	feud.Barony
    BGarmsar加姆萨尔() 	feud.Barony
    BKalar卡拉尔() 	feud.Barony
    BKojoor阔佐尔() 	feud.Barony
    BRostamrood罗斯达摩罗() 	feud.Barony
    BSisangan斯桑干() 	feud.Barony
    BVaramin瓦拉明() 	feud.Barony
}

type 阿拉穆特AlamutCounty struct {
	feud.BaseCounty
	阿拉穆特Alamut 	feud.Barony
	恰卢斯Chaloos 	feud.Barony
	达马万德Damavand 	feud.Barony
	加姆萨尔Garmsar 	feud.Barony
	卡拉尔Kalar 	feud.Barony
	阔佐尔Kojoor 	feud.Barony
	罗斯达摩罗Rostamrood 	feud.Barony
	斯桑干Sisangan 	feud.Barony
	瓦拉明Varamin 	feud.Barony
}

func (c *阿拉穆特AlamutCounty) BAlamut阿拉穆特() feud.Barony {
	return c.阿拉穆特Alamut
}
    
func (c *阿拉穆特AlamutCounty) BChaloos恰卢斯() feud.Barony {
	return c.恰卢斯Chaloos
}
    
func (c *阿拉穆特AlamutCounty) BDamavand达马万德() feud.Barony {
	return c.达马万德Damavand
}
    
func (c *阿拉穆特AlamutCounty) BGarmsar加姆萨尔() feud.Barony {
	return c.加姆萨尔Garmsar
}
    
func (c *阿拉穆特AlamutCounty) BKalar卡拉尔() feud.Barony {
	return c.卡拉尔Kalar
}
    
func (c *阿拉穆特AlamutCounty) BKojoor阔佐尔() feud.Barony {
	return c.阔佐尔Kojoor
}
    
func (c *阿拉穆特AlamutCounty) BRostamrood罗斯达摩罗() feud.Barony {
	return c.罗斯达摩罗Rostamrood
}
    
func (c *阿拉穆特AlamutCounty) BSisangan斯桑干() feud.Barony {
	return c.斯桑干Sisangan
}
    
func (c *阿拉穆特AlamutCounty) BVaramin瓦拉明() feud.Barony {
	return c.瓦拉明Varamin
}
    
var CAlamut阿拉穆特 AlamutCounty = &阿拉穆特AlamutCounty{}

func init() {
	f := CAlamut阿拉穆特.(*阿拉穆特AlamutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "662",
		Title:     "alamut",
		TitleName: "阿拉穆特",
		TitleCode: "c_alamut",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉穆特Alamut = BAlamut阿拉穆特
	f.阿拉穆特Alamut.SetParent(f)
	
	f.恰卢斯Chaloos = BChaloos恰卢斯
	f.恰卢斯Chaloos.SetParent(f)
	
	f.达马万德Damavand = BDamavand达马万德
	f.达马万德Damavand.SetParent(f)
	
	f.加姆萨尔Garmsar = BGarmsar加姆萨尔
	f.加姆萨尔Garmsar.SetParent(f)
	
	f.卡拉尔Kalar = BKalar卡拉尔
	f.卡拉尔Kalar.SetParent(f)
	
	f.阔佐尔Kojoor = BKojoor阔佐尔
	f.阔佐尔Kojoor.SetParent(f)
	
	f.罗斯达摩罗Rostamrood = BRostamrood罗斯达摩罗
	f.罗斯达摩罗Rostamrood.SetParent(f)
	
	f.斯桑干Sisangan = BSisangan斯桑干
	f.斯桑干Sisangan.SetParent(f)
	
	f.瓦拉明Varamin = BVaramin瓦拉明
	f.瓦拉明Varamin.SetParent(f)
	
}
