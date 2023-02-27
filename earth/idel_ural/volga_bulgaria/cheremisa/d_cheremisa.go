package cheremisa

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/cheremisa/chuvash"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/cheremisa/grassland_cheremisa"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/cheremisa/mountain_cheremisa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CheremisaDuke interface {
    feud.Duke
    CChuvash楚瓦什() 	chuvash.ChuvashCounty
    CGrassland_cheremisa克尔热涅茨() 	grassland_cheremisa.Grassland_cheremisaCounty
    CMountain_cheremisa切列米莎() 	mountain_cheremisa.Mountain_cheremisaCounty
}

type 切列米莎CheremisaDuke struct {
	feud.BaseDuke
	楚瓦什Chuvash 	chuvash.ChuvashCounty
	克尔热涅茨Grassland_cheremisa 	grassland_cheremisa.Grassland_cheremisaCounty
	切列米莎Mountain_cheremisa 	mountain_cheremisa.Mountain_cheremisaCounty
}

func (d *切列米莎CheremisaDuke) CChuvash楚瓦什() chuvash.ChuvashCounty {
	return d.楚瓦什Chuvash
}
    
func (d *切列米莎CheremisaDuke) CGrassland_cheremisa克尔热涅茨() grassland_cheremisa.Grassland_cheremisaCounty {
	return d.克尔热涅茨Grassland_cheremisa
}
    
func (d *切列米莎CheremisaDuke) CMountain_cheremisa切列米莎() mountain_cheremisa.Mountain_cheremisaCounty {
	return d.切列米莎Mountain_cheremisa
}
    
var DCheremisa切列米莎 CheremisaDuke = &切列米莎CheremisaDuke{}

func init() {
	f := DCheremisa切列米莎.(*切列米莎CheremisaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cheremisa",
		TitleName: "切列米莎",
		TitleCode: "d_cheremisa",
		Counties:  map[string]feud.County{},
	}

	f.楚瓦什Chuvash = chuvash.CChuvash楚瓦什
	f.楚瓦什Chuvash.SetParent(f)
	
	f.克尔热涅茨Grassland_cheremisa = grassland_cheremisa.CGrassland_cheremisa克尔热涅茨
	f.克尔热涅茨Grassland_cheremisa.SetParent(f)
	
	f.切列米莎Mountain_cheremisa = mountain_cheremisa.CMountain_cheremisa切列米莎
	f.切列米莎Mountain_cheremisa.SetParent(f)
	
}
