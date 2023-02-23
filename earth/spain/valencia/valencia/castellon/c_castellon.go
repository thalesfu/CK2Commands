package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CastellonCounty interface {
	feud.County
	BAlcalaten阿拉卡拉滕() feud.Barony
	BAlpuente阿尔普恩特() feud.Barony
	BBurriana布里亚纳() feud.Barony
	BCastellon卡斯特略() feud.Barony
	BMorella莫雷利亚() feud.Barony
	BNules努莱斯() feud.Barony
	BVilarreal比利亚雷亚尔() feud.Barony
	BVinaros比纳罗斯() feud.Barony
}

type 卡斯特利翁CastellonCounty struct {
	feud.BaseCounty
	阿拉卡拉滕Alcalaten  feud.Barony
	阿尔普恩特Alpuente   feud.Barony
	布里亚纳Burriana    feud.Barony
	卡斯特略Castellon   feud.Barony
	莫雷利亚Morella     feud.Barony
	努莱斯Nules        feud.Barony
	比利亚雷亚尔Vilarreal feud.Barony
	比纳罗斯Vinaros     feud.Barony
}

func (c *卡斯特利翁CastellonCounty) BAlcalaten阿拉卡拉滕() feud.Barony {
	return c.阿拉卡拉滕Alcalaten
}

func (c *卡斯特利翁CastellonCounty) BAlpuente阿尔普恩特() feud.Barony {
	return c.阿尔普恩特Alpuente
}

func (c *卡斯特利翁CastellonCounty) BBurriana布里亚纳() feud.Barony {
	return c.布里亚纳Burriana
}

func (c *卡斯特利翁CastellonCounty) BCastellon卡斯特略() feud.Barony {
	return c.卡斯特略Castellon
}

func (c *卡斯特利翁CastellonCounty) BMorella莫雷利亚() feud.Barony {
	return c.莫雷利亚Morella
}

func (c *卡斯特利翁CastellonCounty) BNules努莱斯() feud.Barony {
	return c.努莱斯Nules
}

func (c *卡斯特利翁CastellonCounty) BVilarreal比利亚雷亚尔() feud.Barony {
	return c.比利亚雷亚尔Vilarreal
}

func (c *卡斯特利翁CastellonCounty) BVinaros比纳罗斯() feud.Barony {
	return c.比纳罗斯Vinaros
}

var CCastellon卡斯特利翁 CastellonCounty = &卡斯特利翁CastellonCounty{}

func init() {
	f := CCastellon卡斯特利翁.(*卡斯特利翁CastellonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "172",
		Title:     "castellon",
		TitleName: "卡斯特利翁",
		TitleCode: "c_castellon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉卡拉滕Alcalaten = BAlcalaten阿拉卡拉滕
	f.阿拉卡拉滕Alcalaten.SetParent(f)

	f.阿尔普恩特Alpuente = BAlpuente阿尔普恩特
	f.阿尔普恩特Alpuente.SetParent(f)

	f.布里亚纳Burriana = BBurriana布里亚纳
	f.布里亚纳Burriana.SetParent(f)

	f.卡斯特略Castellon = BCastellon卡斯特略
	f.卡斯特略Castellon.SetParent(f)

	f.莫雷利亚Morella = BMorella莫雷利亚
	f.莫雷利亚Morella.SetParent(f)

	f.努莱斯Nules = BNules努莱斯
	f.努莱斯Nules.SetParent(f)

	f.比利亚雷亚尔Vilarreal = BVilarreal比利亚雷亚尔
	f.比利亚雷亚尔Vilarreal.SetParent(f)

	f.比纳罗斯Vinaros = BVinaros比纳罗斯
	f.比纳罗斯Vinaros.SetParent(f)

}
