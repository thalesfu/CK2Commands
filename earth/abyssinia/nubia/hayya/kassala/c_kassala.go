package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KassalaCounty interface {
    feud.County
    BAroma阿罗马() 	feud.Barony
    BDinder丁德尔() 	feud.Barony
    BDoka多卡() 	feud.Barony
    BGherset戈尔塞特() 	feud.Barony
    BKagnarti卡格纳尔蒂() 	feud.Barony
    BKassala卡萨拉() 	feud.Barony
    BTebteb特卜特卜() 	feud.Barony
}

type 卡萨拉KassalaCounty struct {
	feud.BaseCounty
	阿罗马Aroma 	feud.Barony
	丁德尔Dinder 	feud.Barony
	多卡Doka 	feud.Barony
	戈尔塞特Gherset 	feud.Barony
	卡格纳尔蒂Kagnarti 	feud.Barony
	卡萨拉Kassala 	feud.Barony
	特卜特卜Tebteb 	feud.Barony
}

func (c *卡萨拉KassalaCounty) BAroma阿罗马() feud.Barony {
	return c.阿罗马Aroma
}
    
func (c *卡萨拉KassalaCounty) BDinder丁德尔() feud.Barony {
	return c.丁德尔Dinder
}
    
func (c *卡萨拉KassalaCounty) BDoka多卡() feud.Barony {
	return c.多卡Doka
}
    
func (c *卡萨拉KassalaCounty) BGherset戈尔塞特() feud.Barony {
	return c.戈尔塞特Gherset
}
    
func (c *卡萨拉KassalaCounty) BKagnarti卡格纳尔蒂() feud.Barony {
	return c.卡格纳尔蒂Kagnarti
}
    
func (c *卡萨拉KassalaCounty) BKassala卡萨拉() feud.Barony {
	return c.卡萨拉Kassala
}
    
func (c *卡萨拉KassalaCounty) BTebteb特卜特卜() feud.Barony {
	return c.特卜特卜Tebteb
}
    
var CKassala卡萨拉 KassalaCounty = &卡萨拉KassalaCounty{}

func init() {
	f := CKassala卡萨拉.(*卡萨拉KassalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "877",
		Title:     "kassala",
		TitleName: "卡萨拉",
		TitleCode: "c_kassala",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿罗马Aroma = BAroma阿罗马
	f.阿罗马Aroma.SetParent(f)
	
	f.丁德尔Dinder = BDinder丁德尔
	f.丁德尔Dinder.SetParent(f)
	
	f.多卡Doka = BDoka多卡
	f.多卡Doka.SetParent(f)
	
	f.戈尔塞特Gherset = BGherset戈尔塞特
	f.戈尔塞特Gherset.SetParent(f)
	
	f.卡格纳尔蒂Kagnarti = BKagnarti卡格纳尔蒂
	f.卡格纳尔蒂Kagnarti.SetParent(f)
	
	f.卡萨拉Kassala = BKassala卡萨拉
	f.卡萨拉Kassala.SetParent(f)
	
	f.特卜特卜Tebteb = BTebteb特卜特卜
	f.特卜特卜Tebteb.SetParent(f)
	
}
