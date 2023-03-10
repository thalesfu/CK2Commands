package mazovia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/mazovia/czersk"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/mazovia/lomzynska"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/mazovia/plock"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/mazovia/sieradzko_leczyckie"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MazoviaDuke interface {
    feud.Duke
    CCzersk利夫() 	czersk.CzerskCounty
    CLomzynska沃姆扎() 	lomzynska.LomzynskaCounty
    CPlock普沃茨克() 	plock.PlockCounty
    CSieradzko_leczyckie戈斯蒂宁() 	sieradzko_leczyckie.Sieradzko_leczyckieCounty
}

type 马佐维亚MazoviaDuke struct {
	feud.BaseDuke
	利夫Czersk 	czersk.CzerskCounty
	沃姆扎Lomzynska 	lomzynska.LomzynskaCounty
	普沃茨克Plock 	plock.PlockCounty
	戈斯蒂宁Sieradzko_leczyckie 	sieradzko_leczyckie.Sieradzko_leczyckieCounty
}

func (d *马佐维亚MazoviaDuke) CCzersk利夫() czersk.CzerskCounty {
	return d.利夫Czersk
}
    
func (d *马佐维亚MazoviaDuke) CLomzynska沃姆扎() lomzynska.LomzynskaCounty {
	return d.沃姆扎Lomzynska
}
    
func (d *马佐维亚MazoviaDuke) CPlock普沃茨克() plock.PlockCounty {
	return d.普沃茨克Plock
}
    
func (d *马佐维亚MazoviaDuke) CSieradzko_leczyckie戈斯蒂宁() sieradzko_leczyckie.Sieradzko_leczyckieCounty {
	return d.戈斯蒂宁Sieradzko_leczyckie
}
    
var DMazovia马佐维亚 MazoviaDuke = &马佐维亚MazoviaDuke{}

func init() {
	f := DMazovia马佐维亚.(*马佐维亚MazoviaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mazovia",
		TitleName: "马佐维亚",
		TitleCode: "d_mazovia",
		Counties:  map[string]feud.County{},
	}

	f.利夫Czersk = czersk.CCzersk利夫
	f.利夫Czersk.SetParent(f)
	
	f.沃姆扎Lomzynska = lomzynska.CLomzynska沃姆扎
	f.沃姆扎Lomzynska.SetParent(f)
	
	f.普沃茨克Plock = plock.CPlock普沃茨克
	f.普沃茨克Plock.SetParent(f)
	
	f.戈斯蒂宁Sieradzko_leczyckie = sieradzko_leczyckie.CSieradzko_leczyckie戈斯蒂宁
	f.戈斯蒂宁Sieradzko_leczyckie.SetParent(f)
	
}
