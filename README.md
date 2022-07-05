# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [gatewayproto/gatewaypb.proto](#gatewayproto_gatewaypb-proto)
    - [CreateUserRequest](#-CreateUserRequest)
    - [CreateUserResponse](#-CreateUserResponse)
    - [CreateUserResponse.Data](#-CreateUserResponse-Data)
    - [PingRequest](#-PingRequest)
    - [PingResponse](#-PingResponse)
  
    - [GrpcGateway](#-GrpcGateway)
  
- [Scalar Value Types](#scalar-value-types)



<a name="gatewayproto_gatewaypb-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gatewayproto/gatewaypb.proto



<a name="-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| age | [int32](#int32) |  |  |






<a name="-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| messsage | [string](#string) |  |  |
| data | [CreateUserResponse.Data](#CreateUserResponse-Data) |  |  |






<a name="-CreateUserResponse-Data"></a>

### CreateUserResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |






<a name="-PingRequest"></a>

### PingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="-PingResponse"></a>

### PingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |





 

 

 


<a name="-GrpcGateway"></a>

### GrpcGateway


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CheckPing | [.PingRequest](#PingRequest) | [.PingResponse](#PingResponse) |  |
| CreateUser | [.CreateUserRequest](#CreateUserRequest) | [.CreateUserResponse](#CreateUserResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |


<h2> For Test</h2>
<p> - curl -X POST "http://localhost:8081/ping" -d '{"message":"Ping"}'  </p>
<p> - curl -d '{"name":"you name", "age": 18}' -X POST 'http://localhost:8081/api/v1/users/add' </p>