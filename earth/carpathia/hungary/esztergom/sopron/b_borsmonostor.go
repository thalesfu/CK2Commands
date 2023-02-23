package sopron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔什莫诺什托尔BorsmonostorBarony struct {
	feud.BaseBarony
}

var BBorsmonostor博尔什莫诺什托尔 feud.Barony = &博尔什莫诺什托尔BorsmonostorBarony{}

func init() {
	f := BBorsmonostor博尔什莫诺什托尔.(*博尔什莫诺什托尔BorsmonostorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borsmonostor",
		TitleName: "博尔什莫诺什托尔",
		TitleCode: "b_borsmonostor",
	}
}
