package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋陀罗BhadraBarony struct {
	feud.BaseBarony
}

var BBhadra跋陀罗 feud.Barony = &跋陀罗BhadraBarony{}

func init() {
	f := BBhadra跋陀罗.(*跋陀罗BhadraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadra",
		TitleName: "跋陀罗",
		TitleCode: "b_bhadra",
	}
}
