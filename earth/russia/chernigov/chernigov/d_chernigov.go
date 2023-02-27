package chernigov

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/chernigov/chernigov"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/chernigov/lyubech"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/chernigov/novgorod_seversky"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/chernigov/starodub"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChernigovDuke interface {
    feud.Duke
    CChernigov切尔尼戈夫() 	chernigov.ChernigovCounty
    CLyubech柳别奇() 	lyubech.LyubechCounty
    CNovgorod_seversky诺夫哥罗德_谢韦尔斯基() 	novgorod_seversky.Novgorod_severskyCounty
    CStarodub斯塔罗杜布() 	starodub.StarodubCounty
}

type 切尔尼戈夫ChernigovDuke struct {
	feud.BaseDuke
	切尔尼戈夫Chernigov 	chernigov.ChernigovCounty
	柳别奇Lyubech 	lyubech.LyubechCounty
	诺夫哥罗德_谢韦尔斯基Novgorod_seversky 	novgorod_seversky.Novgorod_severskyCounty
	斯塔罗杜布Starodub 	starodub.StarodubCounty
}

func (d *切尔尼戈夫ChernigovDuke) CChernigov切尔尼戈夫() chernigov.ChernigovCounty {
	return d.切尔尼戈夫Chernigov
}
    
func (d *切尔尼戈夫ChernigovDuke) CLyubech柳别奇() lyubech.LyubechCounty {
	return d.柳别奇Lyubech
}
    
func (d *切尔尼戈夫ChernigovDuke) CNovgorod_seversky诺夫哥罗德_谢韦尔斯基() novgorod_seversky.Novgorod_severskyCounty {
	return d.诺夫哥罗德_谢韦尔斯基Novgorod_seversky
}
    
func (d *切尔尼戈夫ChernigovDuke) CStarodub斯塔罗杜布() starodub.StarodubCounty {
	return d.斯塔罗杜布Starodub
}
    
var DChernigov切尔尼戈夫 ChernigovDuke = &切尔尼戈夫ChernigovDuke{}

func init() {
	f := DChernigov切尔尼戈夫.(*切尔尼戈夫ChernigovDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "chernigov",
		TitleName: "切尔尼戈夫",
		TitleCode: "d_chernigov",
		Counties:  map[string]feud.County{},
	}

	f.切尔尼戈夫Chernigov = chernigov.CChernigov切尔尼戈夫
	f.切尔尼戈夫Chernigov.SetParent(f)
	
	f.柳别奇Lyubech = lyubech.CLyubech柳别奇
	f.柳别奇Lyubech.SetParent(f)
	
	f.诺夫哥罗德_谢韦尔斯基Novgorod_seversky = novgorod_seversky.CNovgorod_seversky诺夫哥罗德_谢韦尔斯基
	f.诺夫哥罗德_谢韦尔斯基Novgorod_seversky.SetParent(f)
	
	f.斯塔罗杜布Starodub = starodub.CStarodub斯塔罗杜布
	f.斯塔罗杜布Starodub.SetParent(f)
	
}
