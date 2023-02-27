package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KerakCounty interface {
    feud.County
    BBozra波斯拉() 	feud.Barony
    BKerak克拉克() 	feud.Barony
    BKirhaseset基尔哈雷塞特() 	feud.Barony
    BKrakdemoab摩押的吉珥() 	feud.Barony
    BPunon普嫩() 	feud.Barony
    BTamar塔马林() 	feud.Barony
    BZaimona撒摩拿() 	feud.Barony
    BZoar琐珥() 	feud.Barony
}

type 克拉克KerakCounty struct {
	feud.BaseCounty
	波斯拉Bozra 	feud.Barony
	克拉克Kerak 	feud.Barony
	基尔哈雷塞特Kirhaseset 	feud.Barony
	摩押的吉珥Krakdemoab 	feud.Barony
	普嫩Punon 	feud.Barony
	塔马林Tamar 	feud.Barony
	撒摩拿Zaimona 	feud.Barony
	琐珥Zoar 	feud.Barony
}

func (c *克拉克KerakCounty) BBozra波斯拉() feud.Barony {
	return c.波斯拉Bozra
}
    
func (c *克拉克KerakCounty) BKerak克拉克() feud.Barony {
	return c.克拉克Kerak
}
    
func (c *克拉克KerakCounty) BKirhaseset基尔哈雷塞特() feud.Barony {
	return c.基尔哈雷塞特Kirhaseset
}
    
func (c *克拉克KerakCounty) BKrakdemoab摩押的吉珥() feud.Barony {
	return c.摩押的吉珥Krakdemoab
}
    
func (c *克拉克KerakCounty) BPunon普嫩() feud.Barony {
	return c.普嫩Punon
}
    
func (c *克拉克KerakCounty) BTamar塔马林() feud.Barony {
	return c.塔马林Tamar
}
    
func (c *克拉克KerakCounty) BZaimona撒摩拿() feud.Barony {
	return c.撒摩拿Zaimona
}
    
func (c *克拉克KerakCounty) BZoar琐珥() feud.Barony {
	return c.琐珥Zoar
}
    
var CKerak克拉克 KerakCounty = &克拉克KerakCounty{}

func init() {
	f := CKerak克拉克.(*克拉克KerakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "777",
		Title:     "kerak",
		TitleName: "克拉克",
		TitleCode: "c_kerak",
		Baronies:  map[string]feud.Barony{},
	}

	f.波斯拉Bozra = BBozra波斯拉
	f.波斯拉Bozra.SetParent(f)
	
	f.克拉克Kerak = BKerak克拉克
	f.克拉克Kerak.SetParent(f)
	
	f.基尔哈雷塞特Kirhaseset = BKirhaseset基尔哈雷塞特
	f.基尔哈雷塞特Kirhaseset.SetParent(f)
	
	f.摩押的吉珥Krakdemoab = BKrakdemoab摩押的吉珥
	f.摩押的吉珥Krakdemoab.SetParent(f)
	
	f.普嫩Punon = BPunon普嫩
	f.普嫩Punon.SetParent(f)
	
	f.塔马林Tamar = BTamar塔马林
	f.塔马林Tamar.SetParent(f)
	
	f.撒摩拿Zaimona = BZaimona撒摩拿
	f.撒摩拿Zaimona.SetParent(f)
	
	f.琐珥Zoar = BZoar琐珥
	f.琐珥Zoar.SetParent(f)
	
}
