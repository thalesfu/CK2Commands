package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剑苇JanwaiBarony struct {
	feud.BaseBarony
}

var BJanwai剑苇 feud.Barony = &剑苇JanwaiBarony{}

func init() {
    f := BJanwai剑苇.(*剑苇JanwaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "janwai",
		TitleName: "剑苇",
		TitleCode: "b_janwai",
	}
}
