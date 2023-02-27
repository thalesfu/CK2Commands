package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BacauCounty interface {
    feud.County
    BBacau巴克乌() 	feud.Barony
    BBistrita_bacau比斯特里察() 	feud.Barony
    BBodesti博代什蒂() 	feud.Barony
    BNeamnt尼亚姆茨() 	feud.Barony
    BTargu_neamnt特尔古尼亚姆茨() 	feud.Barony
    BTargu_trotus特尔古特罗图什() 	feud.Barony
    BTazlau塔兹勒乌() 	feud.Barony
}

type 巴克乌BacauCounty struct {
	feud.BaseCounty
	巴克乌Bacau 	feud.Barony
	比斯特里察Bistrita_bacau 	feud.Barony
	博代什蒂Bodesti 	feud.Barony
	尼亚姆茨Neamnt 	feud.Barony
	特尔古尼亚姆茨Targu_neamnt 	feud.Barony
	特尔古特罗图什Targu_trotus 	feud.Barony
	塔兹勒乌Tazlau 	feud.Barony
}

func (c *巴克乌BacauCounty) BBacau巴克乌() feud.Barony {
	return c.巴克乌Bacau
}
    
func (c *巴克乌BacauCounty) BBistrita_bacau比斯特里察() feud.Barony {
	return c.比斯特里察Bistrita_bacau
}
    
func (c *巴克乌BacauCounty) BBodesti博代什蒂() feud.Barony {
	return c.博代什蒂Bodesti
}
    
func (c *巴克乌BacauCounty) BNeamnt尼亚姆茨() feud.Barony {
	return c.尼亚姆茨Neamnt
}
    
func (c *巴克乌BacauCounty) BTargu_neamnt特尔古尼亚姆茨() feud.Barony {
	return c.特尔古尼亚姆茨Targu_neamnt
}
    
func (c *巴克乌BacauCounty) BTargu_trotus特尔古特罗图什() feud.Barony {
	return c.特尔古特罗图什Targu_trotus
}
    
func (c *巴克乌BacauCounty) BTazlau塔兹勒乌() feud.Barony {
	return c.塔兹勒乌Tazlau
}
    
var CBacau巴克乌 BacauCounty = &巴克乌BacauCounty{}

func init() {
	f := CBacau巴克乌.(*巴克乌BacauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1639",
		Title:     "bacau",
		TitleName: "巴克乌",
		TitleCode: "c_bacau",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴克乌Bacau = BBacau巴克乌
	f.巴克乌Bacau.SetParent(f)
	
	f.比斯特里察Bistrita_bacau = BBistrita_bacau比斯特里察
	f.比斯特里察Bistrita_bacau.SetParent(f)
	
	f.博代什蒂Bodesti = BBodesti博代什蒂
	f.博代什蒂Bodesti.SetParent(f)
	
	f.尼亚姆茨Neamnt = BNeamnt尼亚姆茨
	f.尼亚姆茨Neamnt.SetParent(f)
	
	f.特尔古尼亚姆茨Targu_neamnt = BTargu_neamnt特尔古尼亚姆茨
	f.特尔古尼亚姆茨Targu_neamnt.SetParent(f)
	
	f.特尔古特罗图什Targu_trotus = BTargu_trotus特尔古特罗图什
	f.特尔古特罗图什Targu_trotus.SetParent(f)
	
	f.塔兹勒乌Tazlau = BTazlau塔兹勒乌
	f.塔兹勒乌Tazlau.SetParent(f)
	
}
