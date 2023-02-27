package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀迦德TekadeeBarony struct {
	feud.BaseBarony
}

var BTekadee陀迦德 feud.Barony = &陀迦德TekadeeBarony{}

func init() {
    f := BTekadee陀迦德.(*陀迦德TekadeeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tekadee",
		TitleName: "陀迦德",
		TitleCode: "b_tekadee",
	}
}
