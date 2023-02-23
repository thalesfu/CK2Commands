package bambuk

import (
	"github.com/thalesfu/CK2Commands/earth/mali/mali/bambuk/bambuk"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/bambuk/diafunu"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/bambuk/kantor"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BambukDuke interface {
	feud.Duke
	CBambuk班布克() bambuk.BambukCounty
	CDiafunu迪亚富努() diafunu.DiafunuCounty
	CKantor坎托尔() kantor.KantorCounty
}

type 班布克BambukDuke struct {
	feud.BaseDuke
	班布克Bambuk   bambuk.BambukCounty
	迪亚富努Diafunu diafunu.DiafunuCounty
	坎托尔Kantor   kantor.KantorCounty
}

func (d *班布克BambukDuke) CBambuk班布克() bambuk.BambukCounty {
	return d.班布克Bambuk
}

func (d *班布克BambukDuke) CDiafunu迪亚富努() diafunu.DiafunuCounty {
	return d.迪亚富努Diafunu
}

func (d *班布克BambukDuke) CKantor坎托尔() kantor.KantorCounty {
	return d.坎托尔Kantor
}

var DBambuk班布克 BambukDuke = &班布克BambukDuke{}

func init() {
	f := DBambuk班布克.(*班布克BambukDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bambuk",
		TitleName: "班布克",
		TitleCode: "d_bambuk",
		Counties:  map[string]feud.County{},
	}

	f.班布克Bambuk = bambuk.CBambuk班布克
	f.班布克Bambuk.SetParent(f)

	f.迪亚富努Diafunu = diafunu.CDiafunu迪亚富努
	f.迪亚富努Diafunu.SetParent(f)

	f.坎托尔Kantor = kantor.CKantor坎托尔
	f.坎托尔Kantor.SetParent(f)

}
