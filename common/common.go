package common

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"io"
	"log"
	"strconv"
	"sync"
)

type (
	ConnectOption struct {
		Host       string `json:"host"`
		Port       uint32 `json:"port"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Key        []byte `json:"key"`
		PassPhrase []byte `json:"pass_phrase"`
	}
	ConnectOptionWithIdentity struct {
		Identity string `json:"identity"`
		ConnectOption
	}
	TunnelOption struct {
		SrcIp   string `json:"src_ip" validate:"required,ip"`
		SrcPort uint   `json:"src_port" validate:"required,numeric"`
		DstIp   string `json:"dst_ip" validate:"required,ip"`
		DstPort uint   `json:"dst_port" validate:"required,numeric"`
	}
	ConfigOption struct {
		Connect map[string]*ConnectOption  `json:"connect"`
		Tunnel  map[string]*[]TunnelOption `json:"tunnel"`
	}
)

var (
	db      *leveldb.DB
	bufpool *sync.Pool
)

// Init buffer pool
func InitBufPool() {
	bufpool = &sync.Pool{}
	bufpool.New = func() interface{} {
		return make([]byte, 64*1024)
	}
}

// Use buffer pool io copy
func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	buf := bufpool.Get().([]byte)
	defer bufpool.Put(buf)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}

// Initialize leveldb
func InitLevelDB(path string) {
	var err error
	db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Set up temporary storage
func SetTemporary(config ConfigOption) (err error) {
	data, err := json.Marshal(config)
	err = db.Put([]byte("temporary"), data, nil)
	return
}

// Get temporary storage
func GetTemporary() (config ConfigOption, err error) {
	exists, err := db.Has([]byte("temporary"), nil)
	if exists == false {
		config = ConfigOption{}
		return
	}
	data, err := db.Get([]byte("temporary"), nil)
	err = json.Unmarshal(data, &config)
	return
}

// Get Addr
func GetAddr(ip string, port uint) string {
	return ip + ":" + strconv.Itoa(int(port))
}
