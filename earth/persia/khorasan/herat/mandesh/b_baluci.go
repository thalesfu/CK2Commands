package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俾路支BaluciBarony struct {
	feud.BaseBarony
}

var BBaluci俾路支 feud.Barony = &俾路支BaluciBarony{}

func init() {
	f := BBaluci俾路支.(*俾路支BaluciBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baluci",
		TitleName: "俾路支",
		TitleCode: "b_baluci",
	}
}
