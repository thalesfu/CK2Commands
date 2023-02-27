package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿加尔AgarBarony struct {
	feud.BaseBarony
}

var BAgar阿加尔 feud.Barony = &阿加尔AgarBarony{}

func init() {
    f := BAgar阿加尔.(*阿加尔AgarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agar",
		TitleName: "阿加尔",
		TitleCode: "b_agar",
	}
}
