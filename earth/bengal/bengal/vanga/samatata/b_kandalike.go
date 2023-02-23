package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乾陀梨艺KandalikeBarony struct {
	feud.BaseBarony
}

var BKandalike乾陀梨艺 feud.Barony = &乾陀梨艺KandalikeBarony{}

func init() {
	f := BKandalike乾陀梨艺.(*乾陀梨艺KandalikeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandalike",
		TitleName: "乾陀梨艺",
		TitleCode: "b_kandalike",
	}
}
