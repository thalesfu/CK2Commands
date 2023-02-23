package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PecsCounty interface {
	feud.County
	BDarda达尔达() feud.Barony
	BKalocsa考洛乔() feud.Barony
	BMohacs莫哈奇() feud.Barony
	BPecs佩奇() feud.Barony
	BPecsvarad佩奇瓦劳德() feud.Barony
	BSasd沙什德() feud.Barony
	BSiklos希克洛什() feud.Barony
	BSzentlorinc圣勒林茨() feud.Barony
}

type 佩奇PecsCounty struct {
	feud.BaseCounty
	达尔达Darda        feud.Barony
	考洛乔Kalocsa      feud.Barony
	莫哈奇Mohacs       feud.Barony
	佩奇Pecs          feud.Barony
	佩奇瓦劳德Pecsvarad  feud.Barony
	沙什德Sasd         feud.Barony
	希克洛什Siklos      feud.Barony
	圣勒林茨Szentlorinc feud.Barony
}

func (c *佩奇PecsCounty) BDarda达尔达() feud.Barony {
	return c.达尔达Darda
}

func (c *佩奇PecsCounty) BKalocsa考洛乔() feud.Barony {
	return c.考洛乔Kalocsa
}

func (c *佩奇PecsCounty) BMohacs莫哈奇() feud.Barony {
	return c.莫哈奇Mohacs
}

func (c *佩奇PecsCounty) BPecs佩奇() feud.Barony {
	return c.佩奇Pecs
}

func (c *佩奇PecsCounty) BPecsvarad佩奇瓦劳德() feud.Barony {
	return c.佩奇瓦劳德Pecsvarad
}

func (c *佩奇PecsCounty) BSasd沙什德() feud.Barony {
	return c.沙什德Sasd
}

func (c *佩奇PecsCounty) BSiklos希克洛什() feud.Barony {
	return c.希克洛什Siklos
}

func (c *佩奇PecsCounty) BSzentlorinc圣勒林茨() feud.Barony {
	return c.圣勒林茨Szentlorinc
}

var CPecs佩奇 PecsCounty = &佩奇PecsCounty{}

func init() {
	f := CPecs佩奇.(*佩奇PecsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "452",
		Title:     "pecs",
		TitleName: "佩奇",
		TitleCode: "c_pecs",
		Baronies:  map[string]feud.Barony{},
	}

	f.达尔达Darda = BDarda达尔达
	f.达尔达Darda.SetParent(f)

	f.考洛乔Kalocsa = BKalocsa考洛乔
	f.考洛乔Kalocsa.SetParent(f)

	f.莫哈奇Mohacs = BMohacs莫哈奇
	f.莫哈奇Mohacs.SetParent(f)

	f.佩奇Pecs = BPecs佩奇
	f.佩奇Pecs.SetParent(f)

	f.佩奇瓦劳德Pecsvarad = BPecsvarad佩奇瓦劳德
	f.佩奇瓦劳德Pecsvarad.SetParent(f)

	f.沙什德Sasd = BSasd沙什德
	f.沙什德Sasd.SetParent(f)

	f.希克洛什Siklos = BSiklos希克洛什
	f.希克洛什Siklos.SetParent(f)

	f.圣勒林茨Szentlorinc = BSzentlorinc圣勒林茨
	f.圣勒林茨Szentlorinc.SetParent(f)

}
