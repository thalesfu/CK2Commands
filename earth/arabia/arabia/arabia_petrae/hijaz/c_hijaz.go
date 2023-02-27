package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HijazCounty interface {
    feud.County
    BAl_ola欧拉() 	feud.Barony
    BHigra希格拉() 	feud.Barony
    BLeuke_kome卢克科梅() 	feud.Barony
    BMadain_salih玛甸沙勒() 	feud.Barony
    BMughayra穆盖拉() 	feud.Barony
    BSamur萨穆尔() 	feud.Barony
    BTawala塔瓦拉() 	feud.Barony
    BTayma泰马() 	feud.Barony
}

type 欧拉HijazCounty struct {
	feud.BaseCounty
	欧拉Al_ola 	feud.Barony
	希格拉Higra 	feud.Barony
	卢克科梅Leuke_kome 	feud.Barony
	玛甸沙勒Madain_salih 	feud.Barony
	穆盖拉Mughayra 	feud.Barony
	萨穆尔Samur 	feud.Barony
	塔瓦拉Tawala 	feud.Barony
	泰马Tayma 	feud.Barony
}

func (c *欧拉HijazCounty) BAl_ola欧拉() feud.Barony {
	return c.欧拉Al_ola
}
    
func (c *欧拉HijazCounty) BHigra希格拉() feud.Barony {
	return c.希格拉Higra
}
    
func (c *欧拉HijazCounty) BLeuke_kome卢克科梅() feud.Barony {
	return c.卢克科梅Leuke_kome
}
    
func (c *欧拉HijazCounty) BMadain_salih玛甸沙勒() feud.Barony {
	return c.玛甸沙勒Madain_salih
}
    
func (c *欧拉HijazCounty) BMughayra穆盖拉() feud.Barony {
	return c.穆盖拉Mughayra
}
    
func (c *欧拉HijazCounty) BSamur萨穆尔() feud.Barony {
	return c.萨穆尔Samur
}
    
func (c *欧拉HijazCounty) BTawala塔瓦拉() feud.Barony {
	return c.塔瓦拉Tawala
}
    
func (c *欧拉HijazCounty) BTayma泰马() feud.Barony {
	return c.泰马Tayma
}
    
var CHijaz欧拉 HijazCounty = &欧拉HijazCounty{}

func init() {
	f := CHijaz欧拉.(*欧拉HijazCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "720",
		Title:     "hijaz",
		TitleName: "欧拉",
		TitleCode: "c_hijaz",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧拉Al_ola = BAl_ola欧拉
	f.欧拉Al_ola.SetParent(f)
	
	f.希格拉Higra = BHigra希格拉
	f.希格拉Higra.SetParent(f)
	
	f.卢克科梅Leuke_kome = BLeuke_kome卢克科梅
	f.卢克科梅Leuke_kome.SetParent(f)
	
	f.玛甸沙勒Madain_salih = BMadain_salih玛甸沙勒
	f.玛甸沙勒Madain_salih.SetParent(f)
	
	f.穆盖拉Mughayra = BMughayra穆盖拉
	f.穆盖拉Mughayra.SetParent(f)
	
	f.萨穆尔Samur = BSamur萨穆尔
	f.萨穆尔Samur.SetParent(f)
	
	f.塔瓦拉Tawala = BTawala塔瓦拉
	f.塔瓦拉Tawala.SetParent(f)
	
	f.泰马Tayma = BTayma泰马
	f.泰马Tayma.SetParent(f)
	
}
