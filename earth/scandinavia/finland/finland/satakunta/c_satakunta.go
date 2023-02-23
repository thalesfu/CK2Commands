package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SatakuntaCounty interface {
	feud.County
	BHahlo哈赫洛() feud.Barony
	BHiittenharju希腾哈里尤() feud.Barony
	BKankaanpaa坎康佩() feud.Barony
	BKiukainen基乌凯宁() feud.Barony
	BLiinmaa林恩马() feud.Barony
	BPori波里() feud.Barony
	BTelja泰列() feud.Barony
	BUlvila乌尔维拉() feud.Barony
}

type 萨塔昆塔SatakuntaCounty struct {
	feud.BaseCounty
	哈赫洛Hahlo          feud.Barony
	希腾哈里尤Hiittenharju feud.Barony
	坎康佩Kankaanpaa     feud.Barony
	基乌凯宁Kiukainen     feud.Barony
	林恩马Liinmaa        feud.Barony
	波里Pori            feud.Barony
	泰列Telja           feud.Barony
	乌尔维拉Ulvila        feud.Barony
}

func (c *萨塔昆塔SatakuntaCounty) BHahlo哈赫洛() feud.Barony {
	return c.哈赫洛Hahlo
}

func (c *萨塔昆塔SatakuntaCounty) BHiittenharju希腾哈里尤() feud.Barony {
	return c.希腾哈里尤Hiittenharju
}

func (c *萨塔昆塔SatakuntaCounty) BKankaanpaa坎康佩() feud.Barony {
	return c.坎康佩Kankaanpaa
}

func (c *萨塔昆塔SatakuntaCounty) BKiukainen基乌凯宁() feud.Barony {
	return c.基乌凯宁Kiukainen
}

func (c *萨塔昆塔SatakuntaCounty) BLiinmaa林恩马() feud.Barony {
	return c.林恩马Liinmaa
}

func (c *萨塔昆塔SatakuntaCounty) BPori波里() feud.Barony {
	return c.波里Pori
}

func (c *萨塔昆塔SatakuntaCounty) BTelja泰列() feud.Barony {
	return c.泰列Telja
}

func (c *萨塔昆塔SatakuntaCounty) BUlvila乌尔维拉() feud.Barony {
	return c.乌尔维拉Ulvila
}

var CSatakunta萨塔昆塔 SatakuntaCounty = &萨塔昆塔SatakuntaCounty{}

func init() {
	f := CSatakunta萨塔昆塔.(*萨塔昆塔SatakuntaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "384",
		Title:     "satakunta",
		TitleName: "萨塔昆塔",
		TitleCode: "c_satakunta",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈赫洛Hahlo = BHahlo哈赫洛
	f.哈赫洛Hahlo.SetParent(f)

	f.希腾哈里尤Hiittenharju = BHiittenharju希腾哈里尤
	f.希腾哈里尤Hiittenharju.SetParent(f)

	f.坎康佩Kankaanpaa = BKankaanpaa坎康佩
	f.坎康佩Kankaanpaa.SetParent(f)

	f.基乌凯宁Kiukainen = BKiukainen基乌凯宁
	f.基乌凯宁Kiukainen.SetParent(f)

	f.林恩马Liinmaa = BLiinmaa林恩马
	f.林恩马Liinmaa.SetParent(f)

	f.波里Pori = BPori波里
	f.波里Pori.SetParent(f)

	f.泰列Telja = BTelja泰列
	f.泰列Telja.SetParent(f)

	f.乌尔维拉Ulvila = BUlvila乌尔维拉
	f.乌尔维拉Ulvila.SetParent(f)

}
