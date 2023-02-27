package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多喀奇DoqaqitBarony struct {
	feud.BaseBarony
}

var BDoqaqit多喀奇 feud.Barony = &多喀奇DoqaqitBarony{}

func init() {
    f := BDoqaqit多喀奇.(*多喀奇DoqaqitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doqaqit",
		TitleName: "多喀奇",
		TitleCode: "b_doqaqit",
	}
}
