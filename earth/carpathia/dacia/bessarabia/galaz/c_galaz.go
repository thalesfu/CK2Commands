package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalazCounty interface {
	feud.County
	BBeresti贝雷什蒂() feud.Barony
	BFalciu弗尔丘() feud.Barony
	BGalaz加拉茨() feud.Barony
	BHusi胡希() feud.Barony
	BMurgeni穆尔杰尼() feud.Barony
	BOancea万恰() feud.Barony
}

type 加拉茨GalazCounty struct {
	feud.BaseCounty
	贝雷什蒂Beresti feud.Barony
	弗尔丘Falciu   feud.Barony
	加拉茨Galaz    feud.Barony
	胡希Husi      feud.Barony
	穆尔杰尼Murgeni feud.Barony
	万恰Oancea    feud.Barony
}

func (c *加拉茨GalazCounty) BBeresti贝雷什蒂() feud.Barony {
	return c.贝雷什蒂Beresti
}

func (c *加拉茨GalazCounty) BFalciu弗尔丘() feud.Barony {
	return c.弗尔丘Falciu
}

func (c *加拉茨GalazCounty) BGalaz加拉茨() feud.Barony {
	return c.加拉茨Galaz
}

func (c *加拉茨GalazCounty) BHusi胡希() feud.Barony {
	return c.胡希Husi
}

func (c *加拉茨GalazCounty) BMurgeni穆尔杰尼() feud.Barony {
	return c.穆尔杰尼Murgeni
}

func (c *加拉茨GalazCounty) BOancea万恰() feud.Barony {
	return c.万恰Oancea
}

var CGalaz加拉茨 GalazCounty = &加拉茨GalazCounty{}

func init() {
	f := CGalaz加拉茨.(*加拉茨GalazCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "511",
		Title:     "galaz",
		TitleName: "加拉茨",
		TitleCode: "c_galaz",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝雷什蒂Beresti = BBeresti贝雷什蒂
	f.贝雷什蒂Beresti.SetParent(f)

	f.弗尔丘Falciu = BFalciu弗尔丘
	f.弗尔丘Falciu.SetParent(f)

	f.加拉茨Galaz = BGalaz加拉茨
	f.加拉茨Galaz.SetParent(f)

	f.胡希Husi = BHusi胡希
	f.胡希Husi.SetParent(f)

	f.穆尔杰尼Murgeni = BMurgeni穆尔杰尼
	f.穆尔杰尼Murgeni.SetParent(f)

	f.万恰Oancea = BOancea万恰
	f.万恰Oancea.SetParent(f)

}
