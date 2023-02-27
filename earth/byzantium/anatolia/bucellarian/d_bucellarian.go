package bucellarian

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/bucellarian/ankyra"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/bucellarian/claudiopolis"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BucellarianDuke interface {
    feud.Duke
    CAnkyra安居拉() 	ankyra.AnkyraCounty
    CClaudiopolis克劳狄奥波利斯() 	claudiopolis.ClaudiopolisCounty
}

type 布凯拉里翁BucellarianDuke struct {
	feud.BaseDuke
	安居拉Ankyra 	ankyra.AnkyraCounty
	克劳狄奥波利斯Claudiopolis 	claudiopolis.ClaudiopolisCounty
}

func (d *布凯拉里翁BucellarianDuke) CAnkyra安居拉() ankyra.AnkyraCounty {
	return d.安居拉Ankyra
}
    
func (d *布凯拉里翁BucellarianDuke) CClaudiopolis克劳狄奥波利斯() claudiopolis.ClaudiopolisCounty {
	return d.克劳狄奥波利斯Claudiopolis
}
    
var DBucellarian布凯拉里翁 BucellarianDuke = &布凯拉里翁BucellarianDuke{}

func init() {
	f := DBucellarian布凯拉里翁.(*布凯拉里翁BucellarianDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bucellarian",
		TitleName: "布凯拉里翁",
		TitleCode: "d_bucellarian",
		Counties:  map[string]feud.County{},
	}

	f.安居拉Ankyra = ankyra.CAnkyra安居拉
	f.安居拉Ankyra.SetParent(f)
	
	f.克劳狄奥波利斯Claudiopolis = claudiopolis.CClaudiopolis克劳狄奥波利斯
	f.克劳狄奥波利斯Claudiopolis.SetParent(f)
	
}
