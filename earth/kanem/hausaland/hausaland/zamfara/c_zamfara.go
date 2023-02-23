package zamfara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZamfaraCounty interface {
	feud.County
	BAnka安卡() feud.Barony
	BDatsi达齐() feud.Barony
	BGusau古绍() feud.Barony
	BJagbada贾格巴巴() feud.Barony
	BKori科里() feud.Barony
	BUburukiti乌布鲁基蒂() feud.Barony
}

type 扎姆法拉ZamfaraCounty struct {
	feud.BaseCounty
	安卡Anka         feud.Barony
	达齐Datsi        feud.Barony
	古绍Gusau        feud.Barony
	贾格巴巴Jagbada    feud.Barony
	科里Kori         feud.Barony
	乌布鲁基蒂Uburukiti feud.Barony
}

func (c *扎姆法拉ZamfaraCounty) BAnka安卡() feud.Barony {
	return c.安卡Anka
}

func (c *扎姆法拉ZamfaraCounty) BDatsi达齐() feud.Barony {
	return c.达齐Datsi
}

func (c *扎姆法拉ZamfaraCounty) BGusau古绍() feud.Barony {
	return c.古绍Gusau
}

func (c *扎姆法拉ZamfaraCounty) BJagbada贾格巴巴() feud.Barony {
	return c.贾格巴巴Jagbada
}

func (c *扎姆法拉ZamfaraCounty) BKori科里() feud.Barony {
	return c.科里Kori
}

func (c *扎姆法拉ZamfaraCounty) BUburukiti乌布鲁基蒂() feud.Barony {
	return c.乌布鲁基蒂Uburukiti
}

var CZamfara扎姆法拉 ZamfaraCounty = &扎姆法拉ZamfaraCounty{}

func init() {
	f := CZamfara扎姆法拉.(*扎姆法拉ZamfaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1750",
		Title:     "zamfara",
		TitleName: "扎姆法拉",
		TitleCode: "c_zamfara",
		Baronies:  map[string]feud.Barony{},
	}

	f.安卡Anka = BAnka安卡
	f.安卡Anka.SetParent(f)

	f.达齐Datsi = BDatsi达齐
	f.达齐Datsi.SetParent(f)

	f.古绍Gusau = BGusau古绍
	f.古绍Gusau.SetParent(f)

	f.贾格巴巴Jagbada = BJagbada贾格巴巴
	f.贾格巴巴Jagbada.SetParent(f)

	f.科里Kori = BKori科里
	f.科里Kori.SetParent(f)

	f.乌布鲁基蒂Uburukiti = BUburukiti乌布鲁基蒂
	f.乌布鲁基蒂Uburukiti.SetParent(f)

}
