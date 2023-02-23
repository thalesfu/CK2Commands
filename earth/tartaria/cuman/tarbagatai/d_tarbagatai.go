package tarbagatai

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/tarbagatai/aylik"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/tarbagatai/saur"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/tarbagatai/tarbagatai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/tarbagatai/urzhar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TarbagataiDuke interface {
	feud.Duke
	CAylik艾里克() aylik.AylikCounty
	CSaur萨吾尔() saur.SaurCounty
	CTarbagatai塔尔巴哈台() tarbagatai.TarbagataiCounty
	CUrzhar阿亚古兹() urzhar.UrzharCounty
}

type 塔尔巴哈台TarbagataiDuke struct {
	feud.BaseDuke
	艾里克Aylik        aylik.AylikCounty
	萨吾尔Saur         saur.SaurCounty
	塔尔巴哈台Tarbagatai tarbagatai.TarbagataiCounty
	阿亚古兹Urzhar      urzhar.UrzharCounty
}

func (d *塔尔巴哈台TarbagataiDuke) CAylik艾里克() aylik.AylikCounty {
	return d.艾里克Aylik
}

func (d *塔尔巴哈台TarbagataiDuke) CSaur萨吾尔() saur.SaurCounty {
	return d.萨吾尔Saur
}

func (d *塔尔巴哈台TarbagataiDuke) CTarbagatai塔尔巴哈台() tarbagatai.TarbagataiCounty {
	return d.塔尔巴哈台Tarbagatai
}

func (d *塔尔巴哈台TarbagataiDuke) CUrzhar阿亚古兹() urzhar.UrzharCounty {
	return d.阿亚古兹Urzhar
}

var DTarbagatai塔尔巴哈台 TarbagataiDuke = &塔尔巴哈台TarbagataiDuke{}

func init() {
	f := DTarbagatai塔尔巴哈台.(*塔尔巴哈台TarbagataiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tarbagatai",
		TitleName: "塔尔巴哈台",
		TitleCode: "d_tarbagatai",
		Counties:  map[string]feud.County{},
	}

	f.艾里克Aylik = aylik.CAylik艾里克
	f.艾里克Aylik.SetParent(f)

	f.萨吾尔Saur = saur.CSaur萨吾尔
	f.萨吾尔Saur.SetParent(f)

	f.塔尔巴哈台Tarbagatai = tarbagatai.CTarbagatai塔尔巴哈台
	f.塔尔巴哈台Tarbagatai.SetParent(f)

	f.阿亚古兹Urzhar = urzhar.CUrzhar阿亚古兹
	f.阿亚古兹Urzhar.SetParent(f)

}
