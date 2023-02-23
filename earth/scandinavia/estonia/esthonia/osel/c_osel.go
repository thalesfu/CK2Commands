package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OselCounty interface {
	feud.County
	BHiiumaa希乌马() feud.Barony
	BKaarma卡尔马() feud.Barony
	BKuressare库雷萨雷() feud.Barony
	BLeisi莱西() feud.Barony
	BMuhu穆胡() feud.Barony
	BOrisaare奥里萨雷() feud.Barony
	BPoide珀伊德() feud.Barony
	BValjala瓦利亚拉() feud.Barony
}

type 厄瑟尔OselCounty struct {
	feud.BaseCounty
	希乌马Hiiumaa    feud.Barony
	卡尔马Kaarma     feud.Barony
	库雷萨雷Kuressare feud.Barony
	莱西Leisi       feud.Barony
	穆胡Muhu        feud.Barony
	奥里萨雷Orisaare  feud.Barony
	珀伊德Poide      feud.Barony
	瓦利亚拉Valjala   feud.Barony
}

func (c *厄瑟尔OselCounty) BHiiumaa希乌马() feud.Barony {
	return c.希乌马Hiiumaa
}

func (c *厄瑟尔OselCounty) BKaarma卡尔马() feud.Barony {
	return c.卡尔马Kaarma
}

func (c *厄瑟尔OselCounty) BKuressare库雷萨雷() feud.Barony {
	return c.库雷萨雷Kuressare
}

func (c *厄瑟尔OselCounty) BLeisi莱西() feud.Barony {
	return c.莱西Leisi
}

func (c *厄瑟尔OselCounty) BMuhu穆胡() feud.Barony {
	return c.穆胡Muhu
}

func (c *厄瑟尔OselCounty) BOrisaare奥里萨雷() feud.Barony {
	return c.奥里萨雷Orisaare
}

func (c *厄瑟尔OselCounty) BPoide珀伊德() feud.Barony {
	return c.珀伊德Poide
}

func (c *厄瑟尔OselCounty) BValjala瓦利亚拉() feud.Barony {
	return c.瓦利亚拉Valjala
}

var COsel厄瑟尔 OselCounty = &厄瑟尔OselCounty{}

func init() {
	f := COsel厄瑟尔.(*厄瑟尔OselCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "376",
		Title:     "osel",
		TitleName: "厄瑟尔",
		TitleCode: "c_osel",
		Baronies:  map[string]feud.Barony{},
	}

	f.希乌马Hiiumaa = BHiiumaa希乌马
	f.希乌马Hiiumaa.SetParent(f)

	f.卡尔马Kaarma = BKaarma卡尔马
	f.卡尔马Kaarma.SetParent(f)

	f.库雷萨雷Kuressare = BKuressare库雷萨雷
	f.库雷萨雷Kuressare.SetParent(f)

	f.莱西Leisi = BLeisi莱西
	f.莱西Leisi.SetParent(f)

	f.穆胡Muhu = BMuhu穆胡
	f.穆胡Muhu.SetParent(f)

	f.奥里萨雷Orisaare = BOrisaare奥里萨雷
	f.奥里萨雷Orisaare.SetParent(f)

	f.珀伊德Poide = BPoide珀伊德
	f.珀伊德Poide.SetParent(f)

	f.瓦利亚拉Valjala = BValjala瓦利亚拉
	f.瓦利亚拉Valjala.SetParent(f)

}
