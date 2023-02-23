package kebbi

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/kebbi/dyamare"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/kebbi/kebbi"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/kebbi/tahoua"
	"github.com/thalesfu/CK2Commands/earth/kanem/hausaland/kebbi/zarma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KebbiDuke interface {
	feud.Duke
	CDyamare迪亚马尔() dyamare.DyamareCounty
	CKebbi凯比() kebbi.KebbiCounty
	CTahoua塔瓦() tahoua.TahouaCounty
	CZarma扎尔马() zarma.ZarmaCounty
}

type 凯比KebbiDuke struct {
	feud.BaseDuke
	迪亚马尔Dyamare dyamare.DyamareCounty
	凯比Kebbi     kebbi.KebbiCounty
	塔瓦Tahoua    tahoua.TahouaCounty
	扎尔马Zarma    zarma.ZarmaCounty
}

func (d *凯比KebbiDuke) CDyamare迪亚马尔() dyamare.DyamareCounty {
	return d.迪亚马尔Dyamare
}

func (d *凯比KebbiDuke) CKebbi凯比() kebbi.KebbiCounty {
	return d.凯比Kebbi
}

func (d *凯比KebbiDuke) CTahoua塔瓦() tahoua.TahouaCounty {
	return d.塔瓦Tahoua
}

func (d *凯比KebbiDuke) CZarma扎尔马() zarma.ZarmaCounty {
	return d.扎尔马Zarma
}

var DKebbi凯比 KebbiDuke = &凯比KebbiDuke{}

func init() {
	f := DKebbi凯比.(*凯比KebbiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kebbi",
		TitleName: "凯比",
		TitleCode: "d_kebbi",
		Counties:  map[string]feud.County{},
	}

	f.迪亚马尔Dyamare = dyamare.CDyamare迪亚马尔
	f.迪亚马尔Dyamare.SetParent(f)

	f.凯比Kebbi = kebbi.CKebbi凯比
	f.凯比Kebbi.SetParent(f)

	f.塔瓦Tahoua = tahoua.CTahoua塔瓦
	f.塔瓦Tahoua.SetParent(f)

	f.扎尔马Zarma = zarma.CZarma扎尔马
	f.扎尔马Zarma.SetParent(f)

}
