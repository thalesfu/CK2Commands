package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SalernoCounty interface {
	feud.County
	BAcenzera阿琴泽拉() feud.Barony
	BAcerno阿切尔诺() feud.Barony
	BAgropoli阿格罗波利() feud.Barony
	BEboli埃博利() feud.Barony
	BLucania卢卡尼亚() feud.Barony
	BNocera诺切拉() feud.Barony
	BSalerno萨莱诺() feud.Barony
	BSarno萨尔诺() feud.Barony
}

type 萨莱诺SalernoCounty struct {
	feud.BaseCounty
	阿琴泽拉Acenzera  feud.Barony
	阿切尔诺Acerno    feud.Barony
	阿格罗波利Agropoli feud.Barony
	埃博利Eboli      feud.Barony
	卢卡尼亚Lucania   feud.Barony
	诺切拉Nocera     feud.Barony
	萨莱诺Salerno    feud.Barony
	萨尔诺Sarno      feud.Barony
}

func (c *萨莱诺SalernoCounty) BAcenzera阿琴泽拉() feud.Barony {
	return c.阿琴泽拉Acenzera
}

func (c *萨莱诺SalernoCounty) BAcerno阿切尔诺() feud.Barony {
	return c.阿切尔诺Acerno
}

func (c *萨莱诺SalernoCounty) BAgropoli阿格罗波利() feud.Barony {
	return c.阿格罗波利Agropoli
}

func (c *萨莱诺SalernoCounty) BEboli埃博利() feud.Barony {
	return c.埃博利Eboli
}

func (c *萨莱诺SalernoCounty) BLucania卢卡尼亚() feud.Barony {
	return c.卢卡尼亚Lucania
}

func (c *萨莱诺SalernoCounty) BNocera诺切拉() feud.Barony {
	return c.诺切拉Nocera
}

func (c *萨莱诺SalernoCounty) BSalerno萨莱诺() feud.Barony {
	return c.萨莱诺Salerno
}

func (c *萨莱诺SalernoCounty) BSarno萨尔诺() feud.Barony {
	return c.萨尔诺Sarno
}

var CSalerno萨莱诺 SalernoCounty = &萨莱诺SalernoCounty{}

func init() {
	f := CSalerno萨莱诺.(*萨莱诺SalernoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "336",
		Title:     "salerno",
		TitleName: "萨莱诺",
		TitleCode: "c_salerno",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿琴泽拉Acenzera = BAcenzera阿琴泽拉
	f.阿琴泽拉Acenzera.SetParent(f)

	f.阿切尔诺Acerno = BAcerno阿切尔诺
	f.阿切尔诺Acerno.SetParent(f)

	f.阿格罗波利Agropoli = BAgropoli阿格罗波利
	f.阿格罗波利Agropoli.SetParent(f)

	f.埃博利Eboli = BEboli埃博利
	f.埃博利Eboli.SetParent(f)

	f.卢卡尼亚Lucania = BLucania卢卡尼亚
	f.卢卡尼亚Lucania.SetParent(f)

	f.诺切拉Nocera = BNocera诺切拉
	f.诺切拉Nocera.SetParent(f)

	f.萨莱诺Salerno = BSalerno萨莱诺
	f.萨莱诺Salerno.SetParent(f)

	f.萨尔诺Sarno = BSarno萨尔诺
	f.萨尔诺Sarno.SetParent(f)

}
