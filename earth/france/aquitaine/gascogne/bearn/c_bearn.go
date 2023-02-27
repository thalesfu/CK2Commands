package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BearnCounty interface {
    feud.County
    BLescar莱斯卡尔() 	feud.Barony
    BMauleonlicharre莫莱翁利沙尔() 	feud.Barony
    BMontaner蒙塔内() 	feud.Barony
    BMorlaas莫尔拉斯() 	feud.Barony
    BOloron奥洛龙() 	feud.Barony
    BOrthez奥尔泰兹() 	feud.Barony
    BPau波城() 	feud.Barony
    BTarbes塔布() 	feud.Barony
}

type 贝阿恩BearnCounty struct {
	feud.BaseCounty
	莱斯卡尔Lescar 	feud.Barony
	莫莱翁利沙尔Mauleonlicharre 	feud.Barony
	蒙塔内Montaner 	feud.Barony
	莫尔拉斯Morlaas 	feud.Barony
	奥洛龙Oloron 	feud.Barony
	奥尔泰兹Orthez 	feud.Barony
	波城Pau 	feud.Barony
	塔布Tarbes 	feud.Barony
}

func (c *贝阿恩BearnCounty) BLescar莱斯卡尔() feud.Barony {
	return c.莱斯卡尔Lescar
}
    
func (c *贝阿恩BearnCounty) BMauleonlicharre莫莱翁利沙尔() feud.Barony {
	return c.莫莱翁利沙尔Mauleonlicharre
}
    
func (c *贝阿恩BearnCounty) BMontaner蒙塔内() feud.Barony {
	return c.蒙塔内Montaner
}
    
func (c *贝阿恩BearnCounty) BMorlaas莫尔拉斯() feud.Barony {
	return c.莫尔拉斯Morlaas
}
    
func (c *贝阿恩BearnCounty) BOloron奥洛龙() feud.Barony {
	return c.奥洛龙Oloron
}
    
func (c *贝阿恩BearnCounty) BOrthez奥尔泰兹() feud.Barony {
	return c.奥尔泰兹Orthez
}
    
func (c *贝阿恩BearnCounty) BPau波城() feud.Barony {
	return c.波城Pau
}
    
func (c *贝阿恩BearnCounty) BTarbes塔布() feud.Barony {
	return c.塔布Tarbes
}
    
var CBearn贝阿恩 BearnCounty = &贝阿恩BearnCounty{}

func init() {
	f := CBearn贝阿恩.(*贝阿恩BearnCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "208",
		Title:     "bearn",
		TitleName: "贝阿恩",
		TitleCode: "c_bearn",
		Baronies:  map[string]feud.Barony{},
	}

	f.莱斯卡尔Lescar = BLescar莱斯卡尔
	f.莱斯卡尔Lescar.SetParent(f)
	
	f.莫莱翁利沙尔Mauleonlicharre = BMauleonlicharre莫莱翁利沙尔
	f.莫莱翁利沙尔Mauleonlicharre.SetParent(f)
	
	f.蒙塔内Montaner = BMontaner蒙塔内
	f.蒙塔内Montaner.SetParent(f)
	
	f.莫尔拉斯Morlaas = BMorlaas莫尔拉斯
	f.莫尔拉斯Morlaas.SetParent(f)
	
	f.奥洛龙Oloron = BOloron奥洛龙
	f.奥洛龙Oloron.SetParent(f)
	
	f.奥尔泰兹Orthez = BOrthez奥尔泰兹
	f.奥尔泰兹Orthez.SetParent(f)
	
	f.波城Pau = BPau波城
	f.波城Pau.SetParent(f)
	
	f.塔布Tarbes = BTarbes塔布
	f.塔布Tarbes.SetParent(f)
	
}
