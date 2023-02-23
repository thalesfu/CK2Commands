package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HastinapuraCounty interface {
	feud.County
	BBaran婆兰() feud.Barony
	BHapur诃补罗() feud.Barony
	BHastinapura象城() feud.Barony
	BIndrasthana因陀罗萨傥那() feud.Barony
	BMirath弥罗多() feud.Barony
	BPandeshwar般滞湿伐罗() feud.Barony
	BPankot般拘吒() feud.Barony
}

type 象城HastinapuraCounty struct {
	feud.BaseCounty
	婆兰Baran           feud.Barony
	诃补罗Hapur          feud.Barony
	象城Hastinapura     feud.Barony
	因陀罗萨傥那Indrasthana feud.Barony
	弥罗多Mirath         feud.Barony
	般滞湿伐罗Pandeshwar   feud.Barony
	般拘吒Pankot         feud.Barony
}

func (c *象城HastinapuraCounty) BBaran婆兰() feud.Barony {
	return c.婆兰Baran
}

func (c *象城HastinapuraCounty) BHapur诃补罗() feud.Barony {
	return c.诃补罗Hapur
}

func (c *象城HastinapuraCounty) BHastinapura象城() feud.Barony {
	return c.象城Hastinapura
}

func (c *象城HastinapuraCounty) BIndrasthana因陀罗萨傥那() feud.Barony {
	return c.因陀罗萨傥那Indrasthana
}

func (c *象城HastinapuraCounty) BMirath弥罗多() feud.Barony {
	return c.弥罗多Mirath
}

func (c *象城HastinapuraCounty) BPandeshwar般滞湿伐罗() feud.Barony {
	return c.般滞湿伐罗Pandeshwar
}

func (c *象城HastinapuraCounty) BPankot般拘吒() feud.Barony {
	return c.般拘吒Pankot
}

var CHastinapura象城 HastinapuraCounty = &象城HastinapuraCounty{}

func init() {
	f := CHastinapura象城.(*象城HastinapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1368",
		Title:     "hastinapura",
		TitleName: "象城",
		TitleCode: "c_hastinapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆兰Baran = BBaran婆兰
	f.婆兰Baran.SetParent(f)

	f.诃补罗Hapur = BHapur诃补罗
	f.诃补罗Hapur.SetParent(f)

	f.象城Hastinapura = BHastinapura象城
	f.象城Hastinapura.SetParent(f)

	f.因陀罗萨傥那Indrasthana = BIndrasthana因陀罗萨傥那
	f.因陀罗萨傥那Indrasthana.SetParent(f)

	f.弥罗多Mirath = BMirath弥罗多
	f.弥罗多Mirath.SetParent(f)

	f.般滞湿伐罗Pandeshwar = BPandeshwar般滞湿伐罗
	f.般滞湿伐罗Pandeshwar.SetParent(f)

	f.般拘吒Pankot = BPankot般拘吒
	f.般拘吒Pankot.SetParent(f)

}
