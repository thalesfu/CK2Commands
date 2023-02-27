package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格根GegeenBarony struct {
	feud.BaseBarony
}

var BGegeen格根 feud.Barony = &格根GegeenBarony{}

func init() {
    f := BGegeen格根.(*格根GegeenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gegeen",
		TitleName: "格根",
		TitleCode: "b_gegeen",
	}
}
