package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhaybarCounty interface {
    feud.County
    BAlgiar阿勒贾尔() 	feud.Barony
    BBadr拜德尔() 	feud.Barony
    BHaura豪拉() 	feud.Barony
    BKhaybar海拜尔() 	feud.Barony
    BRabig拉比格() 	feud.Barony
    BSoridan苏里丹() 	feud.Barony
    BYanbu延布() 	feud.Barony
}

type 海拜尔KhaybarCounty struct {
	feud.BaseCounty
	阿勒贾尔Algiar 	feud.Barony
	拜德尔Badr 	feud.Barony
	豪拉Haura 	feud.Barony
	海拜尔Khaybar 	feud.Barony
	拉比格Rabig 	feud.Barony
	苏里丹Soridan 	feud.Barony
	延布Yanbu 	feud.Barony
}

func (c *海拜尔KhaybarCounty) BAlgiar阿勒贾尔() feud.Barony {
	return c.阿勒贾尔Algiar
}
    
func (c *海拜尔KhaybarCounty) BBadr拜德尔() feud.Barony {
	return c.拜德尔Badr
}
    
func (c *海拜尔KhaybarCounty) BHaura豪拉() feud.Barony {
	return c.豪拉Haura
}
    
func (c *海拜尔KhaybarCounty) BKhaybar海拜尔() feud.Barony {
	return c.海拜尔Khaybar
}
    
func (c *海拜尔KhaybarCounty) BRabig拉比格() feud.Barony {
	return c.拉比格Rabig
}
    
func (c *海拜尔KhaybarCounty) BSoridan苏里丹() feud.Barony {
	return c.苏里丹Soridan
}
    
func (c *海拜尔KhaybarCounty) BYanbu延布() feud.Barony {
	return c.延布Yanbu
}
    
var CKhaybar海拜尔 KhaybarCounty = &海拜尔KhaybarCounty{}

func init() {
	f := CKhaybar海拜尔.(*海拜尔KhaybarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1536",
		Title:     "khaybar",
		TitleName: "海拜尔",
		TitleCode: "c_khaybar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿勒贾尔Algiar = BAlgiar阿勒贾尔
	f.阿勒贾尔Algiar.SetParent(f)
	
	f.拜德尔Badr = BBadr拜德尔
	f.拜德尔Badr.SetParent(f)
	
	f.豪拉Haura = BHaura豪拉
	f.豪拉Haura.SetParent(f)
	
	f.海拜尔Khaybar = BKhaybar海拜尔
	f.海拜尔Khaybar.SetParent(f)
	
	f.拉比格Rabig = BRabig拉比格
	f.拉比格Rabig.SetParent(f)
	
	f.苏里丹Soridan = BSoridan苏里丹
	f.苏里丹Soridan.SetParent(f)
	
	f.延布Yanbu = BYanbu延布
	f.延布Yanbu.SetParent(f)
	
}
