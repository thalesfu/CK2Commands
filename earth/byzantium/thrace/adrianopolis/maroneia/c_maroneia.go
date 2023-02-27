package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaroneiaCounty interface {
    feud.County
    BAnaktoropolis阿纳克托罗波利斯() 	feud.Barony
    BKomotini科莫蒂尼() 	feud.Barony
    BMaroneia马戎尼亚() 	feud.Barony
    BMosynopolis莫西诺波利斯() 	feud.Barony
    BPeritheorion佩里塞奥里翁() 	feud.Barony
    BPolystylon波利斯提伦() 	feud.Barony
    BXantheia桑德雅() 	feud.Barony
}

type 马戎尼亚MaroneiaCounty struct {
	feud.BaseCounty
	阿纳克托罗波利斯Anaktoropolis 	feud.Barony
	科莫蒂尼Komotini 	feud.Barony
	马戎尼亚Maroneia 	feud.Barony
	莫西诺波利斯Mosynopolis 	feud.Barony
	佩里塞奥里翁Peritheorion 	feud.Barony
	波利斯提伦Polystylon 	feud.Barony
	桑德雅Xantheia 	feud.Barony
}

func (c *马戎尼亚MaroneiaCounty) BAnaktoropolis阿纳克托罗波利斯() feud.Barony {
	return c.阿纳克托罗波利斯Anaktoropolis
}
    
func (c *马戎尼亚MaroneiaCounty) BKomotini科莫蒂尼() feud.Barony {
	return c.科莫蒂尼Komotini
}
    
func (c *马戎尼亚MaroneiaCounty) BMaroneia马戎尼亚() feud.Barony {
	return c.马戎尼亚Maroneia
}
    
func (c *马戎尼亚MaroneiaCounty) BMosynopolis莫西诺波利斯() feud.Barony {
	return c.莫西诺波利斯Mosynopolis
}
    
func (c *马戎尼亚MaroneiaCounty) BPeritheorion佩里塞奥里翁() feud.Barony {
	return c.佩里塞奥里翁Peritheorion
}
    
func (c *马戎尼亚MaroneiaCounty) BPolystylon波利斯提伦() feud.Barony {
	return c.波利斯提伦Polystylon
}
    
func (c *马戎尼亚MaroneiaCounty) BXantheia桑德雅() feud.Barony {
	return c.桑德雅Xantheia
}
    
var CMaroneia马戎尼亚 MaroneiaCounty = &马戎尼亚MaroneiaCounty{}

func init() {
	f := CMaroneia马戎尼亚.(*马戎尼亚MaroneiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1885",
		Title:     "maroneia",
		TitleName: "马戎尼亚",
		TitleCode: "c_maroneia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿纳克托罗波利斯Anaktoropolis = BAnaktoropolis阿纳克托罗波利斯
	f.阿纳克托罗波利斯Anaktoropolis.SetParent(f)
	
	f.科莫蒂尼Komotini = BKomotini科莫蒂尼
	f.科莫蒂尼Komotini.SetParent(f)
	
	f.马戎尼亚Maroneia = BMaroneia马戎尼亚
	f.马戎尼亚Maroneia.SetParent(f)
	
	f.莫西诺波利斯Mosynopolis = BMosynopolis莫西诺波利斯
	f.莫西诺波利斯Mosynopolis.SetParent(f)
	
	f.佩里塞奥里翁Peritheorion = BPeritheorion佩里塞奥里翁
	f.佩里塞奥里翁Peritheorion.SetParent(f)
	
	f.波利斯提伦Polystylon = BPolystylon波利斯提伦
	f.波利斯提伦Polystylon.SetParent(f)
	
	f.桑德雅Xantheia = BXantheia桑德雅
	f.桑德雅Xantheia.SetParent(f)
	
}
