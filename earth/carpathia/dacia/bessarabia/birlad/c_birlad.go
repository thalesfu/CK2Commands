package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BirladCounty interface {
	feud.County
	BAdjud阿德茹德() feud.Barony
	BBarlad伯尔拉德() feud.Barony
	BCraciuna克勒丘内什蒂() feud.Barony
	BFocsani福克沙尼() feud.Barony
	BMarasesti默勒谢什蒂() feud.Barony
	BTecuci泰库奇() feud.Barony
}

type 伯尔拉德BirladCounty struct {
	feud.BaseCounty
	阿德茹德Adjud      feud.Barony
	伯尔拉德Barlad     feud.Barony
	克勒丘内什蒂Craciuna feud.Barony
	福克沙尼Focsani    feud.Barony
	默勒谢什蒂Marasesti feud.Barony
	泰库奇Tecuci      feud.Barony
}

func (c *伯尔拉德BirladCounty) BAdjud阿德茹德() feud.Barony {
	return c.阿德茹德Adjud
}

func (c *伯尔拉德BirladCounty) BBarlad伯尔拉德() feud.Barony {
	return c.伯尔拉德Barlad
}

func (c *伯尔拉德BirladCounty) BCraciuna克勒丘内什蒂() feud.Barony {
	return c.克勒丘内什蒂Craciuna
}

func (c *伯尔拉德BirladCounty) BFocsani福克沙尼() feud.Barony {
	return c.福克沙尼Focsani
}

func (c *伯尔拉德BirladCounty) BMarasesti默勒谢什蒂() feud.Barony {
	return c.默勒谢什蒂Marasesti
}

func (c *伯尔拉德BirladCounty) BTecuci泰库奇() feud.Barony {
	return c.泰库奇Tecuci
}

var CBirlad伯尔拉德 BirladCounty = &伯尔拉德BirladCounty{}

func init() {
	f := CBirlad伯尔拉德.(*伯尔拉德BirladCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "513",
		Title:     "birlad",
		TitleName: "伯尔拉德",
		TitleCode: "c_birlad",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德茹德Adjud = BAdjud阿德茹德
	f.阿德茹德Adjud.SetParent(f)

	f.伯尔拉德Barlad = BBarlad伯尔拉德
	f.伯尔拉德Barlad.SetParent(f)

	f.克勒丘内什蒂Craciuna = BCraciuna克勒丘内什蒂
	f.克勒丘内什蒂Craciuna.SetParent(f)

	f.福克沙尼Focsani = BFocsani福克沙尼
	f.福克沙尼Focsani.SetParent(f)

	f.默勒谢什蒂Marasesti = BMarasesti默勒谢什蒂
	f.默勒谢什蒂Marasesti.SetParent(f)

	f.泰库奇Tecuci = BTecuci泰库奇
	f.泰库奇Tecuci.SetParent(f)

}
