package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UjjayiniCounty interface {
    feud.County
    BAgar阿加尔() 	feud.Barony
    BAkodia阿戈迪亚() 	feud.Barony
    BDewas提婆思() 	feud.Barony
    BMahakaleshwar摩诃迦丽湿伐罗() 	feud.Barony
    BMaksi摩克西() 	feud.Barony
    BMehidpur摩醯陀补罗() 	feud.Barony
    BUjjayini邬阇衍那() 	feud.Barony
}

type 邬阇衍那UjjayiniCounty struct {
	feud.BaseCounty
	阿加尔Agar 	feud.Barony
	阿戈迪亚Akodia 	feud.Barony
	提婆思Dewas 	feud.Barony
	摩诃迦丽湿伐罗Mahakaleshwar 	feud.Barony
	摩克西Maksi 	feud.Barony
	摩醯陀补罗Mehidpur 	feud.Barony
	邬阇衍那Ujjayini 	feud.Barony
}

func (c *邬阇衍那UjjayiniCounty) BAgar阿加尔() feud.Barony {
	return c.阿加尔Agar
}
    
func (c *邬阇衍那UjjayiniCounty) BAkodia阿戈迪亚() feud.Barony {
	return c.阿戈迪亚Akodia
}
    
func (c *邬阇衍那UjjayiniCounty) BDewas提婆思() feud.Barony {
	return c.提婆思Dewas
}
    
func (c *邬阇衍那UjjayiniCounty) BMahakaleshwar摩诃迦丽湿伐罗() feud.Barony {
	return c.摩诃迦丽湿伐罗Mahakaleshwar
}
    
func (c *邬阇衍那UjjayiniCounty) BMaksi摩克西() feud.Barony {
	return c.摩克西Maksi
}
    
func (c *邬阇衍那UjjayiniCounty) BMehidpur摩醯陀补罗() feud.Barony {
	return c.摩醯陀补罗Mehidpur
}
    
func (c *邬阇衍那UjjayiniCounty) BUjjayini邬阇衍那() feud.Barony {
	return c.邬阇衍那Ujjayini
}
    
var CUjjayini邬阇衍那 UjjayiniCounty = &邬阇衍那UjjayiniCounty{}

func init() {
	f := CUjjayini邬阇衍那.(*邬阇衍那UjjayiniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1288",
		Title:     "ujjayini",
		TitleName: "邬阇衍那",
		TitleCode: "c_ujjayini",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿加尔Agar = BAgar阿加尔
	f.阿加尔Agar.SetParent(f)
	
	f.阿戈迪亚Akodia = BAkodia阿戈迪亚
	f.阿戈迪亚Akodia.SetParent(f)
	
	f.提婆思Dewas = BDewas提婆思
	f.提婆思Dewas.SetParent(f)
	
	f.摩诃迦丽湿伐罗Mahakaleshwar = BMahakaleshwar摩诃迦丽湿伐罗
	f.摩诃迦丽湿伐罗Mahakaleshwar.SetParent(f)
	
	f.摩克西Maksi = BMaksi摩克西
	f.摩克西Maksi.SetParent(f)
	
	f.摩醯陀补罗Mehidpur = BMehidpur摩醯陀补罗
	f.摩醯陀补罗Mehidpur.SetParent(f)
	
	f.邬阇衍那Ujjayini = BUjjayini邬阇衍那
	f.邬阇衍那Ujjayini.SetParent(f)
	
}
