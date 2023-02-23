package bessarabia

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/bessarabia/belgorod"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/bessarabia/birlad"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/bessarabia/chilia"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/bessarabia/galaz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BessarabiaDuke interface {
	feud.Duke
	CBelgorod别尔哥罗德() belgorod.BelgorodCounty
	CBirlad伯尔拉德() birlad.BirladCounty
	CChilia基利亚() chilia.ChiliaCounty
	CGalaz加拉茨() galaz.GalazCounty
}

type 比萨拉比亚BessarabiaDuke struct {
	feud.BaseDuke
	别尔哥罗德Belgorod belgorod.BelgorodCounty
	伯尔拉德Birlad    birlad.BirladCounty
	基利亚Chilia     chilia.ChiliaCounty
	加拉茨Galaz      galaz.GalazCounty
}

func (d *比萨拉比亚BessarabiaDuke) CBelgorod别尔哥罗德() belgorod.BelgorodCounty {
	return d.别尔哥罗德Belgorod
}

func (d *比萨拉比亚BessarabiaDuke) CBirlad伯尔拉德() birlad.BirladCounty {
	return d.伯尔拉德Birlad
}

func (d *比萨拉比亚BessarabiaDuke) CChilia基利亚() chilia.ChiliaCounty {
	return d.基利亚Chilia
}

func (d *比萨拉比亚BessarabiaDuke) CGalaz加拉茨() galaz.GalazCounty {
	return d.加拉茨Galaz
}

var DBessarabia比萨拉比亚 BessarabiaDuke = &比萨拉比亚BessarabiaDuke{}

func init() {
	f := DBessarabia比萨拉比亚.(*比萨拉比亚BessarabiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bessarabia",
		TitleName: "比萨拉比亚",
		TitleCode: "d_bessarabia",
		Counties:  map[string]feud.County{},
	}

	f.别尔哥罗德Belgorod = belgorod.CBelgorod别尔哥罗德
	f.别尔哥罗德Belgorod.SetParent(f)

	f.伯尔拉德Birlad = birlad.CBirlad伯尔拉德
	f.伯尔拉德Birlad.SetParent(f)

	f.基利亚Chilia = chilia.CChilia基利亚
	f.基利亚Chilia.SetParent(f)

	f.加拉茨Galaz = galaz.CGalaz加拉茨
	f.加拉茨Galaz.SetParent(f)

}
