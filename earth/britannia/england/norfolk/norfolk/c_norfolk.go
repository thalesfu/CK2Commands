package norfolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorfolkCounty interface {
	feud.County
	BBuckenham巴肯纳姆() feud.Barony
	BChatteris查特里斯() feud.Barony
	BElmham埃尔门() feud.Barony
	BLynn林恩() feud.Barony
	BNorwich诺里奇() feud.Barony
	BThetford塞特福特() feud.Barony
}

type 诺福克NorfolkCounty struct {
	feud.BaseCounty
	巴肯纳姆Buckenham feud.Barony
	查特里斯Chatteris feud.Barony
	埃尔门Elmham     feud.Barony
	林恩Lynn        feud.Barony
	诺里奇Norwich    feud.Barony
	塞特福特Thetford  feud.Barony
}

func (c *诺福克NorfolkCounty) BBuckenham巴肯纳姆() feud.Barony {
	return c.巴肯纳姆Buckenham
}

func (c *诺福克NorfolkCounty) BChatteris查特里斯() feud.Barony {
	return c.查特里斯Chatteris
}

func (c *诺福克NorfolkCounty) BElmham埃尔门() feud.Barony {
	return c.埃尔门Elmham
}

func (c *诺福克NorfolkCounty) BLynn林恩() feud.Barony {
	return c.林恩Lynn
}

func (c *诺福克NorfolkCounty) BNorwich诺里奇() feud.Barony {
	return c.诺里奇Norwich
}

func (c *诺福克NorfolkCounty) BThetford塞特福特() feud.Barony {
	return c.塞特福特Thetford
}

var CNorfolk诺福克 NorfolkCounty = &诺福克NorfolkCounty{}

func init() {
	f := CNorfolk诺福克.(*诺福克NorfolkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "70",
		Title:     "norfolk",
		TitleName: "诺福克",
		TitleCode: "c_norfolk",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴肯纳姆Buckenham = BBuckenham巴肯纳姆
	f.巴肯纳姆Buckenham.SetParent(f)

	f.查特里斯Chatteris = BChatteris查特里斯
	f.查特里斯Chatteris.SetParent(f)

	f.埃尔门Elmham = BElmham埃尔门
	f.埃尔门Elmham.SetParent(f)

	f.林恩Lynn = BLynn林恩
	f.林恩Lynn.SetParent(f)

	f.诺里奇Norwich = BNorwich诺里奇
	f.诺里奇Norwich.SetParent(f)

	f.塞特福特Thetford = BThetford塞特福特
	f.塞特福特Thetford.SetParent(f)

}
