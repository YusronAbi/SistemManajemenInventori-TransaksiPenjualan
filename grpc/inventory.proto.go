syntax = "proto3";

package inventory;

service InventoryService {
	rpc GetProperties (Empty) returns (PropertyList);
	rpc CreateTransaction (Transaction) returns (TransactionResponse);
}

message Empty {}

message Property {
	uint32 id = 1;
	string name = 2;
	string location = 3;
	int32 price = 4;
	string description = 5;
}

message PropertyList {
	repeated Property properties = 1;
}

message Transaction {
	uint32 id = 1;
	uint32 property_id = 2;
	string buyer_name = 3;
	string date = 4;
}

message TransactionResponse {
	bool success = 1;
}