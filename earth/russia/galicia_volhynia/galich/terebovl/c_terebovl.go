package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TerebovlCounty interface {
    feud.County
    BBakota巴科塔() 	feud.Barony
    BDolyna多利纳() 	feud.Barony
    BKuchelemyn库切列门() 	feud.Barony
    BMoklekov莫克列科夫() 	feud.Barony
    BMykulyn米库林() 	feud.Barony
    BOnut奥努特() 	feud.Barony
    BTerebovl捷列博夫尔() 	feud.Barony
}

type 捷列博夫尔TerebovlCounty struct {
	feud.BaseCounty
	巴科塔Bakota 	feud.Barony
	多利纳Dolyna 	feud.Barony
	库切列门Kuchelemyn 	feud.Barony
	莫克列科夫Moklekov 	feud.Barony
	米库林Mykulyn 	feud.Barony
	奥努特Onut 	feud.Barony
	捷列博夫尔Terebovl 	feud.Barony
}

func (c *捷列博夫尔TerebovlCounty) BBakota巴科塔() feud.Barony {
	return c.巴科塔Bakota
}
    
func (c *捷列博夫尔TerebovlCounty) BDolyna多利纳() feud.Barony {
	return c.多利纳Dolyna
}
    
func (c *捷列博夫尔TerebovlCounty) BKuchelemyn库切列门() feud.Barony {
	return c.库切列门Kuchelemyn
}
    
func (c *捷列博夫尔TerebovlCounty) BMoklekov莫克列科夫() feud.Barony {
	return c.莫克列科夫Moklekov
}
    
func (c *捷列博夫尔TerebovlCounty) BMykulyn米库林() feud.Barony {
	return c.米库林Mykulyn
}
    
func (c *捷列博夫尔TerebovlCounty) BOnut奥努特() feud.Barony {
	return c.奥努特Onut
}
    
func (c *捷列博夫尔TerebovlCounty) BTerebovl捷列博夫尔() feud.Barony {
	return c.捷列博夫尔Terebovl
}
    
var CTerebovl捷列博夫尔 TerebovlCounty = &捷列博夫尔TerebovlCounty{}

func init() {
	f := CTerebovl捷列博夫尔.(*捷列博夫尔TerebovlCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "546",
		Title:     "terebovl",
		TitleName: "捷列博夫尔",
		TitleCode: "c_terebovl",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴科塔Bakota = BBakota巴科塔
	f.巴科塔Bakota.SetParent(f)
	
	f.多利纳Dolyna = BDolyna多利纳
	f.多利纳Dolyna.SetParent(f)
	
	f.库切列门Kuchelemyn = BKuchelemyn库切列门
	f.库切列门Kuchelemyn.SetParent(f)
	
	f.莫克列科夫Moklekov = BMoklekov莫克列科夫
	f.莫克列科夫Moklekov.SetParent(f)
	
	f.米库林Mykulyn = BMykulyn米库林
	f.米库林Mykulyn.SetParent(f)
	
	f.奥努特Onut = BOnut奥努特
	f.奥努特Onut.SetParent(f)
	
	f.捷列博夫尔Terebovl = BTerebovl捷列博夫尔
	f.捷列博夫尔Terebovl.SetParent(f)
	
}
