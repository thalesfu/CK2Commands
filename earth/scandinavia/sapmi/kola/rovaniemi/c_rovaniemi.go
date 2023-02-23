package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RovaniemiCounty interface {
	feud.County
	BJeesio耶西厄() feud.Barony
	BKitinen基蒂宁() feud.Barony
	BLuiro卢伊罗() feud.Barony
	BNakkala内凯莱() feud.Barony
	BOunas欧纳斯() feud.Barony
	BRaudan劳达() feud.Barony
	BRovaniemi罗瓦涅米() feud.Barony
}

type 罗瓦涅米RovaniemiCounty struct {
	feud.BaseCounty
	耶西厄Jeesio     feud.Barony
	基蒂宁Kitinen    feud.Barony
	卢伊罗Luiro      feud.Barony
	内凯莱Nakkala    feud.Barony
	欧纳斯Ounas      feud.Barony
	劳达Raudan      feud.Barony
	罗瓦涅米Rovaniemi feud.Barony
}

func (c *罗瓦涅米RovaniemiCounty) BJeesio耶西厄() feud.Barony {
	return c.耶西厄Jeesio
}

func (c *罗瓦涅米RovaniemiCounty) BKitinen基蒂宁() feud.Barony {
	return c.基蒂宁Kitinen
}

func (c *罗瓦涅米RovaniemiCounty) BLuiro卢伊罗() feud.Barony {
	return c.卢伊罗Luiro
}

func (c *罗瓦涅米RovaniemiCounty) BNakkala内凯莱() feud.Barony {
	return c.内凯莱Nakkala
}

func (c *罗瓦涅米RovaniemiCounty) BOunas欧纳斯() feud.Barony {
	return c.欧纳斯Ounas
}

func (c *罗瓦涅米RovaniemiCounty) BRaudan劳达() feud.Barony {
	return c.劳达Raudan
}

func (c *罗瓦涅米RovaniemiCounty) BRovaniemi罗瓦涅米() feud.Barony {
	return c.罗瓦涅米Rovaniemi
}

var CRovaniemi罗瓦涅米 RovaniemiCounty = &罗瓦涅米RovaniemiCounty{}

func init() {
	f := CRovaniemi罗瓦涅米.(*罗瓦涅米RovaniemiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1604",
		Title:     "rovaniemi",
		TitleName: "罗瓦涅米",
		TitleCode: "c_rovaniemi",
		Baronies:  map[string]feud.Barony{},
	}

	f.耶西厄Jeesio = BJeesio耶西厄
	f.耶西厄Jeesio.SetParent(f)

	f.基蒂宁Kitinen = BKitinen基蒂宁
	f.基蒂宁Kitinen.SetParent(f)

	f.卢伊罗Luiro = BLuiro卢伊罗
	f.卢伊罗Luiro.SetParent(f)

	f.内凯莱Nakkala = BNakkala内凯莱
	f.内凯莱Nakkala.SetParent(f)

	f.欧纳斯Ounas = BOunas欧纳斯
	f.欧纳斯Ounas.SetParent(f)

	f.劳达Raudan = BRaudan劳达
	f.劳达Raudan.SetParent(f)

	f.罗瓦涅米Rovaniemi = BRovaniemi罗瓦涅米
	f.罗瓦涅米Rovaniemi.SetParent(f)

}
