package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AulonCounty interface {
	feud.County
	BAulon阿夫洛纳斯() feud.Barony
	BHimara希马拉() feud.Barony
	BKanine卡尼内() feud.Barony
	BSaranda萨兰达() feud.Barony
	BSpinarizza斯皮纳里扎() feud.Barony
}

type 阿夫洛纳斯AulonCounty struct {
	feud.BaseCounty
	阿夫洛纳斯Aulon      feud.Barony
	希马拉Himara       feud.Barony
	卡尼内Kanine       feud.Barony
	萨兰达Saranda      feud.Barony
	斯皮纳里扎Spinarizza feud.Barony
}

func (c *阿夫洛纳斯AulonCounty) BAulon阿夫洛纳斯() feud.Barony {
	return c.阿夫洛纳斯Aulon
}

func (c *阿夫洛纳斯AulonCounty) BHimara希马拉() feud.Barony {
	return c.希马拉Himara
}

func (c *阿夫洛纳斯AulonCounty) BKanine卡尼内() feud.Barony {
	return c.卡尼内Kanine
}

func (c *阿夫洛纳斯AulonCounty) BSaranda萨兰达() feud.Barony {
	return c.萨兰达Saranda
}

func (c *阿夫洛纳斯AulonCounty) BSpinarizza斯皮纳里扎() feud.Barony {
	return c.斯皮纳里扎Spinarizza
}

var CAulon阿夫洛纳斯 AulonCounty = &阿夫洛纳斯AulonCounty{}

func init() {
	f := CAulon阿夫洛纳斯.(*阿夫洛纳斯AulonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1886",
		Title:     "aulon",
		TitleName: "阿夫洛纳斯",
		TitleCode: "c_aulon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫洛纳斯Aulon = BAulon阿夫洛纳斯
	f.阿夫洛纳斯Aulon.SetParent(f)

	f.希马拉Himara = BHimara希马拉
	f.希马拉Himara.SetParent(f)

	f.卡尼内Kanine = BKanine卡尼内
	f.卡尼内Kanine.SetParent(f)

	f.萨兰达Saranda = BSaranda萨兰达
	f.萨兰达Saranda.SetParent(f)

	f.斯皮纳里扎Spinarizza = BSpinarizza斯皮纳里扎
	f.斯皮纳里扎Spinarizza.SetParent(f)

}
