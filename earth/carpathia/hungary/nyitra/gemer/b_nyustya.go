package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽什焦NyustyaBarony struct {
	feud.BaseBarony
}

var BNyustya纽什焦 feud.Barony = &纽什焦NyustyaBarony{}

func init() {
    f := BNyustya纽什焦.(*纽什焦NyustyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyustya",
		TitleName: "纽什焦",
		TitleCode: "b_nyustya",
	}
}
