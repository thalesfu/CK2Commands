package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔伊ChoyBarony struct {
	feud.BaseBarony
}

var BChoy乔伊 feud.Barony = &乔伊ChoyBarony{}

func init() {
    f := BChoy乔伊.(*乔伊ChoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "choy",
		TitleName: "乔伊",
		TitleCode: "b_choy",
	}
}
