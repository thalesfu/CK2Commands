package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Novgorod_severskyCounty interface {
    feud.County
    BAltynivka阿尔特诺夫卡() 	feud.Barony
    BAtyshua阿秋沙() 	feud.Barony
    BBakhmatch巴赫马奇() 	feud.Barony
    BBatouryn巴图林() 	feud.Barony
    BGlukhov格卢霍夫() 	feud.Barony
    BKonotop科诺托普() 	feud.Barony
    BNovgorodseversky诺夫哥罗德_谢韦尔斯基() 	feud.Barony
}

type 诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty struct {
	feud.BaseCounty
	阿尔特诺夫卡Altynivka 	feud.Barony
	阿秋沙Atyshua 	feud.Barony
	巴赫马奇Bakhmatch 	feud.Barony
	巴图林Batouryn 	feud.Barony
	格卢霍夫Glukhov 	feud.Barony
	科诺托普Konotop 	feud.Barony
	诺夫哥罗德_谢韦尔斯基Novgorodseversky 	feud.Barony
}

func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BAltynivka阿尔特诺夫卡() feud.Barony {
	return c.阿尔特诺夫卡Altynivka
}
    
func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BAtyshua阿秋沙() feud.Barony {
	return c.阿秋沙Atyshua
}
    
func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BBakhmatch巴赫马奇() feud.Barony {
	return c.巴赫马奇Bakhmatch
}
    
func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BBatouryn巴图林() feud.Barony {
	return c.巴图林Batouryn
}
    
func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BGlukhov格卢霍夫() feud.Barony {
	return c.格卢霍夫Glukhov
}
    
func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BKonotop科诺托普() feud.Barony {
	return c.科诺托普Konotop
}
    
func (c *诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty) BNovgorodseversky诺夫哥罗德_谢韦尔斯基() feud.Barony {
	return c.诺夫哥罗德_谢韦尔斯基Novgorodseversky
}
    
var CNovgorod_seversky诺夫哥罗德_谢韦尔斯基 Novgorod_severskyCounty = &诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty{}

func init() {
	f := CNovgorod_seversky诺夫哥罗德_谢韦尔斯基.(*诺夫哥罗德_谢韦尔斯基Novgorod_severskyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "567",
		Title:     "novgorod_seversky",
		TitleName: "诺夫哥罗德_谢韦尔斯基",
		TitleCode: "c_novgorod_seversky",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔特诺夫卡Altynivka = BAltynivka阿尔特诺夫卡
	f.阿尔特诺夫卡Altynivka.SetParent(f)
	
	f.阿秋沙Atyshua = BAtyshua阿秋沙
	f.阿秋沙Atyshua.SetParent(f)
	
	f.巴赫马奇Bakhmatch = BBakhmatch巴赫马奇
	f.巴赫马奇Bakhmatch.SetParent(f)
	
	f.巴图林Batouryn = BBatouryn巴图林
	f.巴图林Batouryn.SetParent(f)
	
	f.格卢霍夫Glukhov = BGlukhov格卢霍夫
	f.格卢霍夫Glukhov.SetParent(f)
	
	f.科诺托普Konotop = BKonotop科诺托普
	f.科诺托普Konotop.SetParent(f)
	
	f.诺夫哥罗德_谢韦尔斯基Novgorodseversky = BNovgorodseversky诺夫哥罗德_谢韦尔斯基
	f.诺夫哥罗德_谢韦尔斯基Novgorodseversky.SetParent(f)
	
}
