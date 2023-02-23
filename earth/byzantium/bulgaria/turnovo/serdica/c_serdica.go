package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SerdicaCounty interface {
	feud.County
	BBreznik布雷兹尼克() feud.Barony
	BEtropole埃特罗波莱() feud.Barony
	BPernik佩尔尼克() feud.Barony
	BPravets普拉韦茨() feud.Barony
	BRila里拉() feud.Barony
	BSamundzhievo萨蒙吉耶沃() feud.Barony
	BSerdica塞尔迪卡() feud.Barony
	BVelbazhd韦尔伯日德() feud.Barony
}

type 塞尔迪卡SerdicaCounty struct {
	feud.BaseCounty
	布雷兹尼克Breznik      feud.Barony
	埃特罗波莱Etropole     feud.Barony
	佩尔尼克Pernik        feud.Barony
	普拉韦茨Pravets       feud.Barony
	里拉Rila            feud.Barony
	萨蒙吉耶沃Samundzhievo feud.Barony
	塞尔迪卡Serdica       feud.Barony
	韦尔伯日德Velbazhd     feud.Barony
}

func (c *塞尔迪卡SerdicaCounty) BBreznik布雷兹尼克() feud.Barony {
	return c.布雷兹尼克Breznik
}

func (c *塞尔迪卡SerdicaCounty) BEtropole埃特罗波莱() feud.Barony {
	return c.埃特罗波莱Etropole
}

func (c *塞尔迪卡SerdicaCounty) BPernik佩尔尼克() feud.Barony {
	return c.佩尔尼克Pernik
}

func (c *塞尔迪卡SerdicaCounty) BPravets普拉韦茨() feud.Barony {
	return c.普拉韦茨Pravets
}

func (c *塞尔迪卡SerdicaCounty) BRila里拉() feud.Barony {
	return c.里拉Rila
}

func (c *塞尔迪卡SerdicaCounty) BSamundzhievo萨蒙吉耶沃() feud.Barony {
	return c.萨蒙吉耶沃Samundzhievo
}

func (c *塞尔迪卡SerdicaCounty) BSerdica塞尔迪卡() feud.Barony {
	return c.塞尔迪卡Serdica
}

func (c *塞尔迪卡SerdicaCounty) BVelbazhd韦尔伯日德() feud.Barony {
	return c.韦尔伯日德Velbazhd
}

var CSerdica塞尔迪卡 SerdicaCounty = &塞尔迪卡SerdicaCounty{}

func init() {
	f := CSerdica塞尔迪卡.(*塞尔迪卡SerdicaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "500",
		Title:     "serdica",
		TitleName: "塞尔迪卡",
		TitleCode: "c_serdica",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷兹尼克Breznik = BBreznik布雷兹尼克
	f.布雷兹尼克Breznik.SetParent(f)

	f.埃特罗波莱Etropole = BEtropole埃特罗波莱
	f.埃特罗波莱Etropole.SetParent(f)

	f.佩尔尼克Pernik = BPernik佩尔尼克
	f.佩尔尼克Pernik.SetParent(f)

	f.普拉韦茨Pravets = BPravets普拉韦茨
	f.普拉韦茨Pravets.SetParent(f)

	f.里拉Rila = BRila里拉
	f.里拉Rila.SetParent(f)

	f.萨蒙吉耶沃Samundzhievo = BSamundzhievo萨蒙吉耶沃
	f.萨蒙吉耶沃Samundzhievo.SetParent(f)

	f.塞尔迪卡Serdica = BSerdica塞尔迪卡
	f.塞尔迪卡Serdica.SetParent(f)

	f.韦尔伯日德Velbazhd = BVelbazhd韦尔伯日德
	f.韦尔伯日德Velbazhd.SetParent(f)

}
