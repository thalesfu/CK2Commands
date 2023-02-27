package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古林Guulin_khasagtBarony struct {
	feud.BaseBarony
}

var BGuulin_khasagt古林 feud.Barony = &古林Guulin_khasagtBarony{}

func init() {
    f := BGuulin_khasagt古林.(*古林Guulin_khasagtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guulin_khasagt",
		TitleName: "古林",
		TitleCode: "b_guulin_khasagt",
	}
}
