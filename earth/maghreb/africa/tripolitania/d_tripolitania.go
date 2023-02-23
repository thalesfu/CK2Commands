package tripolitania

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tripolitania/djerba"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tripolitania/nafusa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tripolitania/nalut"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/tripolitania/tripolitana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TripolitaniaDuke interface {
	feud.Duke
	CDjerba杰尔巴() djerba.DjerbaCounty
	CNafusa奈富塞() nafusa.NafusaCounty
	CNalut纳卢特() nalut.NalutCounty
	CTripolitana的黎波里塔尼亚() tripolitana.TripolitanaCounty
}

type 的黎波里塔尼亚TripolitaniaDuke struct {
	feud.BaseDuke
	杰尔巴Djerba          djerba.DjerbaCounty
	奈富塞Nafusa          nafusa.NafusaCounty
	纳卢特Nalut           nalut.NalutCounty
	的黎波里塔尼亚Tripolitana tripolitana.TripolitanaCounty
}

func (d *的黎波里塔尼亚TripolitaniaDuke) CDjerba杰尔巴() djerba.DjerbaCounty {
	return d.杰尔巴Djerba
}

func (d *的黎波里塔尼亚TripolitaniaDuke) CNafusa奈富塞() nafusa.NafusaCounty {
	return d.奈富塞Nafusa
}

func (d *的黎波里塔尼亚TripolitaniaDuke) CNalut纳卢特() nalut.NalutCounty {
	return d.纳卢特Nalut
}

func (d *的黎波里塔尼亚TripolitaniaDuke) CTripolitana的黎波里塔尼亚() tripolitana.TripolitanaCounty {
	return d.的黎波里塔尼亚Tripolitana
}

var DTripolitania的黎波里塔尼亚 TripolitaniaDuke = &的黎波里塔尼亚TripolitaniaDuke{}

func init() {
	f := DTripolitania的黎波里塔尼亚.(*的黎波里塔尼亚TripolitaniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tripolitania",
		TitleName: "的黎波里塔尼亚",
		TitleCode: "d_tripolitania",
		Counties:  map[string]feud.County{},
	}

	f.杰尔巴Djerba = djerba.CDjerba杰尔巴
	f.杰尔巴Djerba.SetParent(f)

	f.奈富塞Nafusa = nafusa.CNafusa奈富塞
	f.奈富塞Nafusa.SetParent(f)

	f.纳卢特Nalut = nalut.CNalut纳卢特
	f.纳卢特Nalut.SetParent(f)

	f.的黎波里塔尼亚Tripolitana = tripolitana.CTripolitana的黎波里塔尼亚
	f.的黎波里塔尼亚Tripolitana.SetParent(f)

}
