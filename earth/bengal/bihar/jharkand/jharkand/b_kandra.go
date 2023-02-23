package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建陀罗KandraBarony struct {
	feud.BaseBarony
}

var BKandra建陀罗 feud.Barony = &建陀罗KandraBarony{}

func init() {
	f := BKandra建陀罗.(*建陀罗KandraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kandra",
		TitleName: "建陀罗",
		TitleCode: "b_kandra",
	}
}
