package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那伽梨NagariBarony struct {
	feud.BaseBarony
}

var BNagari那伽梨 feud.Barony = &那伽梨NagariBarony{}

func init() {
	f := BNagari那伽梨.(*那伽梨NagariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagari",
		TitleName: "那伽梨",
		TitleCode: "b_nagari",
	}
}
