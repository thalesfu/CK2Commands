package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SutaiCounty interface {
    feud.County
    BBulgan布尔干() 	feud.Barony
    BDarvi_sutai达尔维() 	feud.Barony
    BIkhes伊赫斯() 	feud.Barony
    BKhulmiin呼尔明() 	feud.Barony
    BMyangan明安() 	feud.Barony
    BSutai苏泰() 	feud.Barony
    BTonkhil吞黑勒() 	feud.Barony
}

type 苏泰SutaiCounty struct {
	feud.BaseCounty
	布尔干Bulgan 	feud.Barony
	达尔维Darvi_sutai 	feud.Barony
	伊赫斯Ikhes 	feud.Barony
	呼尔明Khulmiin 	feud.Barony
	明安Myangan 	feud.Barony
	苏泰Sutai 	feud.Barony
	吞黑勒Tonkhil 	feud.Barony
}

func (c *苏泰SutaiCounty) BBulgan布尔干() feud.Barony {
	return c.布尔干Bulgan
}
    
func (c *苏泰SutaiCounty) BDarvi_sutai达尔维() feud.Barony {
	return c.达尔维Darvi_sutai
}
    
func (c *苏泰SutaiCounty) BIkhes伊赫斯() feud.Barony {
	return c.伊赫斯Ikhes
}
    
func (c *苏泰SutaiCounty) BKhulmiin呼尔明() feud.Barony {
	return c.呼尔明Khulmiin
}
    
func (c *苏泰SutaiCounty) BMyangan明安() feud.Barony {
	return c.明安Myangan
}
    
func (c *苏泰SutaiCounty) BSutai苏泰() feud.Barony {
	return c.苏泰Sutai
}
    
func (c *苏泰SutaiCounty) BTonkhil吞黑勒() feud.Barony {
	return c.吞黑勒Tonkhil
}
    
var CSutai苏泰 SutaiCounty = &苏泰SutaiCounty{}

func init() {
	f := CSutai苏泰.(*苏泰SutaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1910",
		Title:     "sutai",
		TitleName: "苏泰",
		TitleCode: "c_sutai",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔干Bulgan = BBulgan布尔干
	f.布尔干Bulgan.SetParent(f)
	
	f.达尔维Darvi_sutai = BDarvi_sutai达尔维
	f.达尔维Darvi_sutai.SetParent(f)
	
	f.伊赫斯Ikhes = BIkhes伊赫斯
	f.伊赫斯Ikhes.SetParent(f)
	
	f.呼尔明Khulmiin = BKhulmiin呼尔明
	f.呼尔明Khulmiin.SetParent(f)
	
	f.明安Myangan = BMyangan明安
	f.明安Myangan.SetParent(f)
	
	f.苏泰Sutai = BSutai苏泰
	f.苏泰Sutai.SetParent(f)
	
	f.吞黑勒Tonkhil = BTonkhil吞黑勒
	f.吞黑勒Tonkhil.SetParent(f)
	
}
