package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GurganCounty interface {
    feud.County
    BAqqala阿格盖拉() 	feud.Barony
    BGonbadeqabus贡巴德卡武斯() 	feud.Barony
    BGurgan戈尔甘() 	feud.Barony
    BKhanbebin汗贝宾() 	feud.Barony
    BKomishdepa科米什德帕() 	feud.Barony
    BKordkuy科尔德库伊() 	feud.Barony
    BMinudasht米努达什特() 	feud.Barony
    BRamian拉米扬() 	feud.Barony
}

type 戈尔甘GurganCounty struct {
	feud.BaseCounty
	阿格盖拉Aqqala 	feud.Barony
	贡巴德卡武斯Gonbadeqabus 	feud.Barony
	戈尔甘Gurgan 	feud.Barony
	汗贝宾Khanbebin 	feud.Barony
	科米什德帕Komishdepa 	feud.Barony
	科尔德库伊Kordkuy 	feud.Barony
	米努达什特Minudasht 	feud.Barony
	拉米扬Ramian 	feud.Barony
}

func (c *戈尔甘GurganCounty) BAqqala阿格盖拉() feud.Barony {
	return c.阿格盖拉Aqqala
}
    
func (c *戈尔甘GurganCounty) BGonbadeqabus贡巴德卡武斯() feud.Barony {
	return c.贡巴德卡武斯Gonbadeqabus
}
    
func (c *戈尔甘GurganCounty) BGurgan戈尔甘() feud.Barony {
	return c.戈尔甘Gurgan
}
    
func (c *戈尔甘GurganCounty) BKhanbebin汗贝宾() feud.Barony {
	return c.汗贝宾Khanbebin
}
    
func (c *戈尔甘GurganCounty) BKomishdepa科米什德帕() feud.Barony {
	return c.科米什德帕Komishdepa
}
    
func (c *戈尔甘GurganCounty) BKordkuy科尔德库伊() feud.Barony {
	return c.科尔德库伊Kordkuy
}
    
func (c *戈尔甘GurganCounty) BMinudasht米努达什特() feud.Barony {
	return c.米努达什特Minudasht
}
    
func (c *戈尔甘GurganCounty) BRamian拉米扬() feud.Barony {
	return c.拉米扬Ramian
}
    
var CGurgan戈尔甘 GurganCounty = &戈尔甘GurganCounty{}

func init() {
	f := CGurgan戈尔甘.(*戈尔甘GurganCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "633",
		Title:     "gurgan",
		TitleName: "戈尔甘",
		TitleCode: "c_gurgan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格盖拉Aqqala = BAqqala阿格盖拉
	f.阿格盖拉Aqqala.SetParent(f)
	
	f.贡巴德卡武斯Gonbadeqabus = BGonbadeqabus贡巴德卡武斯
	f.贡巴德卡武斯Gonbadeqabus.SetParent(f)
	
	f.戈尔甘Gurgan = BGurgan戈尔甘
	f.戈尔甘Gurgan.SetParent(f)
	
	f.汗贝宾Khanbebin = BKhanbebin汗贝宾
	f.汗贝宾Khanbebin.SetParent(f)
	
	f.科米什德帕Komishdepa = BKomishdepa科米什德帕
	f.科米什德帕Komishdepa.SetParent(f)
	
	f.科尔德库伊Kordkuy = BKordkuy科尔德库伊
	f.科尔德库伊Kordkuy.SetParent(f)
	
	f.米努达什特Minudasht = BMinudasht米努达什特
	f.米努达什特Minudasht.SetParent(f)
	
	f.拉米扬Ramian = BRamian拉米扬
	f.拉米扬Ramian.SetParent(f)
	
}
