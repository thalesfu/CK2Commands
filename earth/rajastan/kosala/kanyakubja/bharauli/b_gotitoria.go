package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿帝都梨耶GotitoriaBarony struct {
	feud.BaseBarony
}

var BGotitoria瞿帝都梨耶 feud.Barony = &瞿帝都梨耶GotitoriaBarony{}

func init() {
	f := BGotitoria瞿帝都梨耶.(*瞿帝都梨耶GotitoriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gotitoria",
		TitleName: "瞿帝都梨耶",
		TitleCode: "b_gotitoria",
	}
}
