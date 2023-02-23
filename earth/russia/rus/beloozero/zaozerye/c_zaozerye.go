package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZaozeryeCounty interface {
	feud.County
	BDevyatiny杰维亚季内() feud.Barony
	BKedra克德拉() feud.Barony
	BOshta奥什塔() feud.Barony
	BRodionovo罗季奥诺沃() feud.Barony
	BVytegra维捷格拉() feud.Barony
	BZarechnaya扎列奇纳亚() feud.Barony
}

type 扎奥焦里耶ZaozeryeCounty struct {
	feud.BaseCounty
	杰维亚季内Devyatiny  feud.Barony
	克德拉Kedra        feud.Barony
	奥什塔Oshta        feud.Barony
	罗季奥诺沃Rodionovo  feud.Barony
	维捷格拉Vytegra     feud.Barony
	扎列奇纳亚Zarechnaya feud.Barony
}

func (c *扎奥焦里耶ZaozeryeCounty) BDevyatiny杰维亚季内() feud.Barony {
	return c.杰维亚季内Devyatiny
}

func (c *扎奥焦里耶ZaozeryeCounty) BKedra克德拉() feud.Barony {
	return c.克德拉Kedra
}

func (c *扎奥焦里耶ZaozeryeCounty) BOshta奥什塔() feud.Barony {
	return c.奥什塔Oshta
}

func (c *扎奥焦里耶ZaozeryeCounty) BRodionovo罗季奥诺沃() feud.Barony {
	return c.罗季奥诺沃Rodionovo
}

func (c *扎奥焦里耶ZaozeryeCounty) BVytegra维捷格拉() feud.Barony {
	return c.维捷格拉Vytegra
}

func (c *扎奥焦里耶ZaozeryeCounty) BZarechnaya扎列奇纳亚() feud.Barony {
	return c.扎列奇纳亚Zarechnaya
}

var CZaozerye扎奥焦里耶 ZaozeryeCounty = &扎奥焦里耶ZaozeryeCounty{}

func init() {
	f := CZaozerye扎奥焦里耶.(*扎奥焦里耶ZaozeryeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "404",
		Title:     "zaozerye",
		TitleName: "扎奥焦里耶",
		TitleCode: "c_zaozerye",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰维亚季内Devyatiny = BDevyatiny杰维亚季内
	f.杰维亚季内Devyatiny.SetParent(f)

	f.克德拉Kedra = BKedra克德拉
	f.克德拉Kedra.SetParent(f)

	f.奥什塔Oshta = BOshta奥什塔
	f.奥什塔Oshta.SetParent(f)

	f.罗季奥诺沃Rodionovo = BRodionovo罗季奥诺沃
	f.罗季奥诺沃Rodionovo.SetParent(f)

	f.维捷格拉Vytegra = BVytegra维捷格拉
	f.维捷格拉Vytegra.SetParent(f)

	f.扎列奇纳亚Zarechnaya = BZarechnaya扎列奇纳亚
	f.扎列奇纳亚Zarechnaya.SetParent(f)

}
