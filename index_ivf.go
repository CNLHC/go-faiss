package faiss

/*
#include <faiss/c_api/IndexIVFFlat_c.h>
#include <faiss/c_api/MetaIndexes_c.h>
#include <faiss/c_api/Index_c.h>
#include <faiss/c_api/IndexIVF_c.h>
#include <faiss/c_api/IndexIVF_c_ex.h>
*/
import "C"
import "fmt"

func (idx *IndexImpl) SetDirectMap(mapType int) (err error) {
	ivfPtr := C.faiss_IndexIVF_cast(idx.cPtr())
	if ivfPtr == nil {
		return fmt.Errorf("index is not of ivf type")
	}
	if c := C.faiss_IndexIVF_set_direct_map(
		ivfPtr,
		C.int(mapType),
	); c != 0 {
		err = getLastError()
	}
	return err
}

func (idx *IndexImpl) GetSubIndex() (*IndexImpl, error) {
	ptr := C.faiss_IndexIDMap2_cast(idx.cPtr())
	if ptr == nil {
		return nil, fmt.Errorf("index is not a id map")
	}

	subIdx := C.faiss_IndexIDMap2_sub_index(ptr)
	if subIdx == nil {
		return nil, fmt.Errorf("couldn't retrieve the sub index")
	}

	return &IndexImpl{&faissIndex{subIdx}}, nil
}
