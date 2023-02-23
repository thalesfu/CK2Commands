package croatia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/croatia/senj"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/croatia/split"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/croatia/veglia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/croatia/zadar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CroatiaDuke interface {
	feud.Duke
	CSenj塞尼() senj.SenjCounty
	CSplit斯普利特() split.SplitCounty
	CVeglia韦利亚() veglia.VegliaCounty
	CZadar扎达尔() zadar.ZadarCounty
}

type 克罗地亚CroatiaDuke struct {
	feud.BaseDuke
	塞尼Senj    senj.SenjCounty
	斯普利特Split split.SplitCounty
	韦利亚Veglia veglia.VegliaCounty
	扎达尔Zadar  zadar.ZadarCounty
}

func (d *克罗地亚CroatiaDuke) CSenj塞尼() senj.SenjCounty {
	return d.塞尼Senj
}

func (d *克罗地亚CroatiaDuke) CSplit斯普利特() split.SplitCounty {
	return d.斯普利特Split
}

func (d *克罗地亚CroatiaDuke) CVeglia韦利亚() veglia.VegliaCounty {
	return d.韦利亚Veglia
}

func (d *克罗地亚CroatiaDuke) CZadar扎达尔() zadar.ZadarCounty {
	return d.扎达尔Zadar
}

var DCroatia克罗地亚 CroatiaDuke = &克罗地亚CroatiaDuke{}

func init() {
	f := DCroatia克罗地亚.(*克罗地亚CroatiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "croatia",
		TitleName: "克罗地亚",
		TitleCode: "d_croatia",
		Counties:  map[string]feud.County{},
	}

	f.塞尼Senj = senj.CSenj塞尼
	f.塞尼Senj.SetParent(f)

	f.斯普利特Split = split.CSplit斯普利特
	f.斯普利特Split.SetParent(f)

	f.韦利亚Veglia = veglia.CVeglia韦利亚
	f.韦利亚Veglia.SetParent(f)

	f.扎达尔Zadar = zadar.CZadar扎达尔
	f.扎达尔Zadar.SetParent(f)

}
