package swabia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/germany/swabia/furstenberg"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/swabia/kempten"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/swabia/schwaben"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/swabia/ulm"
	"github.com/thalesfu/CK2Commands/earth/germany/germany/swabia/wurttemberg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SwabiaDuke interface {
    feud.Duke
    CFurstenberg菲尔斯滕贝格() 	furstenberg.FurstenbergCounty
    CKempten奥格斯堡() 	kempten.KemptenCounty
    CSchwaben施瓦本() 	schwaben.SchwabenCounty
    CUlm乌尔姆() 	ulm.UlmCounty
    CWurttemberg符滕堡() 	wurttemberg.WurttembergCounty
}

type 施瓦本SwabiaDuke struct {
	feud.BaseDuke
	菲尔斯滕贝格Furstenberg 	furstenberg.FurstenbergCounty
	奥格斯堡Kempten 	kempten.KemptenCounty
	施瓦本Schwaben 	schwaben.SchwabenCounty
	乌尔姆Ulm 	ulm.UlmCounty
	符滕堡Wurttemberg 	wurttemberg.WurttembergCounty
}

func (d *施瓦本SwabiaDuke) CFurstenberg菲尔斯滕贝格() furstenberg.FurstenbergCounty {
	return d.菲尔斯滕贝格Furstenberg
}
    
func (d *施瓦本SwabiaDuke) CKempten奥格斯堡() kempten.KemptenCounty {
	return d.奥格斯堡Kempten
}
    
func (d *施瓦本SwabiaDuke) CSchwaben施瓦本() schwaben.SchwabenCounty {
	return d.施瓦本Schwaben
}
    
func (d *施瓦本SwabiaDuke) CUlm乌尔姆() ulm.UlmCounty {
	return d.乌尔姆Ulm
}
    
func (d *施瓦本SwabiaDuke) CWurttemberg符滕堡() wurttemberg.WurttembergCounty {
	return d.符滕堡Wurttemberg
}
    
var DSwabia施瓦本 SwabiaDuke = &施瓦本SwabiaDuke{}

func init() {
	f := DSwabia施瓦本.(*施瓦本SwabiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "swabia",
		TitleName: "施瓦本",
		TitleCode: "d_swabia",
		Counties:  map[string]feud.County{},
	}

	f.菲尔斯滕贝格Furstenberg = furstenberg.CFurstenberg菲尔斯滕贝格
	f.菲尔斯滕贝格Furstenberg.SetParent(f)
	
	f.奥格斯堡Kempten = kempten.CKempten奥格斯堡
	f.奥格斯堡Kempten.SetParent(f)
	
	f.施瓦本Schwaben = schwaben.CSchwaben施瓦本
	f.施瓦本Schwaben.SetParent(f)
	
	f.乌尔姆Ulm = ulm.CUlm乌尔姆
	f.乌尔姆Ulm.SetParent(f)
	
	f.符滕堡Wurttemberg = wurttemberg.CWurttemberg符滕堡
	f.符滕堡Wurttemberg.SetParent(f)
	
}
