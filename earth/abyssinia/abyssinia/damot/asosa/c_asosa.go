package asosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsosaCounty interface {
    feud.County
    BAsosa阿索萨() 	feud.Barony
    BBambasi巴姆巴斯() 	feud.Barony
    BDebrezeyit德布雷塞特() 	feud.Barony
    BDibate迪巴特() 	feud.Barony
    BGenetemariam根特迈利() 	feud.Barony
    BKetema科特马() 	feud.Barony
    BKormuk库尔穆克() 	feud.Barony
    BMankush曼库什() 	feud.Barony
}

type 阿索萨AsosaCounty struct {
	feud.BaseCounty
	阿索萨Asosa 	feud.Barony
	巴姆巴斯Bambasi 	feud.Barony
	德布雷塞特Debrezeyit 	feud.Barony
	迪巴特Dibate 	feud.Barony
	根特迈利Genetemariam 	feud.Barony
	科特马Ketema 	feud.Barony
	库尔穆克Kormuk 	feud.Barony
	曼库什Mankush 	feud.Barony
}

func (c *阿索萨AsosaCounty) BAsosa阿索萨() feud.Barony {
	return c.阿索萨Asosa
}
    
func (c *阿索萨AsosaCounty) BBambasi巴姆巴斯() feud.Barony {
	return c.巴姆巴斯Bambasi
}
    
func (c *阿索萨AsosaCounty) BDebrezeyit德布雷塞特() feud.Barony {
	return c.德布雷塞特Debrezeyit
}
    
func (c *阿索萨AsosaCounty) BDibate迪巴特() feud.Barony {
	return c.迪巴特Dibate
}
    
func (c *阿索萨AsosaCounty) BGenetemariam根特迈利() feud.Barony {
	return c.根特迈利Genetemariam
}
    
func (c *阿索萨AsosaCounty) BKetema科特马() feud.Barony {
	return c.科特马Ketema
}
    
func (c *阿索萨AsosaCounty) BKormuk库尔穆克() feud.Barony {
	return c.库尔穆克Kormuk
}
    
func (c *阿索萨AsosaCounty) BMankush曼库什() feud.Barony {
	return c.曼库什Mankush
}
    
var CAsosa阿索萨 AsosaCounty = &阿索萨AsosaCounty{}

func init() {
	f := CAsosa阿索萨.(*阿索萨AsosaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "881",
		Title:     "asosa",
		TitleName: "阿索萨",
		TitleCode: "c_asosa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿索萨Asosa = BAsosa阿索萨
	f.阿索萨Asosa.SetParent(f)
	
	f.巴姆巴斯Bambasi = BBambasi巴姆巴斯
	f.巴姆巴斯Bambasi.SetParent(f)
	
	f.德布雷塞特Debrezeyit = BDebrezeyit德布雷塞特
	f.德布雷塞特Debrezeyit.SetParent(f)
	
	f.迪巴特Dibate = BDibate迪巴特
	f.迪巴特Dibate.SetParent(f)
	
	f.根特迈利Genetemariam = BGenetemariam根特迈利
	f.根特迈利Genetemariam.SetParent(f)
	
	f.科特马Ketema = BKetema科特马
	f.科特马Ketema.SetParent(f)
	
	f.库尔穆克Kormuk = BKormuk库尔穆克
	f.库尔穆克Kormuk.SetParent(f)
	
	f.曼库什Mankush = BMankush曼库什
	f.曼库什Mankush.SetParent(f)
	
}
