package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KashgarCounty interface {
	feud.County
	BAkto阿克陶() feud.Barony
	BBukou埠口() feud.Barony
	BKashgar可失合儿() feud.Barony
	BTsofa措法() feud.Barony
	BUlugqat乌恰() feud.Barony
	BWeitou尉头() feud.Barony
	BXiuxun休循() feud.Barony
}

type 可失合儿KashgarCounty struct {
	feud.BaseCounty
	阿克陶Akto     feud.Barony
	埠口Bukou     feud.Barony
	可失合儿Kashgar feud.Barony
	措法Tsofa     feud.Barony
	乌恰Ulugqat   feud.Barony
	尉头Weitou    feud.Barony
	休循Xiuxun    feud.Barony
}

func (c *可失合儿KashgarCounty) BAkto阿克陶() feud.Barony {
	return c.阿克陶Akto
}

func (c *可失合儿KashgarCounty) BBukou埠口() feud.Barony {
	return c.埠口Bukou
}

func (c *可失合儿KashgarCounty) BKashgar可失合儿() feud.Barony {
	return c.可失合儿Kashgar
}

func (c *可失合儿KashgarCounty) BTsofa措法() feud.Barony {
	return c.措法Tsofa
}

func (c *可失合儿KashgarCounty) BUlugqat乌恰() feud.Barony {
	return c.乌恰Ulugqat
}

func (c *可失合儿KashgarCounty) BWeitou尉头() feud.Barony {
	return c.尉头Weitou
}

func (c *可失合儿KashgarCounty) BXiuxun休循() feud.Barony {
	return c.休循Xiuxun
}

var CKashgar可失合儿 KashgarCounty = &可失合儿KashgarCounty{}

func init() {
	f := CKashgar可失合儿.(*可失合儿KashgarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1439",
		Title:     "kashgar",
		TitleName: "可失合儿",
		TitleCode: "c_kashgar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克陶Akto = BAkto阿克陶
	f.阿克陶Akto.SetParent(f)

	f.埠口Bukou = BBukou埠口
	f.埠口Bukou.SetParent(f)

	f.可失合儿Kashgar = BKashgar可失合儿
	f.可失合儿Kashgar.SetParent(f)

	f.措法Tsofa = BTsofa措法
	f.措法Tsofa.SetParent(f)

	f.乌恰Ulugqat = BUlugqat乌恰
	f.乌恰Ulugqat.SetParent(f)

	f.尉头Weitou = BWeitou尉头
	f.尉头Weitou.SetParent(f)

	f.休循Xiuxun = BXiuxun休循
	f.休循Xiuxun.SetParent(f)

}
