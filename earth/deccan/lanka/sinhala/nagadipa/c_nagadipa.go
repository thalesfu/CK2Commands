package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NagadipaCounty interface {
	feud.County
	BAgia阿耆阿() feud.Barony
	BManertha摩涅他() feud.Barony
	BNainativu奈纳岛() feud.Barony
	BNallur讷卢尔() feud.Barony
	BSukaratittha苏迦罗咥他() feud.Barony
	BVallipuram婆利浮兰() feud.Barony
	BYapapatuna耶波波都那() feud.Barony
}

type 那伽洲NagadipaCounty struct {
	feud.BaseCounty
	阿耆阿Agia           feud.Barony
	摩涅他Manertha       feud.Barony
	奈纳岛Nainativu      feud.Barony
	讷卢尔Nallur         feud.Barony
	苏迦罗咥他Sukaratittha feud.Barony
	婆利浮兰Vallipuram    feud.Barony
	耶波波都那Yapapatuna   feud.Barony
}

func (c *那伽洲NagadipaCounty) BAgia阿耆阿() feud.Barony {
	return c.阿耆阿Agia
}

func (c *那伽洲NagadipaCounty) BManertha摩涅他() feud.Barony {
	return c.摩涅他Manertha
}

func (c *那伽洲NagadipaCounty) BNainativu奈纳岛() feud.Barony {
	return c.奈纳岛Nainativu
}

func (c *那伽洲NagadipaCounty) BNallur讷卢尔() feud.Barony {
	return c.讷卢尔Nallur
}

func (c *那伽洲NagadipaCounty) BSukaratittha苏迦罗咥他() feud.Barony {
	return c.苏迦罗咥他Sukaratittha
}

func (c *那伽洲NagadipaCounty) BVallipuram婆利浮兰() feud.Barony {
	return c.婆利浮兰Vallipuram
}

func (c *那伽洲NagadipaCounty) BYapapatuna耶波波都那() feud.Barony {
	return c.耶波波都那Yapapatuna
}

var CNagadipa那伽洲 NagadipaCounty = &那伽洲NagadipaCounty{}

func init() {
	f := CNagadipa那伽洲.(*那伽洲NagadipaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1194",
		Title:     "nagadipa",
		TitleName: "那伽洲",
		TitleCode: "c_nagadipa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿耆阿Agia = BAgia阿耆阿
	f.阿耆阿Agia.SetParent(f)

	f.摩涅他Manertha = BManertha摩涅他
	f.摩涅他Manertha.SetParent(f)

	f.奈纳岛Nainativu = BNainativu奈纳岛
	f.奈纳岛Nainativu.SetParent(f)

	f.讷卢尔Nallur = BNallur讷卢尔
	f.讷卢尔Nallur.SetParent(f)

	f.苏迦罗咥他Sukaratittha = BSukaratittha苏迦罗咥他
	f.苏迦罗咥他Sukaratittha.SetParent(f)

	f.婆利浮兰Vallipuram = BVallipuram婆利浮兰
	f.婆利浮兰Vallipuram.SetParent(f)

	f.耶波波都那Yapapatuna = BYapapatuna耶波波都那
	f.耶波波都那Yapapatuna.SetParent(f)

}
