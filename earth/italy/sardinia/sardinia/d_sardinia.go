package sardinia

import (
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/sardinia/arborea"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/sardinia/cagliari"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/sardinia/gallura"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/sardinia/ogliastra"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/sardinia/torres"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SardiniaDuke interface {
    feud.Duke
    CArborea阿尔博雷亚() 	arborea.ArboreaCounty
    CCagliari卡利亚里() 	cagliari.CagliariCounty
    CGallura加卢拉() 	gallura.GalluraCounty
    COgliastra奥利亚斯特拉() 	ogliastra.OgliastraCounty
    CTorres托雷斯() 	torres.TorresCounty
}

type 撒丁SardiniaDuke struct {
	feud.BaseDuke
	阿尔博雷亚Arborea 	arborea.ArboreaCounty
	卡利亚里Cagliari 	cagliari.CagliariCounty
	加卢拉Gallura 	gallura.GalluraCounty
	奥利亚斯特拉Ogliastra 	ogliastra.OgliastraCounty
	托雷斯Torres 	torres.TorresCounty
}

func (d *撒丁SardiniaDuke) CArborea阿尔博雷亚() arborea.ArboreaCounty {
	return d.阿尔博雷亚Arborea
}
    
func (d *撒丁SardiniaDuke) CCagliari卡利亚里() cagliari.CagliariCounty {
	return d.卡利亚里Cagliari
}
    
func (d *撒丁SardiniaDuke) CGallura加卢拉() gallura.GalluraCounty {
	return d.加卢拉Gallura
}
    
func (d *撒丁SardiniaDuke) COgliastra奥利亚斯特拉() ogliastra.OgliastraCounty {
	return d.奥利亚斯特拉Ogliastra
}
    
func (d *撒丁SardiniaDuke) CTorres托雷斯() torres.TorresCounty {
	return d.托雷斯Torres
}
    
var DSardinia撒丁 SardiniaDuke = &撒丁SardiniaDuke{}

func init() {
	f := DSardinia撒丁.(*撒丁SardiniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sardinia",
		TitleName: "撒丁",
		TitleCode: "d_sardinia",
		Counties:  map[string]feud.County{},
	}

	f.阿尔博雷亚Arborea = arborea.CArborea阿尔博雷亚
	f.阿尔博雷亚Arborea.SetParent(f)
	
	f.卡利亚里Cagliari = cagliari.CCagliari卡利亚里
	f.卡利亚里Cagliari.SetParent(f)
	
	f.加卢拉Gallura = gallura.CGallura加卢拉
	f.加卢拉Gallura.SetParent(f)
	
	f.奥利亚斯特拉Ogliastra = ogliastra.COgliastra奥利亚斯特拉
	f.奥利亚斯特拉Ogliastra.SetParent(f)
	
	f.托雷斯Torres = torres.CTorres托雷斯
	f.托雷斯Torres.SetParent(f)
	
}
