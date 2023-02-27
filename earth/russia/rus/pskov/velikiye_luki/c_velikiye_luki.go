package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Velikiye_lukiCounty interface {
    feud.County
    BBely别雷() 	feud.Barony
    BLoknya洛克尼亚() 	feud.Barony
    BNevel涅韦尔() 	feud.Barony
    BSebezh谢别日() 	feud.Barony
    BStaryatoropa旧托罗帕() 	feud.Barony
    BUsvyaty乌斯维亚特() 	feud.Barony
    BVelikiyeluki卢基() 	feud.Barony
}

type 卢基Velikiye_lukiCounty struct {
	feud.BaseCounty
	别雷Bely 	feud.Barony
	洛克尼亚Loknya 	feud.Barony
	涅韦尔Nevel 	feud.Barony
	谢别日Sebezh 	feud.Barony
	旧托罗帕Staryatoropa 	feud.Barony
	乌斯维亚特Usvyaty 	feud.Barony
	卢基Velikiyeluki 	feud.Barony
}

func (c *卢基Velikiye_lukiCounty) BBely别雷() feud.Barony {
	return c.别雷Bely
}
    
func (c *卢基Velikiye_lukiCounty) BLoknya洛克尼亚() feud.Barony {
	return c.洛克尼亚Loknya
}
    
func (c *卢基Velikiye_lukiCounty) BNevel涅韦尔() feud.Barony {
	return c.涅韦尔Nevel
}
    
func (c *卢基Velikiye_lukiCounty) BSebezh谢别日() feud.Barony {
	return c.谢别日Sebezh
}
    
func (c *卢基Velikiye_lukiCounty) BStaryatoropa旧托罗帕() feud.Barony {
	return c.旧托罗帕Staryatoropa
}
    
func (c *卢基Velikiye_lukiCounty) BUsvyaty乌斯维亚特() feud.Barony {
	return c.乌斯维亚特Usvyaty
}
    
func (c *卢基Velikiye_lukiCounty) BVelikiyeluki卢基() feud.Barony {
	return c.卢基Velikiyeluki
}
    
var CVelikiye_luki卢基 Velikiye_lukiCounty = &卢基Velikiye_lukiCounty{}

func init() {
	f := CVelikiye_luki卢基.(*卢基Velikiye_lukiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "415",
		Title:     "velikiye_luki",
		TitleName: "卢基",
		TitleCode: "c_velikiye_luki",
		Baronies:  map[string]feud.Barony{},
	}

	f.别雷Bely = BBely别雷
	f.别雷Bely.SetParent(f)
	
	f.洛克尼亚Loknya = BLoknya洛克尼亚
	f.洛克尼亚Loknya.SetParent(f)
	
	f.涅韦尔Nevel = BNevel涅韦尔
	f.涅韦尔Nevel.SetParent(f)
	
	f.谢别日Sebezh = BSebezh谢别日
	f.谢别日Sebezh.SetParent(f)
	
	f.旧托罗帕Staryatoropa = BStaryatoropa旧托罗帕
	f.旧托罗帕Staryatoropa.SetParent(f)
	
	f.乌斯维亚特Usvyaty = BUsvyaty乌斯维亚特
	f.乌斯维亚特Usvyaty.SetParent(f)
	
	f.卢基Velikiyeluki = BVelikiyeluki卢基
	f.卢基Velikiyeluki.SetParent(f)
	
}
