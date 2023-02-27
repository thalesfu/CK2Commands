package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JharkandCounty interface {
    feud.County
    BAmbikapur菴毗迦城() 	feud.Barony
    BBhagva婆伽缚() 	feud.Barony
    BBhambiya房毗耶() 	feud.Barony
    BBhasma婆娑摩() 	feud.Barony
    BDhamtari陀姆多利() 	feud.Barony
    BKandra建陀罗() 	feud.Barony
    BKoriya憍利耶() 	feud.Barony
}

type 阇罗犍度JharkandCounty struct {
	feud.BaseCounty
	菴毗迦城Ambikapur 	feud.Barony
	婆伽缚Bhagva 	feud.Barony
	房毗耶Bhambiya 	feud.Barony
	婆娑摩Bhasma 	feud.Barony
	陀姆多利Dhamtari 	feud.Barony
	建陀罗Kandra 	feud.Barony
	憍利耶Koriya 	feud.Barony
}

func (c *阇罗犍度JharkandCounty) BAmbikapur菴毗迦城() feud.Barony {
	return c.菴毗迦城Ambikapur
}
    
func (c *阇罗犍度JharkandCounty) BBhagva婆伽缚() feud.Barony {
	return c.婆伽缚Bhagva
}
    
func (c *阇罗犍度JharkandCounty) BBhambiya房毗耶() feud.Barony {
	return c.房毗耶Bhambiya
}
    
func (c *阇罗犍度JharkandCounty) BBhasma婆娑摩() feud.Barony {
	return c.婆娑摩Bhasma
}
    
func (c *阇罗犍度JharkandCounty) BDhamtari陀姆多利() feud.Barony {
	return c.陀姆多利Dhamtari
}
    
func (c *阇罗犍度JharkandCounty) BKandra建陀罗() feud.Barony {
	return c.建陀罗Kandra
}
    
func (c *阇罗犍度JharkandCounty) BKoriya憍利耶() feud.Barony {
	return c.憍利耶Koriya
}
    
var CJharkand阇罗犍度 JharkandCounty = &阇罗犍度JharkandCounty{}

func init() {
	f := CJharkand阇罗犍度.(*阇罗犍度JharkandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1241",
		Title:     "jharkand",
		TitleName: "阇罗犍度",
		TitleCode: "c_jharkand",
		Baronies:  map[string]feud.Barony{},
	}

	f.菴毗迦城Ambikapur = BAmbikapur菴毗迦城
	f.菴毗迦城Ambikapur.SetParent(f)
	
	f.婆伽缚Bhagva = BBhagva婆伽缚
	f.婆伽缚Bhagva.SetParent(f)
	
	f.房毗耶Bhambiya = BBhambiya房毗耶
	f.房毗耶Bhambiya.SetParent(f)
	
	f.婆娑摩Bhasma = BBhasma婆娑摩
	f.婆娑摩Bhasma.SetParent(f)
	
	f.陀姆多利Dhamtari = BDhamtari陀姆多利
	f.陀姆多利Dhamtari.SetParent(f)
	
	f.建陀罗Kandra = BKandra建陀罗
	f.建陀罗Kandra.SetParent(f)
	
	f.憍利耶Koriya = BKoriya憍利耶
	f.憍利耶Koriya.SetParent(f)
	
}
