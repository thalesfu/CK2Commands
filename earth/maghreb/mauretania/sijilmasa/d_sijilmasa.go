package sijilmasa

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sijilmasa/sijilmasa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sijilmasa/taghaza"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sijilmasa/tudgha"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sijilmasa/ziz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SijilmasaDuke interface {
    feud.Duke
    CSijilmasa锡吉勒马萨() 	sijilmasa.SijilmasaCounty
    CTaghaza泰加扎() 	taghaza.TaghazaCounty
    CTudgha图德盖() 	tudgha.TudghaCounty
    CZiz济兹() 	ziz.ZizCounty
}

type 锡吉勒马萨SijilmasaDuke struct {
	feud.BaseDuke
	锡吉勒马萨Sijilmasa 	sijilmasa.SijilmasaCounty
	泰加扎Taghaza 	taghaza.TaghazaCounty
	图德盖Tudgha 	tudgha.TudghaCounty
	济兹Ziz 	ziz.ZizCounty
}

func (d *锡吉勒马萨SijilmasaDuke) CSijilmasa锡吉勒马萨() sijilmasa.SijilmasaCounty {
	return d.锡吉勒马萨Sijilmasa
}
    
func (d *锡吉勒马萨SijilmasaDuke) CTaghaza泰加扎() taghaza.TaghazaCounty {
	return d.泰加扎Taghaza
}
    
func (d *锡吉勒马萨SijilmasaDuke) CTudgha图德盖() tudgha.TudghaCounty {
	return d.图德盖Tudgha
}
    
func (d *锡吉勒马萨SijilmasaDuke) CZiz济兹() ziz.ZizCounty {
	return d.济兹Ziz
}
    
var DSijilmasa锡吉勒马萨 SijilmasaDuke = &锡吉勒马萨SijilmasaDuke{}

func init() {
	f := DSijilmasa锡吉勒马萨.(*锡吉勒马萨SijilmasaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sijilmasa",
		TitleName: "锡吉勒马萨",
		TitleCode: "d_sijilmasa",
		Counties:  map[string]feud.County{},
	}

	f.锡吉勒马萨Sijilmasa = sijilmasa.CSijilmasa锡吉勒马萨
	f.锡吉勒马萨Sijilmasa.SetParent(f)
	
	f.泰加扎Taghaza = taghaza.CTaghaza泰加扎
	f.泰加扎Taghaza.SetParent(f)
	
	f.图德盖Tudgha = tudgha.CTudgha图德盖
	f.图德盖Tudgha.SetParent(f)
	
	f.济兹Ziz = ziz.CZiz济兹
	f.济兹Ziz.SetParent(f)
	
}
