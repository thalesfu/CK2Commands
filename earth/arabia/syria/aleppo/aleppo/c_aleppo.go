package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AleppoCounty interface {
	feud.County
	BAleppo阿勒颇() feud.Barony
	BAzaz阿扎兹() feud.Barony
	BHarim哈里姆() feud.Barony
	BLerminet莱尔米讷() feud.Barony
	BMaaratannuman迈阿赖努阿曼() feud.Barony
	BMashala玛沙拉() feud.Barony
	BQusayr古赛尔() feud.Barony
	BSarmin沙敏() feud.Barony
}

type 阿勒颇AleppoCounty struct {
	feud.BaseCounty
	阿勒颇Aleppo           feud.Barony
	阿扎兹Azaz             feud.Barony
	哈里姆Harim            feud.Barony
	莱尔米讷Lerminet        feud.Barony
	迈阿赖努阿曼Maaratannuman feud.Barony
	玛沙拉Mashala          feud.Barony
	古赛尔Qusayr           feud.Barony
	沙敏Sarmin            feud.Barony
}

func (c *阿勒颇AleppoCounty) BAleppo阿勒颇() feud.Barony {
	return c.阿勒颇Aleppo
}

func (c *阿勒颇AleppoCounty) BAzaz阿扎兹() feud.Barony {
	return c.阿扎兹Azaz
}

func (c *阿勒颇AleppoCounty) BHarim哈里姆() feud.Barony {
	return c.哈里姆Harim
}

func (c *阿勒颇AleppoCounty) BLerminet莱尔米讷() feud.Barony {
	return c.莱尔米讷Lerminet
}

func (c *阿勒颇AleppoCounty) BMaaratannuman迈阿赖努阿曼() feud.Barony {
	return c.迈阿赖努阿曼Maaratannuman
}

func (c *阿勒颇AleppoCounty) BMashala玛沙拉() feud.Barony {
	return c.玛沙拉Mashala
}

func (c *阿勒颇AleppoCounty) BQusayr古赛尔() feud.Barony {
	return c.古赛尔Qusayr
}

func (c *阿勒颇AleppoCounty) BSarmin沙敏() feud.Barony {
	return c.沙敏Sarmin
}

var CAleppo阿勒颇 AleppoCounty = &阿勒颇AleppoCounty{}

func init() {
	f := CAleppo阿勒颇.(*阿勒颇AleppoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "733",
		Title:     "aleppo",
		TitleName: "阿勒颇",
		TitleCode: "c_aleppo",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿勒颇Aleppo = BAleppo阿勒颇
	f.阿勒颇Aleppo.SetParent(f)

	f.阿扎兹Azaz = BAzaz阿扎兹
	f.阿扎兹Azaz.SetParent(f)

	f.哈里姆Harim = BHarim哈里姆
	f.哈里姆Harim.SetParent(f)

	f.莱尔米讷Lerminet = BLerminet莱尔米讷
	f.莱尔米讷Lerminet.SetParent(f)

	f.迈阿赖努阿曼Maaratannuman = BMaaratannuman迈阿赖努阿曼
	f.迈阿赖努阿曼Maaratannuman.SetParent(f)

	f.玛沙拉Mashala = BMashala玛沙拉
	f.玛沙拉Mashala.SetParent(f)

	f.古赛尔Qusayr = BQusayr古赛尔
	f.古赛尔Qusayr.SetParent(f)

	f.沙敏Sarmin = BSarmin沙敏
	f.沙敏Sarmin.SetParent(f)

}
