package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KentCounty interface {
	feud.County
	BCanterbury坎特伯雷() feud.Barony
	BDover多佛尔() feud.Barony
	BFaversham法弗舍姆() feud.Barony
	BLympne利姆() feud.Barony
	BRochester罗切斯特() feud.Barony
	BRomney罗姆尼() feud.Barony
	BSandwich桑威奇() feud.Barony
	BTonbridge汤布里奇() feud.Barony
}

type 肯特KentCounty struct {
	feud.BaseCounty
	坎特伯雷Canterbury feud.Barony
	多佛尔Dover       feud.Barony
	法弗舍姆Faversham  feud.Barony
	利姆Lympne       feud.Barony
	罗切斯特Rochester  feud.Barony
	罗姆尼Romney      feud.Barony
	桑威奇Sandwich    feud.Barony
	汤布里奇Tonbridge  feud.Barony
}

func (c *肯特KentCounty) BCanterbury坎特伯雷() feud.Barony {
	return c.坎特伯雷Canterbury
}

func (c *肯特KentCounty) BDover多佛尔() feud.Barony {
	return c.多佛尔Dover
}

func (c *肯特KentCounty) BFaversham法弗舍姆() feud.Barony {
	return c.法弗舍姆Faversham
}

func (c *肯特KentCounty) BLympne利姆() feud.Barony {
	return c.利姆Lympne
}

func (c *肯特KentCounty) BRochester罗切斯特() feud.Barony {
	return c.罗切斯特Rochester
}

func (c *肯特KentCounty) BRomney罗姆尼() feud.Barony {
	return c.罗姆尼Romney
}

func (c *肯特KentCounty) BSandwich桑威奇() feud.Barony {
	return c.桑威奇Sandwich
}

func (c *肯特KentCounty) BTonbridge汤布里奇() feud.Barony {
	return c.汤布里奇Tonbridge
}

var CKent肯特 KentCounty = &肯特KentCounty{}

func init() {
	f := CKent肯特.(*肯特KentCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "73",
		Title:     "kent",
		TitleName: "肯特",
		TitleCode: "c_kent",
		Baronies:  map[string]feud.Barony{},
	}

	f.坎特伯雷Canterbury = BCanterbury坎特伯雷
	f.坎特伯雷Canterbury.SetParent(f)

	f.多佛尔Dover = BDover多佛尔
	f.多佛尔Dover.SetParent(f)

	f.法弗舍姆Faversham = BFaversham法弗舍姆
	f.法弗舍姆Faversham.SetParent(f)

	f.利姆Lympne = BLympne利姆
	f.利姆Lympne.SetParent(f)

	f.罗切斯特Rochester = BRochester罗切斯特
	f.罗切斯特Rochester.SetParent(f)

	f.罗姆尼Romney = BRomney罗姆尼
	f.罗姆尼Romney.SetParent(f)

	f.桑威奇Sandwich = BSandwich桑威奇
	f.桑威奇Sandwich.SetParent(f)

	f.汤布里奇Tonbridge = BTonbridge汤布里奇
	f.汤布里奇Tonbridge.SetParent(f)

}
