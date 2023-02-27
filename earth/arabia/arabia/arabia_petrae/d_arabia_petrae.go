package arabia_petrae

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/arabia_petrae/al_aqabah"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/arabia_petrae/al_jawf"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/arabia_petrae/hijaz"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/arabia_petrae/maan"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/arabia_petrae/tabuk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Arabia_petraeDuke interface {
    feud.Duke
    CAl_aqabah亚喀巴() 	al_aqabah.Al_aqabahCounty
    CAl_jawf焦夫() 	al_jawf.Al_jawfCounty
    CHijaz欧拉() 	hijaz.HijazCounty
    CMaan马安() 	maan.MaanCounty
    CTabuk塔布克() 	tabuk.TabukCounty
}

type 阿拉比亚Arabia_petraeDuke struct {
	feud.BaseDuke
	亚喀巴Al_aqabah 	al_aqabah.Al_aqabahCounty
	焦夫Al_jawf 	al_jawf.Al_jawfCounty
	欧拉Hijaz 	hijaz.HijazCounty
	马安Maan 	maan.MaanCounty
	塔布克Tabuk 	tabuk.TabukCounty
}

func (d *阿拉比亚Arabia_petraeDuke) CAl_aqabah亚喀巴() al_aqabah.Al_aqabahCounty {
	return d.亚喀巴Al_aqabah
}
    
func (d *阿拉比亚Arabia_petraeDuke) CAl_jawf焦夫() al_jawf.Al_jawfCounty {
	return d.焦夫Al_jawf
}
    
func (d *阿拉比亚Arabia_petraeDuke) CHijaz欧拉() hijaz.HijazCounty {
	return d.欧拉Hijaz
}
    
func (d *阿拉比亚Arabia_petraeDuke) CMaan马安() maan.MaanCounty {
	return d.马安Maan
}
    
func (d *阿拉比亚Arabia_petraeDuke) CTabuk塔布克() tabuk.TabukCounty {
	return d.塔布克Tabuk
}
    
var DArabia_petrae阿拉比亚 Arabia_petraeDuke = &阿拉比亚Arabia_petraeDuke{}

func init() {
	f := DArabia_petrae阿拉比亚.(*阿拉比亚Arabia_petraeDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "arabia_petrae",
		TitleName: "阿拉比亚",
		TitleCode: "d_arabia_petrae",
		Counties:  map[string]feud.County{},
	}

	f.亚喀巴Al_aqabah = al_aqabah.CAl_aqabah亚喀巴
	f.亚喀巴Al_aqabah.SetParent(f)
	
	f.焦夫Al_jawf = al_jawf.CAl_jawf焦夫
	f.焦夫Al_jawf.SetParent(f)
	
	f.欧拉Hijaz = hijaz.CHijaz欧拉
	f.欧拉Hijaz.SetParent(f)
	
	f.马安Maan = maan.CMaan马安
	f.马安Maan.SetParent(f)
	
	f.塔布克Tabuk = tabuk.CTabuk塔布克
	f.塔布克Tabuk.SetParent(f)
	
}
