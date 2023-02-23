package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlexandriaCounty interface {
	feud.County
	BAbukir阿布吉尔() feud.Barony
	BAlexandria亚历山大() feud.Barony
	BBurgelarab阿拉伯堡() feud.Barony
	BDamanhur陀曼那护罗() feud.Barony
	BElkanoun卡努恩() feud.Barony
	BHammam哈马姆() feud.Barony
	BMarabout穆拉比特() feud.Barony
	BNaucratis诺克拉提斯() feud.Barony
}

type 亚历山大AlexandriaCounty struct {
	feud.BaseCounty
	阿布吉尔Abukir     feud.Barony
	亚历山大Alexandria feud.Barony
	阿拉伯堡Burgelarab feud.Barony
	陀曼那护罗Damanhur  feud.Barony
	卡努恩Elkanoun    feud.Barony
	哈马姆Hammam      feud.Barony
	穆拉比特Marabout   feud.Barony
	诺克拉提斯Naucratis feud.Barony
}

func (c *亚历山大AlexandriaCounty) BAbukir阿布吉尔() feud.Barony {
	return c.阿布吉尔Abukir
}

func (c *亚历山大AlexandriaCounty) BAlexandria亚历山大() feud.Barony {
	return c.亚历山大Alexandria
}

func (c *亚历山大AlexandriaCounty) BBurgelarab阿拉伯堡() feud.Barony {
	return c.阿拉伯堡Burgelarab
}

func (c *亚历山大AlexandriaCounty) BDamanhur陀曼那护罗() feud.Barony {
	return c.陀曼那护罗Damanhur
}

func (c *亚历山大AlexandriaCounty) BElkanoun卡努恩() feud.Barony {
	return c.卡努恩Elkanoun
}

func (c *亚历山大AlexandriaCounty) BHammam哈马姆() feud.Barony {
	return c.哈马姆Hammam
}

func (c *亚历山大AlexandriaCounty) BMarabout穆拉比特() feud.Barony {
	return c.穆拉比特Marabout
}

func (c *亚历山大AlexandriaCounty) BNaucratis诺克拉提斯() feud.Barony {
	return c.诺克拉提斯Naucratis
}

var CAlexandria亚历山大 AlexandriaCounty = &亚历山大AlexandriaCounty{}

func init() {
	f := CAlexandria亚历山大.(*亚历山大AlexandriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "802",
		Title:     "alexandria",
		TitleName: "亚历山大",
		TitleCode: "c_alexandria",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布吉尔Abukir = BAbukir阿布吉尔
	f.阿布吉尔Abukir.SetParent(f)

	f.亚历山大Alexandria = BAlexandria亚历山大
	f.亚历山大Alexandria.SetParent(f)

	f.阿拉伯堡Burgelarab = BBurgelarab阿拉伯堡
	f.阿拉伯堡Burgelarab.SetParent(f)

	f.陀曼那护罗Damanhur = BDamanhur陀曼那护罗
	f.陀曼那护罗Damanhur.SetParent(f)

	f.卡努恩Elkanoun = BElkanoun卡努恩
	f.卡努恩Elkanoun.SetParent(f)

	f.哈马姆Hammam = BHammam哈马姆
	f.哈马姆Hammam.SetParent(f)

	f.穆拉比特Marabout = BMarabout穆拉比特
	f.穆拉比特Marabout.SetParent(f)

	f.诺克拉提斯Naucratis = BNaucratis诺克拉提斯
	f.诺克拉提斯Naucratis.SetParent(f)

}
