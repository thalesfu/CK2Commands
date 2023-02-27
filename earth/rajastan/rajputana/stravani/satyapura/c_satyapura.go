package satyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SatyapuraCounty interface {
    feud.County
    BJuna_barmer周那婆摩尔() 	feud.Barony
    BKiradu吉罗菟() 	feud.Barony
    BKiratakupa罽罗多俱波() 	feud.Barony
    BSatyapura娑底也补罗() 	feud.Barony
    BSrungavarapukota室陵伽伐罗补拘吒() 	feud.Barony
    BTiklik置力() 	feud.Barony
    BTimmanayanipela顶摩尼那吠罗() 	feud.Barony
}

type 娑底也补罗SatyapuraCounty struct {
	feud.BaseCounty
	周那婆摩尔Juna_barmer 	feud.Barony
	吉罗菟Kiradu 	feud.Barony
	罽罗多俱波Kiratakupa 	feud.Barony
	娑底也补罗Satyapura 	feud.Barony
	室陵伽伐罗补拘吒Srungavarapukota 	feud.Barony
	置力Tiklik 	feud.Barony
	顶摩尼那吠罗Timmanayanipela 	feud.Barony
}

func (c *娑底也补罗SatyapuraCounty) BJuna_barmer周那婆摩尔() feud.Barony {
	return c.周那婆摩尔Juna_barmer
}
    
func (c *娑底也补罗SatyapuraCounty) BKiradu吉罗菟() feud.Barony {
	return c.吉罗菟Kiradu
}
    
func (c *娑底也补罗SatyapuraCounty) BKiratakupa罽罗多俱波() feud.Barony {
	return c.罽罗多俱波Kiratakupa
}
    
func (c *娑底也补罗SatyapuraCounty) BSatyapura娑底也补罗() feud.Barony {
	return c.娑底也补罗Satyapura
}
    
func (c *娑底也补罗SatyapuraCounty) BSrungavarapukota室陵伽伐罗补拘吒() feud.Barony {
	return c.室陵伽伐罗补拘吒Srungavarapukota
}
    
func (c *娑底也补罗SatyapuraCounty) BTiklik置力() feud.Barony {
	return c.置力Tiklik
}
    
func (c *娑底也补罗SatyapuraCounty) BTimmanayanipela顶摩尼那吠罗() feud.Barony {
	return c.顶摩尼那吠罗Timmanayanipela
}
    
var CSatyapura娑底也补罗 SatyapuraCounty = &娑底也补罗SatyapuraCounty{}

func init() {
	f := CSatyapura娑底也补罗.(*娑底也补罗SatyapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1294",
		Title:     "satyapura",
		TitleName: "娑底也补罗",
		TitleCode: "c_satyapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.周那婆摩尔Juna_barmer = BJuna_barmer周那婆摩尔
	f.周那婆摩尔Juna_barmer.SetParent(f)
	
	f.吉罗菟Kiradu = BKiradu吉罗菟
	f.吉罗菟Kiradu.SetParent(f)
	
	f.罽罗多俱波Kiratakupa = BKiratakupa罽罗多俱波
	f.罽罗多俱波Kiratakupa.SetParent(f)
	
	f.娑底也补罗Satyapura = BSatyapura娑底也补罗
	f.娑底也补罗Satyapura.SetParent(f)
	
	f.室陵伽伐罗补拘吒Srungavarapukota = BSrungavarapukota室陵伽伐罗补拘吒
	f.室陵伽伐罗补拘吒Srungavarapukota.SetParent(f)
	
	f.置力Tiklik = BTiklik置力
	f.置力Tiklik.SetParent(f)
	
	f.顶摩尼那吠罗Timmanayanipela = BTimmanayanipela顶摩尼那吠罗
	f.顶摩尼那吠罗Timmanayanipela.SetParent(f)
	
}
