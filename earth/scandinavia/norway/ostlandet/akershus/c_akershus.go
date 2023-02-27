package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AkershusCounty interface {
    feud.County
    BAkershus阿克什胡斯() 	feud.Barony
    BBaerum拜鲁姆() 	feud.Barony
    BBergheim贝格海姆() 	feud.Barony
    BIsegran伊瑟格兰() 	feud.Barony
    BMoss莫斯() 	feud.Barony
    BNes内斯() 	feud.Barony
    BOslo奥斯陆() 	feud.Barony
}

type 阿克什胡斯AkershusCounty struct {
	feud.BaseCounty
	阿克什胡斯Akershus 	feud.Barony
	拜鲁姆Baerum 	feud.Barony
	贝格海姆Bergheim 	feud.Barony
	伊瑟格兰Isegran 	feud.Barony
	莫斯Moss 	feud.Barony
	内斯Nes 	feud.Barony
	奥斯陆Oslo 	feud.Barony
}

func (c *阿克什胡斯AkershusCounty) BAkershus阿克什胡斯() feud.Barony {
	return c.阿克什胡斯Akershus
}
    
func (c *阿克什胡斯AkershusCounty) BBaerum拜鲁姆() feud.Barony {
	return c.拜鲁姆Baerum
}
    
func (c *阿克什胡斯AkershusCounty) BBergheim贝格海姆() feud.Barony {
	return c.贝格海姆Bergheim
}
    
func (c *阿克什胡斯AkershusCounty) BIsegran伊瑟格兰() feud.Barony {
	return c.伊瑟格兰Isegran
}
    
func (c *阿克什胡斯AkershusCounty) BMoss莫斯() feud.Barony {
	return c.莫斯Moss
}
    
func (c *阿克什胡斯AkershusCounty) BNes内斯() feud.Barony {
	return c.内斯Nes
}
    
func (c *阿克什胡斯AkershusCounty) BOslo奥斯陆() feud.Barony {
	return c.奥斯陆Oslo
}
    
var CAkershus阿克什胡斯 AkershusCounty = &阿克什胡斯AkershusCounty{}

func init() {
	f := CAkershus阿克什胡斯.(*阿克什胡斯AkershusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "272",
		Title:     "akershus",
		TitleName: "阿克什胡斯",
		TitleCode: "c_akershus",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克什胡斯Akershus = BAkershus阿克什胡斯
	f.阿克什胡斯Akershus.SetParent(f)
	
	f.拜鲁姆Baerum = BBaerum拜鲁姆
	f.拜鲁姆Baerum.SetParent(f)
	
	f.贝格海姆Bergheim = BBergheim贝格海姆
	f.贝格海姆Bergheim.SetParent(f)
	
	f.伊瑟格兰Isegran = BIsegran伊瑟格兰
	f.伊瑟格兰Isegran.SetParent(f)
	
	f.莫斯Moss = BMoss莫斯
	f.莫斯Moss.SetParent(f)
	
	f.内斯Nes = BNes内斯
	f.内斯Nes.SetParent(f)
	
	f.奥斯陆Oslo = BOslo奥斯陆
	f.奥斯陆Oslo.SetParent(f)
	
}
