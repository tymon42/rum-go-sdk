package rumgosdk

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/h2non/filetype"
)

const chunkSize = 20 << 10 // 20KB < 25KB = 200Kb
// const chunkSize = 512 << 10 // 0.5MB

type Segment struct {
	TrxId     string
	SplitHash string
}

type Sha256MapValue struct {
	Staut   bool
	Order   int64
	Segment Segment
}

type RumFileInfo struct {
	MeadiaType string
	Title      string
	Sha256     string
	ChunkSize  int
	Size       int64
	Segments   []Segment
}

func (q *Quorum) UploadFile(path string, groupId string) (*RumFileInfo, error) {
	rumFileInfo, err := readFile(path)
	if err != nil {
		return nil, err
	}
	num := math.Ceil(float64(rumFileInfo.Size) / float64(rumFileInfo.ChunkSize))
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	b := make([]byte, chunkSize)
	var i int64 = 1
	for ; i < int64(num); i++ {
		file.Seek(i*chunkSize, 0)
		if len(b) > int(rumFileInfo.Size-i*chunkSize) {
			b = make([]byte, rumFileInfo.Size-i*chunkSize)
		}
		file.Read(b)

		// hash
		sha256byte := sha256.Sum256(b)
		splitHash := hex.EncodeToString(sha256byte[:])

		cont := base64.StdEncoding.EncodeToString(b)

		// upload file split
		trxId, err := q.postFileSplit(cont, splitHash, groupId)
		if err != nil {
			return nil, err
		}
		for {
			ok, err := q.checkSplitTrxId(groupId, trxId)
			if err == nil && ok {
				rumFileInfo.Segments = append(rumFileInfo.Segments, Segment{TrxId: trxId, SplitHash: splitHash})
				break
			} else {
				time.Sleep(time.Second / 2)
			}
		}

	}
	return rumFileInfo, nil
}

func readFile(path string) (*RumFileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file not exist")
		}
		return nil, errors.New("file stat error")
	}

	rumFileInfo := &RumFileInfo{}
	rumFileInfo.ChunkSize = chunkSize
	rumFileInfo.Size = fileInfo.Size()
	rumFileInfo.Title = fileInfo.Name()

	buf, _ := ioutil.ReadFile(path)
	kind, _ := filetype.Match(buf)
	rumFileInfo.MeadiaType = kind.MIME.Value
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sha256byte := sha256.Sum256(f)
	rumFileInfo.Sha256 = hex.EncodeToString(sha256byte[:])
	return rumFileInfo, nil
}
