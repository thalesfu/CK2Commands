package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KirghizCounty interface {
	feud.County
	BIrtysh曳咥() feud.Barony
	BKanbale坎巴列() feud.Barony
	BKurchatov库尔恰托夫() feud.Barony
	BOsinniki奥辛尼基() feud.Barony
	BSinnele辛涅列() feud.Barony
	BTele铁勒() feud.Barony
	BTelengit帖良古惕() feud.Barony
}

type 曳咥KirghizCounty struct {
	feud.BaseCounty
	曳咥Irtysh       feud.Barony
	坎巴列Kanbale     feud.Barony
	库尔恰托夫Kurchatov feud.Barony
	奥辛尼基Osinniki   feud.Barony
	辛涅列Sinnele     feud.Barony
	铁勒Tele         feud.Barony
	帖良古惕Telengit   feud.Barony
}

func (c *曳咥KirghizCounty) BIrtysh曳咥() feud.Barony {
	return c.曳咥Irtysh
}

func (c *曳咥KirghizCounty) BKanbale坎巴列() feud.Barony {
	return c.坎巴列Kanbale
}

func (c *曳咥KirghizCounty) BKurchatov库尔恰托夫() feud.Barony {
	return c.库尔恰托夫Kurchatov
}

func (c *曳咥KirghizCounty) BOsinniki奥辛尼基() feud.Barony {
	return c.奥辛尼基Osinniki
}

func (c *曳咥KirghizCounty) BSinnele辛涅列() feud.Barony {
	return c.辛涅列Sinnele
}

func (c *曳咥KirghizCounty) BTele铁勒() feud.Barony {
	return c.铁勒Tele
}

func (c *曳咥KirghizCounty) BTelengit帖良古惕() feud.Barony {
	return c.帖良古惕Telengit
}

var CKirghiz曳咥 KirghizCounty = &曳咥KirghizCounty{}

func init() {
	f := CKirghiz曳咥.(*曳咥KirghizCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1132",
		Title:     "kirghiz",
		TitleName: "曳咥",
		TitleCode: "c_kirghiz",
		Baronies:  map[string]feud.Barony{},
	}

	f.曳咥Irtysh = BIrtysh曳咥
	f.曳咥Irtysh.SetParent(f)

	f.坎巴列Kanbale = BKanbale坎巴列
	f.坎巴列Kanbale.SetParent(f)

	f.库尔恰托夫Kurchatov = BKurchatov库尔恰托夫
	f.库尔恰托夫Kurchatov.SetParent(f)

	f.奥辛尼基Osinniki = BOsinniki奥辛尼基
	f.奥辛尼基Osinniki.SetParent(f)

	f.辛涅列Sinnele = BSinnele辛涅列
	f.辛涅列Sinnele.SetParent(f)

	f.铁勒Tele = BTele铁勒
	f.铁勒Tele.SetParent(f)

	f.帖良古惕Telengit = BTelengit帖良古惕
	f.帖良古惕Telengit.SetParent(f)

}
