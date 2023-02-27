package hayya

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/hayya/gezira"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/hayya/hayya"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/hayya/kassala"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/hayya/suakin"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/hayya/trinkitat"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HayyaDuke interface {
    feud.Duke
    CGezira杰济拉() 	gezira.GeziraCounty
    CHayya海亚() 	hayya.HayyaCounty
    CKassala卡萨拉() 	kassala.KassalaCounty
    CSuakin萨瓦金() 	suakin.SuakinCounty
    CTrinkitat特林基塔特() 	trinkitat.TrinkitatCounty
}

type 比林美伊亚HayyaDuke struct {
	feud.BaseDuke
	杰济拉Gezira 	gezira.GeziraCounty
	海亚Hayya 	hayya.HayyaCounty
	卡萨拉Kassala 	kassala.KassalaCounty
	萨瓦金Suakin 	suakin.SuakinCounty
	特林基塔特Trinkitat 	trinkitat.TrinkitatCounty
}

func (d *比林美伊亚HayyaDuke) CGezira杰济拉() gezira.GeziraCounty {
	return d.杰济拉Gezira
}
    
func (d *比林美伊亚HayyaDuke) CHayya海亚() hayya.HayyaCounty {
	return d.海亚Hayya
}
    
func (d *比林美伊亚HayyaDuke) CKassala卡萨拉() kassala.KassalaCounty {
	return d.卡萨拉Kassala
}
    
func (d *比林美伊亚HayyaDuke) CSuakin萨瓦金() suakin.SuakinCounty {
	return d.萨瓦金Suakin
}
    
func (d *比林美伊亚HayyaDuke) CTrinkitat特林基塔特() trinkitat.TrinkitatCounty {
	return d.特林基塔特Trinkitat
}
    
var DHayya比林美伊亚 HayyaDuke = &比林美伊亚HayyaDuke{}

func init() {
	f := DHayya比林美伊亚.(*比林美伊亚HayyaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hayya",
		TitleName: "比林美伊亚",
		TitleCode: "d_hayya",
		Counties:  map[string]feud.County{},
	}

	f.杰济拉Gezira = gezira.CGezira杰济拉
	f.杰济拉Gezira.SetParent(f)
	
	f.海亚Hayya = hayya.CHayya海亚
	f.海亚Hayya.SetParent(f)
	
	f.卡萨拉Kassala = kassala.CKassala卡萨拉
	f.卡萨拉Kassala.SetParent(f)
	
	f.萨瓦金Suakin = suakin.CSuakin萨瓦金
	f.萨瓦金Suakin.SetParent(f)
	
	f.特林基塔特Trinkitat = trinkitat.CTrinkitat特林基塔特
	f.特林基塔特Trinkitat.SetParent(f)
	
}
