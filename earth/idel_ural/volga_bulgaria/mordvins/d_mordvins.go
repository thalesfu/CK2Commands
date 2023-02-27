package mordvins

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/mordvins/burtasy"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/mordvins/khopyor"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/mordvins/mordva"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MordvinsDuke interface {
    feud.Duke
    CBurtasy布尔塔瑟() 	burtasy.BurtasyCounty
    CKhopyor霍皮奥尔() 	khopyor.KhopyorCounty
    CMordva莫尔多瓦() 	mordva.MordvaCounty
}

type 莫尔多瓦MordvinsDuke struct {
	feud.BaseDuke
	布尔塔瑟Burtasy 	burtasy.BurtasyCounty
	霍皮奥尔Khopyor 	khopyor.KhopyorCounty
	莫尔多瓦Mordva 	mordva.MordvaCounty
}

func (d *莫尔多瓦MordvinsDuke) CBurtasy布尔塔瑟() burtasy.BurtasyCounty {
	return d.布尔塔瑟Burtasy
}
    
func (d *莫尔多瓦MordvinsDuke) CKhopyor霍皮奥尔() khopyor.KhopyorCounty {
	return d.霍皮奥尔Khopyor
}
    
func (d *莫尔多瓦MordvinsDuke) CMordva莫尔多瓦() mordva.MordvaCounty {
	return d.莫尔多瓦Mordva
}
    
var DMordvins莫尔多瓦 MordvinsDuke = &莫尔多瓦MordvinsDuke{}

func init() {
	f := DMordvins莫尔多瓦.(*莫尔多瓦MordvinsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mordvins",
		TitleName: "莫尔多瓦",
		TitleCode: "d_mordvins",
		Counties:  map[string]feud.County{},
	}

	f.布尔塔瑟Burtasy = burtasy.CBurtasy布尔塔瑟
	f.布尔塔瑟Burtasy.SetParent(f)
	
	f.霍皮奥尔Khopyor = khopyor.CKhopyor霍皮奥尔
	f.霍皮奥尔Khopyor.SetParent(f)
	
	f.莫尔多瓦Mordva = mordva.CMordva莫尔多瓦
	f.莫尔多瓦Mordva.SetParent(f)
	
}
