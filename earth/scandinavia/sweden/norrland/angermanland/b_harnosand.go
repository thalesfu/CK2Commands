package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海讷桑德HarnosandBarony struct {
	feud.BaseBarony
}

var BHarnosand海讷桑德 feud.Barony = &海讷桑德HarnosandBarony{}

func init() {
	f := BHarnosand海讷桑德.(*海讷桑德HarnosandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harnosand",
		TitleName: "海讷桑德",
		TitleCode: "b_harnosand",
	}
}
