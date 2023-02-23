package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀斯米耶QasmiyeBarony struct {
	feud.BaseBarony
}

var BQasmiye喀斯米耶 feud.Barony = &喀斯米耶QasmiyeBarony{}

func init() {
	f := BQasmiye喀斯米耶.(*喀斯米耶QasmiyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qasmiye",
		TitleName: "喀斯米耶",
		TitleCode: "b_qasmiye",
	}
}
