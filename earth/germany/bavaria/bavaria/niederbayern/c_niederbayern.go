package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NiederbayernCounty interface {
	feud.County
	BCham卡姆() feud.Barony
	BKirchroth基希罗特() feud.Barony
	BRegenstauf雷根施陶夫() feud.Barony
	BRoding罗丁() feud.Barony
	BSchwandorf施万多夫() feud.Barony
	BWittelsbach维特尔斯巴赫() feud.Barony
}

type 卡姆NiederbayernCounty struct {
	feud.BaseCounty
	卡姆Cham            feud.Barony
	基希罗特Kirchroth     feud.Barony
	雷根施陶夫Regenstauf   feud.Barony
	罗丁Roding          feud.Barony
	施万多夫Schwandorf    feud.Barony
	维特尔斯巴赫Wittelsbach feud.Barony
}

func (c *卡姆NiederbayernCounty) BCham卡姆() feud.Barony {
	return c.卡姆Cham
}

func (c *卡姆NiederbayernCounty) BKirchroth基希罗特() feud.Barony {
	return c.基希罗特Kirchroth
}

func (c *卡姆NiederbayernCounty) BRegenstauf雷根施陶夫() feud.Barony {
	return c.雷根施陶夫Regenstauf
}

func (c *卡姆NiederbayernCounty) BRoding罗丁() feud.Barony {
	return c.罗丁Roding
}

func (c *卡姆NiederbayernCounty) BSchwandorf施万多夫() feud.Barony {
	return c.施万多夫Schwandorf
}

func (c *卡姆NiederbayernCounty) BWittelsbach维特尔斯巴赫() feud.Barony {
	return c.维特尔斯巴赫Wittelsbach
}

var CNiederbayern卡姆 NiederbayernCounty = &卡姆NiederbayernCounty{}

func init() {
	f := CNiederbayern卡姆.(*卡姆NiederbayernCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "361",
		Title:     "niederbayern",
		TitleName: "卡姆",
		TitleCode: "c_niederbayern",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡姆Cham = BCham卡姆
	f.卡姆Cham.SetParent(f)

	f.基希罗特Kirchroth = BKirchroth基希罗特
	f.基希罗特Kirchroth.SetParent(f)

	f.雷根施陶夫Regenstauf = BRegenstauf雷根施陶夫
	f.雷根施陶夫Regenstauf.SetParent(f)

	f.罗丁Roding = BRoding罗丁
	f.罗丁Roding.SetParent(f)

	f.施万多夫Schwandorf = BSchwandorf施万多夫
	f.施万多夫Schwandorf.SetParent(f)

	f.维特尔斯巴赫Wittelsbach = BWittelsbach维特尔斯巴赫
	f.维特尔斯巴赫Wittelsbach.SetParent(f)

}
