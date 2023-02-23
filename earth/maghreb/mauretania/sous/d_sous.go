package sous

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sous/draa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sous/ifni"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sous/tharasset"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sous/warzazat"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SousDuke interface {
	feud.Duke
	CDraa德拉() draa.DraaCounty
	CIfni伊夫尼() ifni.IfniCounty
	CTharasset苏斯() tharasset.TharassetCounty
	CWarzazat瓦尔扎扎特() warzazat.WarzazatCounty
}

type 苏斯SousDuke struct {
	feud.BaseDuke
	德拉Draa        draa.DraaCounty
	伊夫尼Ifni       ifni.IfniCounty
	苏斯Tharasset   tharasset.TharassetCounty
	瓦尔扎扎特Warzazat warzazat.WarzazatCounty
}

func (d *苏斯SousDuke) CDraa德拉() draa.DraaCounty {
	return d.德拉Draa
}

func (d *苏斯SousDuke) CIfni伊夫尼() ifni.IfniCounty {
	return d.伊夫尼Ifni
}

func (d *苏斯SousDuke) CTharasset苏斯() tharasset.TharassetCounty {
	return d.苏斯Tharasset
}

func (d *苏斯SousDuke) CWarzazat瓦尔扎扎特() warzazat.WarzazatCounty {
	return d.瓦尔扎扎特Warzazat
}

var DSous苏斯 SousDuke = &苏斯SousDuke{}

func init() {
	f := DSous苏斯.(*苏斯SousDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sous",
		TitleName: "苏斯",
		TitleCode: "d_sous",
		Counties:  map[string]feud.County{},
	}

	f.德拉Draa = draa.CDraa德拉
	f.德拉Draa.SetParent(f)

	f.伊夫尼Ifni = ifni.CIfni伊夫尼
	f.伊夫尼Ifni.SetParent(f)

	f.苏斯Tharasset = tharasset.CTharasset苏斯
	f.苏斯Tharasset.SetParent(f)

	f.瓦尔扎扎特Warzazat = warzazat.CWarzazat瓦尔扎扎特
	f.瓦尔扎扎特Warzazat.SetParent(f)

}
