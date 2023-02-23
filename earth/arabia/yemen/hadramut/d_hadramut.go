package hadramut

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/hadramut/dhofar"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/hadramut/kathiri"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/hadramut/mahra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HadramutDuke interface {
	feud.Duke
	CDhofar佐法尔() dhofar.DhofarCounty
	CKathiri卡迪里() kathiri.KathiriCounty
	CMahra马哈拉() mahra.MahraCounty
}

type 哈得拉姆HadramutDuke struct {
	feud.BaseDuke
	佐法尔Dhofar  dhofar.DhofarCounty
	卡迪里Kathiri kathiri.KathiriCounty
	马哈拉Mahra   mahra.MahraCounty
}

func (d *哈得拉姆HadramutDuke) CDhofar佐法尔() dhofar.DhofarCounty {
	return d.佐法尔Dhofar
}

func (d *哈得拉姆HadramutDuke) CKathiri卡迪里() kathiri.KathiriCounty {
	return d.卡迪里Kathiri
}

func (d *哈得拉姆HadramutDuke) CMahra马哈拉() mahra.MahraCounty {
	return d.马哈拉Mahra
}

var DHadramut哈得拉姆 HadramutDuke = &哈得拉姆HadramutDuke{}

func init() {
	f := DHadramut哈得拉姆.(*哈得拉姆HadramutDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hadramut",
		TitleName: "哈得拉姆",
		TitleCode: "d_hadramut",
		Counties:  map[string]feud.County{},
	}

	f.佐法尔Dhofar = dhofar.CDhofar佐法尔
	f.佐法尔Dhofar.SetParent(f)

	f.卡迪里Kathiri = kathiri.CKathiri卡迪里
	f.卡迪里Kathiri.SetParent(f)

	f.马哈拉Mahra = mahra.CMahra马哈拉
	f.马哈拉Mahra.SetParent(f)

}
