package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼荼罗耶罗MandraelBarony struct {
	feud.BaseBarony
}

var BMandrael曼荼罗耶罗 feud.Barony = &曼荼罗耶罗MandraelBarony{}

func init() {
	f := BMandrael曼荼罗耶罗.(*曼荼罗耶罗MandraelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandrael",
		TitleName: "曼荼罗耶罗",
		TitleCode: "b_mandrael",
	}
}
