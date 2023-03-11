package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PeremyshlCounty interface {
    feud.County
    BJaroslav雅罗斯拉夫() 	feud.Barony
    BLezhaisk莱扎伊斯克() 	feud.Barony
    BLiubachiv柳巴维奇() 	feud.Barony
    BPeremyshl佩列梅什利() 	feud.Barony
    BRasiv拉西夫() 	feud.Barony
    BSanok萨诺克() 	feud.Barony
    BUstryky_dolishni下乌斯奇基() 	feud.Barony
}

type 佩列梅什利PeremyshlCounty struct {
	feud.BaseCounty
	雅罗斯拉夫Jaroslav 	feud.Barony
	莱扎伊斯克Lezhaisk 	feud.Barony
	柳巴维奇Liubachiv 	feud.Barony
	佩列梅什利Peremyshl 	feud.Barony
	拉西夫Rasiv 	feud.Barony
	萨诺克Sanok 	feud.Barony
	下乌斯奇基Ustryky_dolishni 	feud.Barony
}

func (c *佩列梅什利PeremyshlCounty) BJaroslav雅罗斯拉夫() feud.Barony {
	return c.雅罗斯拉夫Jaroslav
}
    
func (c *佩列梅什利PeremyshlCounty) BLezhaisk莱扎伊斯克() feud.Barony {
	return c.莱扎伊斯克Lezhaisk
}
    
func (c *佩列梅什利PeremyshlCounty) BLiubachiv柳巴维奇() feud.Barony {
	return c.柳巴维奇Liubachiv
}
    
func (c *佩列梅什利PeremyshlCounty) BPeremyshl佩列梅什利() feud.Barony {
	return c.佩列梅什利Peremyshl
}
    
func (c *佩列梅什利PeremyshlCounty) BRasiv拉西夫() feud.Barony {
	return c.拉西夫Rasiv
}
    
func (c *佩列梅什利PeremyshlCounty) BSanok萨诺克() feud.Barony {
	return c.萨诺克Sanok
}
    
func (c *佩列梅什利PeremyshlCounty) BUstryky_dolishni下乌斯奇基() feud.Barony {
	return c.下乌斯奇基Ustryky_dolishni
}
    
var CPeremyshl佩列梅什利 PeremyshlCounty = &佩列梅什利PeremyshlCounty{}

func init() {
	f := CPeremyshl佩列梅什利.(*佩列梅什利PeremyshlCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "534",
		Title:     "peremyshl",
		TitleName: "佩列梅什利",
		TitleCode: "c_peremyshl",
		Baronies:  map[string]feud.Barony{},
	}

	f.雅罗斯拉夫Jaroslav = BJaroslav雅罗斯拉夫
	f.雅罗斯拉夫Jaroslav.SetParent(f)
	
	f.莱扎伊斯克Lezhaisk = BLezhaisk莱扎伊斯克
	f.莱扎伊斯克Lezhaisk.SetParent(f)
	
	f.柳巴维奇Liubachiv = BLiubachiv柳巴维奇
	f.柳巴维奇Liubachiv.SetParent(f)
	
	f.佩列梅什利Peremyshl = BPeremyshl佩列梅什利
	f.佩列梅什利Peremyshl.SetParent(f)
	
	f.拉西夫Rasiv = BRasiv拉西夫
	f.拉西夫Rasiv.SetParent(f)
	
	f.萨诺克Sanok = BSanok萨诺克
	f.萨诺克Sanok.SetParent(f)
	
	f.下乌斯奇基Ustryky_dolishni = BUstryky_dolishni下乌斯奇基
	f.下乌斯奇基Ustryky_dolishni.SetParent(f)
	
}
