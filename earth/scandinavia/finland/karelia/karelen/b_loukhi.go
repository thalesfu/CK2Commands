package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛乌希LoukhiBarony struct {
	feud.BaseBarony
}

var BLoukhi洛乌希 feud.Barony = &洛乌希LoukhiBarony{}

func init() {
    f := BLoukhi洛乌希.(*洛乌希LoukhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loukhi",
		TitleName: "洛乌希",
		TitleCode: "b_loukhi",
	}
}
