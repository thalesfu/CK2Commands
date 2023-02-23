package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨诃吉罗SathkiraBarony struct {
	feud.BaseBarony
}

var BSathkira萨诃吉罗 feud.Barony = &萨诃吉罗SathkiraBarony{}

func init() {
	f := BSathkira萨诃吉罗.(*萨诃吉罗SathkiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sathkira",
		TitleName: "萨诃吉罗",
		TitleCode: "b_sathkira",
	}
}
