load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/googleapis/gapic-generator-go
gazelle(name = "gazelle")

# Note: Direct gazelle to resolve common proto imports to googleapis and the Go
# targets to the go_proto_library targets.
#
# gazelle:resolve proto proto google/rpc/code.proto @com_google_googleapis//google/rpc:code_proto
# gazelle:resolve proto go google/rpc/code.proto @com_google_googleapis//google/rpc:code_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/rpc/code @com_google_googleapis//google/rpc:code_go_proto
# gazelle:resolve proto proto google/api/annotations.proto @com_google_googleapis//google/api:annotations_proto
# gazelle:resolve proto go google/api/annotations.proto @com_google_googleapis//google/api:annotations_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/api/annotations @com_google_googleapis//google/api:annotations_go_proto
# gazelle:resolve proto proto google/longrunning/operations.proto @com_google_googleapis//google/longrunning:operations_proto
# gazelle:resolve proto go google/longrunning/operations.proto @com_google_googleapis//google/longrunning:longrunning_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/cloud/extendedops @com_google_googleapis//google/cloud:extended_operations_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/cloud/location @com_google_googleapis//google/cloud/location:location_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/gapic/metadata @com_google_googleapis//gapic/metadata:metadata_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/api/httpbody @com_google_googleapis//google/api:httpbody_go_proto
# gazelle:resolve go go google.golang.org/genproto/googleapis/api/serviceconfig @com_google_googleapis//google/api:serviceconfig_go_proto

# Note: Direct gazelle to resolve the protobuf-go v1 module to rules_go well
# known types. This will be removed when we migrate to protobuf-go v2.
#
# gazelle:resolve go go github.com/golang/protobuf/protoc-gen-go/plugin @io_bazel_rules_go//proto/wkt:compiler_plugin_go_proto
# gazelle:resolve go go github.com/golang/protobuf/protoc-gen-go/descriptor @io_bazel_rules_go//proto/wkt:descriptor_go_proto
# gazelle:resolve go go github.com/golang/protobuf/ptypes/duration @io_bazel_rules_go//proto/wkt:duration_go_proto
