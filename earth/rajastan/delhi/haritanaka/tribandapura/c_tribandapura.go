package tribandapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TribandapuraCounty interface {
	feud.County
	BAmin阿闵() feud.Barony
	BChitwye支揭由() feud.Barony
	BDabthala陀卜诃罗() feud.Barony
	BJaintpuri耆那多城() feud.Barony
	BKartikeya迦絺吉夜() feud.Barony
	BSunam须南() feud.Barony
	BTribandapura帝利畔荼城() feud.Barony
}

type 帝利畔荼城TribandapuraCounty struct {
	feud.BaseCounty
	阿闵Amin            feud.Barony
	支揭由Chitwye        feud.Barony
	陀卜诃罗Dabthala      feud.Barony
	耆那多城Jaintpuri     feud.Barony
	迦絺吉夜Kartikeya     feud.Barony
	须南Sunam           feud.Barony
	帝利畔荼城Tribandapura feud.Barony
}

func (c *帝利畔荼城TribandapuraCounty) BAmin阿闵() feud.Barony {
	return c.阿闵Amin
}

func (c *帝利畔荼城TribandapuraCounty) BChitwye支揭由() feud.Barony {
	return c.支揭由Chitwye
}

func (c *帝利畔荼城TribandapuraCounty) BDabthala陀卜诃罗() feud.Barony {
	return c.陀卜诃罗Dabthala
}

func (c *帝利畔荼城TribandapuraCounty) BJaintpuri耆那多城() feud.Barony {
	return c.耆那多城Jaintpuri
}

func (c *帝利畔荼城TribandapuraCounty) BKartikeya迦絺吉夜() feud.Barony {
	return c.迦絺吉夜Kartikeya
}

func (c *帝利畔荼城TribandapuraCounty) BSunam须南() feud.Barony {
	return c.须南Sunam
}

func (c *帝利畔荼城TribandapuraCounty) BTribandapura帝利畔荼城() feud.Barony {
	return c.帝利畔荼城Tribandapura
}

var CTribandapura帝利畔荼城 TribandapuraCounty = &帝利畔荼城TribandapuraCounty{}

func init() {
	f := CTribandapura帝利畔荼城.(*帝利畔荼城TribandapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1364",
		Title:     "tribandapura",
		TitleName: "帝利畔荼城",
		TitleCode: "c_tribandapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿闵Amin = BAmin阿闵
	f.阿闵Amin.SetParent(f)

	f.支揭由Chitwye = BChitwye支揭由
	f.支揭由Chitwye.SetParent(f)

	f.陀卜诃罗Dabthala = BDabthala陀卜诃罗
	f.陀卜诃罗Dabthala.SetParent(f)

	f.耆那多城Jaintpuri = BJaintpuri耆那多城
	f.耆那多城Jaintpuri.SetParent(f)

	f.迦絺吉夜Kartikeya = BKartikeya迦絺吉夜
	f.迦絺吉夜Kartikeya.SetParent(f)

	f.须南Sunam = BSunam须南
	f.须南Sunam.SetParent(f)

	f.帝利畔荼城Tribandapura = BTribandapura帝利畔荼城
	f.帝利畔荼城Tribandapura.SetParent(f)

}
