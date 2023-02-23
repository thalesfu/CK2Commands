package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HollandCounty interface {
	feud.County
	BAmsterdam阿姆斯特丹() feud.Barony
	BDordrecht多德雷赫特() feud.Barony
	BGouda豪达() feud.Barony
	BHaarlem哈勒姆() feud.Barony
	BLeiden雷登() feud.Barony
	BMuiden默伊登() feud.Barony
	BSgravenhage海牙() feud.Barony
	BVlaardingen弗拉丁根() feud.Barony
}

type 荷兰HollandCounty struct {
	feud.BaseCounty
	阿姆斯特丹Amsterdam  feud.Barony
	多德雷赫特Dordrecht  feud.Barony
	豪达Gouda         feud.Barony
	哈勒姆Haarlem      feud.Barony
	雷登Leiden        feud.Barony
	默伊登Muiden       feud.Barony
	海牙Sgravenhage   feud.Barony
	弗拉丁根Vlaardingen feud.Barony
}

func (c *荷兰HollandCounty) BAmsterdam阿姆斯特丹() feud.Barony {
	return c.阿姆斯特丹Amsterdam
}

func (c *荷兰HollandCounty) BDordrecht多德雷赫特() feud.Barony {
	return c.多德雷赫特Dordrecht
}

func (c *荷兰HollandCounty) BGouda豪达() feud.Barony {
	return c.豪达Gouda
}

func (c *荷兰HollandCounty) BHaarlem哈勒姆() feud.Barony {
	return c.哈勒姆Haarlem
}

func (c *荷兰HollandCounty) BLeiden雷登() feud.Barony {
	return c.雷登Leiden
}

func (c *荷兰HollandCounty) BMuiden默伊登() feud.Barony {
	return c.默伊登Muiden
}

func (c *荷兰HollandCounty) BSgravenhage海牙() feud.Barony {
	return c.海牙Sgravenhage
}

func (c *荷兰HollandCounty) BVlaardingen弗拉丁根() feud.Barony {
	return c.弗拉丁根Vlaardingen
}

var CHolland荷兰 HollandCounty = &荷兰HollandCounty{}

func init() {
	f := CHolland荷兰.(*荷兰HollandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "80",
		Title:     "holland",
		TitleName: "荷兰",
		TitleCode: "c_holland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆斯特丹Amsterdam = BAmsterdam阿姆斯特丹
	f.阿姆斯特丹Amsterdam.SetParent(f)

	f.多德雷赫特Dordrecht = BDordrecht多德雷赫特
	f.多德雷赫特Dordrecht.SetParent(f)

	f.豪达Gouda = BGouda豪达
	f.豪达Gouda.SetParent(f)

	f.哈勒姆Haarlem = BHaarlem哈勒姆
	f.哈勒姆Haarlem.SetParent(f)

	f.雷登Leiden = BLeiden雷登
	f.雷登Leiden.SetParent(f)

	f.默伊登Muiden = BMuiden默伊登
	f.默伊登Muiden.SetParent(f)

	f.海牙Sgravenhage = BSgravenhage海牙
	f.海牙Sgravenhage.SetParent(f)

	f.弗拉丁根Vlaardingen = BVlaardingen弗拉丁根
	f.弗拉丁根Vlaardingen.SetParent(f)

}
