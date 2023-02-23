package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽婆梨耶罗GwaliorBarony struct {
	feud.BaseBarony
}

var BGwalior伽婆梨耶罗 feud.Barony = &伽婆梨耶罗GwaliorBarony{}

func init() {
	f := BGwalior伽婆梨耶罗.(*伽婆梨耶罗GwaliorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gwalior",
		TitleName: "伽婆梨耶罗",
		TitleCode: "b_gwalior",
	}
}
