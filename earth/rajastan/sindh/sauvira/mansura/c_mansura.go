package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MansuraCounty interface {
	feud.County
	BHala诃罗() feud.Barony
	BMansura曼苏拉() feud.Barony
	BMatiari摩提阿利() feud.Barony
	BMirman米尔曼() feud.Barony
	BNasarpur那萨补罗() feud.Barony
	BNerunkot内兰乔特() feud.Barony
}

type 曼苏拉MansuraCounty struct {
	feud.BaseCounty
	诃罗Hala       feud.Barony
	曼苏拉Mansura   feud.Barony
	摩提阿利Matiari  feud.Barony
	米尔曼Mirman    feud.Barony
	那萨补罗Nasarpur feud.Barony
	内兰乔特Nerunkot feud.Barony
}

func (c *曼苏拉MansuraCounty) BHala诃罗() feud.Barony {
	return c.诃罗Hala
}

func (c *曼苏拉MansuraCounty) BMansura曼苏拉() feud.Barony {
	return c.曼苏拉Mansura
}

func (c *曼苏拉MansuraCounty) BMatiari摩提阿利() feud.Barony {
	return c.摩提阿利Matiari
}

func (c *曼苏拉MansuraCounty) BMirman米尔曼() feud.Barony {
	return c.米尔曼Mirman
}

func (c *曼苏拉MansuraCounty) BNasarpur那萨补罗() feud.Barony {
	return c.那萨补罗Nasarpur
}

func (c *曼苏拉MansuraCounty) BNerunkot内兰乔特() feud.Barony {
	return c.内兰乔特Nerunkot
}

var CMansura曼苏拉 MansuraCounty = &曼苏拉MansuraCounty{}

func init() {
	f := CMansura曼苏拉.(*曼苏拉MansuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1137",
		Title:     "mansura",
		TitleName: "曼苏拉",
		TitleCode: "c_mansura",
		Baronies:  map[string]feud.Barony{},
	}

	f.诃罗Hala = BHala诃罗
	f.诃罗Hala.SetParent(f)

	f.曼苏拉Mansura = BMansura曼苏拉
	f.曼苏拉Mansura.SetParent(f)

	f.摩提阿利Matiari = BMatiari摩提阿利
	f.摩提阿利Matiari.SetParent(f)

	f.米尔曼Mirman = BMirman米尔曼
	f.米尔曼Mirman.SetParent(f)

	f.那萨补罗Nasarpur = BNasarpur那萨补罗
	f.那萨补罗Nasarpur.SetParent(f)

	f.内兰乔特Nerunkot = BNerunkot内兰乔特
	f.内兰乔特Nerunkot.SetParent(f)

}
