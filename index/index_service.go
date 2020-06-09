package index

import "FalconEngine/tools"

type FalconIndexService interface {

	// 新建mapping
	CreateMappings(mappings *tools.FalconIndexMappings) error

	// 添加文档
	UpdateDocument(documentID string, document map[string]interface{}) error

	// 删除文档
	DeleteDocument(documentID string) error

	// 查询
	//SearchDocuments(query string) error

	//GetDocument(documentID string) (map[string]interface{},error)

}
