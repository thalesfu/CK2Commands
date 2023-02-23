package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔罗GaurBarony struct {
	feud.BaseBarony
}

var BGaur乔罗 feud.Barony = &乔罗GaurBarony{}

func init() {
	f := BGaur乔罗.(*乔罗GaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gaur",
		TitleName: "乔罗",
		TitleCode: "b_gaur",
	}
}
