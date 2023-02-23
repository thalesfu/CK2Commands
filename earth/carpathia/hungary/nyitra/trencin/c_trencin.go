package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrencinCounty interface {
	feud.County
	BBan班() feud.Barony
	BCongsberg孔斯堡() feud.Barony
	BIllava伊洛沃() feud.Barony
	BPovazskabystrica瓦赫河畔比斯特里察() feud.Barony
	BPuho普霍() feud.Barony
	BTrencsen特伦琴() feud.Barony
	BTuroc图罗茨() feud.Barony
	BZilina日利纳() feud.Barony
}

type 特伦钦TrencinCounty struct {
	feud.BaseCounty
	班Ban                      feud.Barony
	孔斯堡Congsberg              feud.Barony
	伊洛沃Illava                 feud.Barony
	瓦赫河畔比斯特里察Povazskabystrica feud.Barony
	普霍Puho                    feud.Barony
	特伦琴Trencsen               feud.Barony
	图罗茨Turoc                  feud.Barony
	日利纳Zilina                 feud.Barony
}

func (c *特伦钦TrencinCounty) BBan班() feud.Barony {
	return c.班Ban
}

func (c *特伦钦TrencinCounty) BCongsberg孔斯堡() feud.Barony {
	return c.孔斯堡Congsberg
}

func (c *特伦钦TrencinCounty) BIllava伊洛沃() feud.Barony {
	return c.伊洛沃Illava
}

func (c *特伦钦TrencinCounty) BPovazskabystrica瓦赫河畔比斯特里察() feud.Barony {
	return c.瓦赫河畔比斯特里察Povazskabystrica
}

func (c *特伦钦TrencinCounty) BPuho普霍() feud.Barony {
	return c.普霍Puho
}

func (c *特伦钦TrencinCounty) BTrencsen特伦琴() feud.Barony {
	return c.特伦琴Trencsen
}

func (c *特伦钦TrencinCounty) BTuroc图罗茨() feud.Barony {
	return c.图罗茨Turoc
}

func (c *特伦钦TrencinCounty) BZilina日利纳() feud.Barony {
	return c.日利纳Zilina
}

var CTrencin特伦钦 TrencinCounty = &特伦钦TrencinCounty{}

func init() {
	f := CTrencin特伦钦.(*特伦钦TrencinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "442",
		Title:     "trencin",
		TitleName: "特伦钦",
		TitleCode: "c_trencin",
		Baronies:  map[string]feud.Barony{},
	}

	f.班Ban = BBan班
	f.班Ban.SetParent(f)

	f.孔斯堡Congsberg = BCongsberg孔斯堡
	f.孔斯堡Congsberg.SetParent(f)

	f.伊洛沃Illava = BIllava伊洛沃
	f.伊洛沃Illava.SetParent(f)

	f.瓦赫河畔比斯特里察Povazskabystrica = BPovazskabystrica瓦赫河畔比斯特里察
	f.瓦赫河畔比斯特里察Povazskabystrica.SetParent(f)

	f.普霍Puho = BPuho普霍
	f.普霍Puho.SetParent(f)

	f.特伦琴Trencsen = BTrencsen特伦琴
	f.特伦琴Trencsen.SetParent(f)

	f.图罗茨Turoc = BTuroc图罗茨
	f.图罗茨Turoc.SetParent(f)

	f.日利纳Zilina = BZilina日利纳
	f.日利纳Zilina.SetParent(f)

}
