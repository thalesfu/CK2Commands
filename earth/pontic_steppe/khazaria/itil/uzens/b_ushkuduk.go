package uzens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌什库杜克UshkudukBarony struct {
	feud.BaseBarony
}

var BUshkuduk乌什库杜克 feud.Barony = &乌什库杜克UshkudukBarony{}

func init() {
    f := BUshkuduk乌什库杜克.(*乌什库杜克UshkudukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ushkuduk",
		TitleName: "乌什库杜克",
		TitleCode: "b_ushkuduk",
	}
}
