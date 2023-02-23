package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BalkhCounty interface {
	feud.County
	BAlkhanoum艾哈努姆() feud.Barony
	BBalkh缚喝() feud.Barony
	BDalverzintepe达利韦尔津捷佩() feud.Barony
	BKhulm忽懔() feud.Barony
	BSurkhkotal苏尔赫科塔尔() feud.Barony
	BTakhtisangin塔赫季桑金() feud.Barony
	BTiliatepe蒂拉丘地() feud.Barony
}

type 缚喝BalkhCounty struct {
	feud.BaseCounty
	艾哈努姆Alkhanoum        feud.Barony
	缚喝Balkh              feud.Barony
	达利韦尔津捷佩Dalverzintepe feud.Barony
	忽懔Khulm              feud.Barony
	苏尔赫科塔尔Surkhkotal     feud.Barony
	塔赫季桑金Takhtisangin    feud.Barony
	蒂拉丘地Tiliatepe        feud.Barony
}

func (c *缚喝BalkhCounty) BAlkhanoum艾哈努姆() feud.Barony {
	return c.艾哈努姆Alkhanoum
}

func (c *缚喝BalkhCounty) BBalkh缚喝() feud.Barony {
	return c.缚喝Balkh
}

func (c *缚喝BalkhCounty) BDalverzintepe达利韦尔津捷佩() feud.Barony {
	return c.达利韦尔津捷佩Dalverzintepe
}

func (c *缚喝BalkhCounty) BKhulm忽懔() feud.Barony {
	return c.忽懔Khulm
}

func (c *缚喝BalkhCounty) BSurkhkotal苏尔赫科塔尔() feud.Barony {
	return c.苏尔赫科塔尔Surkhkotal
}

func (c *缚喝BalkhCounty) BTakhtisangin塔赫季桑金() feud.Barony {
	return c.塔赫季桑金Takhtisangin
}

func (c *缚喝BalkhCounty) BTiliatepe蒂拉丘地() feud.Barony {
	return c.蒂拉丘地Tiliatepe
}

var CBalkh缚喝 BalkhCounty = &缚喝BalkhCounty{}

func init() {
	f := CBalkh缚喝.(*缚喝BalkhCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "904",
		Title:     "balkh",
		TitleName: "缚喝",
		TitleCode: "c_balkh",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾哈努姆Alkhanoum = BAlkhanoum艾哈努姆
	f.艾哈努姆Alkhanoum.SetParent(f)

	f.缚喝Balkh = BBalkh缚喝
	f.缚喝Balkh.SetParent(f)

	f.达利韦尔津捷佩Dalverzintepe = BDalverzintepe达利韦尔津捷佩
	f.达利韦尔津捷佩Dalverzintepe.SetParent(f)

	f.忽懔Khulm = BKhulm忽懔
	f.忽懔Khulm.SetParent(f)

	f.苏尔赫科塔尔Surkhkotal = BSurkhkotal苏尔赫科塔尔
	f.苏尔赫科塔尔Surkhkotal.SetParent(f)

	f.塔赫季桑金Takhtisangin = BTakhtisangin塔赫季桑金
	f.塔赫季桑金Takhtisangin.SetParent(f)

	f.蒂拉丘地Tiliatepe = BTiliatepe蒂拉丘地
	f.蒂拉丘地Tiliatepe.SetParent(f)

}
