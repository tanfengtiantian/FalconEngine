package dict

import (
	"FalconEngine/message"
	"FalconEngine/store"
	"FalconEngine/tools"
)

type FalconDictWriteService interface {
	Put(key interface{}, value interface{}) error
}

// 字典服务
type FalconStringDictService interface {
	Put(key string, dv *message.DictValue) error
	Get(key string) (*message.DictValue, bool)
	Persistence(storeService store.FalconSearchStoreWriteService) (int64, error)
	tools.FalconSearchEncoder
}

type FalconStringDictWriteService interface {
	Put(key string, dv *message.DictValue) error
	Persistence(storeService store.FalconSearchStoreWriteService) (int64, error)
	tools.FalconSearchEncoder
}

type FalconStringDictReadService interface {
	Get(key string) (*message.DictValue, bool)
	LoadDic(storeService store.FalconSearchStoreReadService, offset int64) error
	tools.FalconSearchDecoder
}
