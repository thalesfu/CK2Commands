package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KudymkarCounty interface {
    feud.County
    BGayny盖内() 	feud.Barony
    BKirs基尔斯() 	feud.Barony
    BKudymkar库德姆卡尔() 	feud.Barony
    BLesnoy列斯诺伊() 	feud.Barony
    BOmotninsk奥穆特宁斯克() 	feud.Barony
    BSeyva谢伊瓦() 	feud.Barony
    BYurla尤尔拉() 	feud.Barony
}

type 库德姆卡尔KudymkarCounty struct {
	feud.BaseCounty
	盖内Gayny 	feud.Barony
	基尔斯Kirs 	feud.Barony
	库德姆卡尔Kudymkar 	feud.Barony
	列斯诺伊Lesnoy 	feud.Barony
	奥穆特宁斯克Omotninsk 	feud.Barony
	谢伊瓦Seyva 	feud.Barony
	尤尔拉Yurla 	feud.Barony
}

func (c *库德姆卡尔KudymkarCounty) BGayny盖内() feud.Barony {
	return c.盖内Gayny
}
    
func (c *库德姆卡尔KudymkarCounty) BKirs基尔斯() feud.Barony {
	return c.基尔斯Kirs
}
    
func (c *库德姆卡尔KudymkarCounty) BKudymkar库德姆卡尔() feud.Barony {
	return c.库德姆卡尔Kudymkar
}
    
func (c *库德姆卡尔KudymkarCounty) BLesnoy列斯诺伊() feud.Barony {
	return c.列斯诺伊Lesnoy
}
    
func (c *库德姆卡尔KudymkarCounty) BOmotninsk奥穆特宁斯克() feud.Barony {
	return c.奥穆特宁斯克Omotninsk
}
    
func (c *库德姆卡尔KudymkarCounty) BSeyva谢伊瓦() feud.Barony {
	return c.谢伊瓦Seyva
}
    
func (c *库德姆卡尔KudymkarCounty) BYurla尤尔拉() feud.Barony {
	return c.尤尔拉Yurla
}
    
var CKudymkar库德姆卡尔 KudymkarCounty = &库德姆卡尔KudymkarCounty{}

func init() {
	f := CKudymkar库德姆卡尔.(*库德姆卡尔KudymkarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1720",
		Title:     "kudymkar",
		TitleName: "库德姆卡尔",
		TitleCode: "c_kudymkar",
		Baronies:  map[string]feud.Barony{},
	}

	f.盖内Gayny = BGayny盖内
	f.盖内Gayny.SetParent(f)
	
	f.基尔斯Kirs = BKirs基尔斯
	f.基尔斯Kirs.SetParent(f)
	
	f.库德姆卡尔Kudymkar = BKudymkar库德姆卡尔
	f.库德姆卡尔Kudymkar.SetParent(f)
	
	f.列斯诺伊Lesnoy = BLesnoy列斯诺伊
	f.列斯诺伊Lesnoy.SetParent(f)
	
	f.奥穆特宁斯克Omotninsk = BOmotninsk奥穆特宁斯克
	f.奥穆特宁斯克Omotninsk.SetParent(f)
	
	f.谢伊瓦Seyva = BSeyva谢伊瓦
	f.谢伊瓦Seyva.SetParent(f)
	
	f.尤尔拉Yurla = BYurla尤尔拉
	f.尤尔拉Yurla.SetParent(f)
	
}
