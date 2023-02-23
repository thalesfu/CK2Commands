package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MassawaCounty interface {
	feud.County
	BAdulis阿杜利斯() feud.Barony
	BCheren克伦() feud.Barony
	BDahano达哈努() feud.Barony
	BMassawa马萨瓦() feud.Barony
	BMatara马塔拉() feud.Barony
	BQohaito科哈依托() feud.Barony
	BSembel瑟姆贝() feud.Barony
}

type 马萨瓦MassawaCounty struct {
	feud.BaseCounty
	阿杜利斯Adulis  feud.Barony
	克伦Cheren    feud.Barony
	达哈努Dahano   feud.Barony
	马萨瓦Massawa  feud.Barony
	马塔拉Matara   feud.Barony
	科哈依托Qohaito feud.Barony
	瑟姆贝Sembel   feud.Barony
}

func (c *马萨瓦MassawaCounty) BAdulis阿杜利斯() feud.Barony {
	return c.阿杜利斯Adulis
}

func (c *马萨瓦MassawaCounty) BCheren克伦() feud.Barony {
	return c.克伦Cheren
}

func (c *马萨瓦MassawaCounty) BDahano达哈努() feud.Barony {
	return c.达哈努Dahano
}

func (c *马萨瓦MassawaCounty) BMassawa马萨瓦() feud.Barony {
	return c.马萨瓦Massawa
}

func (c *马萨瓦MassawaCounty) BMatara马塔拉() feud.Barony {
	return c.马塔拉Matara
}

func (c *马萨瓦MassawaCounty) BQohaito科哈依托() feud.Barony {
	return c.科哈依托Qohaito
}

func (c *马萨瓦MassawaCounty) BSembel瑟姆贝() feud.Barony {
	return c.瑟姆贝Sembel
}

var CMassawa马萨瓦 MassawaCounty = &马萨瓦MassawaCounty{}

func init() {
	f := CMassawa马萨瓦.(*马萨瓦MassawaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1343",
		Title:     "massawa",
		TitleName: "马萨瓦",
		TitleCode: "c_massawa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿杜利斯Adulis = BAdulis阿杜利斯
	f.阿杜利斯Adulis.SetParent(f)

	f.克伦Cheren = BCheren克伦
	f.克伦Cheren.SetParent(f)

	f.达哈努Dahano = BDahano达哈努
	f.达哈努Dahano.SetParent(f)

	f.马萨瓦Massawa = BMassawa马萨瓦
	f.马萨瓦Massawa.SetParent(f)

	f.马塔拉Matara = BMatara马塔拉
	f.马塔拉Matara.SetParent(f)

	f.科哈依托Qohaito = BQohaito科哈依托
	f.科哈依托Qohaito.SetParent(f)

	f.瑟姆贝Sembel = BSembel瑟姆贝
	f.瑟姆贝Sembel.SetParent(f)

}
