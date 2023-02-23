package slavonia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/slavonia/krizevci"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/slavonia/varadzin"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia/slavonia/zagreb"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SlavoniaDuke interface {
	feud.Duke
	CKrizevci克里热夫齐() krizevci.KrizevciCounty
	CVaradzin瓦拉日丁() varadzin.VaradzinCounty
	CZagreb扎格雷布() zagreb.ZagrebCounty
}

type 斯拉沃尼亚SlavoniaDuke struct {
	feud.BaseDuke
	克里热夫齐Krizevci krizevci.KrizevciCounty
	瓦拉日丁Varadzin  varadzin.VaradzinCounty
	扎格雷布Zagreb    zagreb.ZagrebCounty
}

func (d *斯拉沃尼亚SlavoniaDuke) CKrizevci克里热夫齐() krizevci.KrizevciCounty {
	return d.克里热夫齐Krizevci
}

func (d *斯拉沃尼亚SlavoniaDuke) CVaradzin瓦拉日丁() varadzin.VaradzinCounty {
	return d.瓦拉日丁Varadzin
}

func (d *斯拉沃尼亚SlavoniaDuke) CZagreb扎格雷布() zagreb.ZagrebCounty {
	return d.扎格雷布Zagreb
}

var DSlavonia斯拉沃尼亚 SlavoniaDuke = &斯拉沃尼亚SlavoniaDuke{}

func init() {
	f := DSlavonia斯拉沃尼亚.(*斯拉沃尼亚SlavoniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "slavonia",
		TitleName: "斯拉沃尼亚",
		TitleCode: "d_slavonia",
		Counties:  map[string]feud.County{},
	}

	f.克里热夫齐Krizevci = krizevci.CKrizevci克里热夫齐
	f.克里热夫齐Krizevci.SetParent(f)

	f.瓦拉日丁Varadzin = varadzin.CVaradzin瓦拉日丁
	f.瓦拉日丁Varadzin.SetParent(f)

	f.扎格雷布Zagreb = zagreb.CZagreb扎格雷布
	f.扎格雷布Zagreb.SetParent(f)

}
