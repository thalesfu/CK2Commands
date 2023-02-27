package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MassatCounty interface {
    feud.County
    BAmerchen埃迈尔尚() 	feud.Barony
    BAnfa安法() 	feud.Barony
    BAzammut艾宰穆尔() 	feud.Barony
    BEl_bakra贝克雷() 	feud.Barony
    BFadala法达拉() 	feud.Barony
    BMazaghan马扎汗() 	feud.Barony
    BZidania齐达尼亚() 	feud.Barony
}

type 安法MassatCounty struct {
	feud.BaseCounty
	埃迈尔尚Amerchen 	feud.Barony
	安法Anfa 	feud.Barony
	艾宰穆尔Azammut 	feud.Barony
	贝克雷El_bakra 	feud.Barony
	法达拉Fadala 	feud.Barony
	马扎汗Mazaghan 	feud.Barony
	齐达尼亚Zidania 	feud.Barony
}

func (c *安法MassatCounty) BAmerchen埃迈尔尚() feud.Barony {
	return c.埃迈尔尚Amerchen
}
    
func (c *安法MassatCounty) BAnfa安法() feud.Barony {
	return c.安法Anfa
}
    
func (c *安法MassatCounty) BAzammut艾宰穆尔() feud.Barony {
	return c.艾宰穆尔Azammut
}
    
func (c *安法MassatCounty) BEl_bakra贝克雷() feud.Barony {
	return c.贝克雷El_bakra
}
    
func (c *安法MassatCounty) BFadala法达拉() feud.Barony {
	return c.法达拉Fadala
}
    
func (c *安法MassatCounty) BMazaghan马扎汗() feud.Barony {
	return c.马扎汗Mazaghan
}
    
func (c *安法MassatCounty) BZidania齐达尼亚() feud.Barony {
	return c.齐达尼亚Zidania
}
    
var CMassat安法 MassatCounty = &安法MassatCounty{}

func init() {
	f := CMassat安法.(*安法MassatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "844",
		Title:     "massat",
		TitleName: "安法",
		TitleCode: "c_massat",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃迈尔尚Amerchen = BAmerchen埃迈尔尚
	f.埃迈尔尚Amerchen.SetParent(f)
	
	f.安法Anfa = BAnfa安法
	f.安法Anfa.SetParent(f)
	
	f.艾宰穆尔Azammut = BAzammut艾宰穆尔
	f.艾宰穆尔Azammut.SetParent(f)
	
	f.贝克雷El_bakra = BEl_bakra贝克雷
	f.贝克雷El_bakra.SetParent(f)
	
	f.法达拉Fadala = BFadala法达拉
	f.法达拉Fadala.SetParent(f)
	
	f.马扎汗Mazaghan = BMazaghan马扎汗
	f.马扎汗Mazaghan.SetParent(f)
	
	f.齐达尼亚Zidania = BZidania齐达尼亚
	f.齐达尼亚Zidania.SetParent(f)
	
}
