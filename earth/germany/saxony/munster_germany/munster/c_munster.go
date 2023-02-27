package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MunsterCounty interface {
    feud.County
    BAhlen阿伦() 	feud.Barony
    BDortmund多特蒙德() 	feud.Barony
    BEssen埃森() 	feud.Barony
    BGreven格雷文() 	feud.Barony
    BGronau格罗瑙() 	feud.Barony
    BGutersloh居特斯洛() 	feud.Barony
    BMunster明斯特() 	feud.Barony
    BSteinfurt施泰因富特() 	feud.Barony
}

type 明斯特MunsterCounty struct {
	feud.BaseCounty
	阿伦Ahlen 	feud.Barony
	多特蒙德Dortmund 	feud.Barony
	埃森Essen 	feud.Barony
	格雷文Greven 	feud.Barony
	格罗瑙Gronau 	feud.Barony
	居特斯洛Gutersloh 	feud.Barony
	明斯特Munster 	feud.Barony
	施泰因富特Steinfurt 	feud.Barony
}

func (c *明斯特MunsterCounty) BAhlen阿伦() feud.Barony {
	return c.阿伦Ahlen
}
    
func (c *明斯特MunsterCounty) BDortmund多特蒙德() feud.Barony {
	return c.多特蒙德Dortmund
}
    
func (c *明斯特MunsterCounty) BEssen埃森() feud.Barony {
	return c.埃森Essen
}
    
func (c *明斯特MunsterCounty) BGreven格雷文() feud.Barony {
	return c.格雷文Greven
}
    
func (c *明斯特MunsterCounty) BGronau格罗瑙() feud.Barony {
	return c.格罗瑙Gronau
}
    
func (c *明斯特MunsterCounty) BGutersloh居特斯洛() feud.Barony {
	return c.居特斯洛Gutersloh
}
    
func (c *明斯特MunsterCounty) BMunster明斯特() feud.Barony {
	return c.明斯特Munster
}
    
func (c *明斯特MunsterCounty) BSteinfurt施泰因富特() feud.Barony {
	return c.施泰因富特Steinfurt
}
    
var CMunster明斯特 MunsterCounty = &明斯特MunsterCounty{}

func init() {
	f := CMunster明斯特.(*明斯特MunsterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "88",
		Title:     "munster",
		TitleName: "明斯特",
		TitleCode: "c_munster",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伦Ahlen = BAhlen阿伦
	f.阿伦Ahlen.SetParent(f)
	
	f.多特蒙德Dortmund = BDortmund多特蒙德
	f.多特蒙德Dortmund.SetParent(f)
	
	f.埃森Essen = BEssen埃森
	f.埃森Essen.SetParent(f)
	
	f.格雷文Greven = BGreven格雷文
	f.格雷文Greven.SetParent(f)
	
	f.格罗瑙Gronau = BGronau格罗瑙
	f.格罗瑙Gronau.SetParent(f)
	
	f.居特斯洛Gutersloh = BGutersloh居特斯洛
	f.居特斯洛Gutersloh.SetParent(f)
	
	f.明斯特Munster = BMunster明斯特
	f.明斯特Munster.SetParent(f)
	
	f.施泰因富特Steinfurt = BSteinfurt施泰因富特
	f.施泰因富特Steinfurt.SetParent(f)
	
}
