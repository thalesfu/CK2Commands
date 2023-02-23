package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AltmarkCounty interface {
	feud.County
	BHalberstedt哈尔伯施塔特() feud.Barony
	BLuchow吕肖() feud.Barony
	BOsterburg奥斯特堡() feud.Barony
	BSalzwedel萨尔茨韦德尔() feud.Barony
	BStendal施滕达尔() feud.Barony
	BTangermunde唐格明德() feud.Barony
	BWalbeck瓦尔贝克() feud.Barony
	BWerben韦本() feud.Barony
}

type 韦尔本AltmarkCounty struct {
	feud.BaseCounty
	哈尔伯施塔特Halberstedt feud.Barony
	吕肖Luchow          feud.Barony
	奥斯特堡Osterburg     feud.Barony
	萨尔茨韦德尔Salzwedel   feud.Barony
	施滕达尔Stendal       feud.Barony
	唐格明德Tangermunde   feud.Barony
	瓦尔贝克Walbeck       feud.Barony
	韦本Werben          feud.Barony
}

func (c *韦尔本AltmarkCounty) BHalberstedt哈尔伯施塔特() feud.Barony {
	return c.哈尔伯施塔特Halberstedt
}

func (c *韦尔本AltmarkCounty) BLuchow吕肖() feud.Barony {
	return c.吕肖Luchow
}

func (c *韦尔本AltmarkCounty) BOsterburg奥斯特堡() feud.Barony {
	return c.奥斯特堡Osterburg
}

func (c *韦尔本AltmarkCounty) BSalzwedel萨尔茨韦德尔() feud.Barony {
	return c.萨尔茨韦德尔Salzwedel
}

func (c *韦尔本AltmarkCounty) BStendal施滕达尔() feud.Barony {
	return c.施滕达尔Stendal
}

func (c *韦尔本AltmarkCounty) BTangermunde唐格明德() feud.Barony {
	return c.唐格明德Tangermunde
}

func (c *韦尔本AltmarkCounty) BWalbeck瓦尔贝克() feud.Barony {
	return c.瓦尔贝克Walbeck
}

func (c *韦尔本AltmarkCounty) BWerben韦本() feud.Barony {
	return c.韦本Werben
}

var CAltmark韦尔本 AltmarkCounty = &韦尔本AltmarkCounty{}

func init() {
	f := CAltmark韦尔本.(*韦尔本AltmarkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "309",
		Title:     "altmark",
		TitleName: "韦尔本",
		TitleCode: "c_altmark",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈尔伯施塔特Halberstedt = BHalberstedt哈尔伯施塔特
	f.哈尔伯施塔特Halberstedt.SetParent(f)

	f.吕肖Luchow = BLuchow吕肖
	f.吕肖Luchow.SetParent(f)

	f.奥斯特堡Osterburg = BOsterburg奥斯特堡
	f.奥斯特堡Osterburg.SetParent(f)

	f.萨尔茨韦德尔Salzwedel = BSalzwedel萨尔茨韦德尔
	f.萨尔茨韦德尔Salzwedel.SetParent(f)

	f.施滕达尔Stendal = BStendal施滕达尔
	f.施滕达尔Stendal.SetParent(f)

	f.唐格明德Tangermunde = BTangermunde唐格明德
	f.唐格明德Tangermunde.SetParent(f)

	f.瓦尔贝克Walbeck = BWalbeck瓦尔贝克
	f.瓦尔贝克Walbeck.SetParent(f)

	f.韦本Werben = BWerben韦本
	f.韦本Werben.SetParent(f)

}
