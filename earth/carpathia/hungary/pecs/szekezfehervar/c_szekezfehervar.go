package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SzekezfehervarCounty interface {
	feud.County
	BBarcs鲍尔奇() feud.Barony
	BCsurgo丘尔戈() feud.Barony
	BKaposvar考波什堡() feud.Barony
	BLengyeltoti伦杰尔托蒂() feud.Barony
	BMarcali毛尔曹利() feud.Barony
	BNagyatad瑙吉奥塔德() feud.Barony
	BSzekezfehervar塞克什白堡() feud.Barony
	BSzigetvar锡盖特堡() feud.Barony
}

type 塞克什白堡SzekezfehervarCounty struct {
	feud.BaseCounty
	鲍尔奇Barcs            feud.Barony
	丘尔戈Csurgo           feud.Barony
	考波什堡Kaposvar        feud.Barony
	伦杰尔托蒂Lengyeltoti    feud.Barony
	毛尔曹利Marcali         feud.Barony
	瑙吉奥塔德Nagyatad       feud.Barony
	塞克什白堡Szekezfehervar feud.Barony
	锡盖特堡Szigetvar       feud.Barony
}

func (c *塞克什白堡SzekezfehervarCounty) BBarcs鲍尔奇() feud.Barony {
	return c.鲍尔奇Barcs
}

func (c *塞克什白堡SzekezfehervarCounty) BCsurgo丘尔戈() feud.Barony {
	return c.丘尔戈Csurgo
}

func (c *塞克什白堡SzekezfehervarCounty) BKaposvar考波什堡() feud.Barony {
	return c.考波什堡Kaposvar
}

func (c *塞克什白堡SzekezfehervarCounty) BLengyeltoti伦杰尔托蒂() feud.Barony {
	return c.伦杰尔托蒂Lengyeltoti
}

func (c *塞克什白堡SzekezfehervarCounty) BMarcali毛尔曹利() feud.Barony {
	return c.毛尔曹利Marcali
}

func (c *塞克什白堡SzekezfehervarCounty) BNagyatad瑙吉奥塔德() feud.Barony {
	return c.瑙吉奥塔德Nagyatad
}

func (c *塞克什白堡SzekezfehervarCounty) BSzekezfehervar塞克什白堡() feud.Barony {
	return c.塞克什白堡Szekezfehervar
}

func (c *塞克什白堡SzekezfehervarCounty) BSzigetvar锡盖特堡() feud.Barony {
	return c.锡盖特堡Szigetvar
}

var CSzekezfehervar塞克什白堡 SzekezfehervarCounty = &塞克什白堡SzekezfehervarCounty{}

func init() {
	f := CSzekezfehervar塞克什白堡.(*塞克什白堡SzekezfehervarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "453",
		Title:     "szekezfehervar",
		TitleName: "塞克什白堡",
		TitleCode: "c_szekezfehervar",
		Baronies:  map[string]feud.Barony{},
	}

	f.鲍尔奇Barcs = BBarcs鲍尔奇
	f.鲍尔奇Barcs.SetParent(f)

	f.丘尔戈Csurgo = BCsurgo丘尔戈
	f.丘尔戈Csurgo.SetParent(f)

	f.考波什堡Kaposvar = BKaposvar考波什堡
	f.考波什堡Kaposvar.SetParent(f)

	f.伦杰尔托蒂Lengyeltoti = BLengyeltoti伦杰尔托蒂
	f.伦杰尔托蒂Lengyeltoti.SetParent(f)

	f.毛尔曹利Marcali = BMarcali毛尔曹利
	f.毛尔曹利Marcali.SetParent(f)

	f.瑙吉奥塔德Nagyatad = BNagyatad瑙吉奥塔德
	f.瑙吉奥塔德Nagyatad.SetParent(f)

	f.塞克什白堡Szekezfehervar = BSzekezfehervar塞克什白堡
	f.塞克什白堡Szekezfehervar.SetParent(f)

	f.锡盖特堡Szigetvar = BSzigetvar锡盖特堡
	f.锡盖特堡Szigetvar.SetParent(f)

}
