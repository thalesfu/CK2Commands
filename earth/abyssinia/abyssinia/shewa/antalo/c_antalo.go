package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AntaloCounty interface {
    feud.County
    BAntsokia安茨基() 	feud.Barony
    BBale巴莱() 	feud.Barony
    BDebreberhan德卜勒伯尔汉() 	feud.Barony
    BDebrelibanos德布腊利巴诺斯() 	feud.Barony
    BEntoto恩托托() 	feud.Barony
    BHayq海格() 	feud.Barony
    BOromos奥罗莫斯() 	feud.Barony
    BShewa绍阿() 	feud.Barony
}

type 安塔洛AntaloCounty struct {
	feud.BaseCounty
	安茨基Antsokia 	feud.Barony
	巴莱Bale 	feud.Barony
	德卜勒伯尔汉Debreberhan 	feud.Barony
	德布腊利巴诺斯Debrelibanos 	feud.Barony
	恩托托Entoto 	feud.Barony
	海格Hayq 	feud.Barony
	奥罗莫斯Oromos 	feud.Barony
	绍阿Shewa 	feud.Barony
}

func (c *安塔洛AntaloCounty) BAntsokia安茨基() feud.Barony {
	return c.安茨基Antsokia
}
    
func (c *安塔洛AntaloCounty) BBale巴莱() feud.Barony {
	return c.巴莱Bale
}
    
func (c *安塔洛AntaloCounty) BDebreberhan德卜勒伯尔汉() feud.Barony {
	return c.德卜勒伯尔汉Debreberhan
}
    
func (c *安塔洛AntaloCounty) BDebrelibanos德布腊利巴诺斯() feud.Barony {
	return c.德布腊利巴诺斯Debrelibanos
}
    
func (c *安塔洛AntaloCounty) BEntoto恩托托() feud.Barony {
	return c.恩托托Entoto
}
    
func (c *安塔洛AntaloCounty) BHayq海格() feud.Barony {
	return c.海格Hayq
}
    
func (c *安塔洛AntaloCounty) BOromos奥罗莫斯() feud.Barony {
	return c.奥罗莫斯Oromos
}
    
func (c *安塔洛AntaloCounty) BShewa绍阿() feud.Barony {
	return c.绍阿Shewa
}
    
var CAntalo安塔洛 AntaloCounty = &安塔洛AntaloCounty{}

func init() {
	f := CAntalo安塔洛.(*安塔洛AntaloCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "884",
		Title:     "antalo",
		TitleName: "安塔洛",
		TitleCode: "c_antalo",
		Baronies:  map[string]feud.Barony{},
	}

	f.安茨基Antsokia = BAntsokia安茨基
	f.安茨基Antsokia.SetParent(f)
	
	f.巴莱Bale = BBale巴莱
	f.巴莱Bale.SetParent(f)
	
	f.德卜勒伯尔汉Debreberhan = BDebreberhan德卜勒伯尔汉
	f.德卜勒伯尔汉Debreberhan.SetParent(f)
	
	f.德布腊利巴诺斯Debrelibanos = BDebrelibanos德布腊利巴诺斯
	f.德布腊利巴诺斯Debrelibanos.SetParent(f)
	
	f.恩托托Entoto = BEntoto恩托托
	f.恩托托Entoto.SetParent(f)
	
	f.海格Hayq = BHayq海格
	f.海格Hayq.SetParent(f)
	
	f.奥罗莫斯Oromos = BOromos奥罗莫斯
	f.奥罗莫斯Oromos.SetParent(f)
	
	f.绍阿Shewa = BShewa绍阿
	f.绍阿Shewa.SetParent(f)
	
}
