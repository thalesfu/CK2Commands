package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DaylamCounty interface {
    feud.County
    BGheydar盖达尔() 	feud.Barony
    BKhalkhal哈勒哈勒() 	feud.Barony
    BKhunaj胡纳杰() 	feud.Barony
    BMahneshan马赫内尚() 	feud.Barony
    BSoltaniyeh苏丹尼耶() 	feud.Barony
    BTarom塔罗姆() 	feud.Barony
    BZanjan赞詹() 	feud.Barony
}

type 德莱木DaylamCounty struct {
	feud.BaseCounty
	盖达尔Gheydar 	feud.Barony
	哈勒哈勒Khalkhal 	feud.Barony
	胡纳杰Khunaj 	feud.Barony
	马赫内尚Mahneshan 	feud.Barony
	苏丹尼耶Soltaniyeh 	feud.Barony
	塔罗姆Tarom 	feud.Barony
	赞詹Zanjan 	feud.Barony
}

func (c *德莱木DaylamCounty) BGheydar盖达尔() feud.Barony {
	return c.盖达尔Gheydar
}
    
func (c *德莱木DaylamCounty) BKhalkhal哈勒哈勒() feud.Barony {
	return c.哈勒哈勒Khalkhal
}
    
func (c *德莱木DaylamCounty) BKhunaj胡纳杰() feud.Barony {
	return c.胡纳杰Khunaj
}
    
func (c *德莱木DaylamCounty) BMahneshan马赫内尚() feud.Barony {
	return c.马赫内尚Mahneshan
}
    
func (c *德莱木DaylamCounty) BSoltaniyeh苏丹尼耶() feud.Barony {
	return c.苏丹尼耶Soltaniyeh
}
    
func (c *德莱木DaylamCounty) BTarom塔罗姆() feud.Barony {
	return c.塔罗姆Tarom
}
    
func (c *德莱木DaylamCounty) BZanjan赞詹() feud.Barony {
	return c.赞詹Zanjan
}
    
var CDaylam德莱木 DaylamCounty = &德莱木DaylamCounty{}

func init() {
	f := CDaylam德莱木.(*德莱木DaylamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "665",
		Title:     "daylam",
		TitleName: "德莱木",
		TitleCode: "c_daylam",
		Baronies:  map[string]feud.Barony{},
	}

	f.盖达尔Gheydar = BGheydar盖达尔
	f.盖达尔Gheydar.SetParent(f)
	
	f.哈勒哈勒Khalkhal = BKhalkhal哈勒哈勒
	f.哈勒哈勒Khalkhal.SetParent(f)
	
	f.胡纳杰Khunaj = BKhunaj胡纳杰
	f.胡纳杰Khunaj.SetParent(f)
	
	f.马赫内尚Mahneshan = BMahneshan马赫内尚
	f.马赫内尚Mahneshan.SetParent(f)
	
	f.苏丹尼耶Soltaniyeh = BSoltaniyeh苏丹尼耶
	f.苏丹尼耶Soltaniyeh.SetParent(f)
	
	f.塔罗姆Tarom = BTarom塔罗姆
	f.塔罗姆Tarom.SetParent(f)
	
	f.赞詹Zanjan = BZanjan赞詹
	f.赞詹Zanjan.SetParent(f)
	
}
