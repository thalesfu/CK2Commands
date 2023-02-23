package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋陀罗迦梨BhadrakaliBarony struct {
	feud.BaseBarony
}

var BBhadrakali跋陀罗迦梨 feud.Barony = &跋陀罗迦梨BhadrakaliBarony{}

func init() {
	f := BBhadrakali跋陀罗迦梨.(*跋陀罗迦梨BhadrakaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadrakali",
		TitleName: "跋陀罗迦梨",
		TitleCode: "b_bhadrakali",
	}
}
