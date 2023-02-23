package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HallandCounty interface {
	feud.County
	BAranaes阿拉内斯() feud.Barony
	BBaastad博斯塔德() feud.Barony
	BFalkenberg法尔肯贝里() feud.Barony
	BGetinge耶廷厄() feud.Barony
	BHalmstad哈尔姆斯塔德() feud.Barony
	BKungsbacka孔斯巴卡() feud.Barony
	BLaholm拉霍尔姆() feud.Barony
	BVarberg瓦尔贝里() feud.Barony
}

type 哈兰HallandCounty struct {
	feud.BaseCounty
	阿拉内斯Aranaes     feud.Barony
	博斯塔德Baastad     feud.Barony
	法尔肯贝里Falkenberg feud.Barony
	耶廷厄Getinge      feud.Barony
	哈尔姆斯塔德Halmstad  feud.Barony
	孔斯巴卡Kungsbacka  feud.Barony
	拉霍尔姆Laholm      feud.Barony
	瓦尔贝里Varberg     feud.Barony
}

func (c *哈兰HallandCounty) BAranaes阿拉内斯() feud.Barony {
	return c.阿拉内斯Aranaes
}

func (c *哈兰HallandCounty) BBaastad博斯塔德() feud.Barony {
	return c.博斯塔德Baastad
}

func (c *哈兰HallandCounty) BFalkenberg法尔肯贝里() feud.Barony {
	return c.法尔肯贝里Falkenberg
}

func (c *哈兰HallandCounty) BGetinge耶廷厄() feud.Barony {
	return c.耶廷厄Getinge
}

func (c *哈兰HallandCounty) BHalmstad哈尔姆斯塔德() feud.Barony {
	return c.哈尔姆斯塔德Halmstad
}

func (c *哈兰HallandCounty) BKungsbacka孔斯巴卡() feud.Barony {
	return c.孔斯巴卡Kungsbacka
}

func (c *哈兰HallandCounty) BLaholm拉霍尔姆() feud.Barony {
	return c.拉霍尔姆Laholm
}

func (c *哈兰HallandCounty) BVarberg瓦尔贝里() feud.Barony {
	return c.瓦尔贝里Varberg
}

var CHalland哈兰 HallandCounty = &哈兰HallandCounty{}

func init() {
	f := CHalland哈兰.(*哈兰HallandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "302",
		Title:     "halland",
		TitleName: "哈兰",
		TitleCode: "c_halland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉内斯Aranaes = BAranaes阿拉内斯
	f.阿拉内斯Aranaes.SetParent(f)

	f.博斯塔德Baastad = BBaastad博斯塔德
	f.博斯塔德Baastad.SetParent(f)

	f.法尔肯贝里Falkenberg = BFalkenberg法尔肯贝里
	f.法尔肯贝里Falkenberg.SetParent(f)

	f.耶廷厄Getinge = BGetinge耶廷厄
	f.耶廷厄Getinge.SetParent(f)

	f.哈尔姆斯塔德Halmstad = BHalmstad哈尔姆斯塔德
	f.哈尔姆斯塔德Halmstad.SetParent(f)

	f.孔斯巴卡Kungsbacka = BKungsbacka孔斯巴卡
	f.孔斯巴卡Kungsbacka.SetParent(f)

	f.拉霍尔姆Laholm = BLaholm拉霍尔姆
	f.拉霍尔姆Laholm.SetParent(f)

	f.瓦尔贝里Varberg = BVarberg瓦尔贝里
	f.瓦尔贝里Varberg.SetParent(f)

}
