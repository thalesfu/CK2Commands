package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MulujaCounty interface {
	feud.County
	BAinechchair艾因舍伊尔() feud.Barony
	BBenitajjit贝尼塔吉特() feud.Barony
	BEladergha阿德加() feud.Barony
	BElkrama克拉马() feud.Barony
	BLemseied姆塞德() feud.Barony
	BMissour米苏尔() feud.Barony
	BSouangountour苏旺贡图尔() feud.Barony
}

type 穆卢耶MulujaCounty struct {
	feud.BaseCounty
	艾因舍伊尔Ainechchair   feud.Barony
	贝尼塔吉特Benitajjit    feud.Barony
	阿德加Eladergha       feud.Barony
	克拉马Elkrama         feud.Barony
	姆塞德Lemseied        feud.Barony
	米苏尔Missour         feud.Barony
	苏旺贡图尔Souangountour feud.Barony
}

func (c *穆卢耶MulujaCounty) BAinechchair艾因舍伊尔() feud.Barony {
	return c.艾因舍伊尔Ainechchair
}

func (c *穆卢耶MulujaCounty) BBenitajjit贝尼塔吉特() feud.Barony {
	return c.贝尼塔吉特Benitajjit
}

func (c *穆卢耶MulujaCounty) BEladergha阿德加() feud.Barony {
	return c.阿德加Eladergha
}

func (c *穆卢耶MulujaCounty) BElkrama克拉马() feud.Barony {
	return c.克拉马Elkrama
}

func (c *穆卢耶MulujaCounty) BLemseied姆塞德() feud.Barony {
	return c.姆塞德Lemseied
}

func (c *穆卢耶MulujaCounty) BMissour米苏尔() feud.Barony {
	return c.米苏尔Missour
}

func (c *穆卢耶MulujaCounty) BSouangountour苏旺贡图尔() feud.Barony {
	return c.苏旺贡图尔Souangountour
}

var CMuluja穆卢耶 MulujaCounty = &穆卢耶MulujaCounty{}

func init() {
	f := CMuluja穆卢耶.(*穆卢耶MulujaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1733",
		Title:     "muluja",
		TitleName: "穆卢耶",
		TitleCode: "c_muluja",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因舍伊尔Ainechchair = BAinechchair艾因舍伊尔
	f.艾因舍伊尔Ainechchair.SetParent(f)

	f.贝尼塔吉特Benitajjit = BBenitajjit贝尼塔吉特
	f.贝尼塔吉特Benitajjit.SetParent(f)

	f.阿德加Eladergha = BEladergha阿德加
	f.阿德加Eladergha.SetParent(f)

	f.克拉马Elkrama = BElkrama克拉马
	f.克拉马Elkrama.SetParent(f)

	f.姆塞德Lemseied = BLemseied姆塞德
	f.姆塞德Lemseied.SetParent(f)

	f.米苏尔Missour = BMissour米苏尔
	f.米苏尔Missour.SetParent(f)

	f.苏旺贡图尔Souangountour = BSouangountour苏旺贡图尔
	f.苏旺贡图尔Souangountour.SetParent(f)

}
