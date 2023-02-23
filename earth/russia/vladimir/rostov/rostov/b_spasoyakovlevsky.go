package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯帕索雅科夫列夫斯基SpasoyakovlevskyBarony struct {
	feud.BaseBarony
}

var BSpasoyakovlevsky斯帕索雅科夫列夫斯基 feud.Barony = &斯帕索雅科夫列夫斯基SpasoyakovlevskyBarony{}

func init() {
	f := BSpasoyakovlevsky斯帕索雅科夫列夫斯基.(*斯帕索雅科夫列夫斯基SpasoyakovlevskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spasoyakovlevsky",
		TitleName: "斯帕索雅科夫列夫斯基",
		TitleCode: "b_spasoyakovlevsky",
	}
}
