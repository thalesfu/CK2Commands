package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IzhmaCounty interface {
    feud.County
    BDiyur季尤尔() 	feud.Barony
    BIzhma伊日马() 	feud.Barony
    BKartayol卡尔塔约尔() 	feud.Barony
    BMokhcha莫赫恰() 	feud.Barony
    BShchell谢利() 	feud.Barony
    BSizyabsk锡贾布斯克() 	feud.Barony
    BUst_izhma乌斯季_伊日马() 	feud.Barony
}

type 伊日马IzhmaCounty struct {
	feud.BaseCounty
	季尤尔Diyur 	feud.Barony
	伊日马Izhma 	feud.Barony
	卡尔塔约尔Kartayol 	feud.Barony
	莫赫恰Mokhcha 	feud.Barony
	谢利Shchell 	feud.Barony
	锡贾布斯克Sizyabsk 	feud.Barony
	乌斯季_伊日马Ust_izhma 	feud.Barony
}

func (c *伊日马IzhmaCounty) BDiyur季尤尔() feud.Barony {
	return c.季尤尔Diyur
}
    
func (c *伊日马IzhmaCounty) BIzhma伊日马() feud.Barony {
	return c.伊日马Izhma
}
    
func (c *伊日马IzhmaCounty) BKartayol卡尔塔约尔() feud.Barony {
	return c.卡尔塔约尔Kartayol
}
    
func (c *伊日马IzhmaCounty) BMokhcha莫赫恰() feud.Barony {
	return c.莫赫恰Mokhcha
}
    
func (c *伊日马IzhmaCounty) BShchell谢利() feud.Barony {
	return c.谢利Shchell
}
    
func (c *伊日马IzhmaCounty) BSizyabsk锡贾布斯克() feud.Barony {
	return c.锡贾布斯克Sizyabsk
}
    
func (c *伊日马IzhmaCounty) BUst_izhma乌斯季_伊日马() feud.Barony {
	return c.乌斯季_伊日马Ust_izhma
}
    
var CIzhma伊日马 IzhmaCounty = &伊日马IzhmaCounty{}

func init() {
	f := CIzhma伊日马.(*伊日马IzhmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1835",
		Title:     "izhma",
		TitleName: "伊日马",
		TitleCode: "c_izhma",
		Baronies:  map[string]feud.Barony{},
	}

	f.季尤尔Diyur = BDiyur季尤尔
	f.季尤尔Diyur.SetParent(f)
	
	f.伊日马Izhma = BIzhma伊日马
	f.伊日马Izhma.SetParent(f)
	
	f.卡尔塔约尔Kartayol = BKartayol卡尔塔约尔
	f.卡尔塔约尔Kartayol.SetParent(f)
	
	f.莫赫恰Mokhcha = BMokhcha莫赫恰
	f.莫赫恰Mokhcha.SetParent(f)
	
	f.谢利Shchell = BShchell谢利
	f.谢利Shchell.SetParent(f)
	
	f.锡贾布斯克Sizyabsk = BSizyabsk锡贾布斯克
	f.锡贾布斯克Sizyabsk.SetParent(f)
	
	f.乌斯季_伊日马Ust_izhma = BUst_izhma乌斯季_伊日马
	f.乌斯季_伊日马Ust_izhma.SetParent(f)
	
}
