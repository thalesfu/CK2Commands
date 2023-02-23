package dublin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DublinCounty interface {
	feud.County
	BClondalkin克朗多金() feud.Barony
	BDublin都柏林() feud.Barony
	BFinglas芬格拉斯() feud.Barony
	BMellifont梅利丰特() feud.Barony
	BTrim特里姆() feud.Barony
	BWicklow威克洛() feud.Barony
}

type 都柏林DublinCounty struct {
	feud.BaseCounty
	克朗多金Clondalkin feud.Barony
	都柏林Dublin      feud.Barony
	芬格拉斯Finglas    feud.Barony
	梅利丰特Mellifont  feud.Barony
	特里姆Trim        feud.Barony
	威克洛Wicklow     feud.Barony
}

func (c *都柏林DublinCounty) BClondalkin克朗多金() feud.Barony {
	return c.克朗多金Clondalkin
}

func (c *都柏林DublinCounty) BDublin都柏林() feud.Barony {
	return c.都柏林Dublin
}

func (c *都柏林DublinCounty) BFinglas芬格拉斯() feud.Barony {
	return c.芬格拉斯Finglas
}

func (c *都柏林DublinCounty) BMellifont梅利丰特() feud.Barony {
	return c.梅利丰特Mellifont
}

func (c *都柏林DublinCounty) BTrim特里姆() feud.Barony {
	return c.特里姆Trim
}

func (c *都柏林DublinCounty) BWicklow威克洛() feud.Barony {
	return c.威克洛Wicklow
}

var CDublin都柏林 DublinCounty = &都柏林DublinCounty{}

func init() {
	f := CDublin都柏林.(*都柏林DublinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "11",
		Title:     "dublin",
		TitleName: "都柏林",
		TitleCode: "c_dublin",
		Baronies:  map[string]feud.Barony{},
	}

	f.克朗多金Clondalkin = BClondalkin克朗多金
	f.克朗多金Clondalkin.SetParent(f)

	f.都柏林Dublin = BDublin都柏林
	f.都柏林Dublin.SetParent(f)

	f.芬格拉斯Finglas = BFinglas芬格拉斯
	f.芬格拉斯Finglas.SetParent(f)

	f.梅利丰特Mellifont = BMellifont梅利丰特
	f.梅利丰特Mellifont.SetParent(f)

	f.特里姆Trim = BTrim特里姆
	f.特里姆Trim.SetParent(f)

	f.威克洛Wicklow = BWicklow威克洛
	f.威克洛Wicklow.SetParent(f)

}
