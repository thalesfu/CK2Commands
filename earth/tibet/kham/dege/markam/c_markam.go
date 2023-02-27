package markam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarkamCounty interface {
    feud.County
    BCuowa措瓦() 	feud.Barony
    BMangling莽岭() 	feud.Barony
    BMarkam马儿敢() 	feud.Barony
    BQudeng曲登() 	feud.Barony
    BWangda旺达() 	feud.Barony
    BZhonglin中林() 	feud.Barony
    BZogong左贡() 	feud.Barony
}

type 马儿敢MarkamCounty struct {
	feud.BaseCounty
	措瓦Cuowa 	feud.Barony
	莽岭Mangling 	feud.Barony
	马儿敢Markam 	feud.Barony
	曲登Qudeng 	feud.Barony
	旺达Wangda 	feud.Barony
	中林Zhonglin 	feud.Barony
	左贡Zogong 	feud.Barony
}

func (c *马儿敢MarkamCounty) BCuowa措瓦() feud.Barony {
	return c.措瓦Cuowa
}
    
func (c *马儿敢MarkamCounty) BMangling莽岭() feud.Barony {
	return c.莽岭Mangling
}
    
func (c *马儿敢MarkamCounty) BMarkam马儿敢() feud.Barony {
	return c.马儿敢Markam
}
    
func (c *马儿敢MarkamCounty) BQudeng曲登() feud.Barony {
	return c.曲登Qudeng
}
    
func (c *马儿敢MarkamCounty) BWangda旺达() feud.Barony {
	return c.旺达Wangda
}
    
func (c *马儿敢MarkamCounty) BZhonglin中林() feud.Barony {
	return c.中林Zhonglin
}
    
func (c *马儿敢MarkamCounty) BZogong左贡() feud.Barony {
	return c.左贡Zogong
}
    
var CMarkam马儿敢 MarkamCounty = &马儿敢MarkamCounty{}

func init() {
	f := CMarkam马儿敢.(*马儿敢MarkamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1505",
		Title:     "markam",
		TitleName: "马儿敢",
		TitleCode: "c_markam",
		Baronies:  map[string]feud.Barony{},
	}

	f.措瓦Cuowa = BCuowa措瓦
	f.措瓦Cuowa.SetParent(f)
	
	f.莽岭Mangling = BMangling莽岭
	f.莽岭Mangling.SetParent(f)
	
	f.马儿敢Markam = BMarkam马儿敢
	f.马儿敢Markam.SetParent(f)
	
	f.曲登Qudeng = BQudeng曲登
	f.曲登Qudeng.SetParent(f)
	
	f.旺达Wangda = BWangda旺达
	f.旺达Wangda.SetParent(f)
	
	f.中林Zhonglin = BZhonglin中林
	f.中林Zhonglin.SetParent(f)
	
	f.左贡Zogong = BZogong左贡
	f.左贡Zogong.SetParent(f)
	
}
