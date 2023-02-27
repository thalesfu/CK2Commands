package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔泰Altai_khasagtBarony struct {
	feud.BaseBarony
}

var BAltai_khasagt阿尔泰 feud.Barony = &阿尔泰Altai_khasagtBarony{}

func init() {
    f := BAltai_khasagt阿尔泰.(*阿尔泰Altai_khasagtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altai_khasagt",
		TitleName: "阿尔泰",
		TitleCode: "b_altai_khasagt",
	}
}
