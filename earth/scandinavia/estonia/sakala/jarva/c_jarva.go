package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JarvaCounty interface {
	feud.County
	BAlempois阿伦波伊斯() feud.Barony
	BJogentagana约根塔加纳() feud.Barony
	BMohu莫胡() feud.Barony
	BNurmekund努尔梅昆德() feud.Barony
	BSackal萨卡尔() feud.Barony
	BVaiga瓦伊加() feud.Barony
	BViljandi维尔扬迪() feud.Barony
}

type 耶尔瓦JarvaCounty struct {
	feud.BaseCounty
	阿伦波伊斯Alempois    feud.Barony
	约根塔加纳Jogentagana feud.Barony
	莫胡Mohu           feud.Barony
	努尔梅昆德Nurmekund   feud.Barony
	萨卡尔Sackal        feud.Barony
	瓦伊加Vaiga         feud.Barony
	维尔扬迪Viljandi     feud.Barony
}

func (c *耶尔瓦JarvaCounty) BAlempois阿伦波伊斯() feud.Barony {
	return c.阿伦波伊斯Alempois
}

func (c *耶尔瓦JarvaCounty) BJogentagana约根塔加纳() feud.Barony {
	return c.约根塔加纳Jogentagana
}

func (c *耶尔瓦JarvaCounty) BMohu莫胡() feud.Barony {
	return c.莫胡Mohu
}

func (c *耶尔瓦JarvaCounty) BNurmekund努尔梅昆德() feud.Barony {
	return c.努尔梅昆德Nurmekund
}

func (c *耶尔瓦JarvaCounty) BSackal萨卡尔() feud.Barony {
	return c.萨卡尔Sackal
}

func (c *耶尔瓦JarvaCounty) BVaiga瓦伊加() feud.Barony {
	return c.瓦伊加Vaiga
}

func (c *耶尔瓦JarvaCounty) BViljandi维尔扬迪() feud.Barony {
	return c.维尔扬迪Viljandi
}

var CJarva耶尔瓦 JarvaCounty = &耶尔瓦JarvaCounty{}

func init() {
	f := CJarva耶尔瓦.(*耶尔瓦JarvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1598",
		Title:     "jarva",
		TitleName: "耶尔瓦",
		TitleCode: "c_jarva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伦波伊斯Alempois = BAlempois阿伦波伊斯
	f.阿伦波伊斯Alempois.SetParent(f)

	f.约根塔加纳Jogentagana = BJogentagana约根塔加纳
	f.约根塔加纳Jogentagana.SetParent(f)

	f.莫胡Mohu = BMohu莫胡
	f.莫胡Mohu.SetParent(f)

	f.努尔梅昆德Nurmekund = BNurmekund努尔梅昆德
	f.努尔梅昆德Nurmekund.SetParent(f)

	f.萨卡尔Sackal = BSackal萨卡尔
	f.萨卡尔Sackal.SetParent(f)

	f.瓦伊加Vaiga = BVaiga瓦伊加
	f.瓦伊加Vaiga.SetParent(f)

	f.维尔扬迪Viljandi = BViljandi维尔扬迪
	f.维尔扬迪Viljandi.SetParent(f)

}
