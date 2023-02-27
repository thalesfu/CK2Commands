package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌良科沃UlyankovoBarony struct {
	feud.BaseBarony
}

var BUlyankovo乌良科沃 feud.Barony = &乌良科沃UlyankovoBarony{}

func init() {
    f := BUlyankovo乌良科沃.(*乌良科沃UlyankovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulyankovo",
		TitleName: "乌良科沃",
		TitleCode: "b_ulyankovo",
	}
}
