package styria

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/styria/graz"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/styria/leoben"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/styria/pitten"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StyriaDuke interface {
    feud.Duke
    CGraz格拉茨() 	graz.GrazCounty
    CLeoben埃彭施泰因() 	leoben.LeobenCounty
    CPitten皮滕() 	pitten.PittenCounty
}

type 施蒂里亚StyriaDuke struct {
	feud.BaseDuke
	格拉茨Graz 	graz.GrazCounty
	埃彭施泰因Leoben 	leoben.LeobenCounty
	皮滕Pitten 	pitten.PittenCounty
}

func (d *施蒂里亚StyriaDuke) CGraz格拉茨() graz.GrazCounty {
	return d.格拉茨Graz
}
    
func (d *施蒂里亚StyriaDuke) CLeoben埃彭施泰因() leoben.LeobenCounty {
	return d.埃彭施泰因Leoben
}
    
func (d *施蒂里亚StyriaDuke) CPitten皮滕() pitten.PittenCounty {
	return d.皮滕Pitten
}
    
var DStyria施蒂里亚 StyriaDuke = &施蒂里亚StyriaDuke{}

func init() {
	f := DStyria施蒂里亚.(*施蒂里亚StyriaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "styria",
		TitleName: "施蒂里亚",
		TitleCode: "d_styria",
		Counties:  map[string]feud.County{},
	}

	f.格拉茨Graz = graz.CGraz格拉茨
	f.格拉茨Graz.SetParent(f)
	
	f.埃彭施泰因Leoben = leoben.CLeoben埃彭施泰因
	f.埃彭施泰因Leoben.SetParent(f)
	
	f.皮滕Pitten = pitten.CPitten皮滕
	f.皮滕Pitten.SetParent(f)
	
}
