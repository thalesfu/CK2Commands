package syrte

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/syrte/ajadabiya"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/syrte/syrte"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/syrte/tadjrift"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyrteDuke interface {
    feud.Duke
    CAjadabiya艾季达比耶() 	ajadabiya.AjadabiyaCounty
    CSyrte苏尔特() 	syrte.SyrteCounty
    CTadjrift塔杰里夫特() 	tadjrift.TadjriftCounty
}

type 苏尔特SyrteDuke struct {
	feud.BaseDuke
	艾季达比耶Ajadabiya 	ajadabiya.AjadabiyaCounty
	苏尔特Syrte 	syrte.SyrteCounty
	塔杰里夫特Tadjrift 	tadjrift.TadjriftCounty
}

func (d *苏尔特SyrteDuke) CAjadabiya艾季达比耶() ajadabiya.AjadabiyaCounty {
	return d.艾季达比耶Ajadabiya
}
    
func (d *苏尔特SyrteDuke) CSyrte苏尔特() syrte.SyrteCounty {
	return d.苏尔特Syrte
}
    
func (d *苏尔特SyrteDuke) CTadjrift塔杰里夫特() tadjrift.TadjriftCounty {
	return d.塔杰里夫特Tadjrift
}
    
var DSyrte苏尔特 SyrteDuke = &苏尔特SyrteDuke{}

func init() {
	f := DSyrte苏尔特.(*苏尔特SyrteDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "syrte",
		TitleName: "苏尔特",
		TitleCode: "d_syrte",
		Counties:  map[string]feud.County{},
	}

	f.艾季达比耶Ajadabiya = ajadabiya.CAjadabiya艾季达比耶
	f.艾季达比耶Ajadabiya.SetParent(f)
	
	f.苏尔特Syrte = syrte.CSyrte苏尔特
	f.苏尔特Syrte.SetParent(f)
	
	f.塔杰里夫特Tadjrift = tadjrift.CTadjrift塔杰里夫特
	f.塔杰里夫特Tadjrift.SetParent(f)
	
}
