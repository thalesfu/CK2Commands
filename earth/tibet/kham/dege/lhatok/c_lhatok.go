package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhatokCounty interface {
	feud.County
	BBolo莫洛() feud.Barony
	BBumgye木协() feud.Barony
	BGonjo馆觉() feud.Barony
	BKorra扩达() feud.Barony
	BLhatok拉多() feud.Barony
	BZeba则巴() feud.Barony
	BZhagyab乍丫() feud.Barony
}

type 拉多LhatokCounty struct {
	feud.BaseCounty
	莫洛Bolo    feud.Barony
	木协Bumgye  feud.Barony
	馆觉Gonjo   feud.Barony
	扩达Korra   feud.Barony
	拉多Lhatok  feud.Barony
	则巴Zeba    feud.Barony
	乍丫Zhagyab feud.Barony
}

func (c *拉多LhatokCounty) BBolo莫洛() feud.Barony {
	return c.莫洛Bolo
}

func (c *拉多LhatokCounty) BBumgye木协() feud.Barony {
	return c.木协Bumgye
}

func (c *拉多LhatokCounty) BGonjo馆觉() feud.Barony {
	return c.馆觉Gonjo
}

func (c *拉多LhatokCounty) BKorra扩达() feud.Barony {
	return c.扩达Korra
}

func (c *拉多LhatokCounty) BLhatok拉多() feud.Barony {
	return c.拉多Lhatok
}

func (c *拉多LhatokCounty) BZeba则巴() feud.Barony {
	return c.则巴Zeba
}

func (c *拉多LhatokCounty) BZhagyab乍丫() feud.Barony {
	return c.乍丫Zhagyab
}

var CLhatok拉多 LhatokCounty = &拉多LhatokCounty{}

func init() {
	f := CLhatok拉多.(*拉多LhatokCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1550",
		Title:     "lhatok",
		TitleName: "拉多",
		TitleCode: "c_lhatok",
		Baronies:  map[string]feud.Barony{},
	}

	f.莫洛Bolo = BBolo莫洛
	f.莫洛Bolo.SetParent(f)

	f.木协Bumgye = BBumgye木协
	f.木协Bumgye.SetParent(f)

	f.馆觉Gonjo = BGonjo馆觉
	f.馆觉Gonjo.SetParent(f)

	f.扩达Korra = BKorra扩达
	f.扩达Korra.SetParent(f)

	f.拉多Lhatok = BLhatok拉多
	f.拉多Lhatok.SetParent(f)

	f.则巴Zeba = BZeba则巴
	f.则巴Zeba.SetParent(f)

	f.乍丫Zhagyab = BZhagyab乍丫
	f.乍丫Zhagyab.SetParent(f)

}
