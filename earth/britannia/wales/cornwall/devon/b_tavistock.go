package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔维斯托克TavistockBarony struct {
	feud.BaseBarony
}

var BTavistock塔维斯托克 feud.Barony = &塔维斯托克TavistockBarony{}

func init() {
	f := BTavistock塔维斯托克.(*塔维斯托克TavistockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tavistock",
		TitleName: "塔维斯托克",
		TitleCode: "b_tavistock",
	}
}
