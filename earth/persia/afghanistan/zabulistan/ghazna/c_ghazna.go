package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GhaznaCounty interface {
	feud.County
	BGardez加德兹() feud.Barony
	BGhazna伽色尼() feud.Barony
	BKatawaz卡塔瓦兹() feud.Barony
	BKhost阔悉多() feud.Barony
	BLoman鲈鳗() feud.Barony
	BSangar桑加尔() feud.Barony
	BSharana娑兰那() feud.Barony
}

type 伽色尼GhaznaCounty struct {
	feud.BaseCounty
	加德兹Gardez   feud.Barony
	伽色尼Ghazna   feud.Barony
	卡塔瓦兹Katawaz feud.Barony
	阔悉多Khost    feud.Barony
	鲈鳗Loman     feud.Barony
	桑加尔Sangar   feud.Barony
	娑兰那Sharana  feud.Barony
}

func (c *伽色尼GhaznaCounty) BGardez加德兹() feud.Barony {
	return c.加德兹Gardez
}

func (c *伽色尼GhaznaCounty) BGhazna伽色尼() feud.Barony {
	return c.伽色尼Ghazna
}

func (c *伽色尼GhaznaCounty) BKatawaz卡塔瓦兹() feud.Barony {
	return c.卡塔瓦兹Katawaz
}

func (c *伽色尼GhaznaCounty) BKhost阔悉多() feud.Barony {
	return c.阔悉多Khost
}

func (c *伽色尼GhaznaCounty) BLoman鲈鳗() feud.Barony {
	return c.鲈鳗Loman
}

func (c *伽色尼GhaznaCounty) BSangar桑加尔() feud.Barony {
	return c.桑加尔Sangar
}

func (c *伽色尼GhaznaCounty) BSharana娑兰那() feud.Barony {
	return c.娑兰那Sharana
}

var CGhazna伽色尼 GhaznaCounty = &伽色尼GhaznaCounty{}

func init() {
	f := CGhazna伽色尼.(*伽色尼GhaznaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1182",
		Title:     "ghazna",
		TitleName: "伽色尼",
		TitleCode: "c_ghazna",
		Baronies:  map[string]feud.Barony{},
	}

	f.加德兹Gardez = BGardez加德兹
	f.加德兹Gardez.SetParent(f)

	f.伽色尼Ghazna = BGhazna伽色尼
	f.伽色尼Ghazna.SetParent(f)

	f.卡塔瓦兹Katawaz = BKatawaz卡塔瓦兹
	f.卡塔瓦兹Katawaz.SetParent(f)

	f.阔悉多Khost = BKhost阔悉多
	f.阔悉多Khost.SetParent(f)

	f.鲈鳗Loman = BLoman鲈鳗
	f.鲈鳗Loman.SetParent(f)

	f.桑加尔Sangar = BSangar桑加尔
	f.桑加尔Sangar.SetParent(f)

	f.娑兰那Sharana = BSharana娑兰那
	f.娑兰那Sharana.SetParent(f)

}
