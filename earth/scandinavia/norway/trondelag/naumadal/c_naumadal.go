package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NaumadalCounty interface {
	feud.County
	BHalsstein哈尔斯泰因() feud.Barony
	BHegra赫格拉() feud.Barony
	BLade拉德() feud.Barony
	BLeksvik莱克斯维克() feud.Barony
	BLevanger莱旺厄尔() feud.Barony
	BLogtun罗格顿() feud.Barony
	BTinghaugen廷海于根() feud.Barony
}

type 瑙马达尔NaumadalCounty struct {
	feud.BaseCounty
	哈尔斯泰因Halsstein feud.Barony
	赫格拉Hegra       feud.Barony
	拉德Lade         feud.Barony
	莱克斯维克Leksvik   feud.Barony
	莱旺厄尔Levanger   feud.Barony
	罗格顿Logtun      feud.Barony
	廷海于根Tinghaugen feud.Barony
}

func (c *瑙马达尔NaumadalCounty) BHalsstein哈尔斯泰因() feud.Barony {
	return c.哈尔斯泰因Halsstein
}

func (c *瑙马达尔NaumadalCounty) BHegra赫格拉() feud.Barony {
	return c.赫格拉Hegra
}

func (c *瑙马达尔NaumadalCounty) BLade拉德() feud.Barony {
	return c.拉德Lade
}

func (c *瑙马达尔NaumadalCounty) BLeksvik莱克斯维克() feud.Barony {
	return c.莱克斯维克Leksvik
}

func (c *瑙马达尔NaumadalCounty) BLevanger莱旺厄尔() feud.Barony {
	return c.莱旺厄尔Levanger
}

func (c *瑙马达尔NaumadalCounty) BLogtun罗格顿() feud.Barony {
	return c.罗格顿Logtun
}

func (c *瑙马达尔NaumadalCounty) BTinghaugen廷海于根() feud.Barony {
	return c.廷海于根Tinghaugen
}

var CNaumadal瑙马达尔 NaumadalCounty = &瑙马达尔NaumadalCounty{}

func init() {
	f := CNaumadal瑙马达尔.(*瑙马达尔NaumadalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "277",
		Title:     "naumadal",
		TitleName: "瑙马达尔",
		TitleCode: "c_naumadal",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈尔斯泰因Halsstein = BHalsstein哈尔斯泰因
	f.哈尔斯泰因Halsstein.SetParent(f)

	f.赫格拉Hegra = BHegra赫格拉
	f.赫格拉Hegra.SetParent(f)

	f.拉德Lade = BLade拉德
	f.拉德Lade.SetParent(f)

	f.莱克斯维克Leksvik = BLeksvik莱克斯维克
	f.莱克斯维克Leksvik.SetParent(f)

	f.莱旺厄尔Levanger = BLevanger莱旺厄尔
	f.莱旺厄尔Levanger.SetParent(f)

	f.罗格顿Logtun = BLogtun罗格顿
	f.罗格顿Logtun.SetParent(f)

	f.廷海于根Tinghaugen = BTinghaugen廷海于根
	f.廷海于根Tinghaugen.SetParent(f)

}
