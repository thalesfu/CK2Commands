package bourbon

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/bourbon/bourbon"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/bourbon/la_marche"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/bourbon/limousin"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BourbonDuke interface {
    feud.Duke
    CBourbon波旁() 	bourbon.BourbonCounty
    CLa_marche拉马什() 	la_marche.La_marcheCounty
    CLimousin利穆赞() 	limousin.LimousinCounty
}

type 波旁BourbonDuke struct {
	feud.BaseDuke
	波旁Bourbon 	bourbon.BourbonCounty
	拉马什La_marche 	la_marche.La_marcheCounty
	利穆赞Limousin 	limousin.LimousinCounty
}

func (d *波旁BourbonDuke) CBourbon波旁() bourbon.BourbonCounty {
	return d.波旁Bourbon
}
    
func (d *波旁BourbonDuke) CLa_marche拉马什() la_marche.La_marcheCounty {
	return d.拉马什La_marche
}
    
func (d *波旁BourbonDuke) CLimousin利穆赞() limousin.LimousinCounty {
	return d.利穆赞Limousin
}
    
var DBourbon波旁 BourbonDuke = &波旁BourbonDuke{}

func init() {
	f := DBourbon波旁.(*波旁BourbonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bourbon",
		TitleName: "波旁",
		TitleCode: "d_bourbon",
		Counties:  map[string]feud.County{},
	}

	f.波旁Bourbon = bourbon.CBourbon波旁
	f.波旁Bourbon.SetParent(f)
	
	f.拉马什La_marche = la_marche.CLa_marche拉马什
	f.拉马什La_marche.SetParent(f)
	
	f.利穆赞Limousin = limousin.CLimousin利穆赞
	f.利穆赞Limousin.SetParent(f)
	
}
