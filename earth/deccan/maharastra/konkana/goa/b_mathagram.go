package goa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩侘伽罗摩MathagramBarony struct {
	feud.BaseBarony
}

var BMathagram摩侘伽罗摩 feud.Barony = &摩侘伽罗摩MathagramBarony{}

func init() {
	f := BMathagram摩侘伽罗摩.(*摩侘伽罗摩MathagramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mathagram",
		TitleName: "摩侘伽罗摩",
		TitleCode: "b_mathagram",
	}
}
