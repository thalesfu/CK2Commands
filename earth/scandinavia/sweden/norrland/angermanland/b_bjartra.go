package angermanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比耶尔特罗BjartraBarony struct {
	feud.BaseBarony
}

var BBjartra比耶尔特罗 feud.Barony = &比耶尔特罗BjartraBarony{}

func init() {
	f := BBjartra比耶尔特罗.(*比耶尔特罗BjartraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bjartra",
		TitleName: "比耶尔特罗",
		TitleCode: "b_bjartra",
	}
}
