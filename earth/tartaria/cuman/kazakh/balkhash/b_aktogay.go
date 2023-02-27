package balkhash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克托盖AktogayBarony struct {
	feud.BaseBarony
}

var BAktogay阿克托盖 feud.Barony = &阿克托盖AktogayBarony{}

func init() {
    f := BAktogay阿克托盖.(*阿克托盖AktogayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aktogay",
		TitleName: "阿克托盖",
		TitleCode: "b_aktogay",
	}
}
