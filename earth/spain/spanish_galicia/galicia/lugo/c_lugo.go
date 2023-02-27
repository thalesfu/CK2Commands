package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LugoCounty interface {
    feud.County
    BBurgas_orense布尔加斯() 	feud.Barony
    BLugo卢戈() 	feud.Barony
    BMino米尼奥() 	feud.Barony
    BNavilubio纳维卢比奥() 	feud.Barony
    BOrense奥伦塞() 	feud.Barony
    BRibadeo里瓦德奥() 	feud.Barony
    BSil西尔() 	feud.Barony
}

type 卢戈LugoCounty struct {
	feud.BaseCounty
	布尔加斯Burgas_orense 	feud.Barony
	卢戈Lugo 	feud.Barony
	米尼奥Mino 	feud.Barony
	纳维卢比奥Navilubio 	feud.Barony
	奥伦塞Orense 	feud.Barony
	里瓦德奥Ribadeo 	feud.Barony
	西尔Sil 	feud.Barony
}

func (c *卢戈LugoCounty) BBurgas_orense布尔加斯() feud.Barony {
	return c.布尔加斯Burgas_orense
}
    
func (c *卢戈LugoCounty) BLugo卢戈() feud.Barony {
	return c.卢戈Lugo
}
    
func (c *卢戈LugoCounty) BMino米尼奥() feud.Barony {
	return c.米尼奥Mino
}
    
func (c *卢戈LugoCounty) BNavilubio纳维卢比奥() feud.Barony {
	return c.纳维卢比奥Navilubio
}
    
func (c *卢戈LugoCounty) BOrense奥伦塞() feud.Barony {
	return c.奥伦塞Orense
}
    
func (c *卢戈LugoCounty) BRibadeo里瓦德奥() feud.Barony {
	return c.里瓦德奥Ribadeo
}
    
func (c *卢戈LugoCounty) BSil西尔() feud.Barony {
	return c.西尔Sil
}
    
var CLugo卢戈 LugoCounty = &卢戈LugoCounty{}

func init() {
	f := CLugo卢戈.(*卢戈LugoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1976",
		Title:     "lugo",
		TitleName: "卢戈",
		TitleCode: "c_lugo",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔加斯Burgas_orense = BBurgas_orense布尔加斯
	f.布尔加斯Burgas_orense.SetParent(f)
	
	f.卢戈Lugo = BLugo卢戈
	f.卢戈Lugo.SetParent(f)
	
	f.米尼奥Mino = BMino米尼奥
	f.米尼奥Mino.SetParent(f)
	
	f.纳维卢比奥Navilubio = BNavilubio纳维卢比奥
	f.纳维卢比奥Navilubio.SetParent(f)
	
	f.奥伦塞Orense = BOrense奥伦塞
	f.奥伦塞Orense.SetParent(f)
	
	f.里瓦德奥Ribadeo = BRibadeo里瓦德奥
	f.里瓦德奥Ribadeo.SetParent(f)
	
	f.西尔Sil = BSil西尔
	f.西尔Sil.SetParent(f)
	
}
