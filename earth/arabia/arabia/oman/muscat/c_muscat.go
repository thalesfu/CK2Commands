package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MuscatCounty interface {
	feud.County
	BAdam亚当() feud.Barony
	BIbra伊卜拉() feud.Barony
	BLizq利兹奇() feud.Barony
	BMuscat马斯喀特() feud.Barony
	BSabt塞卜特() feud.Barony
	BSamail塞马伊勒() feud.Barony
	BShiya什亚() feud.Barony
	BSur苏尔() feud.Barony
}

type 马斯喀特MuscatCounty struct {
	feud.BaseCounty
	亚当Adam     feud.Barony
	伊卜拉Ibra    feud.Barony
	利兹奇Lizq    feud.Barony
	马斯喀特Muscat feud.Barony
	塞卜特Sabt    feud.Barony
	塞马伊勒Samail feud.Barony
	什亚Shiya    feud.Barony
	苏尔Sur      feud.Barony
}

func (c *马斯喀特MuscatCounty) BAdam亚当() feud.Barony {
	return c.亚当Adam
}

func (c *马斯喀特MuscatCounty) BIbra伊卜拉() feud.Barony {
	return c.伊卜拉Ibra
}

func (c *马斯喀特MuscatCounty) BLizq利兹奇() feud.Barony {
	return c.利兹奇Lizq
}

func (c *马斯喀特MuscatCounty) BMuscat马斯喀特() feud.Barony {
	return c.马斯喀特Muscat
}

func (c *马斯喀特MuscatCounty) BSabt塞卜特() feud.Barony {
	return c.塞卜特Sabt
}

func (c *马斯喀特MuscatCounty) BSamail塞马伊勒() feud.Barony {
	return c.塞马伊勒Samail
}

func (c *马斯喀特MuscatCounty) BShiya什亚() feud.Barony {
	return c.什亚Shiya
}

func (c *马斯喀特MuscatCounty) BSur苏尔() feud.Barony {
	return c.苏尔Sur
}

var CMuscat马斯喀特 MuscatCounty = &马斯喀特MuscatCounty{}

func init() {
	f := CMuscat马斯喀特.(*马斯喀特MuscatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "868",
		Title:     "muscat",
		TitleName: "马斯喀特",
		TitleCode: "c_muscat",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚当Adam = BAdam亚当
	f.亚当Adam.SetParent(f)

	f.伊卜拉Ibra = BIbra伊卜拉
	f.伊卜拉Ibra.SetParent(f)

	f.利兹奇Lizq = BLizq利兹奇
	f.利兹奇Lizq.SetParent(f)

	f.马斯喀特Muscat = BMuscat马斯喀特
	f.马斯喀特Muscat.SetParent(f)

	f.塞卜特Sabt = BSabt塞卜特
	f.塞卜特Sabt.SetParent(f)

	f.塞马伊勒Samail = BSamail塞马伊勒
	f.塞马伊勒Samail.SetParent(f)

	f.什亚Shiya = BShiya什亚
	f.什亚Shiya.SetParent(f)

	f.苏尔Sur = BSur苏尔
	f.苏尔Sur.SetParent(f)

}
