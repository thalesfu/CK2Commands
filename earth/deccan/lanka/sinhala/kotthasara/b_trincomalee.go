package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亭可马里TrincomaleeBarony struct {
	feud.BaseBarony
}

var BTrincomalee亭可马里 feud.Barony = &亭可马里TrincomaleeBarony{}

func init() {
	f := BTrincomalee亭可马里.(*亭可马里TrincomaleeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trincomalee",
		TitleName: "亭可马里",
		TitleCode: "b_trincomalee",
	}
}
