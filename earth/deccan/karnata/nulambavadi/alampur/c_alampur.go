package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlampurCounty interface {
	feud.County
	BAdavani阿达瓦尼() feud.Barony
	BAlampur阿蓝补罗() feud.Barony
	BGooty古蒂() feud.Barony
	BHampi汉比() feud.Barony
	BIttagi因塔戈() feud.Barony
	BKampili剑毕离() feud.Barony
	BVijayanagara毗阇耶那揭罗() feud.Barony
}

type 阿蓝补罗AlampurCounty struct {
	feud.BaseCounty
	阿达瓦尼Adavani        feud.Barony
	阿蓝补罗Alampur        feud.Barony
	古蒂Gooty            feud.Barony
	汉比Hampi            feud.Barony
	因塔戈Ittagi          feud.Barony
	剑毕离Kampili         feud.Barony
	毗阇耶那揭罗Vijayanagara feud.Barony
}

func (c *阿蓝补罗AlampurCounty) BAdavani阿达瓦尼() feud.Barony {
	return c.阿达瓦尼Adavani
}

func (c *阿蓝补罗AlampurCounty) BAlampur阿蓝补罗() feud.Barony {
	return c.阿蓝补罗Alampur
}

func (c *阿蓝补罗AlampurCounty) BGooty古蒂() feud.Barony {
	return c.古蒂Gooty
}

func (c *阿蓝补罗AlampurCounty) BHampi汉比() feud.Barony {
	return c.汉比Hampi
}

func (c *阿蓝补罗AlampurCounty) BIttagi因塔戈() feud.Barony {
	return c.因塔戈Ittagi
}

func (c *阿蓝补罗AlampurCounty) BKampili剑毕离() feud.Barony {
	return c.剑毕离Kampili
}

func (c *阿蓝补罗AlampurCounty) BVijayanagara毗阇耶那揭罗() feud.Barony {
	return c.毗阇耶那揭罗Vijayanagara
}

var CAlampur阿蓝补罗 AlampurCounty = &阿蓝补罗AlampurCounty{}

func init() {
	f := CAlampur阿蓝补罗.(*阿蓝补罗AlampurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1218",
		Title:     "alampur",
		TitleName: "阿蓝补罗",
		TitleCode: "c_alampur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿达瓦尼Adavani = BAdavani阿达瓦尼
	f.阿达瓦尼Adavani.SetParent(f)

	f.阿蓝补罗Alampur = BAlampur阿蓝补罗
	f.阿蓝补罗Alampur.SetParent(f)

	f.古蒂Gooty = BGooty古蒂
	f.古蒂Gooty.SetParent(f)

	f.汉比Hampi = BHampi汉比
	f.汉比Hampi.SetParent(f)

	f.因塔戈Ittagi = BIttagi因塔戈
	f.因塔戈Ittagi.SetParent(f)

	f.剑毕离Kampili = BKampili剑毕离
	f.剑毕离Kampili.SetParent(f)

	f.毗阇耶那揭罗Vijayanagara = BVijayanagara毗阇耶那揭罗
	f.毗阇耶那揭罗Vijayanagara.SetParent(f)

}
