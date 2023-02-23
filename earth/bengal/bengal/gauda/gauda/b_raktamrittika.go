package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 络多末知RaktamrittikaBarony struct {
	feud.BaseBarony
}

var BRaktamrittika络多末知 feud.Barony = &络多末知RaktamrittikaBarony{}

func init() {
	f := BRaktamrittika络多末知.(*络多末知RaktamrittikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raktamrittika",
		TitleName: "络多末知",
		TitleCode: "b_raktamrittika",
	}
}
