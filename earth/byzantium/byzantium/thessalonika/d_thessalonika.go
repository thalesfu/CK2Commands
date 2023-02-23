package thessalonika

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/thessalonika/chalkidike"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/thessalonika/seres"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/thessalonika/thessalonike"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThessalonikaDuke interface {
	feud.Duke
	CChalkidike哈尔基季基() chalkidike.ChalkidikeCounty
	CSeres塞雷斯() seres.SeresCounty
	CThessalonike塞萨洛尼基() thessalonike.ThessalonikeCounty
}

type 帖撒罗尼迦ThessalonikaDuke struct {
	feud.BaseDuke
	哈尔基季基Chalkidike   chalkidike.ChalkidikeCounty
	塞雷斯Seres          seres.SeresCounty
	塞萨洛尼基Thessalonike thessalonike.ThessalonikeCounty
}

func (d *帖撒罗尼迦ThessalonikaDuke) CChalkidike哈尔基季基() chalkidike.ChalkidikeCounty {
	return d.哈尔基季基Chalkidike
}

func (d *帖撒罗尼迦ThessalonikaDuke) CSeres塞雷斯() seres.SeresCounty {
	return d.塞雷斯Seres
}

func (d *帖撒罗尼迦ThessalonikaDuke) CThessalonike塞萨洛尼基() thessalonike.ThessalonikeCounty {
	return d.塞萨洛尼基Thessalonike
}

var DThessalonika帖撒罗尼迦 ThessalonikaDuke = &帖撒罗尼迦ThessalonikaDuke{}

func init() {
	f := DThessalonika帖撒罗尼迦.(*帖撒罗尼迦ThessalonikaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "thessalonika",
		TitleName: "帖撒罗尼迦",
		TitleCode: "d_thessalonika",
		Counties:  map[string]feud.County{},
	}

	f.哈尔基季基Chalkidike = chalkidike.CChalkidike哈尔基季基
	f.哈尔基季基Chalkidike.SetParent(f)

	f.塞雷斯Seres = seres.CSeres塞雷斯
	f.塞雷斯Seres.SetParent(f)

	f.塞萨洛尼基Thessalonike = thessalonike.CThessalonike塞萨洛尼基
	f.塞萨洛尼基Thessalonike.SetParent(f)

}
