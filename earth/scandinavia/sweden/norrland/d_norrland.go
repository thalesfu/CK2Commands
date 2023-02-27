package norrland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/norrland/angermanland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/norrland/halsingland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden/norrland/medelpad"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorrlandDuke interface {
    feud.Duke
    CAngermanland翁厄曼兰() 	angermanland.AngermanlandCounty
    CHalsingland海尔辛兰() 	halsingland.HalsinglandCounty
    CMedelpad梅代尔帕德() 	medelpad.MedelpadCounty
}

type 海尔辛兰NorrlandDuke struct {
	feud.BaseDuke
	翁厄曼兰Angermanland 	angermanland.AngermanlandCounty
	海尔辛兰Halsingland 	halsingland.HalsinglandCounty
	梅代尔帕德Medelpad 	medelpad.MedelpadCounty
}

func (d *海尔辛兰NorrlandDuke) CAngermanland翁厄曼兰() angermanland.AngermanlandCounty {
	return d.翁厄曼兰Angermanland
}
    
func (d *海尔辛兰NorrlandDuke) CHalsingland海尔辛兰() halsingland.HalsinglandCounty {
	return d.海尔辛兰Halsingland
}
    
func (d *海尔辛兰NorrlandDuke) CMedelpad梅代尔帕德() medelpad.MedelpadCounty {
	return d.梅代尔帕德Medelpad
}
    
var DNorrland海尔辛兰 NorrlandDuke = &海尔辛兰NorrlandDuke{}

func init() {
	f := DNorrland海尔辛兰.(*海尔辛兰NorrlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "norrland",
		TitleName: "海尔辛兰",
		TitleCode: "d_norrland",
		Counties:  map[string]feud.County{},
	}

	f.翁厄曼兰Angermanland = angermanland.CAngermanland翁厄曼兰
	f.翁厄曼兰Angermanland.SetParent(f)
	
	f.海尔辛兰Halsingland = halsingland.CHalsingland海尔辛兰
	f.海尔辛兰Halsingland.SetParent(f)
	
	f.梅代尔帕德Medelpad = medelpad.CMedelpad梅代尔帕德
	f.梅代尔帕德Medelpad.SetParent(f)
	
}
