package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LoriCounty interface {
	feud.County
	BAgarak阿加拉克() feud.Barony
	BDmanis德曼尼斯() feud.Barony
	BDzegh泽格() feud.Barony
	BErkan厄而罕() feud.Barony
	BHovk何福() feud.Barony
	BSanahin萨那欣() feud.Barony
}

type 洛里LoriCounty struct {
	feud.BaseCounty
	阿加拉克Agarak feud.Barony
	德曼尼斯Dmanis feud.Barony
	泽格Dzegh    feud.Barony
	厄而罕Erkan   feud.Barony
	何福Hovk     feud.Barony
	萨那欣Sanahin feud.Barony
}

func (c *洛里LoriCounty) BAgarak阿加拉克() feud.Barony {
	return c.阿加拉克Agarak
}

func (c *洛里LoriCounty) BDmanis德曼尼斯() feud.Barony {
	return c.德曼尼斯Dmanis
}

func (c *洛里LoriCounty) BDzegh泽格() feud.Barony {
	return c.泽格Dzegh
}

func (c *洛里LoriCounty) BErkan厄而罕() feud.Barony {
	return c.厄而罕Erkan
}

func (c *洛里LoriCounty) BHovk何福() feud.Barony {
	return c.何福Hovk
}

func (c *洛里LoriCounty) BSanahin萨那欣() feud.Barony {
	return c.萨那欣Sanahin
}

var CLori洛里 LoriCounty = &洛里LoriCounty{}

func init() {
	f := CLori洛里.(*洛里LoriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1013",
		Title:     "lori",
		TitleName: "洛里",
		TitleCode: "c_lori",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿加拉克Agarak = BAgarak阿加拉克
	f.阿加拉克Agarak.SetParent(f)

	f.德曼尼斯Dmanis = BDmanis德曼尼斯
	f.德曼尼斯Dmanis.SetParent(f)

	f.泽格Dzegh = BDzegh泽格
	f.泽格Dzegh.SetParent(f)

	f.厄而罕Erkan = BErkan厄而罕
	f.厄而罕Erkan.SetParent(f)

	f.何福Hovk = BHovk何福
	f.何福Hovk.SetParent(f)

	f.萨那欣Sanahin = BSanahin萨那欣
	f.萨那欣Sanahin.SetParent(f)

}
