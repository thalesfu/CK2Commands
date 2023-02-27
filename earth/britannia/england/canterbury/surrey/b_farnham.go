package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法纳姆FarnhamBarony struct {
	feud.BaseBarony
}

var BFarnham法纳姆 feud.Barony = &法纳姆FarnhamBarony{}

func init() {
    f := BFarnham法纳姆.(*法纳姆FarnhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farnham",
		TitleName: "法纳姆",
		TitleCode: "b_farnham",
	}
}
