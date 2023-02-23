package mali

import (
	"github.com/thalesfu/CK2Commands/earth/mali/mali/bambuk"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/mali"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/yatenga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaliKingdom interface {
	feud.Kingdom
	DBambuk班布克() bambuk.BambukDuke
	DMali曼丁() mali.MaliDuke
	DYatenga亚滕加() yatenga.YatengaDuke
}

type 马里MaliKingdom struct {
	feud.BaseKingdom
	班布克Bambuk  bambuk.BambukDuke
	曼丁Mali     mali.MaliDuke
	亚滕加Yatenga yatenga.YatengaDuke
}

func (k *马里MaliKingdom) DBambuk班布克() bambuk.BambukDuke {
	return k.班布克Bambuk
}

func (k *马里MaliKingdom) DMali曼丁() mali.MaliDuke {
	return k.曼丁Mali
}

func (k *马里MaliKingdom) DYatenga亚滕加() yatenga.YatengaDuke {
	return k.亚滕加Yatenga
}

var KMali马里 MaliKingdom = &马里MaliKingdom{}

func init() {
	f := KMali马里.(*马里MaliKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "mali",
		TitleName: "马里",
		TitleCode: "k_mali",
		Dukes:     map[string]feud.Duke{},
	}

	f.班布克Bambuk = bambuk.DBambuk班布克
	f.班布克Bambuk.SetParent(f)

	f.曼丁Mali = mali.DMali曼丁
	f.曼丁Mali.SetParent(f)

	f.亚滕加Yatenga = yatenga.DYatenga亚滕加
	f.亚滕加Yatenga.SetParent(f)

}
