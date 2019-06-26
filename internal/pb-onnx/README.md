This files are generated with gogo-protobuf

Some gadgets are manually added to the basic definition file of onnx: 

```diff
diff --git a/onnx/onnx.proto3 b/onnx/onnx.proto3
index 87bfadc2..99d1fe48 100644
--- a/onnx/onnx.proto3
+++ b/onnx/onnx.proto3
@@ -10,6 +10,15 @@ syntax = "proto3";

 package onnx;

+
+import "github.com/gogo/protobuf/gogoproto/gogo.proto";
+
+option (gogoproto.equal_all) = true;
+option (gogoproto.testgen_all) = true;
+option (gogoproto.benchgen_all) = true;
+option (gogoproto.populate_all) = true;
+
+
 // Overview
 //
 // ONNX is an open specification that is comprised of the following components:
 ```

 Then the files are generated with:

 ```
protoc -I=~/onnx/onnx -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogofaster_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:. \
~/onnx/onnx/onnx.proto3
```
