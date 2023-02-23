package kurdistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kurdistan/kurdistan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/kurdistan/shahrazur"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KurdistanDuke interface {
	feud.Duke
	CKurdistan库尔德斯坦() kurdistan.KurdistanCounty
	CShahrazur沙赫拉祖尔() shahrazur.ShahrazurCounty
}

type 库尔德斯坦KurdistanDuke struct {
	feud.BaseDuke
	库尔德斯坦Kurdistan kurdistan.KurdistanCounty
	沙赫拉祖尔Shahrazur shahrazur.ShahrazurCounty
}

func (d *库尔德斯坦KurdistanDuke) CKurdistan库尔德斯坦() kurdistan.KurdistanCounty {
	return d.库尔德斯坦Kurdistan
}

func (d *库尔德斯坦KurdistanDuke) CShahrazur沙赫拉祖尔() shahrazur.ShahrazurCounty {
	return d.沙赫拉祖尔Shahrazur
}

var DKurdistan库尔德斯坦 KurdistanDuke = &库尔德斯坦KurdistanDuke{}

func init() {
	f := DKurdistan库尔德斯坦.(*库尔德斯坦KurdistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kurdistan",
		TitleName: "库尔德斯坦",
		TitleCode: "d_kurdistan",
		Counties:  map[string]feud.County{},
	}

	f.库尔德斯坦Kurdistan = kurdistan.CKurdistan库尔德斯坦
	f.库尔德斯坦Kurdistan.SetParent(f)

	f.沙赫拉祖尔Shahrazur = shahrazur.CShahrazur沙赫拉祖尔
	f.沙赫拉祖尔Shahrazur.SetParent(f)

}
