// source: common/types.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.containersai.common.ColumnType', null, global);
goog.exportSymbol('proto.containersai.common.DataType', null, global);
/**
 * @enum {number}
 */
proto.containersai.common.DataType = {
  DATATYPE_UNDEFINED: 0,
  DATATYPE_BOOL: 1,
  DATATYPE_INT: 2,
  DATATYPE_INT8: 3,
  DATATYPE_INT16: 4,
  DATATYPE_INT32: 5,
  DATATYPE_INT64: 6,
  DATATYPE_UINT: 7,
  DATATYPE_UINT8: 8,
  DATATYPE_UINT16: 9,
  DATATYPE_UINT32: 10,
  DATATYPE_UTIN64: 11,
  DATATYPE_FLOAT32: 12,
  DATATYPE_FLOAT64: 13,
  DATATYPE_STRING: 14
};

/**
 * @enum {number}
 */
proto.containersai.common.ColumnType = {
  COLUMNTYPE_UDEFINED: 0,
  COLUMNTYPE_TAG: 1,
  COLUMNTYPE_FIELD: 2
};

goog.object.extend(exports, proto.containersai.common);
