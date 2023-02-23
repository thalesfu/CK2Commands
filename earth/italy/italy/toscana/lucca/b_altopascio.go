package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔托帕肖AltopascioBarony struct {
	feud.BaseBarony
}

var BAltopascio阿尔托帕肖 feud.Barony = &阿尔托帕肖AltopascioBarony{}

func init() {
	f := BAltopascio阿尔托帕肖.(*阿尔托帕肖AltopascioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altopascio",
		TitleName: "阿尔托帕肖",
		TitleCode: "b_altopascio",
	}
}
