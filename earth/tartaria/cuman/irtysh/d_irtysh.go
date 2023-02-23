package irtysh

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/irtysh/kimak"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/irtysh/kirghiz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IrtyshDuke interface {
	feud.Duke
	CKimak卡拉索尔() kimak.KimakCounty
	CKirghiz曳咥() kirghiz.KirghizCounty
}

type 曳咥IrtyshDuke struct {
	feud.BaseDuke
	卡拉索尔Kimak kimak.KimakCounty
	曳咥Kirghiz kirghiz.KirghizCounty
}

func (d *曳咥IrtyshDuke) CKimak卡拉索尔() kimak.KimakCounty {
	return d.卡拉索尔Kimak
}

func (d *曳咥IrtyshDuke) CKirghiz曳咥() kirghiz.KirghizCounty {
	return d.曳咥Kirghiz
}

var DIrtysh曳咥 IrtyshDuke = &曳咥IrtyshDuke{}

func init() {
	f := DIrtysh曳咥.(*曳咥IrtyshDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "irtysh",
		TitleName: "曳咥",
		TitleCode: "d_irtysh",
		Counties:  map[string]feud.County{},
	}

	f.卡拉索尔Kimak = kimak.CKimak卡拉索尔
	f.卡拉索尔Kimak.SetParent(f)

	f.曳咥Kirghiz = kirghiz.CKirghiz曳咥
	f.曳咥Kirghiz.SetParent(f)

}
