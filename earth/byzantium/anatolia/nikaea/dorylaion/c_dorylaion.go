package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DorylaionCounty interface {
    feud.County
    BCarura加卢拉() 	feud.Barony
    BDorylaion多律莱翁() 	feud.Barony
    BGermia日尔米亚() 	feud.Barony
    BIustinianopolis优士丁尼波利斯() 	feud.Barony
    BKotiaion科泰延() 	feud.Barony
    BNaeotea奈奥蒂亚() 	feud.Barony
    BOrkistos欧基斯图斯() 	feud.Barony
}

type 多律莱翁DorylaionCounty struct {
	feud.BaseCounty
	加卢拉Carura 	feud.Barony
	多律莱翁Dorylaion 	feud.Barony
	日尔米亚Germia 	feud.Barony
	优士丁尼波利斯Iustinianopolis 	feud.Barony
	科泰延Kotiaion 	feud.Barony
	奈奥蒂亚Naeotea 	feud.Barony
	欧基斯图斯Orkistos 	feud.Barony
}

func (c *多律莱翁DorylaionCounty) BCarura加卢拉() feud.Barony {
	return c.加卢拉Carura
}
    
func (c *多律莱翁DorylaionCounty) BDorylaion多律莱翁() feud.Barony {
	return c.多律莱翁Dorylaion
}
    
func (c *多律莱翁DorylaionCounty) BGermia日尔米亚() feud.Barony {
	return c.日尔米亚Germia
}
    
func (c *多律莱翁DorylaionCounty) BIustinianopolis优士丁尼波利斯() feud.Barony {
	return c.优士丁尼波利斯Iustinianopolis
}
    
func (c *多律莱翁DorylaionCounty) BKotiaion科泰延() feud.Barony {
	return c.科泰延Kotiaion
}
    
func (c *多律莱翁DorylaionCounty) BNaeotea奈奥蒂亚() feud.Barony {
	return c.奈奥蒂亚Naeotea
}
    
func (c *多律莱翁DorylaionCounty) BOrkistos欧基斯图斯() feud.Barony {
	return c.欧基斯图斯Orkistos
}
    
var CDorylaion多律莱翁 DorylaionCounty = &多律莱翁DorylaionCounty{}

func init() {
	f := CDorylaion多律莱翁.(*多律莱翁DorylaionCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "749",
		Title:     "dorylaion",
		TitleName: "多律莱翁",
		TitleCode: "c_dorylaion",
		Baronies:  map[string]feud.Barony{},
	}

	f.加卢拉Carura = BCarura加卢拉
	f.加卢拉Carura.SetParent(f)
	
	f.多律莱翁Dorylaion = BDorylaion多律莱翁
	f.多律莱翁Dorylaion.SetParent(f)
	
	f.日尔米亚Germia = BGermia日尔米亚
	f.日尔米亚Germia.SetParent(f)
	
	f.优士丁尼波利斯Iustinianopolis = BIustinianopolis优士丁尼波利斯
	f.优士丁尼波利斯Iustinianopolis.SetParent(f)
	
	f.科泰延Kotiaion = BKotiaion科泰延
	f.科泰延Kotiaion.SetParent(f)
	
	f.奈奥蒂亚Naeotea = BNaeotea奈奥蒂亚
	f.奈奥蒂亚Naeotea.SetParent(f)
	
	f.欧基斯图斯Orkistos = BOrkistos欧基斯图斯
	f.欧基斯图斯Orkistos.SetParent(f)
	
}
