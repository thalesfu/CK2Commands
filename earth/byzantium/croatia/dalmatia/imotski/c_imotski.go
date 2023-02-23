package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ImotskiCounty interface {
	feud.County
	BBlagaj布拉加伊() feud.Barony
	BDuvno杜夫诺() feud.Barony
	BGrude格鲁德() feud.Barony
	BImotski伊莫茨基() feud.Barony
	BLivno利夫诺() feud.Barony
	BPosusje波苏谢() feud.Barony
	BTopana托帕纳() feud.Barony
}

type 伊莫茨基ImotskiCounty struct {
	feud.BaseCounty
	布拉加伊Blagaj  feud.Barony
	杜夫诺Duvno    feud.Barony
	格鲁德Grude    feud.Barony
	伊莫茨基Imotski feud.Barony
	利夫诺Livno    feud.Barony
	波苏谢Posusje  feud.Barony
	托帕纳Topana   feud.Barony
}

func (c *伊莫茨基ImotskiCounty) BBlagaj布拉加伊() feud.Barony {
	return c.布拉加伊Blagaj
}

func (c *伊莫茨基ImotskiCounty) BDuvno杜夫诺() feud.Barony {
	return c.杜夫诺Duvno
}

func (c *伊莫茨基ImotskiCounty) BGrude格鲁德() feud.Barony {
	return c.格鲁德Grude
}

func (c *伊莫茨基ImotskiCounty) BImotski伊莫茨基() feud.Barony {
	return c.伊莫茨基Imotski
}

func (c *伊莫茨基ImotskiCounty) BLivno利夫诺() feud.Barony {
	return c.利夫诺Livno
}

func (c *伊莫茨基ImotskiCounty) BPosusje波苏谢() feud.Barony {
	return c.波苏谢Posusje
}

func (c *伊莫茨基ImotskiCounty) BTopana托帕纳() feud.Barony {
	return c.托帕纳Topana
}

var CImotski伊莫茨基 ImotskiCounty = &伊莫茨基ImotskiCounty{}

func init() {
	f := CImotski伊莫茨基.(*伊莫茨基ImotskiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1967",
		Title:     "imotski",
		TitleName: "伊莫茨基",
		TitleCode: "c_imotski",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉加伊Blagaj = BBlagaj布拉加伊
	f.布拉加伊Blagaj.SetParent(f)

	f.杜夫诺Duvno = BDuvno杜夫诺
	f.杜夫诺Duvno.SetParent(f)

	f.格鲁德Grude = BGrude格鲁德
	f.格鲁德Grude.SetParent(f)

	f.伊莫茨基Imotski = BImotski伊莫茨基
	f.伊莫茨基Imotski.SetParent(f)

	f.利夫诺Livno = BLivno利夫诺
	f.利夫诺Livno.SetParent(f)

	f.波苏谢Posusje = BPosusje波苏谢
	f.波苏谢Posusje.SetParent(f)

	f.托帕纳Topana = BTopana托帕纳
	f.托帕纳Topana.SetParent(f)

}
