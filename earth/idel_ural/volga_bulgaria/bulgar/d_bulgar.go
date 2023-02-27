package bulgar

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/bulgar/ashli"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/bulgar/bilyar"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/bulgar/bulgar"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/bulgar/naberezhnye"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/bulgar/syrt"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BulgarDuke interface {
    feud.Duke
    CAshli阿什勒() 	ashli.AshliCounty
    CBilyar比利亚尔() 	bilyar.BilyarCounty
    CBulgar保加尔() 	bulgar.BulgarCounty
    CNaberezhnye纳别列日内耶() 	naberezhnye.NaberezhnyeCounty
    CSyrt瑟尔特() 	syrt.SyrtCounty
}

type 保加尔BulgarDuke struct {
	feud.BaseDuke
	阿什勒Ashli 	ashli.AshliCounty
	比利亚尔Bilyar 	bilyar.BilyarCounty
	保加尔Bulgar 	bulgar.BulgarCounty
	纳别列日内耶Naberezhnye 	naberezhnye.NaberezhnyeCounty
	瑟尔特Syrt 	syrt.SyrtCounty
}

func (d *保加尔BulgarDuke) CAshli阿什勒() ashli.AshliCounty {
	return d.阿什勒Ashli
}
    
func (d *保加尔BulgarDuke) CBilyar比利亚尔() bilyar.BilyarCounty {
	return d.比利亚尔Bilyar
}
    
func (d *保加尔BulgarDuke) CBulgar保加尔() bulgar.BulgarCounty {
	return d.保加尔Bulgar
}
    
func (d *保加尔BulgarDuke) CNaberezhnye纳别列日内耶() naberezhnye.NaberezhnyeCounty {
	return d.纳别列日内耶Naberezhnye
}
    
func (d *保加尔BulgarDuke) CSyrt瑟尔特() syrt.SyrtCounty {
	return d.瑟尔特Syrt
}
    
var DBulgar保加尔 BulgarDuke = &保加尔BulgarDuke{}

func init() {
	f := DBulgar保加尔.(*保加尔BulgarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bulgar",
		TitleName: "保加尔",
		TitleCode: "d_bulgar",
		Counties:  map[string]feud.County{},
	}

	f.阿什勒Ashli = ashli.CAshli阿什勒
	f.阿什勒Ashli.SetParent(f)
	
	f.比利亚尔Bilyar = bilyar.CBilyar比利亚尔
	f.比利亚尔Bilyar.SetParent(f)
	
	f.保加尔Bulgar = bulgar.CBulgar保加尔
	f.保加尔Bulgar.SetParent(f)
	
	f.纳别列日内耶Naberezhnye = naberezhnye.CNaberezhnye纳别列日内耶
	f.纳别列日内耶Naberezhnye.SetParent(f)
	
	f.瑟尔特Syrt = syrt.CSyrt瑟尔特
	f.瑟尔特Syrt.SetParent(f)
	
}
