package segment

import (
	"FalconEngine/mlog"
	"FalconEngine/tools"
	"fmt"
	"testing"
)

func Test_Segment(t *testing.T) {

	mlog.Start(mlog.LevelInfo, "iw.log")

	mappings := tools.NewFalconIndexMappings()
	mappings.AddFieldMapping(&tools.FalconMapping{FieldName: "field1", FieldType: tools.TKeywordType})
	mappings.AddFieldMapping(&tools.FalconMapping{FieldName: "field2", FieldType: tools.TKeywordType})

	sg := NewFalconSegment(uint32(0), "test_index", "./test_segment", mappings)

	for i := 0; i < 100; i++ {
		doc := make(map[string]interface{})
		doc["field1"] = fmt.Sprintf("dfdsf%d", i)
		doc["field2"] = []string{fmt.Sprintf("1112"), fmt.Sprintf("2222f%d", i), fmt.Sprintf("3333%d", i)}

		sg.UpdateDocument(doc)

	}

	sg.Persistence()
	l, _, _ := sg.SimpleSearch("field2", "1112")
	mlog.Info("%s", l.ToString())
	sg.Close()
}

func Test_LoadSegment(t *testing.T) {

	mappings := tools.NewFalconIndexMappings()
	mappings.AddFieldMapping(&tools.FalconMapping{FieldName: "field1", FieldType: tools.TKeywordType})
	mappings.AddFieldMapping(&tools.FalconMapping{FieldName: "field2", FieldType: tools.TKeywordType})

	sg := LoadFalconSegment(uint32(0), "test_index", "./test_segment", mappings)

	l, _, _ := sg.SimpleSearch("field2", "1112")
	mlog.Info("%s", l.ToString())
	sg.Close()
}
