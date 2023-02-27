package massat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃迈尔尚AmerchenBarony struct {
	feud.BaseBarony
}

var BAmerchen埃迈尔尚 feud.Barony = &埃迈尔尚AmerchenBarony{}

func init() {
    f := BAmerchen埃迈尔尚.(*埃迈尔尚AmerchenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amerchen",
		TitleName: "埃迈尔尚",
		TitleCode: "b_amerchen",
	}
}
