package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ProvenceCounty interface {
	feud.County
	BAix艾克斯() feud.Barony
	BArles阿尔勒() feud.Barony
	BCastellane卡斯泰朗() feud.Barony
	BFrejus弗雷瑞斯() feud.Barony
	BGrasse格拉斯() feud.Barony
	BGrimaud格里莫() feud.Barony
	BMarseille马赛() feud.Barony
	BTarascon塔拉斯孔() feud.Barony
}

type 普罗旺斯ProvenceCounty struct {
	feud.BaseCounty
	艾克斯Aix         feud.Barony
	阿尔勒Arles       feud.Barony
	卡斯泰朗Castellane feud.Barony
	弗雷瑞斯Frejus     feud.Barony
	格拉斯Grasse      feud.Barony
	格里莫Grimaud     feud.Barony
	马赛Marseille    feud.Barony
	塔拉斯孔Tarascon   feud.Barony
}

func (c *普罗旺斯ProvenceCounty) BAix艾克斯() feud.Barony {
	return c.艾克斯Aix
}

func (c *普罗旺斯ProvenceCounty) BArles阿尔勒() feud.Barony {
	return c.阿尔勒Arles
}

func (c *普罗旺斯ProvenceCounty) BCastellane卡斯泰朗() feud.Barony {
	return c.卡斯泰朗Castellane
}

func (c *普罗旺斯ProvenceCounty) BFrejus弗雷瑞斯() feud.Barony {
	return c.弗雷瑞斯Frejus
}

func (c *普罗旺斯ProvenceCounty) BGrasse格拉斯() feud.Barony {
	return c.格拉斯Grasse
}

func (c *普罗旺斯ProvenceCounty) BGrimaud格里莫() feud.Barony {
	return c.格里莫Grimaud
}

func (c *普罗旺斯ProvenceCounty) BMarseille马赛() feud.Barony {
	return c.马赛Marseille
}

func (c *普罗旺斯ProvenceCounty) BTarascon塔拉斯孔() feud.Barony {
	return c.塔拉斯孔Tarascon
}

var CProvence普罗旺斯 ProvenceCounty = &普罗旺斯ProvenceCounty{}

func init() {
	f := CProvence普罗旺斯.(*普罗旺斯ProvenceCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "221",
		Title:     "provence",
		TitleName: "普罗旺斯",
		TitleCode: "c_provence",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾克斯Aix = BAix艾克斯
	f.艾克斯Aix.SetParent(f)

	f.阿尔勒Arles = BArles阿尔勒
	f.阿尔勒Arles.SetParent(f)

	f.卡斯泰朗Castellane = BCastellane卡斯泰朗
	f.卡斯泰朗Castellane.SetParent(f)

	f.弗雷瑞斯Frejus = BFrejus弗雷瑞斯
	f.弗雷瑞斯Frejus.SetParent(f)

	f.格拉斯Grasse = BGrasse格拉斯
	f.格拉斯Grasse.SetParent(f)

	f.格里莫Grimaud = BGrimaud格里莫
	f.格里莫Grimaud.SetParent(f)

	f.马赛Marseille = BMarseille马赛
	f.马赛Marseille.SetParent(f)

	f.塔拉斯孔Tarascon = BTarascon塔拉斯孔
	f.塔拉斯孔Tarascon.SetParent(f)

}
