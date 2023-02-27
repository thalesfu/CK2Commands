package uvs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UvsCounty interface {
    feud.County
    BDavst达布斯特() 	feud.Barony
    BErzin额尔逊() 	feud.Barony
    BSagil萨吉勒() 	feud.Barony
    BTes特斯() 	feud.Barony
    BTsalgar察勒噶尔() 	feud.Barony
    BUlaangom乌兰固木() 	feud.Barony
    BUvs乌布苏() 	feud.Barony
}

type 乌布苏UvsCounty struct {
	feud.BaseCounty
	达布斯特Davst 	feud.Barony
	额尔逊Erzin 	feud.Barony
	萨吉勒Sagil 	feud.Barony
	特斯Tes 	feud.Barony
	察勒噶尔Tsalgar 	feud.Barony
	乌兰固木Ulaangom 	feud.Barony
	乌布苏Uvs 	feud.Barony
}

func (c *乌布苏UvsCounty) BDavst达布斯特() feud.Barony {
	return c.达布斯特Davst
}
    
func (c *乌布苏UvsCounty) BErzin额尔逊() feud.Barony {
	return c.额尔逊Erzin
}
    
func (c *乌布苏UvsCounty) BSagil萨吉勒() feud.Barony {
	return c.萨吉勒Sagil
}
    
func (c *乌布苏UvsCounty) BTes特斯() feud.Barony {
	return c.特斯Tes
}
    
func (c *乌布苏UvsCounty) BTsalgar察勒噶尔() feud.Barony {
	return c.察勒噶尔Tsalgar
}
    
func (c *乌布苏UvsCounty) BUlaangom乌兰固木() feud.Barony {
	return c.乌兰固木Ulaangom
}
    
func (c *乌布苏UvsCounty) BUvs乌布苏() feud.Barony {
	return c.乌布苏Uvs
}
    
var CUvs乌布苏 UvsCounty = &乌布苏UvsCounty{}

func init() {
	f := CUvs乌布苏.(*乌布苏UvsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1906",
		Title:     "uvs",
		TitleName: "乌布苏",
		TitleCode: "c_uvs",
		Baronies:  map[string]feud.Barony{},
	}

	f.达布斯特Davst = BDavst达布斯特
	f.达布斯特Davst.SetParent(f)
	
	f.额尔逊Erzin = BErzin额尔逊
	f.额尔逊Erzin.SetParent(f)
	
	f.萨吉勒Sagil = BSagil萨吉勒
	f.萨吉勒Sagil.SetParent(f)
	
	f.特斯Tes = BTes特斯
	f.特斯Tes.SetParent(f)
	
	f.察勒噶尔Tsalgar = BTsalgar察勒噶尔
	f.察勒噶尔Tsalgar.SetParent(f)
	
	f.乌兰固木Ulaangom = BUlaangom乌兰固木
	f.乌兰固木Ulaangom.SetParent(f)
	
	f.乌布苏Uvs = BUvs乌布苏
	f.乌布苏Uvs.SetParent(f)
	
}
