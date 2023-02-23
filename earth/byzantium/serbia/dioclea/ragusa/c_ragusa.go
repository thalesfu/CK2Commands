package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RagusaCounty interface {
	feud.County
	BCavtat察夫塔特() feud.Barony
	BKolocep科罗切普() feud.Barony
	BMljet姆列特() feud.Barony
	BNarona那罗那() feud.Barony
	BRagusa拉古萨() feud.Barony
	BSipan希潘() feud.Barony
	BSlano斯拉诺() feud.Barony
	BZaton扎通() feud.Barony
}

type 拉古萨RagusaCounty struct {
	feud.BaseCounty
	察夫塔特Cavtat  feud.Barony
	科罗切普Kolocep feud.Barony
	姆列特Mljet    feud.Barony
	那罗那Narona   feud.Barony
	拉古萨Ragusa   feud.Barony
	希潘Sipan     feud.Barony
	斯拉诺Slano    feud.Barony
	扎通Zaton     feud.Barony
}

func (c *拉古萨RagusaCounty) BCavtat察夫塔特() feud.Barony {
	return c.察夫塔特Cavtat
}

func (c *拉古萨RagusaCounty) BKolocep科罗切普() feud.Barony {
	return c.科罗切普Kolocep
}

func (c *拉古萨RagusaCounty) BMljet姆列特() feud.Barony {
	return c.姆列特Mljet
}

func (c *拉古萨RagusaCounty) BNarona那罗那() feud.Barony {
	return c.那罗那Narona
}

func (c *拉古萨RagusaCounty) BRagusa拉古萨() feud.Barony {
	return c.拉古萨Ragusa
}

func (c *拉古萨RagusaCounty) BSipan希潘() feud.Barony {
	return c.希潘Sipan
}

func (c *拉古萨RagusaCounty) BSlano斯拉诺() feud.Barony {
	return c.斯拉诺Slano
}

func (c *拉古萨RagusaCounty) BZaton扎通() feud.Barony {
	return c.扎通Zaton
}

var CRagusa拉古萨 RagusaCounty = &拉古萨RagusaCounty{}

func init() {
	f := CRagusa拉古萨.(*拉古萨RagusaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "468",
		Title:     "ragusa",
		TitleName: "拉古萨",
		TitleCode: "c_ragusa",
		Baronies:  map[string]feud.Barony{},
	}

	f.察夫塔特Cavtat = BCavtat察夫塔特
	f.察夫塔特Cavtat.SetParent(f)

	f.科罗切普Kolocep = BKolocep科罗切普
	f.科罗切普Kolocep.SetParent(f)

	f.姆列特Mljet = BMljet姆列特
	f.姆列特Mljet.SetParent(f)

	f.那罗那Narona = BNarona那罗那
	f.那罗那Narona.SetParent(f)

	f.拉古萨Ragusa = BRagusa拉古萨
	f.拉古萨Ragusa.SetParent(f)

	f.希潘Sipan = BSipan希潘
	f.希潘Sipan.SetParent(f)

	f.斯拉诺Slano = BSlano斯拉诺
	f.斯拉诺Slano.SetParent(f)

	f.扎通Zaton = BZaton扎通
	f.扎通Zaton.SetParent(f)

}
