package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FarafraCounty interface {
	feud.County
	BFarafra法拉弗拉() feud.Barony
	BMut穆特() feud.Barony
	BQasrfarfra费拉菲拉堡() feud.Barony
}

type c_farafraFarafraCounty struct {
	feud.BaseCounty
	法拉弗拉Farafra     feud.Barony
	穆特Mut           feud.Barony
	费拉菲拉堡Qasrfarfra feud.Barony
}

func (c *c_farafraFarafraCounty) BFarafra法拉弗拉() feud.Barony {
	return c.法拉弗拉Farafra
}

func (c *c_farafraFarafraCounty) BMut穆特() feud.Barony {
	return c.穆特Mut
}

func (c *c_farafraFarafraCounty) BQasrfarfra费拉菲拉堡() feud.Barony {
	return c.费拉菲拉堡Qasrfarfra
}

var CFarafrac_farafra FarafraCounty = &c_farafraFarafraCounty{}

func init() {
	f := CFarafrac_farafra.(*c_farafraFarafraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2007",
		Title:     "farafra",
		TitleName: "c_farafra",
		TitleCode: "c_farafra",
		Baronies:  map[string]feud.Barony{},
	}

	f.法拉弗拉Farafra = BFarafra法拉弗拉
	f.法拉弗拉Farafra.SetParent(f)

	f.穆特Mut = BMut穆特
	f.穆特Mut.SetParent(f)

	f.费拉菲拉堡Qasrfarfra = BQasrfarfra费拉菲拉堡
	f.费拉菲拉堡Qasrfarfra.SetParent(f)

}
