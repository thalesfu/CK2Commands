package gezira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GeziraCounty interface {
    feud.County
    BDahar_jihena达哈尔吉赫纳() 	feud.Barony
    BKarkab卡尔卡布() 	feud.Barony
    BKoijam科伊贾姆() 	feud.Barony
    BShawwaf沙瓦夫() 	feud.Barony
    BShowak舒沃克() 	feud.Barony
    BTangasi唐加西() 	feud.Barony
    BWolqayt瓦尔卡亚特() 	feud.Barony
}

type 杰济拉GeziraCounty struct {
	feud.BaseCounty
	达哈尔吉赫纳Dahar_jihena 	feud.Barony
	卡尔卡布Karkab 	feud.Barony
	科伊贾姆Koijam 	feud.Barony
	沙瓦夫Shawwaf 	feud.Barony
	舒沃克Showak 	feud.Barony
	唐加西Tangasi 	feud.Barony
	瓦尔卡亚特Wolqayt 	feud.Barony
}

func (c *杰济拉GeziraCounty) BDahar_jihena达哈尔吉赫纳() feud.Barony {
	return c.达哈尔吉赫纳Dahar_jihena
}
    
func (c *杰济拉GeziraCounty) BKarkab卡尔卡布() feud.Barony {
	return c.卡尔卡布Karkab
}
    
func (c *杰济拉GeziraCounty) BKoijam科伊贾姆() feud.Barony {
	return c.科伊贾姆Koijam
}
    
func (c *杰济拉GeziraCounty) BShawwaf沙瓦夫() feud.Barony {
	return c.沙瓦夫Shawwaf
}
    
func (c *杰济拉GeziraCounty) BShowak舒沃克() feud.Barony {
	return c.舒沃克Showak
}
    
func (c *杰济拉GeziraCounty) BTangasi唐加西() feud.Barony {
	return c.唐加西Tangasi
}
    
func (c *杰济拉GeziraCounty) BWolqayt瓦尔卡亚特() feud.Barony {
	return c.瓦尔卡亚特Wolqayt
}
    
var CGezira杰济拉 GeziraCounty = &杰济拉GeziraCounty{}

func init() {
	f := CGezira杰济拉.(*杰济拉GeziraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1993",
		Title:     "gezira",
		TitleName: "杰济拉",
		TitleCode: "c_gezira",
		Baronies:  map[string]feud.Barony{},
	}

	f.达哈尔吉赫纳Dahar_jihena = BDahar_jihena达哈尔吉赫纳
	f.达哈尔吉赫纳Dahar_jihena.SetParent(f)
	
	f.卡尔卡布Karkab = BKarkab卡尔卡布
	f.卡尔卡布Karkab.SetParent(f)
	
	f.科伊贾姆Koijam = BKoijam科伊贾姆
	f.科伊贾姆Koijam.SetParent(f)
	
	f.沙瓦夫Shawwaf = BShawwaf沙瓦夫
	f.沙瓦夫Shawwaf.SetParent(f)
	
	f.舒沃克Showak = BShowak舒沃克
	f.舒沃克Showak.SetParent(f)
	
	f.唐加西Tangasi = BTangasi唐加西
	f.唐加西Tangasi.SetParent(f)
	
	f.瓦尔卡亚特Wolqayt = BWolqayt瓦尔卡亚特
	f.瓦尔卡亚特Wolqayt.SetParent(f)
	
}
