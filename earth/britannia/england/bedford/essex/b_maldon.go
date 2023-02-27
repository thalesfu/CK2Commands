package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔登MaldonBarony struct {
	feud.BaseBarony
}

var BMaldon莫尔登 feud.Barony = &莫尔登MaldonBarony{}

func init() {
    f := BMaldon莫尔登.(*莫尔登MaldonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maldon",
		TitleName: "莫尔登",
		TitleCode: "b_maldon",
	}
}
