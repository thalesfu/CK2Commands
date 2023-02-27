package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷戈巴TinngobaBarony struct {
	feud.BaseBarony
}

var BTinngoba廷戈巴 feud.Barony = &廷戈巴TinngobaBarony{}

func init() {
    f := BTinngoba廷戈巴.(*廷戈巴TinngobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tinngoba",
		TitleName: "廷戈巴",
		TitleCode: "b_tinngoba",
	}
}
