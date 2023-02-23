package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZarmaCounty interface {
	feud.County
	BAbetam阿贝塔姆() feud.Barony
	BMangalam芒加拉姆() feud.Barony
	BObora奥博拉() feud.Barony
	BOuallam瓦拉姆() feud.Barony
	BTera特拉() feud.Barony
	BTora托拉() feud.Barony
	BZuru祖鲁() feud.Barony
}

type 扎尔马ZarmaCounty struct {
	feud.BaseCounty
	阿贝塔姆Abetam   feud.Barony
	芒加拉姆Mangalam feud.Barony
	奥博拉Obora     feud.Barony
	瓦拉姆Ouallam   feud.Barony
	特拉Tera       feud.Barony
	托拉Tora       feud.Barony
	祖鲁Zuru       feud.Barony
}

func (c *扎尔马ZarmaCounty) BAbetam阿贝塔姆() feud.Barony {
	return c.阿贝塔姆Abetam
}

func (c *扎尔马ZarmaCounty) BMangalam芒加拉姆() feud.Barony {
	return c.芒加拉姆Mangalam
}

func (c *扎尔马ZarmaCounty) BObora奥博拉() feud.Barony {
	return c.奥博拉Obora
}

func (c *扎尔马ZarmaCounty) BOuallam瓦拉姆() feud.Barony {
	return c.瓦拉姆Ouallam
}

func (c *扎尔马ZarmaCounty) BTera特拉() feud.Barony {
	return c.特拉Tera
}

func (c *扎尔马ZarmaCounty) BTora托拉() feud.Barony {
	return c.托拉Tora
}

func (c *扎尔马ZarmaCounty) BZuru祖鲁() feud.Barony {
	return c.祖鲁Zuru
}

var CZarma扎尔马 ZarmaCounty = &扎尔马ZarmaCounty{}

func init() {
	f := CZarma扎尔马.(*扎尔马ZarmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "927",
		Title:     "zarma",
		TitleName: "扎尔马",
		TitleCode: "c_zarma",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿贝塔姆Abetam = BAbetam阿贝塔姆
	f.阿贝塔姆Abetam.SetParent(f)

	f.芒加拉姆Mangalam = BMangalam芒加拉姆
	f.芒加拉姆Mangalam.SetParent(f)

	f.奥博拉Obora = BObora奥博拉
	f.奥博拉Obora.SetParent(f)

	f.瓦拉姆Ouallam = BOuallam瓦拉姆
	f.瓦拉姆Ouallam.SetParent(f)

	f.特拉Tera = BTera特拉
	f.特拉Tera.SetParent(f)

	f.托拉Tora = BTora托拉
	f.托拉Tora.SetParent(f)

	f.祖鲁Zuru = BZuru祖鲁
	f.祖鲁Zuru.SetParent(f)

}
