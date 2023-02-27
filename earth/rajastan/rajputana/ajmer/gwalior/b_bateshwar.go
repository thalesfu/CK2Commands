package gwalior

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆知湿伐罗BateshwarBarony struct {
	feud.BaseBarony
}

var BBateshwar婆知湿伐罗 feud.Barony = &婆知湿伐罗BateshwarBarony{}

func init() {
    f := BBateshwar婆知湿伐罗.(*婆知湿伐罗BateshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bateshwar",
		TitleName: "婆知湿伐罗",
		TitleCode: "b_bateshwar",
	}
}
