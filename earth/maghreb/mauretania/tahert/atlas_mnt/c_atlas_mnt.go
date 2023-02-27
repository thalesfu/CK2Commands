package atlas_mnt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Atlas_mntCounty interface {
    feud.County
    BAinrich艾因里什() 	feud.Barony
    BAinzmila艾因斯马拉() 	feud.Barony
    BAskedal阿斯克达尔() 	feud.Barony
    BBordj布尔吉() 	feud.Barony
    BDjelfa杰勒法() 	feud.Barony
    BElyourh尤尔赫() 	feud.Barony
    BHassibenmadjeb哈西本马贾() 	feud.Barony
    BSelmana撒玛娜() 	feud.Barony
}

type 阿斯克达尔Atlas_mntCounty struct {
	feud.BaseCounty
	艾因里什Ainrich 	feud.Barony
	艾因斯马拉Ainzmila 	feud.Barony
	阿斯克达尔Askedal 	feud.Barony
	布尔吉Bordj 	feud.Barony
	杰勒法Djelfa 	feud.Barony
	尤尔赫Elyourh 	feud.Barony
	哈西本马贾Hassibenmadjeb 	feud.Barony
	撒玛娜Selmana 	feud.Barony
}

func (c *阿斯克达尔Atlas_mntCounty) BAinrich艾因里什() feud.Barony {
	return c.艾因里什Ainrich
}
    
func (c *阿斯克达尔Atlas_mntCounty) BAinzmila艾因斯马拉() feud.Barony {
	return c.艾因斯马拉Ainzmila
}
    
func (c *阿斯克达尔Atlas_mntCounty) BAskedal阿斯克达尔() feud.Barony {
	return c.阿斯克达尔Askedal
}
    
func (c *阿斯克达尔Atlas_mntCounty) BBordj布尔吉() feud.Barony {
	return c.布尔吉Bordj
}
    
func (c *阿斯克达尔Atlas_mntCounty) BDjelfa杰勒法() feud.Barony {
	return c.杰勒法Djelfa
}
    
func (c *阿斯克达尔Atlas_mntCounty) BElyourh尤尔赫() feud.Barony {
	return c.尤尔赫Elyourh
}
    
func (c *阿斯克达尔Atlas_mntCounty) BHassibenmadjeb哈西本马贾() feud.Barony {
	return c.哈西本马贾Hassibenmadjeb
}
    
func (c *阿斯克达尔Atlas_mntCounty) BSelmana撒玛娜() feud.Barony {
	return c.撒玛娜Selmana
}
    
var CAtlas_mnt阿斯克达尔 Atlas_mntCounty = &阿斯克达尔Atlas_mntCounty{}

func init() {
	f := CAtlas_mnt阿斯克达尔.(*阿斯克达尔Atlas_mntCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "833",
		Title:     "atlas_mnt",
		TitleName: "阿斯克达尔",
		TitleCode: "c_atlas_mnt",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因里什Ainrich = BAinrich艾因里什
	f.艾因里什Ainrich.SetParent(f)
	
	f.艾因斯马拉Ainzmila = BAinzmila艾因斯马拉
	f.艾因斯马拉Ainzmila.SetParent(f)
	
	f.阿斯克达尔Askedal = BAskedal阿斯克达尔
	f.阿斯克达尔Askedal.SetParent(f)
	
	f.布尔吉Bordj = BBordj布尔吉
	f.布尔吉Bordj.SetParent(f)
	
	f.杰勒法Djelfa = BDjelfa杰勒法
	f.杰勒法Djelfa.SetParent(f)
	
	f.尤尔赫Elyourh = BElyourh尤尔赫
	f.尤尔赫Elyourh.SetParent(f)
	
	f.哈西本马贾Hassibenmadjeb = BHassibenmadjeb哈西本马贾
	f.哈西本马贾Hassibenmadjeb.SetParent(f)
	
	f.撒玛娜Selmana = BSelmana撒玛娜
	f.撒玛娜Selmana.SetParent(f)
	
}
