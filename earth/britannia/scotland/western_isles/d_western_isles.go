package western_isles

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/western_isles/argyll"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/western_isles/innse_gall"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/western_isles/iona"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Western_islesDuke interface {
    feud.Duke
    CArgyll阿盖尔() 	argyll.ArgyllCounty
    CInnse_gall赫布里底群岛() 	innse_gall.Innse_gallCounty
    CIona艾奥纳() 	iona.IonaCounty
}

type 西部群岛Western_islesDuke struct {
	feud.BaseDuke
	阿盖尔Argyll 	argyll.ArgyllCounty
	赫布里底群岛Innse_gall 	innse_gall.Innse_gallCounty
	艾奥纳Iona 	iona.IonaCounty
}

func (d *西部群岛Western_islesDuke) CArgyll阿盖尔() argyll.ArgyllCounty {
	return d.阿盖尔Argyll
}
    
func (d *西部群岛Western_islesDuke) CInnse_gall赫布里底群岛() innse_gall.Innse_gallCounty {
	return d.赫布里底群岛Innse_gall
}
    
func (d *西部群岛Western_islesDuke) CIona艾奥纳() iona.IonaCounty {
	return d.艾奥纳Iona
}
    
var DWestern_isles西部群岛 Western_islesDuke = &西部群岛Western_islesDuke{}

func init() {
	f := DWestern_isles西部群岛.(*西部群岛Western_islesDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "western_isles",
		TitleName: "西部群岛",
		TitleCode: "d_western_isles",
		Counties:  map[string]feud.County{},
	}

	f.阿盖尔Argyll = argyll.CArgyll阿盖尔
	f.阿盖尔Argyll.SetParent(f)
	
	f.赫布里底群岛Innse_gall = innse_gall.CInnse_gall赫布里底群岛
	f.赫布里底群岛Innse_gall.SetParent(f)
	
	f.艾奥纳Iona = iona.CIona艾奥纳
	f.艾奥纳Iona.SetParent(f)
	
}
