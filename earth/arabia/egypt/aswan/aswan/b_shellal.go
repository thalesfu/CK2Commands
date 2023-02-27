package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢拉尔ShellalBarony struct {
	feud.BaseBarony
}

var BShellal谢拉尔 feud.Barony = &谢拉尔ShellalBarony{}

func init() {
    f := BShellal谢拉尔.(*谢拉尔ShellalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shellal",
		TitleName: "谢拉尔",
		TitleCode: "b_shellal",
	}
}
