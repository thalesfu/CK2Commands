package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特利维ThreaveBarony struct {
	feud.BaseBarony
}

var BThreave特利维 feud.Barony = &特利维ThreaveBarony{}

func init() {
	f := BThreave特利维.(*特利维ThreaveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "threave",
		TitleName: "特利维",
		TitleCode: "b_threave",
	}
}
