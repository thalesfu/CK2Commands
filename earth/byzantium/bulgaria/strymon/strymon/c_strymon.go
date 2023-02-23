package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StrymonCounty interface {
	feud.County
	BKocane科查内() feud.Barony
	BKratovo克拉托沃() feud.Barony
	BPrilep普里莱普() feud.Barony
	BProsek普罗塞克() feud.Barony
	BRadovish拉多维什() feud.Barony
	BStrumica斯特鲁米察() feud.Barony
	BTrikves蒂克韦什() feud.Barony
}

type 斯特里蒙StrymonCounty struct {
	feud.BaseCounty
	科查内Kocane     feud.Barony
	克拉托沃Kratovo   feud.Barony
	普里莱普Prilep    feud.Barony
	普罗塞克Prosek    feud.Barony
	拉多维什Radovish  feud.Barony
	斯特鲁米察Strumica feud.Barony
	蒂克韦什Trikves   feud.Barony
}

func (c *斯特里蒙StrymonCounty) BKocane科查内() feud.Barony {
	return c.科查内Kocane
}

func (c *斯特里蒙StrymonCounty) BKratovo克拉托沃() feud.Barony {
	return c.克拉托沃Kratovo
}

func (c *斯特里蒙StrymonCounty) BPrilep普里莱普() feud.Barony {
	return c.普里莱普Prilep
}

func (c *斯特里蒙StrymonCounty) BProsek普罗塞克() feud.Barony {
	return c.普罗塞克Prosek
}

func (c *斯特里蒙StrymonCounty) BRadovish拉多维什() feud.Barony {
	return c.拉多维什Radovish
}

func (c *斯特里蒙StrymonCounty) BStrumica斯特鲁米察() feud.Barony {
	return c.斯特鲁米察Strumica
}

func (c *斯特里蒙StrymonCounty) BTrikves蒂克韦什() feud.Barony {
	return c.蒂克韦什Trikves
}

var CStrymon斯特里蒙 StrymonCounty = &斯特里蒙StrymonCounty{}

func init() {
	f := CStrymon斯特里蒙.(*斯特里蒙StrymonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "492",
		Title:     "strymon",
		TitleName: "斯特里蒙",
		TitleCode: "c_strymon",
		Baronies:  map[string]feud.Barony{},
	}

	f.科查内Kocane = BKocane科查内
	f.科查内Kocane.SetParent(f)

	f.克拉托沃Kratovo = BKratovo克拉托沃
	f.克拉托沃Kratovo.SetParent(f)

	f.普里莱普Prilep = BPrilep普里莱普
	f.普里莱普Prilep.SetParent(f)

	f.普罗塞克Prosek = BProsek普罗塞克
	f.普罗塞克Prosek.SetParent(f)

	f.拉多维什Radovish = BRadovish拉多维什
	f.拉多维什Radovish.SetParent(f)

	f.斯特鲁米察Strumica = BStrumica斯特鲁米察
	f.斯特鲁米察Strumica.SetParent(f)

	f.蒂克韦什Trikves = BTrikves蒂克韦什
	f.蒂克韦什Trikves.SetParent(f)

}
