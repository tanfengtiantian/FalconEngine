package invert

import (
	"FalconEngine/index/dict"
	"FalconEngine/message"
	"FalconEngine/mlog"
	"FalconEngine/store"
	"fmt"
	"sync"
)

type InvertWriter struct {
	name string

	tmpInvert    map[string]FalconDocList
	invertLocker *sync.RWMutex
}

func NewStringInvertWriter(name string) FalconStringInvertWriteService {
	writer := &InvertWriter{name: name, tmpInvert: make(map[string]FalconDocList), invertLocker: new(sync.RWMutex)}

	return writer
}

func (iw *InvertWriter) Put(key string, docid *message.DocId) error {
	iw.invertLocker.Lock()
	defer iw.invertLocker.Unlock()
	if _, ok := iw.tmpInvert[key]; !ok {
		iw.tmpInvert[key] = NewMemoryFalconDocList()
	}

	return iw.tmpInvert[key].Push(docid)

}

func (iw *InvertWriter) Persistence(invertListStore, dictStore store.FalconSearchStoreWriteService) (int64, error) {

	//invertListStore := store.NewFalconFileStoreWriteService(iw.path + "/" + iw.name + ".ivt")
	//dictStore := store.NewFalconFileStoreWriteService(iw.path + "/" + iw.name + ".dic")
	dictMap := dict.NewFalconWriteMap()

	for key, v := range iw.tmpInvert {

		by, _ := v.FalconEncoding()
		pos, err := invertListStore.AppendBytes(by)
		if err != nil {
			mlog.Error("Write Error : %v", err)
			return -1, err
		}
		dictMap.Put(key, &message.DictValue{Offset: uint64(pos), Length: uint64(v.GetLength())})
	}

	return dictMap.Persistence(dictStore)

	//offset,err:=dictStore.AppendMessage(dictMap)
	//if err != nil {
	//	mlog.Error("Write Error : %v",err)
	//	return -1,err
	//}

	//invertListStore.Close()
	//dictStore.Close()
	//return offset,err
}

func (iw *InvertWriter) ToString() string {

	result := "\n"
	for key, v := range iw.tmpInvert {
		result = result + fmt.Sprintf("[ %s ] >> %s \n", key, v.ToString())
	}
	return result
}
