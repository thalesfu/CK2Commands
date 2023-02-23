package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WightCounty interface {
	feud.County
	BBonchurch本彻奇() feud.Barony
	BCarisbrooke卡里斯布鲁克() feud.Barony
	BCowes考斯() feud.Barony
	BSandown桑当() feud.Barony
	BYarmouth雅茅斯() feud.Barony
}

type 怀特WightCounty struct {
	feud.BaseCounty
	本彻奇Bonchurch      feud.Barony
	卡里斯布鲁克Carisbrooke feud.Barony
	考斯Cowes           feud.Barony
	桑当Sandown         feud.Barony
	雅茅斯Yarmouth       feud.Barony
}

func (c *怀特WightCounty) BBonchurch本彻奇() feud.Barony {
	return c.本彻奇Bonchurch
}

func (c *怀特WightCounty) BCarisbrooke卡里斯布鲁克() feud.Barony {
	return c.卡里斯布鲁克Carisbrooke
}

func (c *怀特WightCounty) BCowes考斯() feud.Barony {
	return c.考斯Cowes
}

func (c *怀特WightCounty) BSandown桑当() feud.Barony {
	return c.桑当Sandown
}

func (c *怀特WightCounty) BYarmouth雅茅斯() feud.Barony {
	return c.雅茅斯Yarmouth
}

var CWight怀特 WightCounty = &怀特WightCounty{}

func init() {
	f := CWight怀特.(*怀特WightCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1950",
		Title:     "wight",
		TitleName: "怀特",
		TitleCode: "c_wight",
		Baronies:  map[string]feud.Barony{},
	}

	f.本彻奇Bonchurch = BBonchurch本彻奇
	f.本彻奇Bonchurch.SetParent(f)

	f.卡里斯布鲁克Carisbrooke = BCarisbrooke卡里斯布鲁克
	f.卡里斯布鲁克Carisbrooke.SetParent(f)

	f.考斯Cowes = BCowes考斯
	f.考斯Cowes.SetParent(f)

	f.桑当Sandown = BSandown桑当
	f.桑当Sandown.SetParent(f)

	f.雅茅斯Yarmouth = BYarmouth雅茅斯
	f.雅茅斯Yarmouth.SetParent(f)

}
