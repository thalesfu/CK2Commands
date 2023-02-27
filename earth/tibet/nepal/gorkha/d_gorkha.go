package gorkha

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/gorkha/doti"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/gorkha/jumla"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/gorkha/lumbini"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/gorkha/mustang"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/gorkha/pokhara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GorkhaDuke interface {
    feud.Duke
    CDoti多蒂() 	doti.DotiCounty
    CJumla久姆拉() 	jumla.JumlaCounty
    CLumbini岚毗尼() 	lumbini.LumbiniCounty
    CMustang木斯塘() 	mustang.MustangCounty
    CPokhara布佉罗() 	pokhara.PokharaCounty
}

type 廓尔喀GorkhaDuke struct {
	feud.BaseDuke
	多蒂Doti 	doti.DotiCounty
	久姆拉Jumla 	jumla.JumlaCounty
	岚毗尼Lumbini 	lumbini.LumbiniCounty
	木斯塘Mustang 	mustang.MustangCounty
	布佉罗Pokhara 	pokhara.PokharaCounty
}

func (d *廓尔喀GorkhaDuke) CDoti多蒂() doti.DotiCounty {
	return d.多蒂Doti
}
    
func (d *廓尔喀GorkhaDuke) CJumla久姆拉() jumla.JumlaCounty {
	return d.久姆拉Jumla
}
    
func (d *廓尔喀GorkhaDuke) CLumbini岚毗尼() lumbini.LumbiniCounty {
	return d.岚毗尼Lumbini
}
    
func (d *廓尔喀GorkhaDuke) CMustang木斯塘() mustang.MustangCounty {
	return d.木斯塘Mustang
}
    
func (d *廓尔喀GorkhaDuke) CPokhara布佉罗() pokhara.PokharaCounty {
	return d.布佉罗Pokhara
}
    
var DGorkha廓尔喀 GorkhaDuke = &廓尔喀GorkhaDuke{}

func init() {
	f := DGorkha廓尔喀.(*廓尔喀GorkhaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gorkha",
		TitleName: "廓尔喀",
		TitleCode: "d_gorkha",
		Counties:  map[string]feud.County{},
	}

	f.多蒂Doti = doti.CDoti多蒂
	f.多蒂Doti.SetParent(f)
	
	f.久姆拉Jumla = jumla.CJumla久姆拉
	f.久姆拉Jumla.SetParent(f)
	
	f.岚毗尼Lumbini = lumbini.CLumbini岚毗尼
	f.岚毗尼Lumbini.SetParent(f)
	
	f.木斯塘Mustang = mustang.CMustang木斯塘
	f.木斯塘Mustang.SetParent(f)
	
	f.布佉罗Pokhara = pokhara.CPokhara布佉罗
	f.布佉罗Pokhara.SetParent(f)
	
}
