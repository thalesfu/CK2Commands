package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KemiCounty interface {
	feud.County
	BInari伊纳里() feud.Barony
	BKemi凯米() feud.Barony
	BKemijarvi凯米耶尔维() feud.Barony
	BNeiden奈登() feud.Barony
	BSavukoski萨武科斯基() feud.Barony
	BTornio托尔尼奥() feud.Barony
	BUtsjoki乌茨约基() feud.Barony
}

type 凯米KemiCounty struct {
	feud.BaseCounty
	伊纳里Inari       feud.Barony
	凯米Kemi         feud.Barony
	凯米耶尔维Kemijarvi feud.Barony
	奈登Neiden       feud.Barony
	萨武科斯基Savukoski feud.Barony
	托尔尼奥Tornio     feud.Barony
	乌茨约基Utsjoki    feud.Barony
}

func (c *凯米KemiCounty) BInari伊纳里() feud.Barony {
	return c.伊纳里Inari
}

func (c *凯米KemiCounty) BKemi凯米() feud.Barony {
	return c.凯米Kemi
}

func (c *凯米KemiCounty) BKemijarvi凯米耶尔维() feud.Barony {
	return c.凯米耶尔维Kemijarvi
}

func (c *凯米KemiCounty) BNeiden奈登() feud.Barony {
	return c.奈登Neiden
}

func (c *凯米KemiCounty) BSavukoski萨武科斯基() feud.Barony {
	return c.萨武科斯基Savukoski
}

func (c *凯米KemiCounty) BTornio托尔尼奥() feud.Barony {
	return c.托尔尼奥Tornio
}

func (c *凯米KemiCounty) BUtsjoki乌茨约基() feud.Barony {
	return c.乌茨约基Utsjoki
}

var CKemi凯米 KemiCounty = &凯米KemiCounty{}

func init() {
	f := CKemi凯米.(*凯米KemiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "386",
		Title:     "kemi",
		TitleName: "凯米",
		TitleCode: "c_kemi",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊纳里Inari = BInari伊纳里
	f.伊纳里Inari.SetParent(f)

	f.凯米Kemi = BKemi凯米
	f.凯米Kemi.SetParent(f)

	f.凯米耶尔维Kemijarvi = BKemijarvi凯米耶尔维
	f.凯米耶尔维Kemijarvi.SetParent(f)

	f.奈登Neiden = BNeiden奈登
	f.奈登Neiden.SetParent(f)

	f.萨武科斯基Savukoski = BSavukoski萨武科斯基
	f.萨武科斯基Savukoski.SetParent(f)

	f.托尔尼奥Tornio = BTornio托尔尼奥
	f.托尔尼奥Tornio.SetParent(f)

	f.乌茨约基Utsjoki = BUtsjoki乌茨约基
	f.乌茨约基Utsjoki.SetParent(f)

}
