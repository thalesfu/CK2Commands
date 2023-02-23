package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SortavalaCounty interface {
	feud.County
	BKarmalan卡马兰() feud.Barony
	BKitee基泰() feud.Barony
	BLakhdenpokhia拉赫坚波希亚() feud.Barony
	BLappeenranta拉彭兰塔() feud.Barony
	BRuskeala鲁斯克阿拉() feud.Barony
	BSortavala索尔塔瓦拉() feud.Barony
	BYanis亚尼斯() feud.Barony
}

type 索尔塔瓦拉SortavalaCounty struct {
	feud.BaseCounty
	卡马兰Karmalan         feud.Barony
	基泰Kitee             feud.Barony
	拉赫坚波希亚Lakhdenpokhia feud.Barony
	拉彭兰塔Lappeenranta    feud.Barony
	鲁斯克阿拉Ruskeala       feud.Barony
	索尔塔瓦拉Sortavala      feud.Barony
	亚尼斯Yanis            feud.Barony
}

func (c *索尔塔瓦拉SortavalaCounty) BKarmalan卡马兰() feud.Barony {
	return c.卡马兰Karmalan
}

func (c *索尔塔瓦拉SortavalaCounty) BKitee基泰() feud.Barony {
	return c.基泰Kitee
}

func (c *索尔塔瓦拉SortavalaCounty) BLakhdenpokhia拉赫坚波希亚() feud.Barony {
	return c.拉赫坚波希亚Lakhdenpokhia
}

func (c *索尔塔瓦拉SortavalaCounty) BLappeenranta拉彭兰塔() feud.Barony {
	return c.拉彭兰塔Lappeenranta
}

func (c *索尔塔瓦拉SortavalaCounty) BRuskeala鲁斯克阿拉() feud.Barony {
	return c.鲁斯克阿拉Ruskeala
}

func (c *索尔塔瓦拉SortavalaCounty) BSortavala索尔塔瓦拉() feud.Barony {
	return c.索尔塔瓦拉Sortavala
}

func (c *索尔塔瓦拉SortavalaCounty) BYanis亚尼斯() feud.Barony {
	return c.亚尼斯Yanis
}

var CSortavala索尔塔瓦拉 SortavalaCounty = &索尔塔瓦拉SortavalaCounty{}

func init() {
	f := CSortavala索尔塔瓦拉.(*索尔塔瓦拉SortavalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1602",
		Title:     "sortavala",
		TitleName: "索尔塔瓦拉",
		TitleCode: "c_sortavala",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡马兰Karmalan = BKarmalan卡马兰
	f.卡马兰Karmalan.SetParent(f)

	f.基泰Kitee = BKitee基泰
	f.基泰Kitee.SetParent(f)

	f.拉赫坚波希亚Lakhdenpokhia = BLakhdenpokhia拉赫坚波希亚
	f.拉赫坚波希亚Lakhdenpokhia.SetParent(f)

	f.拉彭兰塔Lappeenranta = BLappeenranta拉彭兰塔
	f.拉彭兰塔Lappeenranta.SetParent(f)

	f.鲁斯克阿拉Ruskeala = BRuskeala鲁斯克阿拉
	f.鲁斯克阿拉Ruskeala.SetParent(f)

	f.索尔塔瓦拉Sortavala = BSortavala索尔塔瓦拉
	f.索尔塔瓦拉Sortavala.SetParent(f)

	f.亚尼斯Yanis = BYanis亚尼斯
	f.亚尼斯Yanis.SetParent(f)

}
