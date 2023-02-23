package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThrakeCounty interface {
	feud.County
	BAulaeitichus奥莱梯古斯() feud.Barony
	BChariopolis哈里奥波利斯() feud.Barony
	BDeleus蔽日() feud.Barony
	BPhinopolis菲诺波利斯() feud.Barony
	BSalmydessus撒米德索斯() feud.Barony
	BSestos希斯托斯() feud.Barony
	BSyrallum希拉冷() feud.Barony
	BVerissa维里萨() feud.Barony
}

type 色雷斯ThrakeCounty struct {
	feud.BaseCounty
	奥莱梯古斯Aulaeitichus feud.Barony
	哈里奥波利斯Chariopolis feud.Barony
	蔽日Deleus          feud.Barony
	菲诺波利斯Phinopolis   feud.Barony
	撒米德索斯Salmydessus  feud.Barony
	希斯托斯Sestos        feud.Barony
	希拉冷Syrallum       feud.Barony
	维里萨Verissa        feud.Barony
}

func (c *色雷斯ThrakeCounty) BAulaeitichus奥莱梯古斯() feud.Barony {
	return c.奥莱梯古斯Aulaeitichus
}

func (c *色雷斯ThrakeCounty) BChariopolis哈里奥波利斯() feud.Barony {
	return c.哈里奥波利斯Chariopolis
}

func (c *色雷斯ThrakeCounty) BDeleus蔽日() feud.Barony {
	return c.蔽日Deleus
}

func (c *色雷斯ThrakeCounty) BPhinopolis菲诺波利斯() feud.Barony {
	return c.菲诺波利斯Phinopolis
}

func (c *色雷斯ThrakeCounty) BSalmydessus撒米德索斯() feud.Barony {
	return c.撒米德索斯Salmydessus
}

func (c *色雷斯ThrakeCounty) BSestos希斯托斯() feud.Barony {
	return c.希斯托斯Sestos
}

func (c *色雷斯ThrakeCounty) BSyrallum希拉冷() feud.Barony {
	return c.希拉冷Syrallum
}

func (c *色雷斯ThrakeCounty) BVerissa维里萨() feud.Barony {
	return c.维里萨Verissa
}

var CThrake色雷斯 ThrakeCounty = &色雷斯ThrakeCounty{}

func init() {
	f := CThrake色雷斯.(*色雷斯ThrakeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "497",
		Title:     "thrake",
		TitleName: "色雷斯",
		TitleCode: "c_thrake",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥莱梯古斯Aulaeitichus = BAulaeitichus奥莱梯古斯
	f.奥莱梯古斯Aulaeitichus.SetParent(f)

	f.哈里奥波利斯Chariopolis = BChariopolis哈里奥波利斯
	f.哈里奥波利斯Chariopolis.SetParent(f)

	f.蔽日Deleus = BDeleus蔽日
	f.蔽日Deleus.SetParent(f)

	f.菲诺波利斯Phinopolis = BPhinopolis菲诺波利斯
	f.菲诺波利斯Phinopolis.SetParent(f)

	f.撒米德索斯Salmydessus = BSalmydessus撒米德索斯
	f.撒米德索斯Salmydessus.SetParent(f)

	f.希斯托斯Sestos = BSestos希斯托斯
	f.希斯托斯Sestos.SetParent(f)

	f.希拉冷Syrallum = BSyrallum希拉冷
	f.希拉冷Syrallum.SetParent(f)

	f.维里萨Verissa = BVerissa维里萨
	f.维里萨Verissa.SetParent(f)

}
