package trans-portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博泽尔斯基ObozerskiyBarony struct {
	feud.BaseBarony
}

var BObozerskiy奥博泽尔斯基 feud.Barony = &奥博泽尔斯基ObozerskiyBarony{}

func init() {
    f := BObozerskiy奥博泽尔斯基.(*奥博泽尔斯基ObozerskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obozerskiy",
		TitleName: "奥博泽尔斯基",
		TitleCode: "b_obozerskiy",
	}
}
