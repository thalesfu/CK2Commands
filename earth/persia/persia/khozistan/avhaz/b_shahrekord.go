package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫尔库尔德ShahrekordBarony struct {
	feud.BaseBarony
}

var BShahrekord沙赫尔库尔德 feud.Barony = &沙赫尔库尔德ShahrekordBarony{}

func init() {
    f := BShahrekord沙赫尔库尔德.(*沙赫尔库尔德ShahrekordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahrekord",
		TitleName: "沙赫尔库尔德",
		TitleCode: "b_shahrekord",
	}
}
