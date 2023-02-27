package cyprus

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cyprus/famagusta"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/cyprus/limisol"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CyprusDuke interface {
    feud.Duke
    CFamagusta法马古斯塔() 	famagusta.FamagustaCounty
    CLimisol利马索尔() 	limisol.LimisolCounty
}

type 塞浦路斯CyprusDuke struct {
	feud.BaseDuke
	法马古斯塔Famagusta 	famagusta.FamagustaCounty
	利马索尔Limisol 	limisol.LimisolCounty
}

func (d *塞浦路斯CyprusDuke) CFamagusta法马古斯塔() famagusta.FamagustaCounty {
	return d.法马古斯塔Famagusta
}
    
func (d *塞浦路斯CyprusDuke) CLimisol利马索尔() limisol.LimisolCounty {
	return d.利马索尔Limisol
}
    
var DCyprus塞浦路斯 CyprusDuke = &塞浦路斯CyprusDuke{}

func init() {
	f := DCyprus塞浦路斯.(*塞浦路斯CyprusDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cyprus",
		TitleName: "塞浦路斯",
		TitleCode: "d_cyprus",
		Counties:  map[string]feud.County{},
	}

	f.法马古斯塔Famagusta = famagusta.CFamagusta法马古斯塔
	f.法马古斯塔Famagusta.SetParent(f)
	
	f.利马索尔Limisol = limisol.CLimisol利马索尔
	f.利马索尔Limisol.SetParent(f)
	
}
