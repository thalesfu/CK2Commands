package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SelengeCounty interface {
    feud.County
    BBorgoy博尔戈伊() 	feud.Barony
    BBotsiy博斯() 	feud.Barony
    BDzhida者台() 	feud.Barony
    BKhyagt恰克图() 	feud.Barony
    BNaushki纳乌什基() 	feud.Barony
    BSelenge娑陵() 	feud.Barony
    BYonkhor永霍尔() 	feud.Barony
}

type 娑陵SelengeCounty struct {
	feud.BaseCounty
	博尔戈伊Borgoy 	feud.Barony
	博斯Botsiy 	feud.Barony
	者台Dzhida 	feud.Barony
	恰克图Khyagt 	feud.Barony
	纳乌什基Naushki 	feud.Barony
	娑陵Selenge 	feud.Barony
	永霍尔Yonkhor 	feud.Barony
}

func (c *娑陵SelengeCounty) BBorgoy博尔戈伊() feud.Barony {
	return c.博尔戈伊Borgoy
}
    
func (c *娑陵SelengeCounty) BBotsiy博斯() feud.Barony {
	return c.博斯Botsiy
}
    
func (c *娑陵SelengeCounty) BDzhida者台() feud.Barony {
	return c.者台Dzhida
}
    
func (c *娑陵SelengeCounty) BKhyagt恰克图() feud.Barony {
	return c.恰克图Khyagt
}
    
func (c *娑陵SelengeCounty) BNaushki纳乌什基() feud.Barony {
	return c.纳乌什基Naushki
}
    
func (c *娑陵SelengeCounty) BSelenge娑陵() feud.Barony {
	return c.娑陵Selenge
}
    
func (c *娑陵SelengeCounty) BYonkhor永霍尔() feud.Barony {
	return c.永霍尔Yonkhor
}
    
var CSelenge娑陵 SelengeCounty = &娑陵SelengeCounty{}

func init() {
	f := CSelenge娑陵.(*娑陵SelengeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1921",
		Title:     "selenge",
		TitleName: "娑陵",
		TitleCode: "c_selenge",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔戈伊Borgoy = BBorgoy博尔戈伊
	f.博尔戈伊Borgoy.SetParent(f)
	
	f.博斯Botsiy = BBotsiy博斯
	f.博斯Botsiy.SetParent(f)
	
	f.者台Dzhida = BDzhida者台
	f.者台Dzhida.SetParent(f)
	
	f.恰克图Khyagt = BKhyagt恰克图
	f.恰克图Khyagt.SetParent(f)
	
	f.纳乌什基Naushki = BNaushki纳乌什基
	f.纳乌什基Naushki.SetParent(f)
	
	f.娑陵Selenge = BSelenge娑陵
	f.娑陵Selenge.SetParent(f)
	
	f.永霍尔Yonkhor = BYonkhor永霍尔
	f.永霍尔Yonkhor.SetParent(f)
	
}
