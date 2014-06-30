package network

import (
	"net/http"
	"strings"
	"bytes"
	"crypto/md5"
	"io"
	"encoding/hex"
	"strconv"
	"io/ioutil"
	"log"
)

const SERVER_URL = "https://epicwar-facebook.progrestar.net/rpc/"
const APP_ID     = "1424411677784893"
const NETWORK    = "facebook"

var _unionRequestID int
var _uid string
var _sid string
var _akey string

func createFingerprint(headers map[string]string) string {
	var fingerprint bytes.Buffer
	preparedHeaders := []Pair{}

	for header, _ := range headers{
		if(strings.Index(header, "X-Env") != -1){
			preparedHeaders = append(preparedHeaders, Pair{
		        key   : strings.ToUpper(header[6:len(header)]),
		        value : headers[header]})
		}
	}

	sortByKey(preparedHeaders)

	count := len(preparedHeaders);
	i := 0;
	for i < count {
		fingerprint.WriteString(preparedHeaders[i].key)
		fingerprint.WriteString("=")
		fingerprint.WriteString(preparedHeaders[i].value)
		i++;
	}
	res := fingerprint.String()

	return res
}

func createAuthSignature (headers map[string]string, postData string) string {
    h := md5.New()
	io.WriteString(h, headers["X-Request-Id"])
	io.WriteString(h, ":")
    io.WriteString(h, _akey);
    io.WriteString(h, ":");
    io.WriteString(h, headers["X-Auth-Session-Id"]);
    io.WriteString(h, ":");
    io.WriteString(h, postData);
    io.WriteString(h, ":");
    io.WriteString(h, createFingerprint(headers));
    return hex.EncodeToString(h.Sum(nil))
}

func createHeaders() (map[string]string) {
	_unionRequestID = _unionRequestID + 1
	headers := map[string]string{
		"X-Request-Id":strconv.Itoa(_unionRequestID),
		"X-Auth-Network-Ident":NETWORK,
		"X-Auth-Application-Id":APP_ID,
		"X-Auth-User-Id":_uid,
		"X-Auth-Session-Id":_sid,
		"X-Env-Library-Version": "0"}
	return headers
}

func addHeaders(req *http.Request, postData string) {
	headers := createHeaders();
	headers["X-Auth-Signature"] = createAuthSignature(headers, postData);

	for index, header := range headers {
		req.Header.Add(index, header)
	}
}

func Post(postData []byte) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("POST", SERVER_URL, bytes.NewReader(postData))
	addHeaders(req, string(postData))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if(err != nil){
		log.Fatal(err)
	}

	return body
}

func Init(Uid string, Sid string, Akey string) {
	_unionRequestID = 0
	_uid = Uid
	_sid = Sid
	_akey = Akey
}
