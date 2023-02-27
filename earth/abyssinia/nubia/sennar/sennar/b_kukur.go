package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库库尔KukurBarony struct {
	feud.BaseBarony
}

var BKukur库库尔 feud.Barony = &库库尔KukurBarony{}

func init() {
    f := BKukur库库尔.(*库库尔KukurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukur",
		TitleName: "库库尔",
		TitleCode: "b_kukur",
	}
}
