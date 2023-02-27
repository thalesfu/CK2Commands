package charolais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿鲁河畔土伦ToulonsurarrouxBarony struct {
	feud.BaseBarony
}

var BToulonsurarroux阿鲁河畔土伦 feud.Barony = &阿鲁河畔土伦ToulonsurarrouxBarony{}

func init() {
    f := BToulonsurarroux阿鲁河畔土伦.(*阿鲁河畔土伦ToulonsurarrouxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toulonsurarroux",
		TitleName: "阿鲁河畔土伦",
		TitleCode: "b_toulonsurarroux",
	}
}
