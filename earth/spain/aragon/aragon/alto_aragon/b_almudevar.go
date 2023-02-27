package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔穆德瓦尔AlmudevarBarony struct {
	feud.BaseBarony
}

var BAlmudevar阿尔穆德瓦尔 feud.Barony = &阿尔穆德瓦尔AlmudevarBarony{}

func init() {
    f := BAlmudevar阿尔穆德瓦尔.(*阿尔穆德瓦尔AlmudevarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almudevar",
		TitleName: "阿尔穆德瓦尔",
		TitleCode: "b_almudevar",
	}
}
