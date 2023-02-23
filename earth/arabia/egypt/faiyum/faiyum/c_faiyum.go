package faiyum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FaiyumCounty interface {
	feud.County
	BAhnas阿赫纳斯() feud.Barony
	BBahnasa巴赫纳萨() feud.Barony
	BFaiyum法尤姆() feud.Barony
	BHeracleus赫拉克琉斯() feud.Barony
}

type 法尤姆FaiyumCounty struct {
	feud.BaseCounty
	阿赫纳斯Ahnas      feud.Barony
	巴赫纳萨Bahnasa    feud.Barony
	法尤姆Faiyum      feud.Barony
	赫拉克琉斯Heracleus feud.Barony
}

func (c *法尤姆FaiyumCounty) BAhnas阿赫纳斯() feud.Barony {
	return c.阿赫纳斯Ahnas
}

func (c *法尤姆FaiyumCounty) BBahnasa巴赫纳萨() feud.Barony {
	return c.巴赫纳萨Bahnasa
}

func (c *法尤姆FaiyumCounty) BFaiyum法尤姆() feud.Barony {
	return c.法尤姆Faiyum
}

func (c *法尤姆FaiyumCounty) BHeracleus赫拉克琉斯() feud.Barony {
	return c.赫拉克琉斯Heracleus
}

var CFaiyum法尤姆 FaiyumCounty = &法尤姆FaiyumCounty{}

func init() {
	f := CFaiyum法尤姆.(*法尤姆FaiyumCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2001",
		Title:     "faiyum",
		TitleName: "法尤姆",
		TitleCode: "c_faiyum",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫纳斯Ahnas = BAhnas阿赫纳斯
	f.阿赫纳斯Ahnas.SetParent(f)

	f.巴赫纳萨Bahnasa = BBahnasa巴赫纳萨
	f.巴赫纳萨Bahnasa.SetParent(f)

	f.法尤姆Faiyum = BFaiyum法尤姆
	f.法尤姆Faiyum.SetParent(f)

	f.赫拉克琉斯Heracleus = BHeracleus赫拉克琉斯
	f.赫拉克琉斯Heracleus.SetParent(f)

}
