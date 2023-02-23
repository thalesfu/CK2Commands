package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SplitCounty interface {
	feud.County
	BHvar赫瓦尔() feud.Barony
	BKlis克利斯() feud.Barony
	BMakarska马卡尔斯卡() feud.Barony
	BSinj西尼() feud.Barony
	BSolin索林() feud.Barony
	BSplit斯普利特() feud.Barony
	BTrogir特罗吉尔() feud.Barony
}

type 斯普利特SplitCounty struct {
	feud.BaseCounty
	赫瓦尔Hvar       feud.Barony
	克利斯Klis       feud.Barony
	马卡尔斯卡Makarska feud.Barony
	西尼Sinj        feud.Barony
	索林Solin       feud.Barony
	斯普利特Split     feud.Barony
	特罗吉尔Trogir    feud.Barony
}

func (c *斯普利特SplitCounty) BHvar赫瓦尔() feud.Barony {
	return c.赫瓦尔Hvar
}

func (c *斯普利特SplitCounty) BKlis克利斯() feud.Barony {
	return c.克利斯Klis
}

func (c *斯普利特SplitCounty) BMakarska马卡尔斯卡() feud.Barony {
	return c.马卡尔斯卡Makarska
}

func (c *斯普利特SplitCounty) BSinj西尼() feud.Barony {
	return c.西尼Sinj
}

func (c *斯普利特SplitCounty) BSolin索林() feud.Barony {
	return c.索林Solin
}

func (c *斯普利特SplitCounty) BSplit斯普利特() feud.Barony {
	return c.斯普利特Split
}

func (c *斯普利特SplitCounty) BTrogir特罗吉尔() feud.Barony {
	return c.特罗吉尔Trogir
}

var CSplit斯普利特 SplitCounty = &斯普利特SplitCounty{}

func init() {
	f := CSplit斯普利特.(*斯普利特SplitCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "467",
		Title:     "split",
		TitleName: "斯普利特",
		TitleCode: "c_split",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫瓦尔Hvar = BHvar赫瓦尔
	f.赫瓦尔Hvar.SetParent(f)

	f.克利斯Klis = BKlis克利斯
	f.克利斯Klis.SetParent(f)

	f.马卡尔斯卡Makarska = BMakarska马卡尔斯卡
	f.马卡尔斯卡Makarska.SetParent(f)

	f.西尼Sinj = BSinj西尼
	f.西尼Sinj.SetParent(f)

	f.索林Solin = BSolin索林
	f.索林Solin.SetParent(f)

	f.斯普利特Split = BSplit斯普利特
	f.斯普利特Split.SetParent(f)

	f.特罗吉尔Trogir = BTrogir特罗吉尔
	f.特罗吉尔Trogir.SetParent(f)

}
