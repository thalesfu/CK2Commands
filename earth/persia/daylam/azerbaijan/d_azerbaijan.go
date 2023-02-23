package azerbaijan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/azerbaijan/azerbaijan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/azerbaijan/shemakha"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/azerbaijan/shirvan"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/azerbaijan/suenik"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AzerbaijanDuke interface {
	feud.Duke
	CAzerbaijan阿塞拜疆() azerbaijan.AzerbaijanCounty
	CShemakha舍马哈() shemakha.ShemakhaCounty
	CShirvan希尔凡() shirvan.ShirvanCounty
	CSuenik休尼克() suenik.SuenikCounty
}

type 阿塞拜疆AzerbaijanDuke struct {
	feud.BaseDuke
	阿塞拜疆Azerbaijan azerbaijan.AzerbaijanCounty
	舍马哈Shemakha    shemakha.ShemakhaCounty
	希尔凡Shirvan     shirvan.ShirvanCounty
	休尼克Suenik      suenik.SuenikCounty
}

func (d *阿塞拜疆AzerbaijanDuke) CAzerbaijan阿塞拜疆() azerbaijan.AzerbaijanCounty {
	return d.阿塞拜疆Azerbaijan
}

func (d *阿塞拜疆AzerbaijanDuke) CShemakha舍马哈() shemakha.ShemakhaCounty {
	return d.舍马哈Shemakha
}

func (d *阿塞拜疆AzerbaijanDuke) CShirvan希尔凡() shirvan.ShirvanCounty {
	return d.希尔凡Shirvan
}

func (d *阿塞拜疆AzerbaijanDuke) CSuenik休尼克() suenik.SuenikCounty {
	return d.休尼克Suenik
}

var DAzerbaijan阿塞拜疆 AzerbaijanDuke = &阿塞拜疆AzerbaijanDuke{}

func init() {
	f := DAzerbaijan阿塞拜疆.(*阿塞拜疆AzerbaijanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "azerbaijan",
		TitleName: "阿塞拜疆",
		TitleCode: "d_azerbaijan",
		Counties:  map[string]feud.County{},
	}

	f.阿塞拜疆Azerbaijan = azerbaijan.CAzerbaijan阿塞拜疆
	f.阿塞拜疆Azerbaijan.SetParent(f)

	f.舍马哈Shemakha = shemakha.CShemakha舍马哈
	f.舍马哈Shemakha.SetParent(f)

	f.希尔凡Shirvan = shirvan.CShirvan希尔凡
	f.希尔凡Shirvan.SetParent(f)

	f.休尼克Suenik = suenik.CSuenik休尼克
	f.休尼克Suenik.SetParent(f)

}
