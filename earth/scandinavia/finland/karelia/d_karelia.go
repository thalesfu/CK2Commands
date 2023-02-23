package karelia

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/karelia/karelen"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/karelia/kexholm"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/karelia/onega"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/karelia/sortavala"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KareliaDuke interface {
	feud.Duke
	CKarelen凯姆() karelen.KarelenCounty
	CKexholm科克斯霍姆() kexholm.KexholmCounty
	COnega奥洛涅茨() onega.OnegaCounty
	CSortavala索尔塔瓦拉() sortavala.SortavalaCounty
}

type 卡累利阿KareliaDuke struct {
	feud.BaseDuke
	凯姆Karelen      karelen.KarelenCounty
	科克斯霍姆Kexholm   kexholm.KexholmCounty
	奥洛涅茨Onega      onega.OnegaCounty
	索尔塔瓦拉Sortavala sortavala.SortavalaCounty
}

func (d *卡累利阿KareliaDuke) CKarelen凯姆() karelen.KarelenCounty {
	return d.凯姆Karelen
}

func (d *卡累利阿KareliaDuke) CKexholm科克斯霍姆() kexholm.KexholmCounty {
	return d.科克斯霍姆Kexholm
}

func (d *卡累利阿KareliaDuke) COnega奥洛涅茨() onega.OnegaCounty {
	return d.奥洛涅茨Onega
}

func (d *卡累利阿KareliaDuke) CSortavala索尔塔瓦拉() sortavala.SortavalaCounty {
	return d.索尔塔瓦拉Sortavala
}

var DKarelia卡累利阿 KareliaDuke = &卡累利阿KareliaDuke{}

func init() {
	f := DKarelia卡累利阿.(*卡累利阿KareliaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "karelia",
		TitleName: "卡累利阿",
		TitleCode: "d_karelia",
		Counties:  map[string]feud.County{},
	}

	f.凯姆Karelen = karelen.CKarelen凯姆
	f.凯姆Karelen.SetParent(f)

	f.科克斯霍姆Kexholm = kexholm.CKexholm科克斯霍姆
	f.科克斯霍姆Kexholm.SetParent(f)

	f.奥洛涅茨Onega = onega.COnega奥洛涅茨
	f.奥洛涅茨Onega.SetParent(f)

	f.索尔塔瓦拉Sortavala = sortavala.CSortavala索尔塔瓦拉
	f.索尔塔瓦拉Sortavala.SetParent(f)

}
