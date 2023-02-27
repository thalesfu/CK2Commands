package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝卢尔BelurBarony struct {
	feud.BaseBarony
}

var BBelur贝卢尔 feud.Barony = &贝卢尔BelurBarony{}

func init() {
    f := BBelur贝卢尔.(*贝卢尔BelurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belur",
		TitleName: "贝卢尔",
		TitleCode: "b_belur",
	}
}
