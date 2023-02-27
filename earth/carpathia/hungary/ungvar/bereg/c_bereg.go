package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeregCounty interface {
    feud.County
    BBeregszasz拜赖格萨斯() 	feud.Barony
    BIlosva伊洛什沃() 	feud.Barony
    BKapos考波什() 	feud.Barony
    BMunkacs蒙卡奇() 	feud.Barony
    BPerecseny派赖切尼() 	feud.Barony
    BSzobranc索布兰茨() 	feud.Barony
    BSzolyva索伊沃() 	feud.Barony
    BUngvar温格瓦尔() 	feud.Barony
}

type 拜赖格BeregCounty struct {
	feud.BaseCounty
	拜赖格萨斯Beregszasz 	feud.Barony
	伊洛什沃Ilosva 	feud.Barony
	考波什Kapos 	feud.Barony
	蒙卡奇Munkacs 	feud.Barony
	派赖切尼Perecseny 	feud.Barony
	索布兰茨Szobranc 	feud.Barony
	索伊沃Szolyva 	feud.Barony
	温格瓦尔Ungvar 	feud.Barony
}

func (c *拜赖格BeregCounty) BBeregszasz拜赖格萨斯() feud.Barony {
	return c.拜赖格萨斯Beregszasz
}
    
func (c *拜赖格BeregCounty) BIlosva伊洛什沃() feud.Barony {
	return c.伊洛什沃Ilosva
}
    
func (c *拜赖格BeregCounty) BKapos考波什() feud.Barony {
	return c.考波什Kapos
}
    
func (c *拜赖格BeregCounty) BMunkacs蒙卡奇() feud.Barony {
	return c.蒙卡奇Munkacs
}
    
func (c *拜赖格BeregCounty) BPerecseny派赖切尼() feud.Barony {
	return c.派赖切尼Perecseny
}
    
func (c *拜赖格BeregCounty) BSzobranc索布兰茨() feud.Barony {
	return c.索布兰茨Szobranc
}
    
func (c *拜赖格BeregCounty) BSzolyva索伊沃() feud.Barony {
	return c.索伊沃Szolyva
}
    
func (c *拜赖格BeregCounty) BUngvar温格瓦尔() feud.Barony {
	return c.温格瓦尔Ungvar
}
    
var CBereg拜赖格 BeregCounty = &拜赖格BeregCounty{}

func init() {
	f := CBereg拜赖格.(*拜赖格BeregCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "537",
		Title:     "bereg",
		TitleName: "拜赖格",
		TitleCode: "c_bereg",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜赖格萨斯Beregszasz = BBeregszasz拜赖格萨斯
	f.拜赖格萨斯Beregszasz.SetParent(f)
	
	f.伊洛什沃Ilosva = BIlosva伊洛什沃
	f.伊洛什沃Ilosva.SetParent(f)
	
	f.考波什Kapos = BKapos考波什
	f.考波什Kapos.SetParent(f)
	
	f.蒙卡奇Munkacs = BMunkacs蒙卡奇
	f.蒙卡奇Munkacs.SetParent(f)
	
	f.派赖切尼Perecseny = BPerecseny派赖切尼
	f.派赖切尼Perecseny.SetParent(f)
	
	f.索布兰茨Szobranc = BSzobranc索布兰茨
	f.索布兰茨Szobranc.SetParent(f)
	
	f.索伊沃Szolyva = BSzolyva索伊沃
	f.索伊沃Szolyva.SetParent(f)
	
	f.温格瓦尔Ungvar = BUngvar温格瓦尔
	f.温格瓦尔Ungvar.SetParent(f)
	
}
