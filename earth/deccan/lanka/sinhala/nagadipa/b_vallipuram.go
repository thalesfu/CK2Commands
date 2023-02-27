package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆利浮兰VallipuramBarony struct {
	feud.BaseBarony
}

var BVallipuram婆利浮兰 feud.Barony = &婆利浮兰VallipuramBarony{}

func init() {
    f := BVallipuram婆利浮兰.(*婆利浮兰VallipuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vallipuram",
		TitleName: "婆利浮兰",
		TitleCode: "b_vallipuram",
	}
}
