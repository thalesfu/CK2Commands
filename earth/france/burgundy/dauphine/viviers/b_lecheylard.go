package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒谢拉尔LecheylardBarony struct {
	feud.BaseBarony
}

var BLecheylard勒谢拉尔 feud.Barony = &勒谢拉尔LecheylardBarony{}

func init() {
	f := BLecheylard勒谢拉尔.(*勒谢拉尔LecheylardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lecheylard",
		TitleName: "勒谢拉尔",
		TitleCode: "b_lecheylard",
	}
}
