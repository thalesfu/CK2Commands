package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BacauCounty interface {
	feud.County
	BBacau巴克乌() feud.Barony
	BBodesti博代什蒂() feud.Barony
	BNeamnt尼亚姆茨() feud.Barony
	BTazlau塔兹勒乌() feud.Barony
}

type 巴克乌BacauCounty struct {
	feud.BaseCounty
	巴克乌Bacau    feud.Barony
	博代什蒂Bodesti feud.Barony
	尼亚姆茨Neamnt  feud.Barony
	塔兹勒乌Tazlau  feud.Barony
}

func (c *巴克乌BacauCounty) BBacau巴克乌() feud.Barony {
	return c.巴克乌Bacau
}

func (c *巴克乌BacauCounty) BBodesti博代什蒂() feud.Barony {
	return c.博代什蒂Bodesti
}

func (c *巴克乌BacauCounty) BNeamnt尼亚姆茨() feud.Barony {
	return c.尼亚姆茨Neamnt
}

func (c *巴克乌BacauCounty) BTazlau塔兹勒乌() feud.Barony {
	return c.塔兹勒乌Tazlau
}

var CBacau巴克乌 BacauCounty = &巴克乌BacauCounty{}

func init() {
	f := CBacau巴克乌.(*巴克乌BacauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1639",
		Title:     "bacau",
		TitleName: "巴克乌",
		TitleCode: "c_bacau",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴克乌Bacau = BBacau巴克乌
	f.巴克乌Bacau.SetParent(f)

	f.博代什蒂Bodesti = BBodesti博代什蒂
	f.博代什蒂Bodesti.SetParent(f)

	f.尼亚姆茨Neamnt = BNeamnt尼亚姆茨
	f.尼亚姆茨Neamnt.SetParent(f)

	f.塔兹勒乌Tazlau = BTazlau塔兹勒乌
	f.塔兹勒乌Tazlau.SetParent(f)

}
