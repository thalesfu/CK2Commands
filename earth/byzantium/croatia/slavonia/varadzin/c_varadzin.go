package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VaradzinCounty interface {
    feud.County
    BCakovec查科韦茨() 	feud.Barony
    BDonjastubica下斯图比察() 	feud.Barony
    BKrapina克拉皮纳() 	feud.Barony
    BLepoglava莱波格拉瓦() 	feud.Barony
    BLudbreg卢德布雷格() 	feud.Barony
    BOroslavje奥罗斯拉维() 	feud.Barony
    BToplice托普利采() 	feud.Barony
    BVarazdin瓦拉日丁() 	feud.Barony
}

type 瓦拉日丁VaradzinCounty struct {
	feud.BaseCounty
	查科韦茨Cakovec 	feud.Barony
	下斯图比察Donjastubica 	feud.Barony
	克拉皮纳Krapina 	feud.Barony
	莱波格拉瓦Lepoglava 	feud.Barony
	卢德布雷格Ludbreg 	feud.Barony
	奥罗斯拉维Oroslavje 	feud.Barony
	托普利采Toplice 	feud.Barony
	瓦拉日丁Varazdin 	feud.Barony
}

func (c *瓦拉日丁VaradzinCounty) BCakovec查科韦茨() feud.Barony {
	return c.查科韦茨Cakovec
}
    
func (c *瓦拉日丁VaradzinCounty) BDonjastubica下斯图比察() feud.Barony {
	return c.下斯图比察Donjastubica
}
    
func (c *瓦拉日丁VaradzinCounty) BKrapina克拉皮纳() feud.Barony {
	return c.克拉皮纳Krapina
}
    
func (c *瓦拉日丁VaradzinCounty) BLepoglava莱波格拉瓦() feud.Barony {
	return c.莱波格拉瓦Lepoglava
}
    
func (c *瓦拉日丁VaradzinCounty) BLudbreg卢德布雷格() feud.Barony {
	return c.卢德布雷格Ludbreg
}
    
func (c *瓦拉日丁VaradzinCounty) BOroslavje奥罗斯拉维() feud.Barony {
	return c.奥罗斯拉维Oroslavje
}
    
func (c *瓦拉日丁VaradzinCounty) BToplice托普利采() feud.Barony {
	return c.托普利采Toplice
}
    
func (c *瓦拉日丁VaradzinCounty) BVarazdin瓦拉日丁() feud.Barony {
	return c.瓦拉日丁Varazdin
}
    
var CVaradzin瓦拉日丁 VaradzinCounty = &瓦拉日丁VaradzinCounty{}

func init() {
	f := CVaradzin瓦拉日丁.(*瓦拉日丁VaradzinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "460",
		Title:     "varadzin",
		TitleName: "瓦拉日丁",
		TitleCode: "c_varadzin",
		Baronies:  map[string]feud.Barony{},
	}

	f.查科韦茨Cakovec = BCakovec查科韦茨
	f.查科韦茨Cakovec.SetParent(f)
	
	f.下斯图比察Donjastubica = BDonjastubica下斯图比察
	f.下斯图比察Donjastubica.SetParent(f)
	
	f.克拉皮纳Krapina = BKrapina克拉皮纳
	f.克拉皮纳Krapina.SetParent(f)
	
	f.莱波格拉瓦Lepoglava = BLepoglava莱波格拉瓦
	f.莱波格拉瓦Lepoglava.SetParent(f)
	
	f.卢德布雷格Ludbreg = BLudbreg卢德布雷格
	f.卢德布雷格Ludbreg.SetParent(f)
	
	f.奥罗斯拉维Oroslavje = BOroslavje奥罗斯拉维
	f.奥罗斯拉维Oroslavje.SetParent(f)
	
	f.托普利采Toplice = BToplice托普利采
	f.托普利采Toplice.SetParent(f)
	
	f.瓦拉日丁Varazdin = BVarazdin瓦拉日丁
	f.瓦拉日丁Varazdin.SetParent(f)
	
}
