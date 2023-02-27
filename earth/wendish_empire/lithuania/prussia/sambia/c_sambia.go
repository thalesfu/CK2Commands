package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SambiaCounty interface {
    feud.County
    BCranz克兰茨() 	feud.Barony
    BFischhausen菲施豪森() 	feud.Barony
    BFrombork弗龙堡() 	feud.Barony
    BKonigsberg柯尼斯堡() 	feud.Barony
    BLabiau拉比奥() 	feud.Barony
    BSambrandenburg勃兰登堡() 	feud.Barony
    BTapiau塔皮奥() 	feud.Barony
    BWiskiauten维斯基奥滕() 	feud.Barony
}

type 桑比亚SambiaCounty struct {
	feud.BaseCounty
	克兰茨Cranz 	feud.Barony
	菲施豪森Fischhausen 	feud.Barony
	弗龙堡Frombork 	feud.Barony
	柯尼斯堡Konigsberg 	feud.Barony
	拉比奥Labiau 	feud.Barony
	勃兰登堡Sambrandenburg 	feud.Barony
	塔皮奥Tapiau 	feud.Barony
	维斯基奥滕Wiskiauten 	feud.Barony
}

func (c *桑比亚SambiaCounty) BCranz克兰茨() feud.Barony {
	return c.克兰茨Cranz
}
    
func (c *桑比亚SambiaCounty) BFischhausen菲施豪森() feud.Barony {
	return c.菲施豪森Fischhausen
}
    
func (c *桑比亚SambiaCounty) BFrombork弗龙堡() feud.Barony {
	return c.弗龙堡Frombork
}
    
func (c *桑比亚SambiaCounty) BKonigsberg柯尼斯堡() feud.Barony {
	return c.柯尼斯堡Konigsberg
}
    
func (c *桑比亚SambiaCounty) BLabiau拉比奥() feud.Barony {
	return c.拉比奥Labiau
}
    
func (c *桑比亚SambiaCounty) BSambrandenburg勃兰登堡() feud.Barony {
	return c.勃兰登堡Sambrandenburg
}
    
func (c *桑比亚SambiaCounty) BTapiau塔皮奥() feud.Barony {
	return c.塔皮奥Tapiau
}
    
func (c *桑比亚SambiaCounty) BWiskiauten维斯基奥滕() feud.Barony {
	return c.维斯基奥滕Wiskiauten
}
    
var CSambia桑比亚 SambiaCounty = &桑比亚SambiaCounty{}

func init() {
	f := CSambia桑比亚.(*桑比亚SambiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "371",
		Title:     "sambia",
		TitleName: "桑比亚",
		TitleCode: "c_sambia",
		Baronies:  map[string]feud.Barony{},
	}

	f.克兰茨Cranz = BCranz克兰茨
	f.克兰茨Cranz.SetParent(f)
	
	f.菲施豪森Fischhausen = BFischhausen菲施豪森
	f.菲施豪森Fischhausen.SetParent(f)
	
	f.弗龙堡Frombork = BFrombork弗龙堡
	f.弗龙堡Frombork.SetParent(f)
	
	f.柯尼斯堡Konigsberg = BKonigsberg柯尼斯堡
	f.柯尼斯堡Konigsberg.SetParent(f)
	
	f.拉比奥Labiau = BLabiau拉比奥
	f.拉比奥Labiau.SetParent(f)
	
	f.勃兰登堡Sambrandenburg = BSambrandenburg勃兰登堡
	f.勃兰登堡Sambrandenburg.SetParent(f)
	
	f.塔皮奥Tapiau = BTapiau塔皮奥
	f.塔皮奥Tapiau.SetParent(f)
	
	f.维斯基奥滕Wiskiauten = BWiskiauten维斯基奥滕
	f.维斯基奥滕Wiskiauten.SetParent(f)
	
}
