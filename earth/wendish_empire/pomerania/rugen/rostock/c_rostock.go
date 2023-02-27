package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RostockCounty interface {
    feud.County
    BAhrensberg阿伦斯堡() 	feud.Barony
    BDamgarten达姆加滕() 	feud.Barony
    BGustrow居斯特罗() 	feud.Barony
    BMalchin马尔欣() 	feud.Barony
    BPenzlin彭茨林() 	feud.Barony
    BRostock罗斯托克() 	feud.Barony
    BStrelitz施特雷利茨() 	feud.Barony
}

type 罗斯托克RostockCounty struct {
	feud.BaseCounty
	阿伦斯堡Ahrensberg 	feud.Barony
	达姆加滕Damgarten 	feud.Barony
	居斯特罗Gustrow 	feud.Barony
	马尔欣Malchin 	feud.Barony
	彭茨林Penzlin 	feud.Barony
	罗斯托克Rostock 	feud.Barony
	施特雷利茨Strelitz 	feud.Barony
}

func (c *罗斯托克RostockCounty) BAhrensberg阿伦斯堡() feud.Barony {
	return c.阿伦斯堡Ahrensberg
}
    
func (c *罗斯托克RostockCounty) BDamgarten达姆加滕() feud.Barony {
	return c.达姆加滕Damgarten
}
    
func (c *罗斯托克RostockCounty) BGustrow居斯特罗() feud.Barony {
	return c.居斯特罗Gustrow
}
    
func (c *罗斯托克RostockCounty) BMalchin马尔欣() feud.Barony {
	return c.马尔欣Malchin
}
    
func (c *罗斯托克RostockCounty) BPenzlin彭茨林() feud.Barony {
	return c.彭茨林Penzlin
}
    
func (c *罗斯托克RostockCounty) BRostock罗斯托克() feud.Barony {
	return c.罗斯托克Rostock
}
    
func (c *罗斯托克RostockCounty) BStrelitz施特雷利茨() feud.Barony {
	return c.施特雷利茨Strelitz
}
    
var CRostock罗斯托克 RostockCounty = &罗斯托克RostockCounty{}

func init() {
	f := CRostock罗斯托克.(*罗斯托克RostockCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "306",
		Title:     "rostock",
		TitleName: "罗斯托克",
		TitleCode: "c_rostock",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伦斯堡Ahrensberg = BAhrensberg阿伦斯堡
	f.阿伦斯堡Ahrensberg.SetParent(f)
	
	f.达姆加滕Damgarten = BDamgarten达姆加滕
	f.达姆加滕Damgarten.SetParent(f)
	
	f.居斯特罗Gustrow = BGustrow居斯特罗
	f.居斯特罗Gustrow.SetParent(f)
	
	f.马尔欣Malchin = BMalchin马尔欣
	f.马尔欣Malchin.SetParent(f)
	
	f.彭茨林Penzlin = BPenzlin彭茨林
	f.彭茨林Penzlin.SetParent(f)
	
	f.罗斯托克Rostock = BRostock罗斯托克
	f.罗斯托克Rostock.SetParent(f)
	
	f.施特雷利茨Strelitz = BStrelitz施特雷利茨
	f.施特雷利茨Strelitz.SetParent(f)
	
}
