package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KakhetiCounty interface {
	feud.County
	BBochorma博乔尔马() feud.Barony
	BBodbe巴德比() feud.Barony
	BDavidgareja大卫加列贾() feud.Barony
	BDzveligalavani兹韦利加拉瓦尼() feud.Barony
	BGremi吉米() feud.Barony
	BSighnaghi西格纳吉() feud.Barony
	BTelavi泰拉维() feud.Barony
	BZedazaden泽达泽登() feud.Barony
}

type 卡赫季KakhetiCounty struct {
	feud.BaseCounty
	博乔尔马Bochorma          feud.Barony
	巴德比Bodbe              feud.Barony
	大卫加列贾Davidgareja      feud.Barony
	兹韦利加拉瓦尼Dzveligalavani feud.Barony
	吉米Gremi               feud.Barony
	西格纳吉Sighnaghi         feud.Barony
	泰拉维Telavi             feud.Barony
	泽达泽登Zedazaden         feud.Barony
}

func (c *卡赫季KakhetiCounty) BBochorma博乔尔马() feud.Barony {
	return c.博乔尔马Bochorma
}

func (c *卡赫季KakhetiCounty) BBodbe巴德比() feud.Barony {
	return c.巴德比Bodbe
}

func (c *卡赫季KakhetiCounty) BDavidgareja大卫加列贾() feud.Barony {
	return c.大卫加列贾Davidgareja
}

func (c *卡赫季KakhetiCounty) BDzveligalavani兹韦利加拉瓦尼() feud.Barony {
	return c.兹韦利加拉瓦尼Dzveligalavani
}

func (c *卡赫季KakhetiCounty) BGremi吉米() feud.Barony {
	return c.吉米Gremi
}

func (c *卡赫季KakhetiCounty) BSighnaghi西格纳吉() feud.Barony {
	return c.西格纳吉Sighnaghi
}

func (c *卡赫季KakhetiCounty) BTelavi泰拉维() feud.Barony {
	return c.泰拉维Telavi
}

func (c *卡赫季KakhetiCounty) BZedazaden泽达泽登() feud.Barony {
	return c.泽达泽登Zedazaden
}

var CKakheti卡赫季 KakhetiCounty = &卡赫季KakhetiCounty{}

func init() {
	f := CKakheti卡赫季.(*卡赫季KakhetiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "676",
		Title:     "kakheti",
		TitleName: "卡赫季",
		TitleCode: "c_kakheti",
		Baronies:  map[string]feud.Barony{},
	}

	f.博乔尔马Bochorma = BBochorma博乔尔马
	f.博乔尔马Bochorma.SetParent(f)

	f.巴德比Bodbe = BBodbe巴德比
	f.巴德比Bodbe.SetParent(f)

	f.大卫加列贾Davidgareja = BDavidgareja大卫加列贾
	f.大卫加列贾Davidgareja.SetParent(f)

	f.兹韦利加拉瓦尼Dzveligalavani = BDzveligalavani兹韦利加拉瓦尼
	f.兹韦利加拉瓦尼Dzveligalavani.SetParent(f)

	f.吉米Gremi = BGremi吉米
	f.吉米Gremi.SetParent(f)

	f.西格纳吉Sighnaghi = BSighnaghi西格纳吉
	f.西格纳吉Sighnaghi.SetParent(f)

	f.泰拉维Telavi = BTelavi泰拉维
	f.泰拉维Telavi.SetParent(f)

	f.泽达泽登Zedazaden = BZedazaden泽达泽登
	f.泽达泽登Zedazaden.SetParent(f)

}
