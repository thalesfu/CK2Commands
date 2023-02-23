package norfolk

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/norfolk/norfolk"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/norfolk/suffolk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorfolkDuke interface {
	feud.Duke
	CNorfolk诺福克() norfolk.NorfolkCounty
	CSuffolk萨福克() suffolk.SuffolkCounty
}

type 东盎格利亚NorfolkDuke struct {
	feud.BaseDuke
	诺福克Norfolk norfolk.NorfolkCounty
	萨福克Suffolk suffolk.SuffolkCounty
}

func (d *东盎格利亚NorfolkDuke) CNorfolk诺福克() norfolk.NorfolkCounty {
	return d.诺福克Norfolk
}

func (d *东盎格利亚NorfolkDuke) CSuffolk萨福克() suffolk.SuffolkCounty {
	return d.萨福克Suffolk
}

var DNorfolk东盎格利亚 NorfolkDuke = &东盎格利亚NorfolkDuke{}

func init() {
	f := DNorfolk东盎格利亚.(*东盎格利亚NorfolkDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "norfolk",
		TitleName: "东盎格利亚",
		TitleCode: "d_norfolk",
		Counties:  map[string]feud.County{},
	}

	f.诺福克Norfolk = norfolk.CNorfolk诺福克
	f.诺福克Norfolk.SetParent(f)

	f.萨福克Suffolk = suffolk.CSuffolk萨福克
	f.萨福克Suffolk.SetParent(f)

}
