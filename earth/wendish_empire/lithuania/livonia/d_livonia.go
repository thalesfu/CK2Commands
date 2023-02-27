package livonia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/livonia/riga"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/livonia/tukums"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/livonia/zemigalians"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LivoniaDuke interface {
    feud.Duke
    CRiga里加() 	riga.RigaCounty
    CTukums图库姆斯() 	tukums.TukumsCounty
    CZemigalians泽姆加莱() 	zemigalians.ZemigaliansCounty
}

type 利沃尼亚LivoniaDuke struct {
	feud.BaseDuke
	里加Riga 	riga.RigaCounty
	图库姆斯Tukums 	tukums.TukumsCounty
	泽姆加莱Zemigalians 	zemigalians.ZemigaliansCounty
}

func (d *利沃尼亚LivoniaDuke) CRiga里加() riga.RigaCounty {
	return d.里加Riga
}
    
func (d *利沃尼亚LivoniaDuke) CTukums图库姆斯() tukums.TukumsCounty {
	return d.图库姆斯Tukums
}
    
func (d *利沃尼亚LivoniaDuke) CZemigalians泽姆加莱() zemigalians.ZemigaliansCounty {
	return d.泽姆加莱Zemigalians
}
    
var DLivonia利沃尼亚 LivoniaDuke = &利沃尼亚LivoniaDuke{}

func init() {
	f := DLivonia利沃尼亚.(*利沃尼亚LivoniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "livonia",
		TitleName: "利沃尼亚",
		TitleCode: "d_livonia",
		Counties:  map[string]feud.County{},
	}

	f.里加Riga = riga.CRiga里加
	f.里加Riga.SetParent(f)
	
	f.图库姆斯Tukums = tukums.CTukums图库姆斯
	f.图库姆斯Tukums.SetParent(f)
	
	f.泽姆加莱Zemigalians = zemigalians.CZemigalians泽姆加莱
	f.泽姆加莱Zemigalians.SetParent(f)
	
}
