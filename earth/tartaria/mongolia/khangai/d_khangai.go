package khangai

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/khangai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/otgontenger"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/tarvagatai"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhangaiDuke interface {
	feud.Duke
	CKhangai杭海() khangai.KhangaiCounty
	COtgontenger鄂特冈腾格里() otgontenger.OtgontengerCounty
	CTarvagatai塔尔巴嘎泰() tarvagatai.TarvagataiCounty
}

type 杭海KhangaiDuke struct {
	feud.BaseDuke
	杭海Khangai         khangai.KhangaiCounty
	鄂特冈腾格里Otgontenger otgontenger.OtgontengerCounty
	塔尔巴嘎泰Tarvagatai   tarvagatai.TarvagataiCounty
}

func (d *杭海KhangaiDuke) CKhangai杭海() khangai.KhangaiCounty {
	return d.杭海Khangai
}

func (d *杭海KhangaiDuke) COtgontenger鄂特冈腾格里() otgontenger.OtgontengerCounty {
	return d.鄂特冈腾格里Otgontenger
}

func (d *杭海KhangaiDuke) CTarvagatai塔尔巴嘎泰() tarvagatai.TarvagataiCounty {
	return d.塔尔巴嘎泰Tarvagatai
}

var DKhangai杭海 KhangaiDuke = &杭海KhangaiDuke{}

func init() {
	f := DKhangai杭海.(*杭海KhangaiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khangai",
		TitleName: "杭海",
		TitleCode: "d_khangai",
		Counties:  map[string]feud.County{},
	}

	f.杭海Khangai = khangai.CKhangai杭海
	f.杭海Khangai.SetParent(f)

	f.鄂特冈腾格里Otgontenger = otgontenger.COtgontenger鄂特冈腾格里
	f.鄂特冈腾格里Otgontenger.SetParent(f)

	f.塔尔巴嘎泰Tarvagatai = tarvagatai.CTarvagatai塔尔巴嘎泰
	f.塔尔巴嘎泰Tarvagatai.SetParent(f)

}
