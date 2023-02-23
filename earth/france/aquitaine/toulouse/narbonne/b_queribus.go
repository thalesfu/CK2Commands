package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯里比QueribusBarony struct {
	feud.BaseBarony
}

var BQueribus凯里比 feud.Barony = &凯里比QueribusBarony{}

func init() {
	f := BQueribus凯里比.(*凯里比QueribusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "queribus",
		TitleName: "凯里比",
		TitleCode: "b_queribus",
	}
}
