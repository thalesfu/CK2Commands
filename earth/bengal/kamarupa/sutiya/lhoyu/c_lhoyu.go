package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhoyuCounty interface {
	feud.County
	BAnini安尼尼() feud.Barony
	BDambuk达木布克() feud.Barony
	BDiyun蒂芸() feud.Barony
	BHawai哈威() feud.Barony
	BLhoyu珞隅() feud.Barony
	BNamsai南赛() feud.Barony
	BYingkiong因基扬() feud.Barony
}

type 珞隅LhoyuCounty struct {
	feud.BaseCounty
	安尼尼Anini     feud.Barony
	达木布克Dambuk   feud.Barony
	蒂芸Diyun      feud.Barony
	哈威Hawai      feud.Barony
	珞隅Lhoyu      feud.Barony
	南赛Namsai     feud.Barony
	因基扬Yingkiong feud.Barony
}

func (c *珞隅LhoyuCounty) BAnini安尼尼() feud.Barony {
	return c.安尼尼Anini
}

func (c *珞隅LhoyuCounty) BDambuk达木布克() feud.Barony {
	return c.达木布克Dambuk
}

func (c *珞隅LhoyuCounty) BDiyun蒂芸() feud.Barony {
	return c.蒂芸Diyun
}

func (c *珞隅LhoyuCounty) BHawai哈威() feud.Barony {
	return c.哈威Hawai
}

func (c *珞隅LhoyuCounty) BLhoyu珞隅() feud.Barony {
	return c.珞隅Lhoyu
}

func (c *珞隅LhoyuCounty) BNamsai南赛() feud.Barony {
	return c.南赛Namsai
}

func (c *珞隅LhoyuCounty) BYingkiong因基扬() feud.Barony {
	return c.因基扬Yingkiong
}

var CLhoyu珞隅 LhoyuCounty = &珞隅LhoyuCounty{}

func init() {
	f := CLhoyu珞隅.(*珞隅LhoyuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1484",
		Title:     "lhoyu",
		TitleName: "珞隅",
		TitleCode: "c_lhoyu",
		Baronies:  map[string]feud.Barony{},
	}

	f.安尼尼Anini = BAnini安尼尼
	f.安尼尼Anini.SetParent(f)

	f.达木布克Dambuk = BDambuk达木布克
	f.达木布克Dambuk.SetParent(f)

	f.蒂芸Diyun = BDiyun蒂芸
	f.蒂芸Diyun.SetParent(f)

	f.哈威Hawai = BHawai哈威
	f.哈威Hawai.SetParent(f)

	f.珞隅Lhoyu = BLhoyu珞隅
	f.珞隅Lhoyu.SetParent(f)

	f.南赛Namsai = BNamsai南赛
	f.南赛Namsai.SetParent(f)

	f.因基扬Yingkiong = BYingkiong因基扬
	f.因基扬Yingkiong.SetParent(f)

}
