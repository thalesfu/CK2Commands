package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰于西内IvanovskayaBarony struct {
	feud.BaseBarony
}

var BIvanovskaya泰于西内 feud.Barony = &泰于西内IvanovskayaBarony{}

func init() {
    f := BIvanovskaya泰于西内.(*泰于西内IvanovskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivanovskaya",
		TitleName: "泰于西内",
		TitleCode: "b_ivanovskaya",
	}
}
