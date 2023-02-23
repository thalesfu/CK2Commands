package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PuriCounty interface {
	feud.County
	BAska阿斯加() feud.Barony
	BBuguda菩瞿陀() feud.Barony
	BGanjam甘贾姆() feud.Barony
	BGopalpur瞿波罗补罗() feud.Barony
	BJaugada贾戈达() feud.Barony
	BPuri补梨() feud.Barony
	BTaratarini塔拉塔里尼() feud.Barony
}

type 补梨PuriCounty struct {
	feud.BaseCounty
	阿斯加Aska         feud.Barony
	菩瞿陀Buguda       feud.Barony
	甘贾姆Ganjam       feud.Barony
	瞿波罗补罗Gopalpur   feud.Barony
	贾戈达Jaugada      feud.Barony
	补梨Puri          feud.Barony
	塔拉塔里尼Taratarini feud.Barony
}

func (c *补梨PuriCounty) BAska阿斯加() feud.Barony {
	return c.阿斯加Aska
}

func (c *补梨PuriCounty) BBuguda菩瞿陀() feud.Barony {
	return c.菩瞿陀Buguda
}

func (c *补梨PuriCounty) BGanjam甘贾姆() feud.Barony {
	return c.甘贾姆Ganjam
}

func (c *补梨PuriCounty) BGopalpur瞿波罗补罗() feud.Barony {
	return c.瞿波罗补罗Gopalpur
}

func (c *补梨PuriCounty) BJaugada贾戈达() feud.Barony {
	return c.贾戈达Jaugada
}

func (c *补梨PuriCounty) BPuri补梨() feud.Barony {
	return c.补梨Puri
}

func (c *补梨PuriCounty) BTaratarini塔拉塔里尼() feud.Barony {
	return c.塔拉塔里尼Taratarini
}

var CPuri补梨 PuriCounty = &补梨PuriCounty{}

func init() {
	f := CPuri补梨.(*补梨PuriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1225",
		Title:     "puri",
		TitleName: "补梨",
		TitleCode: "c_puri",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯加Aska = BAska阿斯加
	f.阿斯加Aska.SetParent(f)

	f.菩瞿陀Buguda = BBuguda菩瞿陀
	f.菩瞿陀Buguda.SetParent(f)

	f.甘贾姆Ganjam = BGanjam甘贾姆
	f.甘贾姆Ganjam.SetParent(f)

	f.瞿波罗补罗Gopalpur = BGopalpur瞿波罗补罗
	f.瞿波罗补罗Gopalpur.SetParent(f)

	f.贾戈达Jaugada = BJaugada贾戈达
	f.贾戈达Jaugada.SetParent(f)

	f.补梨Puri = BPuri补梨
	f.补梨Puri.SetParent(f)

	f.塔拉塔里尼Taratarini = BTaratarini塔拉塔里尼
	f.塔拉塔里尼Taratarini.SetParent(f)

}
