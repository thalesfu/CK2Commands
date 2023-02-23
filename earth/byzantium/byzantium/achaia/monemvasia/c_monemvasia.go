package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MonemvasiaCounty interface {
	feud.County
	BArkadia阿卡迪亚() feud.Barony
	BElos艾洛斯() feud.Barony
	BGythio吉雄() feud.Barony
	BLacedaemonia拉凯戴孟尼亚() feud.Barony
	BMistra米斯特拉斯() feud.Barony
	BMonemvasia莫奈姆瓦夏() feud.Barony
	BNikli尼克拉斯() feud.Barony
	BSparta斯巴达() feud.Barony
}

type 莫奈姆瓦夏MonemvasiaCounty struct {
	feud.BaseCounty
	阿卡迪亚Arkadia        feud.Barony
	艾洛斯Elos            feud.Barony
	吉雄Gythio           feud.Barony
	拉凯戴孟尼亚Lacedaemonia feud.Barony
	米斯特拉斯Mistra        feud.Barony
	莫奈姆瓦夏Monemvasia    feud.Barony
	尼克拉斯Nikli          feud.Barony
	斯巴达Sparta          feud.Barony
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BArkadia阿卡迪亚() feud.Barony {
	return c.阿卡迪亚Arkadia
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BElos艾洛斯() feud.Barony {
	return c.艾洛斯Elos
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BGythio吉雄() feud.Barony {
	return c.吉雄Gythio
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BLacedaemonia拉凯戴孟尼亚() feud.Barony {
	return c.拉凯戴孟尼亚Lacedaemonia
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BMistra米斯特拉斯() feud.Barony {
	return c.米斯特拉斯Mistra
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BMonemvasia莫奈姆瓦夏() feud.Barony {
	return c.莫奈姆瓦夏Monemvasia
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BNikli尼克拉斯() feud.Barony {
	return c.尼克拉斯Nikli
}

func (c *莫奈姆瓦夏MonemvasiaCounty) BSparta斯巴达() feud.Barony {
	return c.斯巴达Sparta
}

var CMonemvasia莫奈姆瓦夏 MonemvasiaCounty = &莫奈姆瓦夏MonemvasiaCounty{}

func init() {
	f := CMonemvasia莫奈姆瓦夏.(*莫奈姆瓦夏MonemvasiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "478",
		Title:     "monemvasia",
		TitleName: "莫奈姆瓦夏",
		TitleCode: "c_monemvasia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卡迪亚Arkadia = BArkadia阿卡迪亚
	f.阿卡迪亚Arkadia.SetParent(f)

	f.艾洛斯Elos = BElos艾洛斯
	f.艾洛斯Elos.SetParent(f)

	f.吉雄Gythio = BGythio吉雄
	f.吉雄Gythio.SetParent(f)

	f.拉凯戴孟尼亚Lacedaemonia = BLacedaemonia拉凯戴孟尼亚
	f.拉凯戴孟尼亚Lacedaemonia.SetParent(f)

	f.米斯特拉斯Mistra = BMistra米斯特拉斯
	f.米斯特拉斯Mistra.SetParent(f)

	f.莫奈姆瓦夏Monemvasia = BMonemvasia莫奈姆瓦夏
	f.莫奈姆瓦夏Monemvasia.SetParent(f)

	f.尼克拉斯Nikli = BNikli尼克拉斯
	f.尼克拉斯Nikli.SetParent(f)

	f.斯巴达Sparta = BSparta斯巴达
	f.斯巴达Sparta.SetParent(f)

}
