package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TamraliptiCounty interface {
    feud.County
    BChandraketugarh旃陀罗计都姞利呬() 	feud.Barony
    BIshwaripur伊湿伐梨城() 	feud.Barony
    BJahajghata耶诃耶迦陀() 	feud.Barony
    BMeltraipara摩帝利阿波罗() 	feud.Barony
    BNarota那卢多() 	feud.Barony
    BSagardwip娑竭罗洲() 	feud.Barony
    BTamralipti多摩梨帝() 	feud.Barony
}

type 多摩梨帝TamraliptiCounty struct {
	feud.BaseCounty
	旃陀罗计都姞利呬Chandraketugarh 	feud.Barony
	伊湿伐梨城Ishwaripur 	feud.Barony
	耶诃耶迦陀Jahajghata 	feud.Barony
	摩帝利阿波罗Meltraipara 	feud.Barony
	那卢多Narota 	feud.Barony
	娑竭罗洲Sagardwip 	feud.Barony
	多摩梨帝Tamralipti 	feud.Barony
}

func (c *多摩梨帝TamraliptiCounty) BChandraketugarh旃陀罗计都姞利呬() feud.Barony {
	return c.旃陀罗计都姞利呬Chandraketugarh
}
    
func (c *多摩梨帝TamraliptiCounty) BIshwaripur伊湿伐梨城() feud.Barony {
	return c.伊湿伐梨城Ishwaripur
}
    
func (c *多摩梨帝TamraliptiCounty) BJahajghata耶诃耶迦陀() feud.Barony {
	return c.耶诃耶迦陀Jahajghata
}
    
func (c *多摩梨帝TamraliptiCounty) BMeltraipara摩帝利阿波罗() feud.Barony {
	return c.摩帝利阿波罗Meltraipara
}
    
func (c *多摩梨帝TamraliptiCounty) BNarota那卢多() feud.Barony {
	return c.那卢多Narota
}
    
func (c *多摩梨帝TamraliptiCounty) BSagardwip娑竭罗洲() feud.Barony {
	return c.娑竭罗洲Sagardwip
}
    
func (c *多摩梨帝TamraliptiCounty) BTamralipti多摩梨帝() feud.Barony {
	return c.多摩梨帝Tamralipti
}
    
var CTamralipti多摩梨帝 TamraliptiCounty = &多摩梨帝TamraliptiCounty{}

func init() {
	f := CTamralipti多摩梨帝.(*多摩梨帝TamraliptiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1235",
		Title:     "tamralipti",
		TitleName: "多摩梨帝",
		TitleCode: "c_tamralipti",
		Baronies:  map[string]feud.Barony{},
	}

	f.旃陀罗计都姞利呬Chandraketugarh = BChandraketugarh旃陀罗计都姞利呬
	f.旃陀罗计都姞利呬Chandraketugarh.SetParent(f)
	
	f.伊湿伐梨城Ishwaripur = BIshwaripur伊湿伐梨城
	f.伊湿伐梨城Ishwaripur.SetParent(f)
	
	f.耶诃耶迦陀Jahajghata = BJahajghata耶诃耶迦陀
	f.耶诃耶迦陀Jahajghata.SetParent(f)
	
	f.摩帝利阿波罗Meltraipara = BMeltraipara摩帝利阿波罗
	f.摩帝利阿波罗Meltraipara.SetParent(f)
	
	f.那卢多Narota = BNarota那卢多
	f.那卢多Narota.SetParent(f)
	
	f.娑竭罗洲Sagardwip = BSagardwip娑竭罗洲
	f.娑竭罗洲Sagardwip.SetParent(f)
	
	f.多摩梨帝Tamralipti = BTamralipti多摩梨帝
	f.多摩梨帝Tamralipti.SetParent(f)
	
}
