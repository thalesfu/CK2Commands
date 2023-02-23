package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JamtlandCounty interface {
	feud.County
	BBorgvattnet博里瓦特讷() feud.Barony
	BBrunflo布伦弗卢() feud.Barony
	BGene耶内() feud.Barony
	BHusan许松() feud.Barony
	BMjalleborgen姆耶勒博尔根() feud.Barony
	BOfferdal奥弗达尔() feud.Barony
	BVasterhus韦斯特胡斯() feud.Barony
}

type 耶姆特兰JamtlandCounty struct {
	feud.BaseCounty
	博里瓦特讷Borgvattnet   feud.Barony
	布伦弗卢Brunflo        feud.Barony
	耶内Gene             feud.Barony
	许松Husan            feud.Barony
	姆耶勒博尔根Mjalleborgen feud.Barony
	奥弗达尔Offerdal       feud.Barony
	韦斯特胡斯Vasterhus     feud.Barony
}

func (c *耶姆特兰JamtlandCounty) BBorgvattnet博里瓦特讷() feud.Barony {
	return c.博里瓦特讷Borgvattnet
}

func (c *耶姆特兰JamtlandCounty) BBrunflo布伦弗卢() feud.Barony {
	return c.布伦弗卢Brunflo
}

func (c *耶姆特兰JamtlandCounty) BGene耶内() feud.Barony {
	return c.耶内Gene
}

func (c *耶姆特兰JamtlandCounty) BHusan许松() feud.Barony {
	return c.许松Husan
}

func (c *耶姆特兰JamtlandCounty) BMjalleborgen姆耶勒博尔根() feud.Barony {
	return c.姆耶勒博尔根Mjalleborgen
}

func (c *耶姆特兰JamtlandCounty) BOfferdal奥弗达尔() feud.Barony {
	return c.奥弗达尔Offerdal
}

func (c *耶姆特兰JamtlandCounty) BVasterhus韦斯特胡斯() feud.Barony {
	return c.韦斯特胡斯Vasterhus
}

var CJamtland耶姆特兰 JamtlandCounty = &耶姆特兰JamtlandCounty{}

func init() {
	f := CJamtland耶姆特兰.(*耶姆特兰JamtlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "282",
		Title:     "jamtland",
		TitleName: "耶姆特兰",
		TitleCode: "c_jamtland",
		Baronies:  map[string]feud.Barony{},
	}

	f.博里瓦特讷Borgvattnet = BBorgvattnet博里瓦特讷
	f.博里瓦特讷Borgvattnet.SetParent(f)

	f.布伦弗卢Brunflo = BBrunflo布伦弗卢
	f.布伦弗卢Brunflo.SetParent(f)

	f.耶内Gene = BGene耶内
	f.耶内Gene.SetParent(f)

	f.许松Husan = BHusan许松
	f.许松Husan.SetParent(f)

	f.姆耶勒博尔根Mjalleborgen = BMjalleborgen姆耶勒博尔根
	f.姆耶勒博尔根Mjalleborgen.SetParent(f)

	f.奥弗达尔Offerdal = BOfferdal奥弗达尔
	f.奥弗达尔Offerdal.SetParent(f)

	f.韦斯特胡斯Vasterhus = BVasterhus韦斯特胡斯
	f.韦斯特胡斯Vasterhus.SetParent(f)

}
