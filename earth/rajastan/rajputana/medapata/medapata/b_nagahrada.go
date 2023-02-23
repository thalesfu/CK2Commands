package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那伽贺逻驮NagahradaBarony struct {
	feud.BaseBarony
}

var BNagahrada那伽贺逻驮 feud.Barony = &那伽贺逻驮NagahradaBarony{}

func init() {
	f := BNagahrada那伽贺逻驮.(*那伽贺逻驮NagahradaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagahrada",
		TitleName: "那伽贺逻驮",
		TitleCode: "b_nagahrada",
	}
}
