package uvs

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/uvs/kharkhiraa"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/uvs/kyzyl"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/uvs/tsambagarav"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/uvs/uvs"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UvsDuke interface {
	feud.Duke
	CKharkhiraa哈尔齐拉() kharkhiraa.KharkhiraaCounty
	CKyzyl哈尔嘎斯() kyzyl.KyzylCounty
	CTsambagarav查木巴嘎拉布() tsambagarav.TsambagaravCounty
	CUvs乌布苏() uvs.UvsCounty
}

type 乌布苏UvsDuke struct {
	feud.BaseDuke
	哈尔齐拉Kharkhiraa    kharkhiraa.KharkhiraaCounty
	哈尔嘎斯Kyzyl         kyzyl.KyzylCounty
	查木巴嘎拉布Tsambagarav tsambagarav.TsambagaravCounty
	乌布苏Uvs            uvs.UvsCounty
}

func (d *乌布苏UvsDuke) CKharkhiraa哈尔齐拉() kharkhiraa.KharkhiraaCounty {
	return d.哈尔齐拉Kharkhiraa
}

func (d *乌布苏UvsDuke) CKyzyl哈尔嘎斯() kyzyl.KyzylCounty {
	return d.哈尔嘎斯Kyzyl
}

func (d *乌布苏UvsDuke) CTsambagarav查木巴嘎拉布() tsambagarav.TsambagaravCounty {
	return d.查木巴嘎拉布Tsambagarav
}

func (d *乌布苏UvsDuke) CUvs乌布苏() uvs.UvsCounty {
	return d.乌布苏Uvs
}

var DUvs乌布苏 UvsDuke = &乌布苏UvsDuke{}

func init() {
	f := DUvs乌布苏.(*乌布苏UvsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "uvs",
		TitleName: "乌布苏",
		TitleCode: "d_uvs",
		Counties:  map[string]feud.County{},
	}

	f.哈尔齐拉Kharkhiraa = kharkhiraa.CKharkhiraa哈尔齐拉
	f.哈尔齐拉Kharkhiraa.SetParent(f)

	f.哈尔嘎斯Kyzyl = kyzyl.CKyzyl哈尔嘎斯
	f.哈尔嘎斯Kyzyl.SetParent(f)

	f.查木巴嘎拉布Tsambagarav = tsambagarav.CTsambagarav查木巴嘎拉布
	f.查木巴嘎拉布Tsambagarav.SetParent(f)

	f.乌布苏Uvs = uvs.CUvs乌布苏
	f.乌布苏Uvs.SetParent(f)

}
