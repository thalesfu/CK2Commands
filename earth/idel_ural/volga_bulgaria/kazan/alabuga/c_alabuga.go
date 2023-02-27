package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlabugaCounty interface {
    feud.County
    BAgryz阿格雷兹() 	feud.Barony
    BAlabuga阿拉布加() 	feud.Barony
    BBima比马() 	feud.Barony
    BLyali利亚利() 	feud.Barony
    BMozhga莫日加() 	feud.Barony
    BSarapul萨拉普尔() 	feud.Barony
    BSiva锡瓦() 	feud.Barony
}

type 阿拉布加AlabugaCounty struct {
	feud.BaseCounty
	阿格雷兹Agryz 	feud.Barony
	阿拉布加Alabuga 	feud.Barony
	比马Bima 	feud.Barony
	利亚利Lyali 	feud.Barony
	莫日加Mozhga 	feud.Barony
	萨拉普尔Sarapul 	feud.Barony
	锡瓦Siva 	feud.Barony
}

func (c *阿拉布加AlabugaCounty) BAgryz阿格雷兹() feud.Barony {
	return c.阿格雷兹Agryz
}
    
func (c *阿拉布加AlabugaCounty) BAlabuga阿拉布加() feud.Barony {
	return c.阿拉布加Alabuga
}
    
func (c *阿拉布加AlabugaCounty) BBima比马() feud.Barony {
	return c.比马Bima
}
    
func (c *阿拉布加AlabugaCounty) BLyali利亚利() feud.Barony {
	return c.利亚利Lyali
}
    
func (c *阿拉布加AlabugaCounty) BMozhga莫日加() feud.Barony {
	return c.莫日加Mozhga
}
    
func (c *阿拉布加AlabugaCounty) BSarapul萨拉普尔() feud.Barony {
	return c.萨拉普尔Sarapul
}
    
func (c *阿拉布加AlabugaCounty) BSiva锡瓦() feud.Barony {
	return c.锡瓦Siva
}
    
var CAlabuga阿拉布加 AlabugaCounty = &阿拉布加AlabugaCounty{}

func init() {
	f := CAlabuga阿拉布加.(*阿拉布加AlabugaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1717",
		Title:     "alabuga",
		TitleName: "阿拉布加",
		TitleCode: "c_alabuga",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格雷兹Agryz = BAgryz阿格雷兹
	f.阿格雷兹Agryz.SetParent(f)
	
	f.阿拉布加Alabuga = BAlabuga阿拉布加
	f.阿拉布加Alabuga.SetParent(f)
	
	f.比马Bima = BBima比马
	f.比马Bima.SetParent(f)
	
	f.利亚利Lyali = BLyali利亚利
	f.利亚利Lyali.SetParent(f)
	
	f.莫日加Mozhga = BMozhga莫日加
	f.莫日加Mozhga.SetParent(f)
	
	f.萨拉普尔Sarapul = BSarapul萨拉普尔
	f.萨拉普尔Sarapul.SetParent(f)
	
	f.锡瓦Siva = BSiva锡瓦
	f.锡瓦Siva.SetParent(f)
	
}
