package applications
 
import (
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
)

type hashService struct { }

func NewHashService() *hashService {
	return &hashService{}
}

func (a *hashService) GetSqlHash(sqlFilePath string) (string, error) {
	content, err := ioutil.ReadFile(sqlFilePath)  // ioutil close file after reading
	if err != nil {
        	return "", err
    	}
	hash_md5 := md5.Sum([]byte(string(content)))
	string_hash_md5 := hex.EncodeToString(hash_md5[:])
	return string_hash_md5, nil
}