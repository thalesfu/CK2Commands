package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolhapurCounty interface {
	feud.County
	BKarahataka卡拉荷塔克() feud.Barony
	BKolhapur拘罗诃补罗() feud.Barony
	BKurundaka鸠伦陀迦() feud.Barony
	BManapura马纳普拉() feud.Barony
	BMiraj米勒杰() feud.Barony
	BPranala普罗纳拉() feud.Barony
	BSatara萨达拉() feud.Barony
}

type 拘罗诃补罗KolhapurCounty struct {
	feud.BaseCounty
	卡拉荷塔克Karahataka feud.Barony
	拘罗诃补罗Kolhapur   feud.Barony
	鸠伦陀迦Kurundaka   feud.Barony
	马纳普拉Manapura    feud.Barony
	米勒杰Miraj        feud.Barony
	普罗纳拉Pranala     feud.Barony
	萨达拉Satara       feud.Barony
}

func (c *拘罗诃补罗KolhapurCounty) BKarahataka卡拉荷塔克() feud.Barony {
	return c.卡拉荷塔克Karahataka
}

func (c *拘罗诃补罗KolhapurCounty) BKolhapur拘罗诃补罗() feud.Barony {
	return c.拘罗诃补罗Kolhapur
}

func (c *拘罗诃补罗KolhapurCounty) BKurundaka鸠伦陀迦() feud.Barony {
	return c.鸠伦陀迦Kurundaka
}

func (c *拘罗诃补罗KolhapurCounty) BManapura马纳普拉() feud.Barony {
	return c.马纳普拉Manapura
}

func (c *拘罗诃补罗KolhapurCounty) BMiraj米勒杰() feud.Barony {
	return c.米勒杰Miraj
}

func (c *拘罗诃补罗KolhapurCounty) BPranala普罗纳拉() feud.Barony {
	return c.普罗纳拉Pranala
}

func (c *拘罗诃补罗KolhapurCounty) BSatara萨达拉() feud.Barony {
	return c.萨达拉Satara
}

var CKolhapur拘罗诃补罗 KolhapurCounty = &拘罗诃补罗KolhapurCounty{}

func init() {
	f := CKolhapur拘罗诃补罗.(*拘罗诃补罗KolhapurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1416",
		Title:     "kolhapur",
		TitleName: "拘罗诃补罗",
		TitleCode: "c_kolhapur",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉荷塔克Karahataka = BKarahataka卡拉荷塔克
	f.卡拉荷塔克Karahataka.SetParent(f)

	f.拘罗诃补罗Kolhapur = BKolhapur拘罗诃补罗
	f.拘罗诃补罗Kolhapur.SetParent(f)

	f.鸠伦陀迦Kurundaka = BKurundaka鸠伦陀迦
	f.鸠伦陀迦Kurundaka.SetParent(f)

	f.马纳普拉Manapura = BManapura马纳普拉
	f.马纳普拉Manapura.SetParent(f)

	f.米勒杰Miraj = BMiraj米勒杰
	f.米勒杰Miraj.SetParent(f)

	f.普罗纳拉Pranala = BPranala普罗纳拉
	f.普罗纳拉Pranala.SetParent(f)

	f.萨达拉Satara = BSatara萨达拉
	f.萨达拉Satara.SetParent(f)

}
