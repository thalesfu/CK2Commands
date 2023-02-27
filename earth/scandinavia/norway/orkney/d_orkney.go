package orkney

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/orkney/faereyar"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/orkney/orkney"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/orkney/shetland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrkneyDuke interface {
    feud.Duke
    CFaereyar法伊雷亚尔() 	faereyar.FaereyarCounty
    COrkney奥克尼() 	orkney.OrkneyCounty
    CShetland设得兰() 	shetland.ShetlandCounty
}

type 奥克尼OrkneyDuke struct {
	feud.BaseDuke
	法伊雷亚尔Faereyar 	faereyar.FaereyarCounty
	奥克尼Orkney 	orkney.OrkneyCounty
	设得兰Shetland 	shetland.ShetlandCounty
}

func (d *奥克尼OrkneyDuke) CFaereyar法伊雷亚尔() faereyar.FaereyarCounty {
	return d.法伊雷亚尔Faereyar
}
    
func (d *奥克尼OrkneyDuke) COrkney奥克尼() orkney.OrkneyCounty {
	return d.奥克尼Orkney
}
    
func (d *奥克尼OrkneyDuke) CShetland设得兰() shetland.ShetlandCounty {
	return d.设得兰Shetland
}
    
var DOrkney奥克尼 OrkneyDuke = &奥克尼OrkneyDuke{}

func init() {
	f := DOrkney奥克尼.(*奥克尼OrkneyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "orkney",
		TitleName: "奥克尼",
		TitleCode: "d_orkney",
		Counties:  map[string]feud.County{},
	}

	f.法伊雷亚尔Faereyar = faereyar.CFaereyar法伊雷亚尔
	f.法伊雷亚尔Faereyar.SetParent(f)
	
	f.奥克尼Orkney = orkney.COrkney奥克尼
	f.奥克尼Orkney.SetParent(f)
	
	f.设得兰Shetland = shetland.CShetland设得兰
	f.设得兰Shetland.SetParent(f)
	
}
