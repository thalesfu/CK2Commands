package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NyssaCounty interface {
	feud.County
	BArchelais阿耳刻莱斯() feud.Barony
	BKoron科龙() feud.Barony
	BMalakopeia马拉科皮亚() feud.Barony
	BNazianus纳齐安() feud.Barony
	BNyssa尼撒() feud.Barony
	BParnassus帕耳那索斯() feud.Barony
	BThebasa塞巴萨() feud.Barony
}

type 尼撒NyssaCounty struct {
	feud.BaseCounty
	阿耳刻莱斯Archelais  feud.Barony
	科龙Koron         feud.Barony
	马拉科皮亚Malakopeia feud.Barony
	纳齐安Nazianus     feud.Barony
	尼撒Nyssa         feud.Barony
	帕耳那索斯Parnassus  feud.Barony
	塞巴萨Thebasa      feud.Barony
}

func (c *尼撒NyssaCounty) BArchelais阿耳刻莱斯() feud.Barony {
	return c.阿耳刻莱斯Archelais
}

func (c *尼撒NyssaCounty) BKoron科龙() feud.Barony {
	return c.科龙Koron
}

func (c *尼撒NyssaCounty) BMalakopeia马拉科皮亚() feud.Barony {
	return c.马拉科皮亚Malakopeia
}

func (c *尼撒NyssaCounty) BNazianus纳齐安() feud.Barony {
	return c.纳齐安Nazianus
}

func (c *尼撒NyssaCounty) BNyssa尼撒() feud.Barony {
	return c.尼撒Nyssa
}

func (c *尼撒NyssaCounty) BParnassus帕耳那索斯() feud.Barony {
	return c.帕耳那索斯Parnassus
}

func (c *尼撒NyssaCounty) BThebasa塞巴萨() feud.Barony {
	return c.塞巴萨Thebasa
}

var CNyssa尼撒 NyssaCounty = &尼撒NyssaCounty{}

func init() {
	f := CNyssa尼撒.(*尼撒NyssaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1933",
		Title:     "nyssa",
		TitleName: "尼撒",
		TitleCode: "c_nyssa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿耳刻莱斯Archelais = BArchelais阿耳刻莱斯
	f.阿耳刻莱斯Archelais.SetParent(f)

	f.科龙Koron = BKoron科龙
	f.科龙Koron.SetParent(f)

	f.马拉科皮亚Malakopeia = BMalakopeia马拉科皮亚
	f.马拉科皮亚Malakopeia.SetParent(f)

	f.纳齐安Nazianus = BNazianus纳齐安
	f.纳齐安Nazianus.SetParent(f)

	f.尼撒Nyssa = BNyssa尼撒
	f.尼撒Nyssa.SetParent(f)

	f.帕耳那索斯Parnassus = BParnassus帕耳那索斯
	f.帕耳那索斯Parnassus.SetParent(f)

	f.塞巴萨Thebasa = BThebasa塞巴萨
	f.塞巴萨Thebasa.SetParent(f)

}
