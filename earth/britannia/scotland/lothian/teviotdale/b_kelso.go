package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔索KelsoBarony struct {
	feud.BaseBarony
}

var BKelso凯尔索 feud.Barony = &凯尔索KelsoBarony{}

func init() {
	f := BKelso凯尔索.(*凯尔索KelsoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kelso",
		TitleName: "凯尔索",
		TitleCode: "b_kelso",
	}
}
