package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DjenneCounty interface {
	feud.County
	BBandiagara邦贾加拉() feud.Barony
	BBindugu班杜古() feud.Barony
	BDerary德拉里() feud.Barony
	BDjenne杰内() feud.Barony
	BFemaye费马耶() feud.Barony
	BKorienza科里安扎() feud.Barony
	BMopti莫普提() feud.Barony
	BPondori旁多里() feud.Barony
}

type 杰内DjenneCounty struct {
	feud.BaseCounty
	邦贾加拉Bandiagara feud.Barony
	班杜古Bindugu     feud.Barony
	德拉里Derary      feud.Barony
	杰内Djenne       feud.Barony
	费马耶Femaye      feud.Barony
	科里安扎Korienza   feud.Barony
	莫普提Mopti       feud.Barony
	旁多里Pondori     feud.Barony
}

func (c *杰内DjenneCounty) BBandiagara邦贾加拉() feud.Barony {
	return c.邦贾加拉Bandiagara
}

func (c *杰内DjenneCounty) BBindugu班杜古() feud.Barony {
	return c.班杜古Bindugu
}

func (c *杰内DjenneCounty) BDerary德拉里() feud.Barony {
	return c.德拉里Derary
}

func (c *杰内DjenneCounty) BDjenne杰内() feud.Barony {
	return c.杰内Djenne
}

func (c *杰内DjenneCounty) BFemaye费马耶() feud.Barony {
	return c.费马耶Femaye
}

func (c *杰内DjenneCounty) BKorienza科里安扎() feud.Barony {
	return c.科里安扎Korienza
}

func (c *杰内DjenneCounty) BMopti莫普提() feud.Barony {
	return c.莫普提Mopti
}

func (c *杰内DjenneCounty) BPondori旁多里() feud.Barony {
	return c.旁多里Pondori
}

var CDjenne杰内 DjenneCounty = &杰内DjenneCounty{}

func init() {
	f := CDjenne杰内.(*杰内DjenneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "915",
		Title:     "djenne",
		TitleName: "杰内",
		TitleCode: "c_djenne",
		Baronies:  map[string]feud.Barony{},
	}

	f.邦贾加拉Bandiagara = BBandiagara邦贾加拉
	f.邦贾加拉Bandiagara.SetParent(f)

	f.班杜古Bindugu = BBindugu班杜古
	f.班杜古Bindugu.SetParent(f)

	f.德拉里Derary = BDerary德拉里
	f.德拉里Derary.SetParent(f)

	f.杰内Djenne = BDjenne杰内
	f.杰内Djenne.SetParent(f)

	f.费马耶Femaye = BFemaye费马耶
	f.费马耶Femaye.SetParent(f)

	f.科里安扎Korienza = BKorienza科里安扎
	f.科里安扎Korienza.SetParent(f)

	f.莫普提Mopti = BMopti莫普提
	f.莫普提Mopti.SetParent(f)

	f.旁多里Pondori = BPondori旁多里
	f.旁多里Pondori.SetParent(f)

}
