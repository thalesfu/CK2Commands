package builth

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BuilthCounty interface {
	feud.County
	BAberedw阿伯雷杜() feud.Barony
	BAbergwesyn阿伯圭辛() feud.Barony
	BBuilth比尔斯() feud.Barony
	BLlandrindod兰德林多德() feud.Barony
	BLlangammarch兰加马赫() feud.Barony
	BLlanwrtyd拉努蒂德() feud.Barony
}

type 比尔斯BuilthCounty struct {
	feud.BaseCounty
	阿伯雷杜Aberedw      feud.Barony
	阿伯圭辛Abergwesyn   feud.Barony
	比尔斯Builth        feud.Barony
	兰德林多德Llandrindod feud.Barony
	兰加马赫Llangammarch feud.Barony
	拉努蒂德Llanwrtyd    feud.Barony
}

func (c *比尔斯BuilthCounty) BAberedw阿伯雷杜() feud.Barony {
	return c.阿伯雷杜Aberedw
}

func (c *比尔斯BuilthCounty) BAbergwesyn阿伯圭辛() feud.Barony {
	return c.阿伯圭辛Abergwesyn
}

func (c *比尔斯BuilthCounty) BBuilth比尔斯() feud.Barony {
	return c.比尔斯Builth
}

func (c *比尔斯BuilthCounty) BLlandrindod兰德林多德() feud.Barony {
	return c.兰德林多德Llandrindod
}

func (c *比尔斯BuilthCounty) BLlangammarch兰加马赫() feud.Barony {
	return c.兰加马赫Llangammarch
}

func (c *比尔斯BuilthCounty) BLlanwrtyd拉努蒂德() feud.Barony {
	return c.拉努蒂德Llanwrtyd
}

var CBuilth比尔斯 BuilthCounty = &比尔斯BuilthCounty{}

func init() {
	f := CBuilth比尔斯.(*比尔斯BuilthCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1947",
		Title:     "builth",
		TitleName: "比尔斯",
		TitleCode: "c_builth",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯雷杜Aberedw = BAberedw阿伯雷杜
	f.阿伯雷杜Aberedw.SetParent(f)

	f.阿伯圭辛Abergwesyn = BAbergwesyn阿伯圭辛
	f.阿伯圭辛Abergwesyn.SetParent(f)

	f.比尔斯Builth = BBuilth比尔斯
	f.比尔斯Builth.SetParent(f)

	f.兰德林多德Llandrindod = BLlandrindod兰德林多德
	f.兰德林多德Llandrindod.SetParent(f)

	f.兰加马赫Llangammarch = BLlangammarch兰加马赫
	f.兰加马赫Llangammarch.SetParent(f)

	f.拉努蒂德Llanwrtyd = BLlanwrtyd拉努蒂德
	f.拉努蒂德Llanwrtyd.SetParent(f)

}
