package ob

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ob/buqtirma"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ob/katun"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ob/ket"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ob/ob"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/ob/zaysan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ObDuke interface {
	feud.Duke
	CBuqtirma布赫塔尔马() buqtirma.BuqtirmaCounty
	CKatun卡通() katun.KatunCounty
	CKet阿努伊() ket.KetCounty
	COb鄂毕() ob.ObCounty
	CZaysan斋桑() zaysan.ZaysanCounty
}

type 鄂毕ObDuke struct {
	feud.BaseDuke
	布赫塔尔马Buqtirma buqtirma.BuqtirmaCounty
	卡通Katun       katun.KatunCounty
	阿努伊Ket        ket.KetCounty
	鄂毕Ob          ob.ObCounty
	斋桑Zaysan      zaysan.ZaysanCounty
}

func (d *鄂毕ObDuke) CBuqtirma布赫塔尔马() buqtirma.BuqtirmaCounty {
	return d.布赫塔尔马Buqtirma
}

func (d *鄂毕ObDuke) CKatun卡通() katun.KatunCounty {
	return d.卡通Katun
}

func (d *鄂毕ObDuke) CKet阿努伊() ket.KetCounty {
	return d.阿努伊Ket
}

func (d *鄂毕ObDuke) COb鄂毕() ob.ObCounty {
	return d.鄂毕Ob
}

func (d *鄂毕ObDuke) CZaysan斋桑() zaysan.ZaysanCounty {
	return d.斋桑Zaysan
}

var DOb鄂毕 ObDuke = &鄂毕ObDuke{}

func init() {
	f := DOb鄂毕.(*鄂毕ObDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ob",
		TitleName: "鄂毕",
		TitleCode: "d_ob",
		Counties:  map[string]feud.County{},
	}

	f.布赫塔尔马Buqtirma = buqtirma.CBuqtirma布赫塔尔马
	f.布赫塔尔马Buqtirma.SetParent(f)

	f.卡通Katun = katun.CKatun卡通
	f.卡通Katun.SetParent(f)

	f.阿努伊Ket = ket.CKet阿努伊
	f.阿努伊Ket.SetParent(f)

	f.鄂毕Ob = ob.COb鄂毕
	f.鄂毕Ob.SetParent(f)

	f.斋桑Zaysan = zaysan.CZaysan斋桑
	f.斋桑Zaysan.SetParent(f)

}
