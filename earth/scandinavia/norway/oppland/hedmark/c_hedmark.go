package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HedmarkCounty interface {
	feud.County
	BEidskog埃兹库格() feud.Barony
	BElverum埃尔沃吕姆() feud.Barony
	BHamar哈马尔() feud.Barony
	BHamarhus哈马尔胡斯() feud.Barony
	BKongsvinger孔斯温厄尔() feud.Barony
	BLoten勒滕() feud.Barony
	BStange斯唐厄() feud.Barony
	BVang旺() feud.Barony
}

type 厄斯特谷地HedmarkCounty struct {
	feud.BaseCounty
	埃兹库格Eidskog      feud.Barony
	埃尔沃吕姆Elverum     feud.Barony
	哈马尔Hamar         feud.Barony
	哈马尔胡斯Hamarhus    feud.Barony
	孔斯温厄尔Kongsvinger feud.Barony
	勒滕Loten          feud.Barony
	斯唐厄Stange        feud.Barony
	旺Vang            feud.Barony
}

func (c *厄斯特谷地HedmarkCounty) BEidskog埃兹库格() feud.Barony {
	return c.埃兹库格Eidskog
}

func (c *厄斯特谷地HedmarkCounty) BElverum埃尔沃吕姆() feud.Barony {
	return c.埃尔沃吕姆Elverum
}

func (c *厄斯特谷地HedmarkCounty) BHamar哈马尔() feud.Barony {
	return c.哈马尔Hamar
}

func (c *厄斯特谷地HedmarkCounty) BHamarhus哈马尔胡斯() feud.Barony {
	return c.哈马尔胡斯Hamarhus
}

func (c *厄斯特谷地HedmarkCounty) BKongsvinger孔斯温厄尔() feud.Barony {
	return c.孔斯温厄尔Kongsvinger
}

func (c *厄斯特谷地HedmarkCounty) BLoten勒滕() feud.Barony {
	return c.勒滕Loten
}

func (c *厄斯特谷地HedmarkCounty) BStange斯唐厄() feud.Barony {
	return c.斯唐厄Stange
}

func (c *厄斯特谷地HedmarkCounty) BVang旺() feud.Barony {
	return c.旺Vang
}

var CHedmark厄斯特谷地 HedmarkCounty = &厄斯特谷地HedmarkCounty{}

func init() {
	f := CHedmark厄斯特谷地.(*厄斯特谷地HedmarkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "276",
		Title:     "hedmark",
		TitleName: "厄斯特谷地",
		TitleCode: "c_hedmark",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃兹库格Eidskog = BEidskog埃兹库格
	f.埃兹库格Eidskog.SetParent(f)

	f.埃尔沃吕姆Elverum = BElverum埃尔沃吕姆
	f.埃尔沃吕姆Elverum.SetParent(f)

	f.哈马尔Hamar = BHamar哈马尔
	f.哈马尔Hamar.SetParent(f)

	f.哈马尔胡斯Hamarhus = BHamarhus哈马尔胡斯
	f.哈马尔胡斯Hamarhus.SetParent(f)

	f.孔斯温厄尔Kongsvinger = BKongsvinger孔斯温厄尔
	f.孔斯温厄尔Kongsvinger.SetParent(f)

	f.勒滕Loten = BLoten勒滕
	f.勒滕Loten.SetParent(f)

	f.斯唐厄Stange = BStange斯唐厄
	f.斯唐厄Stange.SetParent(f)

	f.旺Vang = BVang旺
	f.旺Vang.SetParent(f)

}
