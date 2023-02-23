package magadha

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/magadha/gaya"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/magadha/magadha"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/magadha/mudgagiri"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/magadha/sasaram"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MagadhaDuke interface {
	feud.Duke
	CGaya伽耶() gaya.GayaCounty
	CMagadha摩揭陀() magadha.MagadhaCounty
	CMudgagiri勿伽耆厘() mudgagiri.MudgagiriCounty
	CSasaram萨娑蓝() sasaram.SasaramCounty
}

type 摩揭陀MagadhaDuke struct {
	feud.BaseDuke
	伽耶Gaya        gaya.GayaCounty
	摩揭陀Magadha    magadha.MagadhaCounty
	勿伽耆厘Mudgagiri mudgagiri.MudgagiriCounty
	萨娑蓝Sasaram    sasaram.SasaramCounty
}

func (d *摩揭陀MagadhaDuke) CGaya伽耶() gaya.GayaCounty {
	return d.伽耶Gaya
}

func (d *摩揭陀MagadhaDuke) CMagadha摩揭陀() magadha.MagadhaCounty {
	return d.摩揭陀Magadha
}

func (d *摩揭陀MagadhaDuke) CMudgagiri勿伽耆厘() mudgagiri.MudgagiriCounty {
	return d.勿伽耆厘Mudgagiri
}

func (d *摩揭陀MagadhaDuke) CSasaram萨娑蓝() sasaram.SasaramCounty {
	return d.萨娑蓝Sasaram
}

var DMagadha摩揭陀 MagadhaDuke = &摩揭陀MagadhaDuke{}

func init() {
	f := DMagadha摩揭陀.(*摩揭陀MagadhaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "magadha",
		TitleName: "摩揭陀",
		TitleCode: "d_magadha",
		Counties:  map[string]feud.County{},
	}

	f.伽耶Gaya = gaya.CGaya伽耶
	f.伽耶Gaya.SetParent(f)

	f.摩揭陀Magadha = magadha.CMagadha摩揭陀
	f.摩揭陀Magadha.SetParent(f)

	f.勿伽耆厘Mudgagiri = mudgagiri.CMudgagiri勿伽耆厘
	f.勿伽耆厘Mudgagiri.SetParent(f)

	f.萨娑蓝Sasaram = sasaram.CSasaram萨娑蓝
	f.萨娑蓝Sasaram.SetParent(f)

}
