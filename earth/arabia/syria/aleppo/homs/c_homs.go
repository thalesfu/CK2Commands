package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HomsCounty interface {
    feud.County
    BAlkhazandar哈赞达() 	feud.Barony
    BEmesa埃梅萨() 	feud.Barony
    BHoms霍姆斯() 	feud.Barony
    BMarmarita马尔马里塔() 	feud.Barony
    BQadesh卡迭石() 	feud.Barony
    BQatna卡特那() 	feud.Barony
    BSadad塞代德() 	feud.Barony
    BZweitina祖韦提奈() 	feud.Barony
}

type 霍姆斯HomsCounty struct {
	feud.BaseCounty
	哈赞达Alkhazandar 	feud.Barony
	埃梅萨Emesa 	feud.Barony
	霍姆斯Homs 	feud.Barony
	马尔马里塔Marmarita 	feud.Barony
	卡迭石Qadesh 	feud.Barony
	卡特那Qatna 	feud.Barony
	塞代德Sadad 	feud.Barony
	祖韦提奈Zweitina 	feud.Barony
}

func (c *霍姆斯HomsCounty) BAlkhazandar哈赞达() feud.Barony {
	return c.哈赞达Alkhazandar
}
    
func (c *霍姆斯HomsCounty) BEmesa埃梅萨() feud.Barony {
	return c.埃梅萨Emesa
}
    
func (c *霍姆斯HomsCounty) BHoms霍姆斯() feud.Barony {
	return c.霍姆斯Homs
}
    
func (c *霍姆斯HomsCounty) BMarmarita马尔马里塔() feud.Barony {
	return c.马尔马里塔Marmarita
}
    
func (c *霍姆斯HomsCounty) BQadesh卡迭石() feud.Barony {
	return c.卡迭石Qadesh
}
    
func (c *霍姆斯HomsCounty) BQatna卡特那() feud.Barony {
	return c.卡特那Qatna
}
    
func (c *霍姆斯HomsCounty) BSadad塞代德() feud.Barony {
	return c.塞代德Sadad
}
    
func (c *霍姆斯HomsCounty) BZweitina祖韦提奈() feud.Barony {
	return c.祖韦提奈Zweitina
}
    
var CHoms霍姆斯 HomsCounty = &霍姆斯HomsCounty{}

func init() {
	f := CHoms霍姆斯.(*霍姆斯HomsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "731",
		Title:     "homs",
		TitleName: "霍姆斯",
		TitleCode: "c_homs",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈赞达Alkhazandar = BAlkhazandar哈赞达
	f.哈赞达Alkhazandar.SetParent(f)
	
	f.埃梅萨Emesa = BEmesa埃梅萨
	f.埃梅萨Emesa.SetParent(f)
	
	f.霍姆斯Homs = BHoms霍姆斯
	f.霍姆斯Homs.SetParent(f)
	
	f.马尔马里塔Marmarita = BMarmarita马尔马里塔
	f.马尔马里塔Marmarita.SetParent(f)
	
	f.卡迭石Qadesh = BQadesh卡迭石
	f.卡迭石Qadesh.SetParent(f)
	
	f.卡特那Qatna = BQatna卡特那
	f.卡特那Qatna.SetParent(f)
	
	f.塞代德Sadad = BSadad塞代德
	f.塞代德Sadad.SetParent(f)
	
	f.祖韦提奈Zweitina = BZweitina祖韦提奈
	f.祖韦提奈Zweitina.SetParent(f)
	
}
