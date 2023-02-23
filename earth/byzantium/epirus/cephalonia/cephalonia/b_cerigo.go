package cephalonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切里戈CerigoBarony struct {
	feud.BaseBarony
}

var BCerigo切里戈 feud.Barony = &切里戈CerigoBarony{}

func init() {
	f := BCerigo切里戈.(*切里戈CerigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cerigo",
		TitleName: "切里戈",
		TitleCode: "b_cerigo",
	}
}
