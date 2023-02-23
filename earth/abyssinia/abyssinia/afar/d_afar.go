package afar

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/afar/asayita"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/afar/assab"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/afar/harer"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/afar/tadjoura"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AfarDuke interface {
	feud.Duke
	CAsayita阿塞塔() asayita.AsayitaCounty
	CAssab阿萨布() assab.AssabCounty
	CHarer哈勒尔() harer.HarerCounty
	CTadjoura塔朱拉() tadjoura.TadjouraCounty
}

type 阿法尔AfarDuke struct {
	feud.BaseDuke
	阿塞塔Asayita  asayita.AsayitaCounty
	阿萨布Assab    assab.AssabCounty
	哈勒尔Harer    harer.HarerCounty
	塔朱拉Tadjoura tadjoura.TadjouraCounty
}

func (d *阿法尔AfarDuke) CAsayita阿塞塔() asayita.AsayitaCounty {
	return d.阿塞塔Asayita
}

func (d *阿法尔AfarDuke) CAssab阿萨布() assab.AssabCounty {
	return d.阿萨布Assab
}

func (d *阿法尔AfarDuke) CHarer哈勒尔() harer.HarerCounty {
	return d.哈勒尔Harer
}

func (d *阿法尔AfarDuke) CTadjoura塔朱拉() tadjoura.TadjouraCounty {
	return d.塔朱拉Tadjoura
}

var DAfar阿法尔 AfarDuke = &阿法尔AfarDuke{}

func init() {
	f := DAfar阿法尔.(*阿法尔AfarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "afar",
		TitleName: "阿法尔",
		TitleCode: "d_afar",
		Counties:  map[string]feud.County{},
	}

	f.阿塞塔Asayita = asayita.CAsayita阿塞塔
	f.阿塞塔Asayita.SetParent(f)

	f.阿萨布Assab = assab.CAssab阿萨布
	f.阿萨布Assab.SetParent(f)

	f.哈勒尔Harer = harer.CHarer哈勒尔
	f.哈勒尔Harer.SetParent(f)

	f.塔朱拉Tadjoura = tadjoura.CTadjoura塔朱拉
	f.塔朱拉Tadjoura.SetParent(f)

}
