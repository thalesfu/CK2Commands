package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NovogrodekCounty interface {
	feud.County
	BDziatlava佳特洛沃() feud.Barony
	BIwie伊维耶() feud.Barony
	BKarelichy科列利奇() feud.Barony
	BLida利达() feud.Barony
	BNovogrodek新格鲁多克() feud.Barony
	BShchuchyn休钦() feud.Barony
	BSlonim斯洛尼姆() feud.Barony
}

type 新格鲁多克NovogrodekCounty struct {
	feud.BaseCounty
	佳特洛沃Dziatlava   feud.Barony
	伊维耶Iwie         feud.Barony
	科列利奇Karelichy   feud.Barony
	利达Lida          feud.Barony
	新格鲁多克Novogrodek feud.Barony
	休钦Shchuchyn     feud.Barony
	斯洛尼姆Slonim      feud.Barony
}

func (c *新格鲁多克NovogrodekCounty) BDziatlava佳特洛沃() feud.Barony {
	return c.佳特洛沃Dziatlava
}

func (c *新格鲁多克NovogrodekCounty) BIwie伊维耶() feud.Barony {
	return c.伊维耶Iwie
}

func (c *新格鲁多克NovogrodekCounty) BKarelichy科列利奇() feud.Barony {
	return c.科列利奇Karelichy
}

func (c *新格鲁多克NovogrodekCounty) BLida利达() feud.Barony {
	return c.利达Lida
}

func (c *新格鲁多克NovogrodekCounty) BNovogrodek新格鲁多克() feud.Barony {
	return c.新格鲁多克Novogrodek
}

func (c *新格鲁多克NovogrodekCounty) BShchuchyn休钦() feud.Barony {
	return c.休钦Shchuchyn
}

func (c *新格鲁多克NovogrodekCounty) BSlonim斯洛尼姆() feud.Barony {
	return c.斯洛尼姆Slonim
}

var CNovogrodek新格鲁多克 NovogrodekCounty = &新格鲁多克NovogrodekCounty{}

func init() {
	f := CNovogrodek新格鲁多克.(*新格鲁多克NovogrodekCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1654",
		Title:     "novogrodek",
		TitleName: "新格鲁多克",
		TitleCode: "c_novogrodek",
		Baronies:  map[string]feud.Barony{},
	}

	f.佳特洛沃Dziatlava = BDziatlava佳特洛沃
	f.佳特洛沃Dziatlava.SetParent(f)

	f.伊维耶Iwie = BIwie伊维耶
	f.伊维耶Iwie.SetParent(f)

	f.科列利奇Karelichy = BKarelichy科列利奇
	f.科列利奇Karelichy.SetParent(f)

	f.利达Lida = BLida利达
	f.利达Lida.SetParent(f)

	f.新格鲁多克Novogrodek = BNovogrodek新格鲁多克
	f.新格鲁多克Novogrodek.SetParent(f)

	f.休钦Shchuchyn = BShchuchyn休钦
	f.休钦Shchuchyn.SetParent(f)

	f.斯洛尼姆Slonim = BSlonim斯洛尼姆
	f.斯洛尼姆Slonim.SetParent(f)

}
