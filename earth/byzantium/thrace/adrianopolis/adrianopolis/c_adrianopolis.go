package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AdrianopolisCounty interface {
	feud.County
	BAdrianopolis哈德良波利斯() feud.Barony
	BBerat培拉特() feud.Barony
	BDemotika德莫提卡() feud.Barony
	BDidymoteichon迪蒂莫迪索() feud.Barony
	BKypsela基普瑟拉() feud.Barony
	BOrestias奥雷斯蒂阿斯() feud.Barony
	BSkalothe斯科洛特() feud.Barony
}

type 哈德良波利斯AdrianopolisCounty struct {
	feud.BaseCounty
	哈德良波利斯Adrianopolis feud.Barony
	培拉特Berat           feud.Barony
	德莫提卡Demotika       feud.Barony
	迪蒂莫迪索Didymoteichon feud.Barony
	基普瑟拉Kypsela        feud.Barony
	奥雷斯蒂阿斯Orestias     feud.Barony
	斯科洛特Skalothe       feud.Barony
}

func (c *哈德良波利斯AdrianopolisCounty) BAdrianopolis哈德良波利斯() feud.Barony {
	return c.哈德良波利斯Adrianopolis
}

func (c *哈德良波利斯AdrianopolisCounty) BBerat培拉特() feud.Barony {
	return c.培拉特Berat
}

func (c *哈德良波利斯AdrianopolisCounty) BDemotika德莫提卡() feud.Barony {
	return c.德莫提卡Demotika
}

func (c *哈德良波利斯AdrianopolisCounty) BDidymoteichon迪蒂莫迪索() feud.Barony {
	return c.迪蒂莫迪索Didymoteichon
}

func (c *哈德良波利斯AdrianopolisCounty) BKypsela基普瑟拉() feud.Barony {
	return c.基普瑟拉Kypsela
}

func (c *哈德良波利斯AdrianopolisCounty) BOrestias奥雷斯蒂阿斯() feud.Barony {
	return c.奥雷斯蒂阿斯Orestias
}

func (c *哈德良波利斯AdrianopolisCounty) BSkalothe斯科洛特() feud.Barony {
	return c.斯科洛特Skalothe
}

var CAdrianopolis哈德良波利斯 AdrianopolisCounty = &哈德良波利斯AdrianopolisCounty{}

func init() {
	f := CAdrianopolis哈德良波利斯.(*哈德良波利斯AdrianopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "494",
		Title:     "adrianopolis",
		TitleName: "哈德良波利斯",
		TitleCode: "c_adrianopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈德良波利斯Adrianopolis = BAdrianopolis哈德良波利斯
	f.哈德良波利斯Adrianopolis.SetParent(f)

	f.培拉特Berat = BBerat培拉特
	f.培拉特Berat.SetParent(f)

	f.德莫提卡Demotika = BDemotika德莫提卡
	f.德莫提卡Demotika.SetParent(f)

	f.迪蒂莫迪索Didymoteichon = BDidymoteichon迪蒂莫迪索
	f.迪蒂莫迪索Didymoteichon.SetParent(f)

	f.基普瑟拉Kypsela = BKypsela基普瑟拉
	f.基普瑟拉Kypsela.SetParent(f)

	f.奥雷斯蒂阿斯Orestias = BOrestias奥雷斯蒂阿斯
	f.奥雷斯蒂阿斯Orestias.SetParent(f)

	f.斯科洛特Skalothe = BSkalothe斯科洛特
	f.斯科洛特Skalothe.SetParent(f)

}
