package albret

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlbretCounty interface {
	feud.County
	BBazas巴扎斯() feud.Barony
	BGabarret加巴雷() feud.Barony
	BLabrit拉布里() feud.Barony
	BLatestedebuch拉泰斯特德比克() feud.Barony
	BMimizan米米藏() feud.Barony
	BMontdemarsan蒙德马桑() feud.Barony
	BRoquefort罗克福尔() feud.Barony
	BTartas塔尔塔斯() feud.Barony
}

type 马桑AlbretCounty struct {
	feud.BaseCounty
	巴扎斯Bazas             feud.Barony
	加巴雷Gabarret          feud.Barony
	拉布里Labrit            feud.Barony
	拉泰斯特德比克Latestedebuch feud.Barony
	米米藏Mimizan           feud.Barony
	蒙德马桑Montdemarsan     feud.Barony
	罗克福尔Roquefort        feud.Barony
	塔尔塔斯Tartas           feud.Barony
}

func (c *马桑AlbretCounty) BBazas巴扎斯() feud.Barony {
	return c.巴扎斯Bazas
}

func (c *马桑AlbretCounty) BGabarret加巴雷() feud.Barony {
	return c.加巴雷Gabarret
}

func (c *马桑AlbretCounty) BLabrit拉布里() feud.Barony {
	return c.拉布里Labrit
}

func (c *马桑AlbretCounty) BLatestedebuch拉泰斯特德比克() feud.Barony {
	return c.拉泰斯特德比克Latestedebuch
}

func (c *马桑AlbretCounty) BMimizan米米藏() feud.Barony {
	return c.米米藏Mimizan
}

func (c *马桑AlbretCounty) BMontdemarsan蒙德马桑() feud.Barony {
	return c.蒙德马桑Montdemarsan
}

func (c *马桑AlbretCounty) BRoquefort罗克福尔() feud.Barony {
	return c.罗克福尔Roquefort
}

func (c *马桑AlbretCounty) BTartas塔尔塔斯() feud.Barony {
	return c.塔尔塔斯Tartas
}

var CAlbret马桑 AlbretCounty = &马桑AlbretCounty{}

func init() {
	f := CAlbret马桑.(*马桑AlbretCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "150",
		Title:     "albret",
		TitleName: "马桑",
		TitleCode: "c_albret",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴扎斯Bazas = BBazas巴扎斯
	f.巴扎斯Bazas.SetParent(f)

	f.加巴雷Gabarret = BGabarret加巴雷
	f.加巴雷Gabarret.SetParent(f)

	f.拉布里Labrit = BLabrit拉布里
	f.拉布里Labrit.SetParent(f)

	f.拉泰斯特德比克Latestedebuch = BLatestedebuch拉泰斯特德比克
	f.拉泰斯特德比克Latestedebuch.SetParent(f)

	f.米米藏Mimizan = BMimizan米米藏
	f.米米藏Mimizan.SetParent(f)

	f.蒙德马桑Montdemarsan = BMontdemarsan蒙德马桑
	f.蒙德马桑Montdemarsan.SetParent(f)

	f.罗克福尔Roquefort = BRoquefort罗克福尔
	f.罗克福尔Roquefort.SetParent(f)

	f.塔尔塔斯Tartas = BTartas塔尔塔斯
	f.塔尔塔斯Tartas.SetParent(f)

}
