package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ManyakhetaCounty interface {
	feud.County
	BFiruzabad菲鲁扎巴迪() feud.Barony
	BGulbarga崛跋伽() feud.Barony
	BKovilkonda拘毗罗军荼() feud.Barony
	BManyakheta摩若契吒() feud.Barony
	BSaongi娑翁耆() feud.Barony
	BSedam塞达姆() feud.Barony
	BYadavagiri耶陀婆耆厘() feud.Barony
}

type 摩若契吒ManyakhetaCounty struct {
	feud.BaseCounty
	菲鲁扎巴迪Firuzabad  feud.Barony
	崛跋伽Gulbarga     feud.Barony
	拘毗罗军荼Kovilkonda feud.Barony
	摩若契吒Manyakheta  feud.Barony
	娑翁耆Saongi       feud.Barony
	塞达姆Sedam        feud.Barony
	耶陀婆耆厘Yadavagiri feud.Barony
}

func (c *摩若契吒ManyakhetaCounty) BFiruzabad菲鲁扎巴迪() feud.Barony {
	return c.菲鲁扎巴迪Firuzabad
}

func (c *摩若契吒ManyakhetaCounty) BGulbarga崛跋伽() feud.Barony {
	return c.崛跋伽Gulbarga
}

func (c *摩若契吒ManyakhetaCounty) BKovilkonda拘毗罗军荼() feud.Barony {
	return c.拘毗罗军荼Kovilkonda
}

func (c *摩若契吒ManyakhetaCounty) BManyakheta摩若契吒() feud.Barony {
	return c.摩若契吒Manyakheta
}

func (c *摩若契吒ManyakhetaCounty) BSaongi娑翁耆() feud.Barony {
	return c.娑翁耆Saongi
}

func (c *摩若契吒ManyakhetaCounty) BSedam塞达姆() feud.Barony {
	return c.塞达姆Sedam
}

func (c *摩若契吒ManyakhetaCounty) BYadavagiri耶陀婆耆厘() feud.Barony {
	return c.耶陀婆耆厘Yadavagiri
}

var CManyakheta摩若契吒 ManyakhetaCounty = &摩若契吒ManyakhetaCounty{}

func init() {
	f := CManyakheta摩若契吒.(*摩若契吒ManyakhetaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1210",
		Title:     "manyakheta",
		TitleName: "摩若契吒",
		TitleCode: "c_manyakheta",
		Baronies:  map[string]feud.Barony{},
	}

	f.菲鲁扎巴迪Firuzabad = BFiruzabad菲鲁扎巴迪
	f.菲鲁扎巴迪Firuzabad.SetParent(f)

	f.崛跋伽Gulbarga = BGulbarga崛跋伽
	f.崛跋伽Gulbarga.SetParent(f)

	f.拘毗罗军荼Kovilkonda = BKovilkonda拘毗罗军荼
	f.拘毗罗军荼Kovilkonda.SetParent(f)

	f.摩若契吒Manyakheta = BManyakheta摩若契吒
	f.摩若契吒Manyakheta.SetParent(f)

	f.娑翁耆Saongi = BSaongi娑翁耆
	f.娑翁耆Saongi.SetParent(f)

	f.塞达姆Sedam = BSedam塞达姆
	f.塞达姆Sedam.SetParent(f)

	f.耶陀婆耆厘Yadavagiri = BYadavagiri耶陀婆耆厘
	f.耶陀婆耆厘Yadavagiri.SetParent(f)

}
