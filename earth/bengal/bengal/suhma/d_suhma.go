package suhma

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/suhma/mallabhum"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/suhma/midnapore"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/suhma/tamralipti"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuhmaDuke interface {
    feud.Duke
    CMallabhum末罗浮摩() 	mallabhum.MallabhumCounty
    CMidnapore迷地尼城() 	midnapore.MidnaporeCounty
    CTamralipti多摩梨帝() 	tamralipti.TamraliptiCounty
}

type 苏吸摩SuhmaDuke struct {
	feud.BaseDuke
	末罗浮摩Mallabhum 	mallabhum.MallabhumCounty
	迷地尼城Midnapore 	midnapore.MidnaporeCounty
	多摩梨帝Tamralipti 	tamralipti.TamraliptiCounty
}

func (d *苏吸摩SuhmaDuke) CMallabhum末罗浮摩() mallabhum.MallabhumCounty {
	return d.末罗浮摩Mallabhum
}
    
func (d *苏吸摩SuhmaDuke) CMidnapore迷地尼城() midnapore.MidnaporeCounty {
	return d.迷地尼城Midnapore
}
    
func (d *苏吸摩SuhmaDuke) CTamralipti多摩梨帝() tamralipti.TamraliptiCounty {
	return d.多摩梨帝Tamralipti
}
    
var DSuhma苏吸摩 SuhmaDuke = &苏吸摩SuhmaDuke{}

func init() {
	f := DSuhma苏吸摩.(*苏吸摩SuhmaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "suhma",
		TitleName: "苏吸摩",
		TitleCode: "d_suhma",
		Counties:  map[string]feud.County{},
	}

	f.末罗浮摩Mallabhum = mallabhum.CMallabhum末罗浮摩
	f.末罗浮摩Mallabhum.SetParent(f)
	
	f.迷地尼城Midnapore = midnapore.CMidnapore迷地尼城
	f.迷地尼城Midnapore.SetParent(f)
	
	f.多摩梨帝Tamralipti = tamralipti.CTamralipti多摩梨帝
	f.多摩梨帝Tamralipti.SetParent(f)
	
}
