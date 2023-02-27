package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 居特斯洛GuterslohBarony struct {
	feud.BaseBarony
}

var BGutersloh居特斯洛 feud.Barony = &居特斯洛GuterslohBarony{}

func init() {
    f := BGutersloh居特斯洛.(*居特斯洛GuterslohBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gutersloh",
		TitleName: "居特斯洛",
		TitleCode: "b_gutersloh",
	}
}
