package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThalnerCounty interface {
	feud.County
	BAmalner阿默尔内尔() feud.Barony
	BBalasane婆罗娑内() feud.Barony
	BBhamer婆摩尔() feud.Barony
	BJhodga殊迦() feud.Barony
	BLaling罗尼恩() feud.Barony
	BSindkheda申陀契荼() feud.Barony
	BThalner闼罗内() feud.Barony
}

type 闼罗内ThalnerCounty struct {
	feud.BaseCounty
	阿默尔内尔Amalner  feud.Barony
	婆罗娑内Balasane  feud.Barony
	婆摩尔Bhamer     feud.Barony
	殊迦Jhodga      feud.Barony
	罗尼恩Laling     feud.Barony
	申陀契荼Sindkheda feud.Barony
	闼罗内Thalner    feud.Barony
}

func (c *闼罗内ThalnerCounty) BAmalner阿默尔内尔() feud.Barony {
	return c.阿默尔内尔Amalner
}

func (c *闼罗内ThalnerCounty) BBalasane婆罗娑内() feud.Barony {
	return c.婆罗娑内Balasane
}

func (c *闼罗内ThalnerCounty) BBhamer婆摩尔() feud.Barony {
	return c.婆摩尔Bhamer
}

func (c *闼罗内ThalnerCounty) BJhodga殊迦() feud.Barony {
	return c.殊迦Jhodga
}

func (c *闼罗内ThalnerCounty) BLaling罗尼恩() feud.Barony {
	return c.罗尼恩Laling
}

func (c *闼罗内ThalnerCounty) BSindkheda申陀契荼() feud.Barony {
	return c.申陀契荼Sindkheda
}

func (c *闼罗内ThalnerCounty) BThalner闼罗内() feud.Barony {
	return c.闼罗内Thalner
}

var CThalner闼罗内 ThalnerCounty = &闼罗内ThalnerCounty{}

func init() {
	f := CThalner闼罗内.(*闼罗内ThalnerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1262",
		Title:     "thalner",
		TitleName: "闼罗内",
		TitleCode: "c_thalner",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿默尔内尔Amalner = BAmalner阿默尔内尔
	f.阿默尔内尔Amalner.SetParent(f)

	f.婆罗娑内Balasane = BBalasane婆罗娑内
	f.婆罗娑内Balasane.SetParent(f)

	f.婆摩尔Bhamer = BBhamer婆摩尔
	f.婆摩尔Bhamer.SetParent(f)

	f.殊迦Jhodga = BJhodga殊迦
	f.殊迦Jhodga.SetParent(f)

	f.罗尼恩Laling = BLaling罗尼恩
	f.罗尼恩Laling.SetParent(f)

	f.申陀契荼Sindkheda = BSindkheda申陀契荼
	f.申陀契荼Sindkheda.SetParent(f)

	f.闼罗内Thalner = BThalner闼罗内
	f.闼罗内Thalner.SetParent(f)

}
