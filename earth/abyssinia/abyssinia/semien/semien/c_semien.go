package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SemienCounty interface {
	feud.County
	BAmbaras阿姆巴拉斯() feud.Barony
	BBelesa比勒萨() feud.Barony
	BCiarveta奇阿韦塔() feud.Barony
	BDabat达巴特() feud.Barony
	BDebark德巴尔格() feud.Barony
}

type 塞米恩SemienCounty struct {
	feud.BaseCounty
	阿姆巴拉斯Ambaras feud.Barony
	比勒萨Belesa    feud.Barony
	奇阿韦塔Ciarveta feud.Barony
	达巴特Dabat     feud.Barony
	德巴尔格Debark   feud.Barony
}

func (c *塞米恩SemienCounty) BAmbaras阿姆巴拉斯() feud.Barony {
	return c.阿姆巴拉斯Ambaras
}

func (c *塞米恩SemienCounty) BBelesa比勒萨() feud.Barony {
	return c.比勒萨Belesa
}

func (c *塞米恩SemienCounty) BCiarveta奇阿韦塔() feud.Barony {
	return c.奇阿韦塔Ciarveta
}

func (c *塞米恩SemienCounty) BDabat达巴特() feud.Barony {
	return c.达巴特Dabat
}

func (c *塞米恩SemienCounty) BDebark德巴尔格() feud.Barony {
	return c.德巴尔格Debark
}

var CSemien塞米恩 SemienCounty = &塞米恩SemienCounty{}

func init() {
	f := CSemien塞米恩.(*塞米恩SemienCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1380",
		Title:     "semien",
		TitleName: "塞米恩",
		TitleCode: "c_semien",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆巴拉斯Ambaras = BAmbaras阿姆巴拉斯
	f.阿姆巴拉斯Ambaras.SetParent(f)

	f.比勒萨Belesa = BBelesa比勒萨
	f.比勒萨Belesa.SetParent(f)

	f.奇阿韦塔Ciarveta = BCiarveta奇阿韦塔
	f.奇阿韦塔Ciarveta.SetParent(f)

	f.达巴特Dabat = BDabat达巴特
	f.达巴特Dabat.SetParent(f)

	f.德巴尔格Debark = BDebark德巴尔格
	f.德巴尔格Debark.SetParent(f)

}
