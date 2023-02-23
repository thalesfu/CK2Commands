package saargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaargauCounty interface {
	feud.County
	BDudweiler杜德韦勒() feud.Barony
	BEschringen埃施林根() feud.Barony
	BFechingen费兴根() feud.Barony
	BMalstatt马尔施塔特() feud.Barony
	BMerkingen梅尔克林根() feud.Barony
	BSaarbrucken萨尔布吕肯() feud.Barony
}

type 萨尔布吕肯SaargauCounty struct {
	feud.BaseCounty
	杜德韦勒Dudweiler    feud.Barony
	埃施林根Eschringen   feud.Barony
	费兴根Fechingen     feud.Barony
	马尔施塔特Malstatt    feud.Barony
	梅尔克林根Merkingen   feud.Barony
	萨尔布吕肯Saarbrucken feud.Barony
}

func (c *萨尔布吕肯SaargauCounty) BDudweiler杜德韦勒() feud.Barony {
	return c.杜德韦勒Dudweiler
}

func (c *萨尔布吕肯SaargauCounty) BEschringen埃施林根() feud.Barony {
	return c.埃施林根Eschringen
}

func (c *萨尔布吕肯SaargauCounty) BFechingen费兴根() feud.Barony {
	return c.费兴根Fechingen
}

func (c *萨尔布吕肯SaargauCounty) BMalstatt马尔施塔特() feud.Barony {
	return c.马尔施塔特Malstatt
}

func (c *萨尔布吕肯SaargauCounty) BMerkingen梅尔克林根() feud.Barony {
	return c.梅尔克林根Merkingen
}

func (c *萨尔布吕肯SaargauCounty) BSaarbrucken萨尔布吕肯() feud.Barony {
	return c.萨尔布吕肯Saarbrucken
}

var CSaargau萨尔布吕肯 SaargauCounty = &萨尔布吕肯SaargauCounty{}

func init() {
	f := CSaargau萨尔布吕肯.(*萨尔布吕肯SaargauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1986",
		Title:     "saargau",
		TitleName: "萨尔布吕肯",
		TitleCode: "c_saargau",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜德韦勒Dudweiler = BDudweiler杜德韦勒
	f.杜德韦勒Dudweiler.SetParent(f)

	f.埃施林根Eschringen = BEschringen埃施林根
	f.埃施林根Eschringen.SetParent(f)

	f.费兴根Fechingen = BFechingen费兴根
	f.费兴根Fechingen.SetParent(f)

	f.马尔施塔特Malstatt = BMalstatt马尔施塔特
	f.马尔施塔特Malstatt.SetParent(f)

	f.梅尔克林根Merkingen = BMerkingen梅尔克林根
	f.梅尔克林根Merkingen.SetParent(f)

	f.萨尔布吕肯Saarbrucken = BSaarbrucken萨尔布吕肯
	f.萨尔布吕肯Saarbrucken.SetParent(f)

}
