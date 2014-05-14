package main

import (
	http
)

const SERVER_URL = "https://epicwar-facebook.progrestar.net/rpc/"
const SESSION_ID = "0n5kd0n1p0k2z1"
const USER_ID    = "675063875"
const APP_ID     = "1424411677784893"
const AUTH_KEY   = "6514a97ae525f196b8060337380e0cbb"

public function getData() : Object {
         return {"calls":this._bodies};
      }

public function getFormattedData() : Object {
         return JSONS.stringify(this.getData());
      }

private function createFingerprint(param1:Object) : String {
 var _loc4_:String = null;
 var _loc5_:* = 0;
 var _loc6_:* = 0;
 var _loc7_:String = null;
 var _loc2_:Array = [];
 var _loc3_:* = "";
 for(_loc4_ in param1)
 {
    if(_loc4_.indexOf("X-Env") != -1)
    {
       _loc7_ = _loc4_.substr(6);
       _loc2_.push(
          {
             "key":_loc7_.toUpperCase(),
             "value":param1[_loc4_]
          });
    }
 }
 _loc2_.sortOn("key");
 _loc5_ = _loc2_.length;
 _loc6_ = 0;
 while(_loc6_ < _loc5_)
 {
    _loc3_ = _loc3_ + (_loc2_[_loc6_].key + "=" + _loc2_[_loc6_].value);
    _loc6_++;
 }
 return _loc3_;
}

func createAuthSignature (headers, rpc *RpcEntryBase) {
	var _loc5_:ByteArray = null;
         var _loc3_:Object = rpc.request.getFormattedData();
         var _loc4_:ByteArray = new ByteArray();
         _loc4_.writeUTFBytes(headers["X-Request-Id"]);
         _loc4_.writeUTFBytes(":");
         _loc4_.writeUTFBytes(AUTH_KEY);
         _loc4_.writeUTFBytes(":");
         _loc4_.writeUTFBytes(headers["X-Auth-Session-Id"]);
         _loc4_.writeUTFBytes(":");
         if(_loc3_ is ByteArray)
         {
            _loc5_ = _loc3_ as ByteArray;
            _loc5_.position = 0;
            _loc4_.writeBytes(_loc5_,0,_loc5_.length);
         }
         else if(_loc3_ is String)
         {
            _loc4_.writeUTFBytes(_loc3_ as String);
         }
         
         _loc4_.writeUTFBytes(":");
         _loc4_.writeUTFBytes(this.createFingerprint(headers));
         return _loc4_;
}

protected function createHeaders(rpc RpcEntryBase) Object {
	var _loc5_:String = null;
	var _loc2_:String = SocialAdapter.instance.flashVars["session_key"];
	var _loc3_:Object = 
	{
	   "Content-Type":"application/json; charset=UTF-8",
	   "X-Request-Id":++this.unionRequestID,
	   "X-Auth-Network-Ident":Env.NETWORK,
	   "X-Auth-Application-Id":SocialAdapter.instance.app_id,
	   "X-Auth-User-Id":SocialAdapter.instance.getPlayer().id,
	   "X-Auth-Session-Id":Env.sessionKey
	};
	if(_loc2_ != null)
	{
		_loc3_["X-Auth-Session-Key"] = _loc2_;
	}
	var _loc4_:Object = rpc.headers;
	if(_loc4_ != null)
	{
		for(_loc5_ in _loc4_)
		{
		   _loc3_[_loc5_] = _loc4_[_loc5_];
		}
	}
	return _loc3_;
}

func addHeaders(req Request, rpc RpcEntryBase) : void {
	var _loc4_:String = null;
	var _loc3_:Object = this.createHeaders(rpc);
	var _loc5_:ByteArray = this.createAuthSignature(_loc3_,rpc);
	_loc3_["X-Auth-Signature"] = MD5.hashBytes(_loc5_);
	for(_loc4_ in _loc3_)
	{
		req.requestHeaders.push(new URLRequestHeader(_loc4_,_loc3_[_loc4_].toString()));
	}
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("POST", SERVER_URL, bytes.NewReader(postData))
	addHeaders(req, rpc)
	resp, err := client.Do(req)
	defer resp.Body.Close()
}